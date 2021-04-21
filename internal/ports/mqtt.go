package ports

import (
	"bytes"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/cesarFuhr/mqttSubscriber/internal/app"
	"github.com/cesarFuhr/mqttSubscriber/internal/app/command"
	mqtt "github.com/eclipse/paho.mqtt.golang"
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
	b := bytes.NewBuffer(msg.Payload())
	log.Println(string(msg.Payload()))

	var o PID
	if err := decodeJSONBody(b, &o); err != nil {
		var mr *malformedPayload
		if errors.As(err, &mr) {
			msg.Ack()
			log.Println("message malformed: ", err)
			return
		}
		return
	}

	license := strings.Split(strings.TrimPrefix(msg.Topic(), "carMon/"), "/")[0]
	pid := strings.TrimPrefix(msg.Topic(), "carMon/"+license+"/param/")

	err := h.application.Commands.StorePIDs.Handle(license, command.StorePIDCommand{
		EventID: o.EventID,
		At:      o.At,
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
	b := bytes.NewBuffer(msg.Payload())

	var o Status
	if err := decodeJSONBody(b, &o); err != nil {
		var mr *malformedPayload
		if errors.As(err, &mr) {
			msg.Ack()
			log.Println("message malformed: ", err)
			return
		}
		return
	}

	license := strings.TrimRight(strings.TrimPrefix(msg.Topic(), "carMon/"), "/")

	err := h.application.Commands.LogStatus.Handle(license, mqttToCommand(o))
	if err != nil {
		return
	}

	msg.Ack()
}

func mqttToCommand(s Status) command.LogStatusCommand {
	return command.LogStatusCommand{
		At:     s.At,
		Status: s.Status,
	}
}
