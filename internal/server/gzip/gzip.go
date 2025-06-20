package gzip

import (
	"compress/gzip"
	"io"
	"net/http"
	"slices"
)

/* Writer */

type compressWriter struct {
	w            http.ResponseWriter
	zw           *gzip.Writer
	contentTypes []string
}

func NewCompressWriter(w http.ResponseWriter) *compressWriter {
	return &compressWriter{
		w:            w,
		zw:           gzip.NewWriter(w),
		contentTypes: []string{"application/json", "text/html"},
	}
}

func (c *compressWriter) Header() http.Header {
	return c.w.Header()
}

func (c *compressWriter) Write(p []byte) (int, error) {
	if c.compressibleContent() {
		return c.zw.Write(p)
	} else {
		return c.w.Write(p)
	}
}

func (c *compressWriter) WriteHeader(statusCode int) {
	if c.compressibleContent() {
		c.w.Header().Set("Content-Encoding", "gzip")
	}
	c.w.WriteHeader(statusCode)
}

func (c *compressWriter) Close() error {
	if c.compressibleContent() {
		return c.zw.Close()
	}

	return nil
}

// Helpers

// TODO: memorize the result
func (c *compressWriter) compressibleContent() bool {
	return slices.Contains(c.contentTypes, c.w.Header().Get("content-type"))
}

/* Reader */

type compressReader struct {
	r  io.ReadCloser
	zr *gzip.Reader
}

func NewCompressReader(r io.ReadCloser) (*compressReader, error) {
	zr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	return &compressReader{
		r:  r,
		zr: zr,
	}, nil
}

func (c compressReader) Read(p []byte) (n int, err error) {
	return c.zr.Read(p)
}

func (c *compressReader) Close() error {
	if err := c.r.Close(); err != nil {
		return err
	}
	return c.zr.Close()
}
