package adapters

import (
	"database/sql"
	"log"
	"time"

	"github.com/cesarFuhr/mqttSubscriber/internal/domain/dtc"
	"github.com/cesarFuhr/mqttSubscriber/internal/domain/pid"

	// Is there any other way?
	_ "github.com/lib/pq"
)

func NewPIDRepository(db *sql.DB) PIDRepository {
	return PIDRepository{
		db: db,
	}
}

type PIDRepository struct {
	db *sql.DB
}

var insertPIDStatement = `
	INSERT INTO pids (event_id, registered_at, license, pid, description, unit, reading, read_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

func (r *PIDRepository) InsertPID(id string, p pid.PID) error {
	log.Println(time.Now().UTC(), p.At)
	_, err := r.db.Exec(
		insertPIDStatement,
		p.EventID,
		time.Now().UTC(),
		id,
		p.PID,
		p.Description,
		p.Unit,
		p.Value,
		p.At,
	)
	return err
}

func NewDTCRepository(db *sql.DB) DTCRepository {
	return DTCRepository{
		db: db,
	}
}

type DTCRepository struct {
	db *sql.DB
}

var insertDTCStatement = `
	INSERT INTO dtcs (event_id, registered_at, license, dtc, description, read_at)
		VALUES ($1, $2, $3, $4, $5, $6)`

func (r *DTCRepository) InsertDTC(id string, d dtc.DTC) error {
	_, err := r.db.Exec(
		insertDTCStatement,
		d.EventID,
		time.Now().UTC(),
		id,
		d.DTC,
		d.Description,
		d.At,
	)
	return err
}
