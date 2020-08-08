package careregimen

import (
	"database/sql"

	"github.com/wonesy/plantparenthood/graph/model"
)

// Handler manages a member
type Handler struct {
	conn *sql.DB
}

// NewHandler construct for Handler
func NewHandler(conn *sql.DB) *Handler {
	return &Handler{
		conn: conn,
	}
}

// Save insert into database
func (h *Handler) Save() (string, error) {
	return "", nil
}

// GetAll fetch all care regimens from database
func (h *Handler) GetAll() ([]*model.CareRegimen, error) {
	return nil, nil
}