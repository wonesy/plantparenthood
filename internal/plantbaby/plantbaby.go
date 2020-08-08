package plantbaby

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
func (h *Handler) Create(ownerID string, n *model.NewNurseryAddition) (*model.PlantBaby, error) {
	insertSQL := `
	INSERT INTO plant_baby (id, owner, plant, care_regimen, nickname, location)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
	`

	id := util.GenerateID()

	returnedID := ""
	if err := h.conn.QueryRow(insertSQL,
		id, ownerID, n.PlantID,
		n.CareRegimenID, n.Nickname, n.Location,
	).Scan(&returnedID); err != nil {
		return nil, err
	}

	return h.GetByID(id)
}

// GetByID fetch a plantbaby by ID
func (h *Handler) GetByID(id string) (*model.PlantBaby, error) {
	selectSQL := `
	SELECT *
	FROM plant_baby
	WHERE id=$1`

	stmt, err := h.conn.Prepare(selectSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	pb := &model.PlantBaby{}
	if err := stmt.QueryRow(id).Scan(pb); err != nil {
		return nil, &db.NoSuchEntity{Type: "plantbaby"}
	}

	return pb, nil
}

// GetAllByOwnerID fetch all plant babies from database for a particular owner
func (h *Handler) GetAllByOwnerID(ownerID string) ([]*model.PlantBaby, error) {
	selectSQL := `
	SELECT
		pb.id, pb.owner, pb.nickname, pb.location, pb.added_on,
		p.id, p.botanical_name, p.common_name, p.sun_preference, p.water_preference, p.soil_preference,
		cr.id, cr.water_ml, cr.water_hr
	FROM plant_baby pb
		INNER JOIN plant p ON pb.plant=p.id
		INNER JOIN care_regimen cr ON pb.care_regimen=cr.id
	WHERE pb.owner=$1
	`

	stmt, err := h.conn.Prepare(selectSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	nursery := []*model.PlantBaby{}
	for rows.Next() {
		pbaby := &model.PlantBaby{}
		plant := &model.Plant{}
		care := &model.CareRegimen{}

		if err := rows.Scan(
			&pbaby.ID,
			&pbaby.OwnerID,
			&pbaby.Nickname,
			&pbaby.Location,
			&pbaby.AddedOn,
			//
			&plant.ID,
			&plant.BotanicalName,
			&plant.CommonName,
			&plant.SunPreference,
			&plant.WaterPreference,
			&plant.SoilPreference,
			//
			&care.ID,
			&care.Waterml,
			&care.Waterhr,
		); err != nil {
			return nil, err
		}

		pbaby.Plant = plant
		pbaby.CareRegimen = care
		nursery = append(nursery, pbaby)
	}

	return nursery, nil
}

// MemberOwnsPlant returns true if member has an associated owned plant baby
func (h *Handler) MemberOwnsPlant(ownerID, plantBabyID string) error {
	selectSQL := `
	SELECT id
	FROM plant_baby
	WHERE owner=$1 AND id=$2
	`

	stmt, err := h.conn.Prepare(selectSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	id := ""
	if err := stmt.QueryRow(ownerID, plantBabyID).Scan(&id); err != nil {
		return &db.NoSuchEntity{Type: "plantbaby"}
	}

	return nil
}
