-- name: CleanUpDB :exec
TRUNCATE accounts, users RESTART IDENTITY CASCADE;