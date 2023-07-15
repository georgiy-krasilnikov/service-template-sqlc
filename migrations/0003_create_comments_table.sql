-- +goose Up
CREATE TABLE comments (
    id int GENERATED ALWAYS AS IDENTITY NOT NULL,
    post_id int,
    user_id int,
    text text NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_posts
    FOREIGN KEY(post_id)
    REFERENCES posts(id),
    CONSTRAINT fk_users
    FOREIGN KEY (user_id)
    REFERENCES users(id)
);

-- +goose Down
DROP TABLE comments;