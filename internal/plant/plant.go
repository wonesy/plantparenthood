package plant

import (
	"database/sql"

	"github.com/wonesy/plantparenthood/graph/model"
	"github.com/wonesy/plantparenthood/internal/pkg/db"
	"github.com/wonesy/plantparenthood/util"
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
func (h *Handler) Create(p *model.NewPlant) (*model.Plant, error) {
	insertSQL := `
	INSERT INTO plant (id, common_name, botanical_name, water_preference, sun_preference, soil_preference)
	VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`

	id := util.GenerateID()

	returnedID := ""
	if err := h.conn.QueryRow(
		insertSQL,
		id,
		p.CommonName,
		p.BotanicalName,
		p.WaterPreference,
		p.SunPreference,
		p.SoilPreference,
	).Scan(&returnedID); err != nil {
		return nil, err
	}

	return h.GetByID(returnedID)
}

// GetByID fetch a plant by ID
func (h *Handler) GetByID(id string) (*model.Plant, error) {
	selectSQL := `
	SELECT id, common_name, botanical_name, sun_preference,
	  water_preference, soil_preference
	FROM plant WHERE id=$1`

	stmt, err := h.conn.Prepare(selectSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res := stmt.QueryRow(id)
	if err != nil {
		return nil, err
	}

	p := &model.Plant{}
	if err := res.Scan(
		&p.ID,
		&p.CommonName,
		&p.BotanicalName,
		&p.SunPreference,
		&p.WaterPreference,
		&p.SoilPreference,
	); err != nil {
		return nil, &db.NoSuchEntity{Type: "plant"}
	}

	return p, nil
}

// GetAll fetch all plants from database
func (h *Handler) GetAll() ([]*model.Plant, error) {
	selectSQL := `
	SELECT id, common_name, botanical_name, sun_preference,
	  water_preference, soil_preference
	FROM plant`

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

	plants := []*model.Plant{}
	for rows.Next() {
		p := &model.Plant{}
		rows.Scan(
			&p.ID,
			&p.CommonName,
			&p.BotanicalName,
			&p.SunPreference,
			&p.WaterPreference,
			&p.SoilPreference,
		)
		plants = append(plants, p)
	}

	return plants, nil
}
