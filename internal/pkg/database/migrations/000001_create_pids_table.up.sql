CREATE TABLE IF NOT EXISTS pids(
  event_id uuid PRIMARY KEY,
  registered_at TIMESTAMP,
  license VARCHAR(10),
  pid VARCHAR(10),
  reading VARCHAR(10),
  read_at TIMESTAMP
)