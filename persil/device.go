package persil

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
	"henkel-somtec-devscript/devscript"
	iotpersilpb "henkel-somtec-devscript/proto/gen/henkel/device/v3/persil"
	"reflect"
	"time"
)

type Device struct {
	id    string
	props *devscript.DeviceProperties
	conn  *devscript.DeviceConnection

	washCycle *iotpersilpb.WashCycleStartedEvent
	washLoad  int
}

func newDevice(connection *devscript.DeviceConnection) *Device {
	dev := &Device{
		id: connection.ID,
		props: &devscript.DeviceProperties{
			Properties: map[string]*devscript.Property{
				"battery_level": {
					Name:  "persil.ball.battery_level",
					Value: 100,
					Kind:  reflect.Int,
				},
				"battery_status": {
					Name:  "persil.ball.battery_status",
					Value: "ok",
					Kind:  reflect.String,
				},
				"detergent_fill_level": {
					Name:  "persil.ball.detergent_fill_level",
					Value: 100,
					Kind:  reflect.Int,
				},
				"perfume_fill_level": {
					Name:  "persil.ball.perfume_fill_level",
					Value: 100,
					Kind:  reflect.Int,
				},
				"bluetooth_status": {
					Name:  "persil.ball.bluetooth_status",
					Value: "connected",
					Kind:  reflect.String,
				},
				"charging_status": {
					Name:  "persil.ball.charging_status",
					Value: "charging",
					Kind:  reflect.String,
				},
				"wash_cycle_status": {
					Name:  "persil.ball.wash_cycle_status",
					Value: "idle",
					Kind:  reflect.String,
				},
				"refill_status": {
					Name:  "persil.ball.refill_status",
					Value: "inserted",
					Kind:  reflect.String,
				},
				"soiling_level": {
					Name:      "persil.ball.soiling_level",
					Writeable: true,
					Value:     "medium",
					Kind:      reflect.String,
				},
				"perfume_level": {
					Name:      "persil.ball.perfume_level",
					Writeable: true,
					Value:     "normal",
					Kind:      reflect.String,
				},
				"wash_cycle_id": {
					Name:  "persil.ball.wash_cycle_id",
					Value: uuid.NewString(),
					Kind:  reflect.String,
				},
				"refill_cycle_id": {
					Name:  "persil.ball.refill_cycle_id",
					Value: uuid.NewString(),
					Kind:  reflect.String,
				},
			},
		},
		conn: connection,
	}

	dev.props.OnPropertyChanged = dev.onPropertyChanged

	return dev
}

func (dev *Device) onPropertyChanged(prop *devscript.Property) {
	if err := dev.conn.PublishPropertyUpdate(prop); err != nil {
		panic(err)
	}
}

func (dev *Device) Connect() {
	if err := dev.conn.Connect(); err != nil {
		panic(err)
	}
	if err := dev.conn.PublishConnectedEvent(dev.props); err != nil {
		panic(err)
	}
	logrus.Infof("device connected")
}

func (dev *Device) Disconnect() {
	if err := dev.conn.Disconnect(); err != nil {
		panic(err)
	}
	logrus.Infof("device disconnected")
}

func (dev *Device) Set(name string, value any) {
	dev.props.Set(name, value)
}

func (dev *Device) Get(name string) any {
	return dev.props.Get(name)
}

func (dev *Device) SetSoilingLevel(level string) {
	dev.props.Set("soiling_level", level)
}

func (dev *Device) SetPerfumeLevel(level string) {
	dev.props.Set("perfume_level", level)
}

func (dev *Device) SetBatteryLevel(level int) {
	dev.props.Set("battery_level", level)
	if level > 30 {
		dev.props.Set("battery_status", "ok")
	} else if level > 10 {
		dev.props.Set("battery_status", "low")
	} else {
		dev.props.Set("battery_status", "critical")
	}
}

func (dev *Device) SetBatteryStatus(status string) {
	dev.props.Set("battery_status", status)
	if status == "ok" {
		dev.props.Set("battery_level", 100)
	} else if status == "low" {
		dev.props.Set("battery_level", 20)
	} else {
		dev.props.Set("battery_level", 0)
	}
}

func (dev *Device) SetWashLoad(load int) {
	dev.washLoad = load
	logrus.Infof("wash load set to %d", dev.washLoad)
}

