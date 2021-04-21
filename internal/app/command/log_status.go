package command

import (
	"github.com/cesarFuhr/mqttSubscriber/internal/domain/status"
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

func (h *LogStatusHandler) Handle(id string, s status.Status) {
	h.logger.Info(
		"Received status",
		zap.String("id", id),
		zap.Stringer("at", s.At),
		zap.Bool("status", s.Status),
	)
}
