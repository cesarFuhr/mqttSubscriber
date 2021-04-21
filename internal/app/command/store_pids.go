package command

import (
	"time"

	"github.com/cesarFuhr/mqttSubscriber/internal/domain/pid"
)

func NewStorePIDsHandler(r pid.Repository) StorePIDsHandler {
	return StorePIDsHandler{
		Repo: r,
	}
}

type StorePIDsHandler struct {
	Repo pid.Repository
}

type StorePIDCommand struct {
	PID   string
	Value string
	At    time.Time
}

func (h *StorePIDsHandler) Handle(id string, p StorePIDCommand) error {
	err := h.Repo.Store(id, pid.PID{
		At:    p.At,
		PID:   p.PID,
		Value: p.Value,
	})

	return err
}
