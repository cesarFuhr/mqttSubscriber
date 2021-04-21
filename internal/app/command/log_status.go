package command

import (
	"time"

	"github.com/cesarFuhr/mqttSubscriber/internal/pkg/logger"
	"go.uber.org/zap"
)

func NewLogStatusHandler(l logger.Logger) LogStatusHandler {
	return LogStatusHandler{
		logger: l,
	}
}

type LogStatusHandler struct {
	logger logger.Logger
}

type LogStatusCommand struct {
	At     time.Time
	Status bool
}

func (h *LogStatusHandler) Handle(id string, s LogStatusCommand) error {
	h.logger.Info(
		"Received status",
		zap.String("id", id),
		zap.Stringer("at", s.At),
		zap.Bool("status", s.Status),
	)
	return nil
}
