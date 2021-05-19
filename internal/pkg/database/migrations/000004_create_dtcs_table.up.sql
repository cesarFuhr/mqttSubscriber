CREATE TABLE IF NOT EXISTS dtcs(
  event_id uuid PRIMARY KEY,
  registered_at TIMESTAMP,
  license VARCHAR(50),
  dtc VARCHAR(50),
  description VARCHAR(100),
  read_at TIMESTAMP
)