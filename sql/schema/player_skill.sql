CREATE TABLE IF NOT EXISTS player_skill (
    player_id uuid NOT NULL,
    skill_id uuid NOT NULL,
    created_at timestamp NOT NULL DEFAULT NOW(),
    value int NOT NULL,
    PRIMARY KEY (player_id, skill_id),
    FOREIGN KEY (player_id) REFERENCES player(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (skill_id) REFERENCES skill(id) ON DELETE CASCADE ON UPDATE CASCADE
);