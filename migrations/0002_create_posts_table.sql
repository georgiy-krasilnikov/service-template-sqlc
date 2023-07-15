-- +goose Up
CREATE TABLE posts (
    id int GENERATED ALWAYS AS IDENTITY NOT NULL,
    user_id int,
    title varchar(64) NOT NULL,
    text text NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_users
    FOREIGN KEY (user_id)
    REFERENCES users(id)
);

-- +goose Down
DROP TABLE posts;