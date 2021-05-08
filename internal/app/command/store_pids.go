package command

import (
	"log"
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
	EventID     string
	At          time.Time
	PID         string
	Description string
	Value       string
	Unit        string
}

func (h *StorePIDsHandler) Handle(id string, p StorePIDCommand) error {
	err := h.Repo.InsertPID(id, pid.PID{
		EventID:     p.EventID,
		At:          p.At,
		License:     id,
		PID:         p.PID,
		Description: p.Description,
		Value:       p.Value,
		Unit:        p.Unit,
	})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
