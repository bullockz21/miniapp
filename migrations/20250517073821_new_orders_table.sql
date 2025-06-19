-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders
(
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    delivery_type_id INT NOT NULL,
    menu_id INT NOT NULL,
    order_number INT NOT NULL,
    order_state_id INT DEFAULT 1,
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO orders (
    customer_id, delivery_type_id, menu_id, order_number
)
    VALUES 
    (
        1, 2, 3, 1
    ),
    (
        1, 2, 4, 1
    ),
    (
        1, 2, 5, 1
    ),
    (
        2, 1, 1, 2
    ),
    (
        2, 1, 2, 2
    ),
    (
        3, 2, 1, 1
    ),
    (
        3, 2, 2, 1
    )
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
