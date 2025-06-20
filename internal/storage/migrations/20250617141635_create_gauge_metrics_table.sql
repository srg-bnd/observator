-- +goose Up
CREATE TABLE IF NOT EXISTS gauge_metrics(
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL UNIQUE,
  value DOUBLE PRECISION
);

-- +goose Down
DROP TABLE gauge_metrics;
