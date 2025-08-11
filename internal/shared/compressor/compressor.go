package compressor

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"fmt"
)

type Compressor struct {
	buffer *bytes.Buffer
}

func NewCompressor() *Compressor {
	return &Compressor{
		buffer: new(bytes.Buffer),
	}
}

// Compress data
func (c *Compressor) Compress(data []byte) ([]byte, error) {
	c.buffer.Reset()
	writer := gzip.NewWriter(c.buffer)

	_, err := writer.Write(data)
	if err != nil {
		return nil, fmt.Errorf("failed to write data to compress temporary buffer: %v", err)
	}

	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to compress data: %v", err)
	}

	return c.buffer.Bytes(), nil
}

// Decompress data
func (c *Compressor) Decompress(compressedData []byte) ([]byte, error) {
	r := flate.NewReader(bytes.NewReader(compressedData))
	defer r.Close()

	var b bytes.Buffer
	_, err := b.ReadFrom(r)
	if err != nil {
		return nil, fmt.Errorf("failed decompress data: %v", err)
	}

	return b.Bytes(), nil
}
