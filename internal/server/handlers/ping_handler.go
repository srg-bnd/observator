// Ping Handler
package handlers

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"
)

type PingHandler struct {
	db *sql.DB
}

// Returns new Handler
func NewPingHandler(db *sql.DB) *PingHandler {
	return &PingHandler{
		db: db,
	}
}

// GET /ping
func (h *PingHandler) Handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := h.db.PingContext(ctx); err != nil {
		log.Println("error:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
