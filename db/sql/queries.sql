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
-- name: GetResults :many
SELECT p.id,
    p.name,
    COUNT(
        CASE
            WHEN v.voted_for_id = p.id THEN 1
        END
    ) AS votes_for,
    COUNT(
        CASE
            WHEN v.voted_against_id = p.id THEN 1
        END
    ) AS votes_against
FROM pokemon p
    LEFT JOIN vote v ON p.id = v.voted_for_id
    OR p.id = v.voted_against_id
GROUP BY p.id,
    p.name
ORDER BY p.id;