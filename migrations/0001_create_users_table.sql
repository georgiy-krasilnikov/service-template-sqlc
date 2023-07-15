-- +goose Up
CREATE TABLE users (
    id int GENERATED ALWAYS AS IDENTITY NOT NULL,
    name varchar(64) NOT NULL,
    last_name varchar(64) NOT NULL,
    email varchar(64) NOT NULL UNIQUE,
    age int NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE users;