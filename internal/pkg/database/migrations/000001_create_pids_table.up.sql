CREATE TABLE IF NOT EXISTS pids(
  event_id uuid PRIMARY KEY,
  registered_at TIMESTAMP,
  license VARCHAR(50),
  pid VARCHAR(50),
  description VARCHAR(50),
  reading VARCHAR(50),
  unit VARCHAR(50),
  read_at TIMESTAMP
)