// Ping Handler
package handlers

import (
	"context"
	"database/sql"
	"fmt"
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, repetitionInterval := range [4]int{0, 1, 3, 5} {
		time.Sleep(time.Duration(repetitionInterval) * time.Second)

		if err := h.db.PingContext(ctx); err != nil {
			// TODO: check type of error
			log.Println(fmt.Errorf("%s: %w", "error", err))
		} else {
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	w.WriteHeader(http.StatusInternalServerError)
}
