CREATE TABLE IF NOT EXISTS skill (
  id   uuid PRIMARY KEY,
  name text     UNIQUE NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp,
  description text NOT NULL
);