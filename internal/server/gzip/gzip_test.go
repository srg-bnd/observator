package gzip

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompressWriter(t *testing.T) {
	w := httptest.NewRecorder()
	cw := NewCompressWriter(w)

	t.Run("TestHeader", func(t *testing.T) {
		assert.NotNil(t, cw.Header())
	})

	t.Run("TestWrite", func(t *testing.T) {
		data := []byte("test data")
		n, err := cw.Write(data)
		assert.NoError(t, err)
		assert.Equal(t, len(data), n)
	})

	t.Run("TestClose", func(t *testing.T) {
		err := cw.Close()
		assert.NoError(t, err)
	})
}

func TestCompressReader(t *testing.T) {
	data := []byte("compressed data")
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	_, _ = zw.Write(data)
	_ = zw.Close()

	cr, err := NewCompressReader(io.NopCloser(&buf))
	assert.NoError(t, err)

	t.Run("TestRead", func(t *testing.T) {
		var result bytes.Buffer
		_, err := io.Copy(&result, cr)
		assert.NoError(t, err)
		assert.Equal(t, string(data), result.String())
	})

	t.Run("TestClose", func(t *testing.T) {
		err := cr.Close()
		assert.NoError(t, err)
	})
}

func TestCompression(t *testing.T) {
	w := httptest.NewRecorder()
	cw := NewCompressWriter(w)

	cw.Header().Set("Content-Type", "application/json")

	data := []byte("some data to compress")
	cw.Write(data)
	cw.Close()

	assert.NotEqual(t, string(w.Body.Bytes()), string(data))
}

func TestNoCompression(t *testing.T) {
	w := httptest.NewRecorder()
	cw := NewCompressWriter(w)

	cw.Header().Set("Content-Type", "image/png")

	data := []byte("some data")
	cw.Write(data)
	cw.Close()

	assert.Equal(t, data, w.Body.Bytes())
	assert.NotContains(t, w.Header().Get("Content-Encoding"), "gzip")
}
