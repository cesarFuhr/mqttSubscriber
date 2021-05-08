package adapters

import (
	"database/sql"
	"log"
	"time"

	"github.com/cesarFuhr/mqttSubscriber/internal/domain/fault"
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

type PIDRepositoryModel struct {
	EventID      string
	Registration time.Time
	License      string
	PID          string
	Description  string
	Reading      string
	Unit         string
}

var insertPIDStatement = `
	INSERT INTO pids (event_id, registered_at, license, pid, description, unit, reading, read_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

// InsertKey Inserts a key into the repository
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

func NewFaultRepository(db *sql.DB) PIDRepository {
	return PIDRepository{
		db: db,
	}
}

type FaultRepository struct {
	db *sql.DB
}

type FaultRepositoryModel struct {
	EventID      string
	Registration time.Time
	License      string
	PID          string
	Reading      string
}

var insertFaultStatement = `
	INSERT INTO faults (event_id, registered_at, license, fault, read_at)
		VALUES ($1, $2, $3, $4, $5, $6)`

// InsertKey Inserts a key into the repository
func (r *FaultRepository) InsertFault(id string, f fault.Fault) error {
	_, err := r.db.Exec(
		insertFaultStatement,
		f.EventID,
		time.Now().UTC(),
		id,
		f.Fault,
		f.At,
	)
	return err
}
