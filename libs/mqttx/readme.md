# mqttx
--
    import "."


## Usage

#### func  MarshalProto

```go
func MarshalProto(payload proto.Message) ([]byte, error)
```

#### func  UnmarshalProto

```go
func UnmarshalProto(message []byte) (proto.Message, error)
```

#### type Client

```go
type Client interface {
	Connect() error
	Publish(topic string, qos byte, retained bool, payload []byte) error
	Subscribe(topic string, qos byte) (*Subscription, error)
	SubscribeMultiple(filters map[string]byte) (*Subscription, error)
}
```


#### func  NewClient

```go
func NewClient(options *ClientOptions) Client
```

#### type ClientImpl

```go
type ClientImpl struct {
}
```


#### func (*ClientImpl) Connect

```go
func (client *ClientImpl) Connect() error
```

#### func (*ClientImpl) Publish

```go
func (client *ClientImpl) Publish(topic string, qos byte, retained bool, bytes []byte) error
```

#### func (*ClientImpl) Subscribe

```go
func (client *ClientImpl) Subscribe(topic string, qos byte) (*Subscription, error)
```

#### func (*ClientImpl) SubscribeMultiple

```go
func (client *ClientImpl) SubscribeMultiple(filters map[string]byte) (*Subscription, error)
```

#### type ClientOptions

```go
type ClientOptions struct {
	Servers             []*url.URL
	ClientID            string
	Username            string
	Password            string
	CredentialsProvider mqtt.CredentialsProvider
	CleanSession        bool
	Order               bool
	WillEnabled         bool
	WillTopic           string
	WillPayload         []byte
	WillQos             byte
	WillRetained        bool
	ProtocolVersion     uint

	TLSConfig             *tls.Config
	KeepAlive             int64
	PingTimeout           time.Duration
	ConnectTimeout        time.Duration
	MaxReconnectInterval  time.Duration
	AutoReconnect         bool
	ConnectRetryInterval  time.Duration
	ConnectRetry          bool
	Store                 mqtt.Store
	DefaultPublishHandler mqtt.MessageHandler
	OnConnect             ConnectHandler
	OnConnectionLost      ConnectionLostHandler
	OnReconnecting        ReconnectHandler
	WriteTimeout          time.Duration
	MessageChannelDepth   uint
	ResumeSubs            bool
	HTTPHeaders           http.Header
	WebsocketOptions      *mqtt.WebsocketOptions
}
```


#### type ConnectHandler

```go
type ConnectHandler func(client Client)
```


#### type ConnectionLostHandler

```go
type ConnectionLostHandler func(client Client, err error)
```


#### type Message

```go
type Message interface {
	mqtt.Message
	UnmarshalProto() (proto.Message, error)
}
```


#### type ReconnectHandler

```go
type ReconnectHandler func(client Client, opts *ClientOptions)
```


#### type Subscription

```go
type Subscription struct {
}
```


#### func (*Subscription) Messages

```go
func (sub *Subscription) Messages() chan Message
```

#### func (*Subscription) Unsubscribe

```go
func (sub *Subscription) Unsubscribe() error
```
