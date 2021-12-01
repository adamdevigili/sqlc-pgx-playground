CREATE TABLE IF NOT EXISTS team (
  id   uuid PRIMARY KEY,
  name text      NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp
);