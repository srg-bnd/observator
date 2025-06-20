// Database storage for metrics
package storage

import (
	"context"
	"database/sql"
	"embed"

	"errors"

	"github.com/pressly/goose/v3"
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

//go:embed migrations/*.sql
var embedMigrations embed.FS

var (
	// Errors
	ErrUnknown = errors.New("unknown")
)

const (
	// SQL queries
	allCounterSQL = `SELECT name, value FROM counter_metrics`
	setCounterSQL = `INSERT INTO counter_metrics (name, value)
				VALUES ($1, $2)
				ON CONFLICT (name)
				DO UPDATE SET name = $1, value = counter_metrics.value + $2;`
	getCounterSQL = `SELECT value FROM counter_metrics WHERE name = $1`
	allGaugesSQL  = `SELECT name, value FROM gauge_metrics`
	setGaugeSQL   = `INSERT INTO gauge_metrics (name, value)
				VALUES ($1, $2)
				ON CONFLICT (name)
				DO UPDATE SET name = $1, value = $2;`
	getGaugeSQL                 = `SELECT value FROM gauge_metrics WHERE name = $1`
	setBatchOfCounterMetricsSQL = `INSERT INTO counter_metrics (name, value)
				VALUES ($1, $2)
				ON CONFLICT (name)
				DO UPDATE SET name = $1, value = counter_metrics.value + $2;`
	setBatchOfGaugeMetricsSQL = `INSERT INTO gauge_metrics (name, value)
				VALUES ($1, $2)
				ON CONFLICT (name)
				DO UPDATE SET name = $1, value = $2;`
)

// Changes gauge by key
func (dbStore *DBStorage) SetGauge(ctx context.Context, key string, value float64) error {
	tx, err := dbStore.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	tx.ExecContext(ctx, setGaugeSQL, key, value)
	tx.Commit()

	return nil
}

// Returns gauge by key
func (dbStore *DBStorage) GetGauge(ctx context.Context, key string) (float64, error) {
	row := dbStore.db.QueryRowContext(ctx, getGaugeSQL, key)

	var value sql.NullFloat64
	err := row.Scan(&value)
	if err != nil {
		return -1, err
	}

	if value.Valid {
		return value.Float64, nil
	} else {
		return -1, ErrUnknown
	}
}

// Changes counter by key
func (dbStore *DBStorage) SetCounter(ctx context.Context, key string, value int64) error {
	tx, err := dbStore.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, setCounterSQL, key, value)

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
	row := dbStore.db.QueryRowContext(ctx, getCounterSQL, key)

	var value sql.NullInt64
	err := row.Scan(&value)
	if err != nil {
		return -1, err
	}

	if value.Valid {
		return value.Int64, nil
	} else {
		return -1, ErrUnknown
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
		_, err := tx.ExecContext(ctx, setBatchOfCounterMetricsSQL, key, value)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	for key, value := range gaugeMetrics {
		_, err := tx.ExecContext(ctx, setBatchOfGaugeMetricsSQL, key, value)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}

func (dbStore *DBStorage) AllCounterMetrics(ctx context.Context) (map[string]int64, error) {
	rows, err := dbStore.db.QueryContext(ctx, allCounterSQL)
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	defer rows.Close()

	result := make(map[string]int64, 0)
	for rows.Next() {
		var name sql.NullString
		var value sql.NullInt64

		err := rows.Scan(&name, &value)
		if err != nil {
			return nil, err
		}

		if value.Valid && name.Valid {
			result[name.String] = value.Int64
		} else {
			return nil, ErrUnknown
		}
	}

	return result, nil
}

func (dbStore *DBStorage) AllGaugeMetrics(ctx context.Context) (map[string]float64, error) {
	rows, err := dbStore.db.QueryContext(ctx, allGaugesSQL)
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	defer rows.Close()

	result := make(map[string]float64, 0)
	for rows.Next() {
		var name sql.NullString
		var value sql.NullFloat64

		err := rows.Scan(&name, &value)
		if err != nil {
			return nil, err
		}

		if value.Valid && name.Valid {
			result[name.String] = value.Float64
		} else {
			return nil, ErrUnknown
		}
	}

	return result, nil
}

// Helpers

// Closes connection to DB
func (dbStore *DBStorage) Close() {
	dbStore.db.Close()
}

// Executes migrations
func (dbStore *DBStorage) ExecMigrations() error {
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(dbStore.db, "migrations"); err != nil {
		panic(err)
	}

	return nil
}
