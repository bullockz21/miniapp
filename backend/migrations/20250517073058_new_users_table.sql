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
    deleted_at TIMESTAMP
);

INSERT INTO users (
    user_name, password_hash, email, deleted_at
)
    VALUES (
      'firstTestUser', 'asdasdqwe123asd', 'firstUser@email.com', '1999-01-01 00:00:00'
    ),
    (
      'secondTestUser', 'glkgiouioioff2', 'secondUser@email.com',  '1999-01-01 00:00:00'
    ),
    (
      'thirdTestUser', 'kauiwuwuwydydyd', 'thirdUser@email.com',  '1999-01-01 00:00:00'
    )
  
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd