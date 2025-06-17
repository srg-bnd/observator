-- +goose Up
CREATE TABLE IF NOT EXISTS counter_metrics(
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL UNIQUE,
  value BIGINT
);

-- +goose Down
DROP TABLE counter_metrics;
