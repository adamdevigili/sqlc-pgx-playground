-- name: GetPlayer :one
SELECT *
  FROM player
  JOIN player_team ON player_team.player_id = player.id
  JOIN team ON player_team.team_id = team.id
  WHERE player.id = $1;

-- name: ListPlayers :many
SELECT * FROM player
ORDER BY name;

-- name: CreatePlayer :one
INSERT INTO player (
  id, first_name, last_name, name, skills, power_scores
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: AddPlayerToTeamIDList :exec
UPDATE player
	SET teams = array_append(teams, sqlc.arg(team_id))
	WHERE id = sqlc.arg(player_id);

-- name: DeletePlayer :exec
DELETE FROM player
WHERE id = $1;

-- name: DeleteAllPlayers :exec
DELETE FROM player;