func (dev *Device) SetDetergentFillLevel(level int) {
	dev.props.Set("detergent_fill_level", level)
}

func (dev *Device) SetPerfumeFillLevel(level int) {
	dev.props.Set("perfume_fill_level", level)
}

func (dev *Device) DisconnectBluetooth() {
	dev.props.Set("bluetooth_status", "disconnected")
}

func (dev *Device) ReconnectBluetooth() {
	dev.props.Set("bluetooth_status", "connected")
}

func (dev *Device) PutOnStation() {
	dev.props.Set("charging_status", "charging")
}

func (dev *Device) RemoveFromStation() {
	dev.props.Set("charging_status", "discharging")
}

func (dev *Device) SetWashCycleStatus(status string) {
	dev.props.Set("wash_cycle_status", status)
}

func (dev *Device) RemoveRefill() {
	if dev.props.Get("refill_status") != "inserted" {
		panic("refill needs to be inserted first")
	}

	dev.props.Set("refill_status", "removed")

	logrus.Infof("refill removed")
}

func (dev *Device) InsertRefill() {
	if dev.props.Get("refill_status") != "removed" {
		panic("refill needs to be removed first")
	}

	dev.props.Set("refill_cycle_id", uuid.NewString())
	dev.props.Set("detergent_fill_level", 100)
	dev.props.Set("perfume_fill_level", 100)
	dev.props.Set("refill_status", "inserted")

	if err := dev.conn.Publish(&iotpersilpb.RefillChangedEvent{
		EventId:  uuid.NewString(),
		DeviceId: dev.id,
		Time:     timestamppb.Now(),
		RefillId: dev.props.Get("refill_cycle_id").(string),
	}); err != nil {
		panic(err)
	}

	logrus.Infof("refill inserted")
}

func (dev *Device) StartWashCycle() {
	dev.props.Set("wash_cycle_id", uuid.NewString())
	dev.washCycle = &iotpersilpb.WashCycleStartedEvent{
		EventId:            uuid.NewString(),
		DeviceId:           dev.id,
		Time:               timestamppb.Now(),
		WashCycleId:        dev.props.Get("wash_cycle_id").(string),
		SoilingLevel:       dev.props.Get("soiling_level").(string),
		PerfumeLevel:       dev.props.Get("perfume_level").(string),
		DetergentFillLevel: int32(dev.props.Get("detergent_fill_level").(int)),
		BatteryLevel:       int32(dev.props.Get("battery_level").(int)),
		RefillId:           dev.props.Get("refill_cycle_id").(string),
	}

	if err := dev.conn.Publish(dev.washCycle); err != nil {
		panic(err)
	}

	logrus.Infof("wash cycle started")
}

func (dev *Device) StopWashCycle(status string) {
	if dev.washCycle == nil {
		panic("need to start wash cycle first")
	}

	now := time.Now()

	var pbStatus iotpersilpb.WashCycleCompletedEvent_Status
	switch status {
	case "done":
		pbStatus = iotpersilpb.WashCycleCompletedEvent_STATUS_DONE
	case "aborted":
		pbStatus = iotpersilpb.WashCycleCompletedEvent_STATUS_ABORTED
	default:
		panic("unknown status: " + status)
	}

	if err := dev.conn.Publish(&iotpersilpb.WashCycleCompletedEvent{
		EventId:                   uuid.NewString(),
		DeviceId:                  dev.id,
		Time:                      timestamppb.New(now),
		Status:                    pbStatus,
		WashCycleId:               dev.props.Get("wash_cycle_id").(string),
		Duration:                  int32(now.Sub(dev.washCycle.Time.AsTime()).Milliseconds()),
		SoilingLevel:              dev.props.Get("soiling_level").(string),
		PerfumeLevel:              dev.props.Get("perfume_level").(string),
		WashLoad:                  int32(dev.washLoad),
		InitialDetergentFillLevel: dev.washCycle.DetergentFillLevel,
		FinalDetergentFillLevel:   int32(dev.props.Get("detergent_fill_level").(int)),
		InitialBatteryLevel:       dev.washCycle.BatteryLevel,
		FinalBatteryLevel:         int32(dev.props.Get("battery_level").(int)),
		RefillId:                  dev.props.Get("refill_cycle_id").(string),
	}); err != nil {
		panic(err)
	}

	dev.washCycle = nil

	logrus.Infof("wash cycle stopped")
}
