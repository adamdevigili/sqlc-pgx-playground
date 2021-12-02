CREATE TABLE IF NOT EXISTS team (
  id   uuid PRIMARY KEY,
  name text     UNIQUE NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp,
  sport_name text NOT NULL,
  power_score float4 NOT NULL,
  wins smallint NOT NULL DEFAULT 0,
  losses smallint NOT NULL DEFAULT 0,
  CONSTRAINT fk_sport FOREIGN KEY(sport_name) REFERENCES sport(name)
);