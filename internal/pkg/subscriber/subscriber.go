package subscriber

import (
	"time"

	"github.com/cesarFuhr/mqttSubscriber/internal/pkg/logger"
	"github.com/cesarFuhr/mqttSubscriber/internal/ports"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
)

func NewSubscriber(l logger.Logger, c mqtt.Client, p ports.MQTT, qos byte) Subscriber {
	return Subscriber{
		client: c,
		port:   p,
		logger: l,
		qos:    qos,
	}
}

type Subscriber struct {
	client mqtt.Client
	port   ports.MQTT
	logger logger.Logger
	qos    byte
}

func (s *Subscriber) ListenAndHandle() error {
	if !s.client.IsConnected() {
		return mqtt.ErrNotConnected
	}

	if token := s.client.Subscribe("carMon/+/param/+", s.qos, mqttLogger(s.logger, s.port.StorePIDHandler)); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	if token := s.client.Subscribe("carMon/+/status", s.qos, mqttLogger(s.logger, s.port.LogStatusHandler)); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func mqttLogger(l logger.Logger, handle mqtt.MessageHandler) mqtt.MessageHandler {
	return func(c mqtt.Client, m mqtt.Message) {
		l.Info("Incoming... ", zap.String("Topic: ", m.Topic()), zap.Time("At", time.Now()))
		handle(c, m)
	}
}
