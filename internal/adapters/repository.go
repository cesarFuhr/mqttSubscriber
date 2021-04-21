package adapters

import (
	"database/sql"
	"time"

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
	Reading      string
}

var insertPIDStatement = `
	INSERT INTO pids (event_id, registered_at, license, pid, reading, read_at)
		VALUES ($1, $2, $3, $4, $5, $6)`

// InsertKey Inserts a key into the repository
func (r *PIDRepository) InsertPID(id string, p pid.PID) error {
	_, err := r.db.Exec(
		insertPIDStatement,
		p.EventID,
		time.Now().UTC(),
		id,
		p.PID,
		p.Value,
		p.At,
	)
	return err
}
