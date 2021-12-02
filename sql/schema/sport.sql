CREATE TABLE IF NOT EXISTS sport (
  id   uuid PRIMARY KEY,
  name text  UNIQUE NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp,
  description text NOT NULL,
  skill_weights jsonb,
  max_active_players_per_team smallint,
  max_players_per_team smallint
);