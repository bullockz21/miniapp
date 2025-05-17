-- +goose Up
-- +goose StatementBegin
CREATE TABLE order_states
(
    id SERIAL PRIMARY KEY,
    description VARCHAR(100)
);

INSERT INTO order_states (
    description
)
    VALUES (
        'new order'
    ),
    (
        'in progress'
    ),
    (
        'in delivery'
    ),
    (
        'self delivery'
    ),
    (
        'delivered'
    ),
    (
        'done'
    ),
    (
        'cancelled'
    )
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_states;
-- +goose StatementEnd
