-- name: CheckIsSeeded :one
SELECT EXISTS (
        SELECT 1
        FROM pokemon
    );
-- name: CreatePokemon :exec
INSERT INTO pokemon (id, name)
VALUES ($1, $2);
-- name: GetTwoRandomPokemon :many
SELECT *
FROM pokemon
ORDER BY RANDOM()
LIMIT 2;
-- name: InsertVote :exec
INSERT INTO vote (voted_for_id, voted_against_id)
VALUES ($1, $2);