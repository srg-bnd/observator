package compressor

import (
	"fmt"
	"testing"
)

func TestNewCompressor(t *testing.T) {
	c := NewCompressor()
	if c == nil {
		t.Error("Expected non-nil Compressor instance")
	}
}

func TestCompress(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:  "Empty input",
			input: []byte{},
		},
		{
			name:  "Simple text",
			input: []byte("Hello, World!"),
		},
		{
			name:  "Repeating data",
			input: []byte("aaaaaa"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCompressor()
			compressed, err := c.Compress(tt.input)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if len(compressed) == 0 {
				t.Errorf("Expected non-empty compressed data")
			}

			fmt.Println(string(compressed))
			if string(compressed) != tt.expected {
				t.Logf("TODO")
			}
		})
	}
}

func TestDecompress(t *testing.T) {
	t.Logf("TODO")
}
