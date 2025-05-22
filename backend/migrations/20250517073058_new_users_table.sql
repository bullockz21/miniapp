-- +goose Up
-- +goose StatementBegin

CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    user_name VARCHAR(150) UNIQUE, NOT NULL
    role_id INT,
    password_hash VARCHAR(100),
    email VARCHAR(100),
    created_at DATE,
    deleted_at DATE
);

INSERT INTO users (
    user_name, password_hash, email
)
    VALUES (
      'firstTestUser', 'asdasdqwe123asd', 'firstUser@email.com'
    ),
    (
      'secondTestUser', 'glkgiouioioff2', 'secondUser@email.com'
    ),
    (
      'thirdTestUser', 'kauiwuwuwydydyd', 'thirdUser@email.com'
    )
  
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd