-- +goose Up
-- +goose StatementBegin
CREATE TABLE position_types
(
    id SERIAL PRIMARY KEY,
    description VARCHAR(100)
);

INSERT INTO position_types (
    description
)
    VALUES (
        'pizza'
    ),
    (
        'rolls'
    ),
    (
        'drinks'
    ),
    (
        'misc'
    )
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS position_types;
-- +goose StatementEnd
