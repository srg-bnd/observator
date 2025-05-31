// Models for agent
package models

// Metrics
type Metrics struct {
	ID    string   `json:"id"`              // name
	MType string   `json:"type"`            // type
	Delta *int64   `json:"delta,omitempty"` // counter
	Value *float64 `json:"value,omitempty"` // gauge
}
