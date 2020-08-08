package watering

import (
	"database/sql"

	"github.com/wonesy/plantparenthood/graph/model"
	"github.com/wonesy/plantparenthood/pkg/datetime"
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

// Create inserts a watering instance to the DB
func (h *Handler) Create(w *model.NewWatering) (*model.Watering, error) {
	insertSQL := `INSERT INTO watering (id,amount_ml,plant_baby_id,watered_on) VALUES ($1,$2,$3,$4) RETURNING id`

	id := util.GenerateID()

	// if provided time is empty, then default to Now
	if w.WateredOn == "" {
		w.WateredOn = datetime.Now()
	}

	if err := datetime.Validate(w.WateredOn); err != nil {
		return nil, err
	}

	returnedID := ""
	if err := h.conn.QueryRow(insertSQL, id, w.Amountml, w.PlantBabyID, w.WateredOn).Scan(&returnedID); err != nil {
		return nil, err
	}

	return &model.Watering{
		ID:          returnedID,
		PlantBabyID: w.PlantBabyID,
		Amountml:    w.Amountml,
		WateredOn:   w.WateredOn,
	}, nil
}

// GetByPlantBabyID returns a list of waterings by the plant baby ID
func (h *Handler) GetByPlantBabyID(plantBabyID string) ([]*model.Watering, error) {
	selectSQL := `
	SELECT w.id, w.plant_baby_id, w.amount_ml, w.watered_on
	FROM watering w
	WHERE w.plant_baby_id=$1

	`

	stmt, err := h.conn.Prepare(selectSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(plantBabyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	waterings := []*model.Watering{}
	for rows.Next() {
		w := &model.Watering{}
		rows.Scan(&w.ID, &w.PlantBabyID, &w.Amountml, &w.WateredOn)
		waterings = append(waterings, w)
	}

	return waterings, nil
}
