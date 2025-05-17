-- +goose Up
-- +goose StatementBegin
CREATE TABLE menu
(
    id SERIAL PRIMARY KEY,
    type_id INT NOT NULL,
    name VARCHAR(50) UNIQUE NOT NULL,
    price NUMERIC,
    description VARCHAR(1000)
);

INSERT INTO menu (
    type_id, name, price, description
)
    VALUES (
        1, 'pizza1', 1250, 'ohuennaya pizza #1'
    ),
    (
        1, 'pizza2', 1000, 'menee ohuennaya pizza#2'
    ),
    (
        2, 'krab_rolls', 500, 'californication'
    ),
    (
        2, 'krevetki rolls', 450, 'unagi maki'
    ),
    (
        3, 'koka_kola', 100, 'gazirovka'
    ),
    (
        3, 'sprite', 100, 'prozrachnaya gazirovka'
    )
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS menu;
-- +goose StatementEnd
