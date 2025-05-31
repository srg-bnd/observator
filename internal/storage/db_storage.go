// Database storage for metrics
package storage

import (
	"context"
	"database/sql"

	"errors"
)

// Database storage
type DBStorage struct {
	db *sql.DB
}

// Returns a new database storage
func NewDBStorage(db *sql.DB) *DBStorage {
	return &DBStorage{
		db: db,
	}
}

// Changes gauge by key
func (dbStore *DBStorage) SetGauge(ctx context.Context, key string, value float64) error {
	tx, err := dbStore.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	tx.ExecContext(ctx,
		`INSERT INTO gauge_metrics (name, value)
				VALUES ($1, $2)
				ON CONFLICT (name)
				DO UPDATE SET name = $1, value = $2;`, key, value)
	tx.Commit()

	return nil
}

// Returns gauge by key
func (dbStore *DBStorage) GetGauge(ctx context.Context, key string) (float64, error) {
	row := dbStore.db.QueryRowContext(ctx,
		"SELECT value FROM gauge_metrics WHERE name = $1", key)

	var value sql.NullFloat64
	err := row.Scan(&value)
	if err != nil {
		return -1, err
	}

	if value.Valid {
		return value.Float64, nil
	} else {
		return -1, errors.New("unknown")
	}
}

// Changes counter by key
func (dbStore *DBStorage) SetCounter(ctx context.Context, key string, value int64) error {
	tx, err := dbStore.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx,
		`INSERT INTO counter_metrics (name, value)
				VALUES ($1, $2)
				ON CONFLICT (name)
				DO UPDATE SET name = $1, value = counter_metrics.value + $2;`, key, value)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

// Returns gauge by counter
func (dbStore *DBStorage) GetCounter(ctx context.Context, key string) (int64, error) {
	// TODO: Add transaction
	row := dbStore.db.QueryRowContext(ctx,
		"SELECT value FROM counter_metrics WHERE name = $1", key)

	var value sql.NullInt64
	err := row.Scan(&value)
	if err != nil {
		return -1, err
	}

	if value.Valid {
		return value.Int64, nil
	} else {
		return -1, errors.New("unknown")
	}
}

// Batch update batch of metrics
func (dbStore *DBStorage) SetBatchOfMetrics(ctx context.Context, counterMetrics map[string]int64, gaugeMetrics map[string]float64) error {
	tx, err := dbStore.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for key, value := range counterMetrics {
		_, err := tx.ExecContext(ctx,
			`INSERT INTO counter_metrics (name, value)
				VALUES ($1, $2)
				ON CONFLICT (name)
				DO UPDATE SET name = $1, value = counter_metrics.value + $2;`, key, value)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	for key, value := range gaugeMetrics {
		_, err := tx.ExecContext(ctx,
			`INSERT INTO gauge_metrics (name, value)
				VALUES ($1, $2)
				ON CONFLICT (name)
				DO UPDATE SET name = $1, value = $2;`, key, value)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}

// Helpers

// Closes connection to DB
func (dbStore *DBStorage) Close() {
	dbStore.db.Close()
}

// Executes migrations
func (dbStore *DBStorage) ExecMigrations() error {
	migrations := [2]string{
		`CREATE TABLE IF NOT EXISTS gauge_metrics(
				id SERIAL PRIMARY KEY, name VARCHAR NOT NULL UNIQUE, value DOUBLE PRECISION
			);`,
		`CREATE TABLE IF NOT EXISTS counter_metrics(
				id SERIAL PRIMARY KEY, name VARCHAR NOT NULL UNIQUE, value BIGINT
			);`,
	}

	for _, migration := range migrations {
		_, err := dbStore.db.ExecContext(context.Background(), string(migration))
		if err != nil {
			return err
		}
	}

	return nil
}
