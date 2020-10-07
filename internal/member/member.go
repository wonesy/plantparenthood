package member

import (
	"context"
	"database/sql"

	"github.com/wonesy/plantparenthood/internal/pkg/db"

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
func (h *Handler) Create(m *model.NewMember) (*model.Member, error) {
	id := util.GenerateID()

	hashPass, err := auth.HashPassword(m.Password)
	if err != nil {
		return nil, err
	}

	tx, err := h.conn.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(sqlCreateMember)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(id, m.Email, string(hashPass))
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return h.GetByID(id)
}

// GetAll fetch all members from database
func (h *Handler) GetAll() ([]*model.Member, error) {
	stmt, err := h.conn.Prepare(sqlGetAll)
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
	stmt, err := h.conn.Prepare(sqlGetMemberByID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res := stmt.QueryRow(id)
	if err != nil {
		return nil, err
	}

	m := &model.Member{}
	if err := res.Scan(&m.ID, &m.Email, &m.CreatedOn); err != nil {
		return nil, &db.NoSuchEntity{Type: "member"}
	}

	return m, nil
}

// Login logs a member in
func (h *Handler) Login(credentials *model.Login) (string, error) {
	stmt, err := h.conn.Prepare(sqlLogin)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	id := ""
	email := ""
	hashedPass := ""

	if err := stmt.QueryRow(credentials.Email).Scan(&id, &email, &hashedPass); err != nil {
		return "", &AuthenticationError{}
	}

	if err := auth.Authenticate(hashedPass, credentials.Password); err != nil {
		return "", &AuthenticationError{}
	}

	return id, nil
}

// ValidateMemberFromContext auths member from ctx token and verifies it exists in DB
func (h *Handler) ValidateMemberFromContext(ctx context.Context) (string, error) {
	memberID := auth.IDFromContext(ctx)
	if memberID == "" {
		return "", &auth.AuthenticationError{}
	}

	// ensure that this user exists before adding to the database
	_, err := h.GetByID(memberID)
	if err != nil {
		return "", &db.NoSuchEntity{Type: "member"}
	}

	return memberID, nil
}
