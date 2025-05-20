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
func (dbStore *DBStorage) SetGauge(key string, value float64) error {
	_, err := dbStore.GetGauge(key)
	if err != nil {
		dbStore.db.ExecContext(context.TODO(),
			"INSERT INTO gauge_metrics (name, value) VALUES($1, $2)", key, value)
	} else {
		tx, _ := dbStore.db.BeginTx(context.TODO(), nil)
		tx.ExecContext(context.TODO(),
			"UPDATE gauge_metrics SET value=$1 WHERE name=$2", value, key)
		tx.Commit()
	}

	return nil
}

// Returns gauge by key
func (dbStore *DBStorage) GetGauge(key string) (float64, error) {
	row := dbStore.db.QueryRowContext(context.TODO(),
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
func (dbStore *DBStorage) SetCounter(key string, value int64) error {
	lastValue, err := dbStore.GetCounter(key)
	if err != nil {
		dbStore.db.ExecContext(context.TODO(),
			"INSERT INTO counter_metrics (name, value) VALUES($1, $2)", key, lastValue+value)
	} else {
		dbStore.db.ExecContext(context.TODO(),
			"UPDATE counter_metrics SET value=$1 WHERE name=$2", lastValue+value, key)
	}

	return nil
}

// Returns gauge by counter
func (dbStore *DBStorage) GetCounter(key string) (int64, error) {
	// TODO: Add transaction
	row := dbStore.db.QueryRowContext(context.TODO(),
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
func (dbStore *DBStorage) SetBatchOfMetrics(counterMetrics map[string]int64, gaugeMetrics map[string]float64) error {
	// TODO: needs optimization (saves batch metrics)

	for key, value := range counterMetrics {
		dbStore.SetCounter(key, value)
	}

	for key, value := range gaugeMetrics {
		dbStore.SetGauge(key, value)
	}

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
				id SERIAL PRIMARY KEY, name VARCHAR NOT NULL, value DOUBLE PRECISION
			);`,
		`CREATE TABLE IF NOT EXISTS counter_metrics(
				id SERIAL PRIMARY KEY, name VARCHAR NOT NULL, value BIGINT
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
