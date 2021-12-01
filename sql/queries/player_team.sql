-- name: ListPlayersOnTeam :many
SELECT * 
    FROM player_team
    WHERE team_id = $1
    ORDER BY name;

-- name: AddPlayerToTeam :exec
INSERT INTO player_team (
  player_id, team_id
) VALUES (
  $1, $2
);

-- name: ChangePlayerTeam :exec
UPDATE player_team 
	SET team_id = $1 
	WHERE player_id = $2;

-- name: DeleteAllPlayerTeams :exec
DELETE FROM player_team;