-- +goose Up
-- +goose StatementBegin
CREATE TABLE customers
(
    id SERIAL PRIMARY KEY,
    tg_name VARCHAR(100) UNIQUE NOT NULL,
    tg_id INT UNIQUE NOT NULL,
    phone_number VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO customers (
    tg_name, tg_id, phone_number
)
    VALUES 
    (
        'testUser1', 123456, '+7 (924) 123-45-67'
    ),
    (
        'testUser2', 908767, '+7 (123) 654-77-66'
    ),
    (
        'testUser3', 912983, '+7 (098) 123-66-98'
    )
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS customers;
-- +goose StatementEnd
