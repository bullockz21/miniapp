-- +goose Up
-- +goose StatementBegin

CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    user_name VARCHAR(150) UNIQUE NOT NULL,
    role_id INT DEFAULT 0,
    password_hash VARCHAR(100),
    email VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT '1999-01-01 00:00:00'
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