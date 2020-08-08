package careregimen

import (
	"database/sql"

	"github.com/wonesy/plantparenthood/internal/pkg/db"
	"github.com/wonesy/plantparenthood/util"

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

// Create insert into database
func (h *Handler) Create(c *model.NewCareRegimen) (*model.CareRegimen, error) {
	insertSQL := `
	INSERT INTO care_regimen (id, water_ml, water_hr)
	VALUES ($1, $2, $3) RETURNING id`

	id := util.GenerateID()

	returnedID := ""
	if err := h.conn.QueryRow(insertSQL, id, c.Waterml, c.Waterhr).Scan(&returnedID); err != nil {
		return nil, err
	}

	return h.GetByID(id)
}

// GetByID fetch care regiment by ID
func (h *Handler) GetByID(id string) (*model.CareRegimen, error) {
	selectSQL := `
	SELECT id, water_ml, water_hr
	FROM care_regimen
	WHERE id=$1`

	stmt, err := h.conn.Prepare(selectSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	c := &model.CareRegimen{}
	if err := stmt.QueryRow(id).Scan(&c.ID, &c.Waterml, &c.Waterhr); err != nil {
		return nil, &db.NoSuchEntity{Type: "careregimen"}
	}

	return c, nil
}

// GetAll fetch all care regimens from database
func (h *Handler) GetAll() ([]*model.CareRegimen, error) {
	selectSQL := `SELECT * FROM care_regimen`

	stmt, err := h.conn.Prepare(selectSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	careRegimens := []*model.CareRegimen{}
	for rows.Next() {
		c := &model.CareRegimen{}
		if err := rows.Scan(&c.ID, &c.Waterml, &c.Waterhr); err != nil {
			return nil, err
		}
		careRegimens = append(careRegimens, c)
	}

	return careRegimens, nil
}
