-- name: GetAccount :one
SELECT * FROM accounts
WHERE id=$1;

-- name: AddAccount :one
INSERT INTO accounts (username, ammount, currency) 
VALUES ($1, $2, $3)
RETURNING *;
