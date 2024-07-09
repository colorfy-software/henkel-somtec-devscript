package mqttx

import (
	"crypto/tls"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const messageBufferSize = 10

type ConnectHandler func(client Client)
type ReconnectHandler func(client Client, opts *ClientOptions)
type ConnectionLostHandler func(client Client, err error)

type ClientOptions struct {
	Servers                 []*url.URL
	ClientID                string
	Username                string
	Password                string
	CredentialsProvider     mqtt.CredentialsProvider
	CleanSession            bool
	Order                   bool
	WillEnabled             bool
	WillTopic               string
	WillPayload             []byte
	WillQos                 byte
	WillRetained            bool
	ProtocolVersion         uint
	protocolVersionExplicit bool
	TLSConfig               *tls.Config
	KeepAlive               int64
	PingTimeout             time.Duration
	ConnectTimeout          time.Duration
	MaxReconnectInterval    time.Duration
	AutoReconnect           bool
	ConnectRetryInterval    time.Duration
	ConnectRetry            bool
	Store                   mqtt.Store
	DefaultPublishHandler   mqtt.MessageHandler
	OnConnect               ConnectHandler
	OnConnectionLost        ConnectionLostHandler
	OnReconnecting          ReconnectHandler
	WriteTimeout            time.Duration
	MessageChannelDepth     uint
	ResumeSubs              bool
	HTTPHeaders             http.Header
	WebsocketOptions        *mqtt.WebsocketOptions
}

type Client interface {
	Connect() error
	Disconnect() error
	Publish(topic string, qos byte, retained bool, payload []byte) error
	Subscribe(topic string, qos byte) (*Subscription, error)
	SubscribeMultiple(filters map[string]byte) (*Subscription, error)
}

type ClientImpl struct {
	internal       mqtt.Client
	options        *ClientOptions
	subscriptions  []*Subscription
	reconnectTimer *time.Timer
	mutex          sync.Mutex
}

func NewClient(options *ClientOptions) Client {
	client := &ClientImpl{
		options: options,
	}
	client.internal = mqtt.NewClient(&mqtt.ClientOptions{
		Servers:               options.Servers,
		ClientID:              options.ClientID,
		Username:              options.Username,
		Password:              options.Password,
		CredentialsProvider:   options.CredentialsProvider,
		CleanSession:          options.CleanSession,
		Order:                 options.Order,
		WillEnabled:           options.WillEnabled,
		WillTopic:             options.WillTopic,
		WillPayload:           options.WillPayload,
		WillQos:               options.WillQos,
		WillRetained:          options.WillRetained,
		ProtocolVersion:       options.ProtocolVersion,
		TLSConfig:             options.TLSConfig,
		KeepAlive:             options.KeepAlive,
		PingTimeout:           options.PingTimeout,
		ConnectTimeout:        options.ConnectTimeout,
		MaxReconnectInterval:  options.MaxReconnectInterval,
		AutoReconnect:         false,
		ConnectRetryInterval:  options.ConnectRetryInterval,
		ConnectRetry:          options.ConnectRetry,
		Store:                 options.Store,
		DefaultPublishHandler: options.DefaultPublishHandler,
		OnConnect:             client.onConnect,
		OnConnectionLost:      client.onConnectionLost,
		OnReconnecting:        nil,
		WriteTimeout:          options.WriteTimeout,
		MessageChannelDepth:   options.MessageChannelDepth,
		ResumeSubs:            options.ResumeSubs,
		HTTPHeaders:           options.HTTPHeaders,
		WebsocketOptions:      options.WebsocketOptions,
	})
	return client
}

func awaitToken(token mqtt.Token) error {
	_ = token.Wait()
	return token.Error()
}

func (client *ClientImpl) Connect() error {
	return awaitToken(client.internal.Connect())
}

func (client *ClientImpl) Disconnect() error {
	client.internal.Disconnect(3000)
	return nil
}

func (client *ClientImpl) onConnect(c mqtt.Client) {
	log.Infof("MQTT connected")
	if onConnect := client.options.OnConnect; onConnect != nil {
		onConnect(client)
	}
	for _, sub := range client.subscriptions {
		err := sub.subscribe()
		if err != nil {
			log.Errorf("failed to resubscribe: %s", err)
		}
	}
}

func (client *ClientImpl) onConnectionLost(c mqtt.Client, err error) {
	log.Errorf("MQTT connection lost: %s", err)
	if onConnectionLost := client.options.OnConnectionLost; onConnectionLost != nil {
		onConnectionLost(client, err)
	}
	if client.options.AutoReconnect {
		client.mutex.Lock()
		if client.reconnectTimer == nil {
			client.reconnectTimer = time.AfterFunc(client.options.ConnectRetryInterval, func() {
				client.mutex.Lock()
				client.reconnectTimer = nil
				client.mutex.Unlock()

				if onReconnecting := client.options.OnReconnecting; onReconnecting != nil {
					onReconnecting(client, client.options)
				}
				_ = client.Connect()
			})
		}
		client.mutex.Unlock()
	}
}

func (client *ClientImpl) Publish(topic string, qos byte, retained bool, bytes []byte) error {
	err := awaitToken(client.internal.Publish(topic, qos, retained, bytes))
	if err != nil {
		return err
	}
	log.Debugf("published MQTT message to topic '%s'", topic)
	return nil
}

func (client *ClientImpl) subscribe(sub *Subscription) (*Subscription, error) {
	client.mutex.Lock()
	defer client.mutex.Unlock()
	err := sub.subscribe()
	if err != nil {
		return nil, err
	}
	client.subscriptions = append(client.subscriptions, sub)
	return sub, nil
}

func (client *ClientImpl) Subscribe(topic string, qos byte) (*Subscription, error) {
	return client.subscribe(&Subscription{
		client: client,
		ch:     make(chan Message, messageBufferSize),
		topic:  topic,
		qos:    qos,
	})
}

func (client *ClientImpl) SubscribeMultiple(filters map[string]byte) (*Subscription, error) {
	return client.subscribe(&Subscription{
		client:  client,
		ch:      make(chan Message, messageBufferSize),
		filters: filters,
	})
}

func (client *ClientImpl) removeSubscription(sub *Subscription) {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	subs := make([]*Subscription, 0, len(client.subscriptions))
	for _, existing := range client.subscriptions {
		if sub == existing {
			continue
		}
		subs = append(subs, existing)
	}
	client.subscriptions = subs
}

type Subscription struct {
	client  *ClientImpl
	ch      chan Message
	filters map[string]byte
	topic   string
	qos     byte
}

func (sub *Subscription) Messages() chan Message {
	return sub.ch
}

func (sub *Subscription) subscribe() error {
	if sub.filters != nil {
		err := awaitToken(sub.client.internal.SubscribeMultiple(sub.filters, sub.handleMessage))
		if err != nil {
			return err
		}
		for _, topic := range sub.filters {
			log.Debugf("subscribed to MQTT topic '%s'", string(topic))
		}
	} else {
		err := awaitToken(sub.client.internal.Subscribe(sub.topic, sub.qos, sub.handleMessage))
		if err != nil {
			return err
		}
		log.Debugf("subscribed to MQTT topic '%s'", sub.topic)
	}
	return nil
}

func (sub *Subscription) handleMessage(client mqtt.Client, msg mqtt.Message) {
	select {
	case sub.ch <- &message{msg}:
	default:
	}
}

func (sub *Subscription) Unsubscribe() error {
	topics := make([]string, 0, len(sub.filters))
	for key, _ := range sub.filters {
		topics = append(topics, key)
	}

	var err error
	if sub.filters != nil {
		err = awaitToken(sub.client.internal.Unsubscribe(topics...))
	} else {
		err = awaitToken(sub.client.internal.Unsubscribe(sub.topic))
	}
	if err != nil {
		return err
	}

	sub.client.removeSubscription(sub)

	return nil
}
