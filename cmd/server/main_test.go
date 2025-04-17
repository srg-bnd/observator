package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	app := newApp()
	assert.IsType(t, app, &App{})
}
