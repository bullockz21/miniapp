-- +goose Up
-- +goose StatementBegin
CREATE TABLE delivery_types
(
    id SERIAL PRIMARY KEY,
    description VARCHAR(100)
);

INSERT INTO delivery_types (
    description
)
    VALUES (
        'delivery'
    ),
    (
        'self delivery'
    )
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS delivery_types;
-- +goose StatementEnd
