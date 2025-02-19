-- name: ListTasks :many 
SELECT
    *
FROM
    tasks;

-- name: GetTaskByID :one
SELECT
    *
FROM
    tasks
WHERE
    id = $1;

-- name: CreateTask :one 
INSERT INTO tasks (
    title,
    description
) VALUES (
    $1,
    $2
) 
RETURNING id; 

-- name: UpdateTask :exec
UPDATE tasks
SET
    title = $2,
    description = $3,
    status = $4, 
    updated_at = $5
WHERE 
    id = $1;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE 
    id = $1;