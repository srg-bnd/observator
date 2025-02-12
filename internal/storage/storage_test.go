package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMemStorage(t *testing.T) {
	got := NewMemStorage()
	want := &MemStorage{}

	assert.IsType(t, got, want)
}

func TestSetGauge(t *testing.T) {
	memStore := NewMemStorage()
	want := float64(1)
	memStore.SetGauge("key", want)
	got := float64(memStore.gauges["key"])

	if got != want {
		t.Errorf(`Incorrect 'SetGauge("key", %#v)' method behavior, got "%v", want "%v"`, want, got, want)
	}
}

func TestGetGauge(t *testing.T) {
	memStore := NewMemStorage()
	want := float64(1)
	memStore.gauges["key"] = gauge(want)
	got := memStore.GetGauge("key")

	if want != float64(got) {
		t.Errorf(`Incorrect 'GetGauge("key")' method behavior, got "%v", want "%v"`, got, want)
	}
}

func TestSetCounter(t *testing.T) {
	memStore := NewMemStorage()
	counter := int64(10)
	want := int64(memStore.counters["key"]) + counter
	memStore.SetCounter("key", counter)
	got := int64(memStore.counters["key"])

	if got != want {
		t.Errorf(`Incorrect 'SetCounter("key", %#v)' method behavior, got "%v", want "%v"`, want, got, want)
	}
}

func TestGetCounter(t *testing.T) {
	memStore := NewMemStorage()
	want := int64(1)
	memStore.counters["key"] = counter(want)
	got := int64(memStore.GetCounter("key"))

	if want != got {
		t.Errorf(`Incorrect 'GetCounter("key")' method behavior, got "%v", want "%v"`, got, want)
	}
}
