package member

import (
	"database/sql"

	"github.com/wonesy/plantparenthood/graph/model"
	"github.com/wonesy/plantparenthood/internal/auth"
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
func (h *Handler) Create(m *model.NewMember) (string, error) {
	insertSQL := `INSERT INTO member (id,email,password) VALUES ($1,$2,$3) RETURNING id`

	id := util.GenerateID()

	hashPass, err := auth.HashPassword(m.Password)
	if err != nil {
		return "", err
	}

	returnedID := ""
	err = h.conn.QueryRow(insertSQL, id, m.Email, string(hashPass)).Scan(&returnedID)
	if err != nil {
		return "", err
	}

	return returnedID, nil
}

// GetAll fetch all members from database
func (h *Handler) GetAll() ([]*model.Member, error) {
	selectSQL := "SELECT id, email, created_on FROM member"

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

	members := []*model.Member{}
	for rows.Next() {
		m := &model.Member{}
		rows.Scan(&m.ID, &m.Email, &m.CreatedOn)
		members = append(members, m)
	}

	return members, nil
}

// GetByID fetch a member by their ID
func (h *Handler) GetByID(id string) (*model.Member, error) {
	selectSQL := `
	SELECT id, email, created_on
	FROM member
	WHERE id=?`

	stmt, err := h.conn.Prepare(selectSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	m := &model.Member{}
	if err := rows.Scan(&m.ID, &m.Email, &m.CreatedOn); err != nil {
		return nil, err
	}

	return m, nil
}

// Login logs a member in
func (h *Handler) Login(credentials *model.Login) (string, error) {
	selectSQL := "SELECT id, email, password FROM member WHERE email=$1"

	stmt, err := h.conn.Prepare(selectSQL)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	res := stmt.QueryRow(credentials.Email)
	if err != nil {
		return "", err
	}

	id := ""
	email := ""
	hashedPass := ""
	if err := res.Scan(&id, &email, &hashedPass); err != nil {
		return "", err
	}

	if err := auth.Authenticate(hashedPass, credentials.Password); err != nil {
		return "", err
	}

	return id, nil
}
