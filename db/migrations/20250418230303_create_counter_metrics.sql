CREATE TABLE IF NOT EXISTS counter_metrics(
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL,
  value BIGINT
  /*created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL*/
);
