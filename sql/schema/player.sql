CREATE TABLE IF NOT EXISTS player (
  id   uuid PRIMARY KEY,
  name text      NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp,
  first_name text NOT NULL,
  last_name  text NOT NULL,
  skills jsonb
);