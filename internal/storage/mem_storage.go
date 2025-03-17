// Mem Storage
package storage

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	gauge   float64
	counter int64
)

type MemStorage struct {
	gauges        map[string]gauge
	counters      map[string]counter
	fileStorage   *os.File
	storeInterval time.Duration // seconds
	restore       bool
	sync          bool
}

// Create MemStorage instance
func NewMemStorage(fileStoragePath string, storeInterval int, restore bool) *MemStorage {
	var fileStorage *os.File

	if fileStoragePath != "" {
		var err error

		flag := os.O_RDWR | os.O_CREATE
		if !restore {
			flag = os.O_RDWR | os.O_CREATE | os.O_TRUNC
		} else {
		}

		fileStorage, err = os.OpenFile(fileStoragePath, flag, 0666)
		if err != nil {
			log.Fatal(err)
		}

	}

	return &MemStorage{
		gauges:        make(map[string]gauge),
		counters:      make(map[string]counter),
		fileStorage:   fileStorage,
		storeInterval: time.Duration(storeInterval) * time.Second,
		sync:          storeInterval != 0,
		restore:       restore,
	}
}

// Change gauge by key
func (mStore *MemStorage) SetGauge(key string, value float64) error {
	mStore.gauges[key] = gauge(value)
	if err := mStore.Save(); err != nil {
		return err
	}

	return nil
}

// Return gauge by key
func (mStore *MemStorage) GetGauge(key string) (float64, error) {
	value, ok := mStore.gauges[key]
	if !ok {
		return -1, errors.New("unknown")
	}

	return float64(value), nil
}

// Change counter by key
func (mStore *MemStorage) SetCounter(key string, value int64) error {
	mStore.counters[key] += counter(value)
	if err := mStore.Save(); err != nil {
		return err
	}
	return nil
}

// Return gauge by counter
func (mStore *MemStorage) GetCounter(key string) (int64, error) {
	value, ok := mStore.counters[key]
	if !ok {
		return -1, errors.New("unknown")
	}

	return int64(value), nil
}

// Load data
func (mStore *MemStorage) Load() error {
	if !mStore.restore {
		return nil
	}

	scanner := bufio.NewScanner(mStore.fileStorage)

	for scanner.Scan() {
		data := scanner.Text()
		values := strings.Split(data, ",")
		switch values[0] {
		case "counter":
			number, _ := strconv.ParseInt(values[2], 10, 64)
			mStore.counters[values[1]] = counter(number)
		case "gauge":
			number, _ := strconv.ParseFloat(values[2], 64)
			mStore.gauges[values[1]] = gauge(number)
		}
	}

	return nil
}

// Sync data
func (mStore *MemStorage) Sync() {
	if !mStore.sync {
		return
	}

	go func() {
		for {
			time.Sleep(mStore.storeInterval)
			mStore.SaveAll()
		}
	}()
}

// Save all data
func (mStore *MemStorage) Save() error {
	if !mStore.sync {
		return mStore.SaveAll()
	}

	return nil
}

// Save all data
func (mStore *MemStorage) SaveAll() error {
	if mStore.fileStorage == nil {
		return nil
	}

	// HACK
	_ = mStore.fileStorage.Truncate(0)
	_, _ = mStore.fileStorage.Seek(0, 0)

	for name, value := range mStore.counters {
		_, err := mStore.fileStorage.Write([]byte(fmt.Sprintf("counter,%s,%d\n", name, value)))
		if err != nil {
			log.Fatal(err)
		}
	}

	for name, value := range mStore.gauges {
		_, err := mStore.fileStorage.Write([]byte(fmt.Sprintf("gauge,%s,%f\n", name, value)))
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

// Close file Storage
func (mStore *MemStorage) Close() error {
	return mStore.fileStorage.Close()
}
