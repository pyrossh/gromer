-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1 LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY id LIMIT $1 OFFSET $2;

-- name: CreateTodo :one
INSERT INTO todos (
  id, text, completed, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: UpdateTodo :one
UPDATE todos 
SET
  completed = $2,
  updated_at = $3
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1;