-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users LIMIT $1 OFFSET $2;

-- name: CreateUser :one
INSERT INTO users (
  email, group_id, full_name, global_role, salt, hash
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET full_name=$2, updated_at=CURRENT_TIMESTAMP
WHERE id=$1
RETURNING *;

-- name: UpdateUserPassword :exec
UPDATE users
SET salt=$2, hash=$3, updated_at=CURRENT_TIMESTAMP
WHERE id=$1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: GetGroup :one
SELECT * FROM groups
WHERE id = $1 LIMIT 1;

-- name: CreateGroup :one
INSERT INTO groups (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: CreateUserGroup :one
INSERT INTO groups (
  name,
  user_group
) VALUES (
  $1,
  TRUE
)
RETURNING *;

-- name: UpdateGroup :one
UPDATE groups
SET name = $2, updated_at=CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: GetGroups :many
SELECT * FROM groups
LIMIT $1
OFFSET $2;

-- name: DeleteGroup :exec
DELETE FROM groups
WHERE id = $1;

-- name: AddGroupMember :exec
INSERT INTO group_members (
  user_id, group_id, role
) VALUES (
  $1, $2, $3
);

-- name: DeleteGroupMember :exec
DELETE FROM group_members
WHERE user_id = $1 AND group_id = $2;

-- name: UpdateGroupMember :exec
UPDATE group_members
SET role = $3, updated_at=CURRENT_TIMESTAMP
WHERE user_id = $1 AND group_id = $2;

-- name: GetGroupMembers :many
SELECT user_id, full_name, role, group_members.created_at AS added_at, group_members.updated_at
FROM "group_members" JOIN "users" ON user_id=users.id
WHERE group_members.group_id = $3
LIMIT $1
OFFSET $2;

-- name: GetUserGroups :many
SELECT groups.*
FROM "group_members" JOIN "groups" ON group_id=groups.id
WHERE user_id = $3
LIMIT $1
OFFSET $2;

-- name: CountGroupMembers :one
SELECT COUNT(*)
FROM "group_members"
WHERE group_id = $1;

-- name: CreateSession :one
INSERT INTO sessions (
  user_id, created_at, expired_at
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE id = $1 AND user_id = $2;

-- name: GetSession :one
SELECT *
FROM sessions
WHERE "id" = $1
LIMIT 1;

-- name: GetSessions :many
SELECT *
FROM sessions
WHERE user_id = $3
LIMIT $1
OFFSET $2;

-- name: GetUserGroupRoles :many
SELECT group_id, role
FROM "group_members"
WHERE user_id = $1;
