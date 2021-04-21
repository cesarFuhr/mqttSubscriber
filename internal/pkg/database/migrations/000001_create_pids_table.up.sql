CREATE TABLE IF NOT EXISTS pids(
  event_id uuid PRIMARY KEY,
  registration TIMESTAMP,
  license VARCHAR(10),
  pid VARCHAR(10),
  reading VARCHAR(10)
)