-- +goose Up
-- +goose StatementBegin

CREATE TABLE bucket
(
    id SERIAL PRIMARY KEY,
    customer_id INT,
    menu_id INT,
    date DATE,
    time TIME
);

INSERT INTO bucket (
    customer_id, menu_id, time, date
)
    VALUES (
      1, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
    ),
    (
      1, 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
    ),
    (
      2, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
    ),
    (
      3, 3, CURRENT_TIMESTAMP,CURRENT_TIMESTAMP
    )
  
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS bucket;
-- +goose StatementEnd