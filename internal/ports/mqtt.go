package ports

import (
	"log"
	"strings"
	"time"

	"github.com/cesarFuhr/mqttSubscriber/internal/app"
	"github.com/cesarFuhr/mqttSubscriber/internal/app/command"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"google.golang.org/protobuf/proto"
)

type MQTT struct {
	application app.Application
}

func NewMQTTPort(a app.Application) MQTT {
	return MQTT{
		application: a,
	}
}

type PID struct {
	EventID string
	At      time.Time
	PID     string
	Value   string
}

func (h *MQTT) StorePIDHandler(cli mqtt.Client, msg mqtt.Message) {
	o := &PIDNotification{}

	if err := proto.Unmarshal(msg.Payload(), o); err != nil {
		log.Println("Could not unmarshal")
		return
	}

	license := strings.Split(strings.TrimPrefix(msg.Topic(), "carMon/"), "/")[0]
	pid := strings.TrimPrefix(msg.Topic(), "carMon/"+license+"/param/")

	err := h.application.Commands.StorePIDs.Handle(license, command.StorePIDCommand{
		EventID: o.EventID,
		At:      o.At.AsTime(),
		PID:     pid,
		Value:   o.Value,
	})
	if err != nil {
		return
	}

	msg.Ack()
}

type Status struct {
	At     time.Time
	Status bool
}

func (h *MQTT) LogStatusHandler(cli mqtt.Client, msg mqtt.Message) {
	o := &StatusNotification{}

	if err := proto.Unmarshal(msg.Payload(), o); err != nil {
		log.Println("Could not unmarshal")
		return
	}

	license := strings.TrimRight(strings.TrimPrefix(msg.Topic(), "carMon/"), "/")

	err := h.application.Commands.LogStatus.Handle(license, command.LogStatusCommand{
		At:     o.At.AsTime(),
		Status: o.Status,
	})
	if err != nil {
		return
	}

	msg.Ack()
}
