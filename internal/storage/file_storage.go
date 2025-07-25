// File storage for metrics
package storage

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// File storage
type FileStorage struct {
	gauges        map[string]gauge
	counters      map[string]counter
	fileStorage   *os.File
	storeInterval time.Duration // seconds
	restore       bool
	sync          bool
	mtx           sync.RWMutex
}

// Returns a new file storage
func NewFileStorage(fileStoragePath string, storeInterval int, restore bool) *FileStorage {
	var fileStorage *os.File

	if fileStoragePath != "" {
		var err error

		flag := os.O_RDWR | os.O_CREATE
		if !restore {
			flag = os.O_RDWR | os.O_CREATE | os.O_TRUNC
		}

		fileStorage, err = os.OpenFile(fileStoragePath, flag, 0666)
		if err != nil {
			log.Fatal(err)
		}

	}

	return &FileStorage{
		gauges:        make(map[string]gauge),
		counters:      make(map[string]counter),
		fileStorage:   fileStorage,
		storeInterval: time.Duration(storeInterval) * time.Second,
		sync:          storeInterval != 0,
		restore:       restore,
		mtx:           sync.RWMutex{},
	}
}

// Changes gauge by key
func (fStore *FileStorage) SetGauge(ctx context.Context, key string, value float64) error {
	fStore.gauges[key] = gauge(value)
	if err := fStore.Save(); err != nil {
		return err
	}

	return nil
}

// Returns gauge by key
func (fStore *FileStorage) GetGauge(ctx context.Context, key string) (float64, error) {
	value, ok := fStore.gauges[key]
	if !ok {
		return -1, errors.New("unknown")
	}

	return float64(value), nil
}

// Changes counter by key
func (fStore *FileStorage) SetCounter(ctx context.Context, key string, value int64) error {
	fStore.counters[key] += counter(value)
	if err := fStore.Save(); err != nil {
		return err
	}
	return nil
}

// Returns gauge by counter
func (fStore *FileStorage) GetCounter(ctx context.Context, key string) (int64, error) {
	value, ok := fStore.counters[key]
	if !ok {
		return -1, errors.New("unknown")
	}

	return int64(value), nil
}

// Batch update batch of metrics
func (fStore *FileStorage) SetBatchOfMetrics(ctx context.Context, counterMetrics map[string]int64, gaugeMetrics map[string]float64) error {
	// TODO: needs optimization (all parameters are re-saved in each iteration)
	for key, value := range counterMetrics {
		fStore.SetCounter(ctx, key, value)
	}

	// TODO: needs optimization (all parameters are re-saved in each iteration)
	for key, value := range gaugeMetrics {
		fStore.SetGauge(ctx, key, value)
	}

	return nil
}

// Helpers

// Closes file
func (fStore *FileStorage) Close() error {
	return fStore.fileStorage.Close()
}

// Loads data
func (fStore *FileStorage) Load() error {
	if !fStore.restore {
		return nil
	}

	scanner := bufio.NewScanner(fStore.fileStorage)

	for scanner.Scan() {
		data := scanner.Text()
		values := strings.Split(data, ",")
		switch values[0] {
		case "counter":
			number, _ := strconv.ParseInt(values[2], 10, 64)
			fStore.counters[values[1]] = counter(number)
		case "gauge":
			number, _ := strconv.ParseFloat(values[2], 64)
			fStore.gauges[values[1]] = gauge(number)
		}
	}

	return nil
}

// Syncs data
func (fStore *FileStorage) Sync() {
	if !fStore.sync {
		return
	}

	go func() {
		for {
			time.Sleep(fStore.storeInterval)
			fStore.SaveAll()
		}
	}()
}

// Saves data
func (fStore *FileStorage) Save() error {
	if !fStore.sync {
		return fStore.SaveAll()
	}

	return nil
}

// Saves all data
func (fStore *FileStorage) SaveAll() error {
	if fStore.fileStorage == nil {
		return nil
	}

	// HACK
	_ = fStore.fileStorage.Truncate(0)
	_, _ = fStore.fileStorage.Seek(0, 0)

	for name, value := range fStore.counters {
		_, err := fStore.fileStorage.Write([]byte(fmt.Sprintf("counter,%s,%d\n", name, value)))
		if err != nil {
			log.Fatal(err)
		}
	}

	for name, value := range fStore.gauges {

		_, err := fStore.fileStorage.Write([]byte(fmt.Sprintf("gauge,%s,%s\n", name, strconv.FormatFloat(float64(value), 'f', -1, 64))))
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func (fStore *FileStorage) AllCounterMetrics(ctx context.Context) (map[string]int64, error) {
	fStore.mtx.RLock()
	defer fStore.mtx.RUnlock()

	counters := make(map[string]int64)
	for id, delta := range fStore.counters {
		counters[id] = int64(delta)
	}

	return counters, nil
}

func (fStore *FileStorage) AllGaugeMetrics(ctx context.Context) (map[string]float64, error) {
	fStore.mtx.RLock()
	defer fStore.mtx.RUnlock()

	gauges := make(map[string]float64)
	for id, value := range fStore.gauges {
		gauges[id] = float64(value)
	}

	return gauges, nil
}
