CREATE TABLE IF NOT EXISTS player_team (
    player_id uuid NOT NULL,
    team_id uuid NOT NULL,
    joined_at timestamp NOT NULL DEFAULT NOW(),
    PRIMARY KEY (player_id, team_id),
    FOREIGN KEY (player_id) REFERENCES player(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (team_id) REFERENCES team(id) ON DELETE CASCADE ON UPDATE CASCADE
);