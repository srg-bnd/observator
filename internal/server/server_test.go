package server

import (
	"testing"
)

func TestStart(t *testing.T) {
	t.Logf("TODO")
}

func TestGetHost(t *testing.T) {
	got, _ := getHost()
	want := ":8080"

	if got != want {
		t.Errorf(`Incorrect 'getHost()' method behavior, got "%v", want "%v"`, got, want)
	}
}
