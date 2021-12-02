-- name: ListPlayersByTeamID :many
SELECT p.*
  FROM player_team pt
  JOIN player p ON p.id = pt.player_id
  WHERE pt.team_id = $1;

-- name: ListTeamsByPlayerID :many
SELECT t.*
  FROM player_team pt
  JOIN team t ON t.id = pt.team_id
  WHERE pt.player_id = $1;

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