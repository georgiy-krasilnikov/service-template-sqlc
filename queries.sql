-- name: CreateUser :one
INSERT INTO users (
    name, last_name, email, age
) VALUES (
    $1, $2, $3, $4
)
RETURNING id;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetListOfUsersByIDs :many
SELECT * FROM users
WHERE id = ANY($1::int[]);

-- name: CreatePost :one
INSERT INTO posts (
    user_id, title, text
) VALUES (
    $1, $2, $3
)
RETURNING id;

-- name: GetPostsOfUser :many
SELECT * FROM posts
WHERE user_id = $1;

-- name: CreateCommentForPost :one
INSERT INTO comments (
    post_id, user_id, text
) VALUES (
    $1, $2, $3
)
RETURNING post_id;

-- name: JoinPostsAndUsersTables :one
SELECT posts.id, title, text, name, last_name
FROM posts JOIN users ON posts.user_id = users.id
WHERE posts.id = $1;

-- name: JoinCommentsAndUsersTables :one
SELECT comments.id, text, name, last_name 
FROM comments JOIN users ON user_id = users.id 
WHERE comments.post_id = $1;

-- name: DeleteCommentFromPost :exec
DELETE FROM comments WHERE id = $1;

-- name: GetPostByID :one
SELECT * FROM posts WHERE id = $1;

-- name: GetCommentsOfPostByID :many
SELECT * FROM comments 
WHERE post_id = $1;
