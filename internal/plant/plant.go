package plant

import (
	"database/sql"

	"github.com/wonesy/plantparenthood/graph/model"
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
func (h *Handler) Create(p *model.NewPlant) (string, error) {
	insertSQL := `
	INSERT INTO plant (id, common_name, botanical_name, water_preference, sun_preference, soil_preference)
	VALUES ($1,$2,$3,$4,$5, $6) RETURNING id`

	id := util.GenerateID()

	returnedID := ""
	if err := h.conn.QueryRow(
		insertSQL,
		id,
		p.BotanicalName,
		p.CommonName,
		p.WaterPreference,
		p.SunPreference,
		p.SoilPreference,
	).Scan(&returnedID); err != nil {
		return "", err
	}

	return returnedID, nil
}

// GetAll fetch all plants from database
func (h *Handler) GetAll() ([]*model.Plant, error) {
	selectSQL := "SELECT * FROM plant"

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
		rows.Scan(&p.ID, &p.CommonName, &p.BotanicalName, &p.WaterPreference, &p.SunPreference, &p.SoilPreference)
		plants = append(plants, p)
	}

	return plants, nil
}
