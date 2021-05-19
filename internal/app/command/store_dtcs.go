package command

import (
	"log"
	"time"

	"github.com/cesarFuhr/mqttSubscriber/internal/domain/dtc"
)

func NewStoreDTCHandler(r dtc.Repository) StoreDTCHandler {
	return StoreDTCHandler{
		Repo: r,
	}
}

type StoreDTCHandler struct {
	Repo dtc.Repository
}

type StoreDTCCOmmand struct {
	EventID     string
	At          time.Time
	DTC         string
	Description string
}

func (h *StoreDTCHandler) Handle(id string, d StoreDTCCOmmand) error {
	err := h.Repo.InsertDTC(id, dtc.DTC{
		EventID:     d.EventID,
		At:          d.At,
		License:     id,
		DTC:         d.DTC,
		Description: d.Description,
	})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
