-- name: FindAllTasks :many
SELECT * FROM tasks ORDER BY created_at;

-- name: CreateTask :one
INSERT INTO tasks (title)
VALUES ($1)
RETURNING *;

-- name: UpdateTask :one
UPDATE tasks
SET title = $2, completed = $3
WHERE id = $1
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1;
