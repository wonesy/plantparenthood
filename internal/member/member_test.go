package member

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"github.com/wonesy/plantparenthood/graph/model"
	"github.com/wonesy/plantparenthood/internal/auth"
	"github.com/wonesy/plantparenthood/internal/pkg/testhelper"
	"github.com/wonesy/plantparenthood/util"
)

type MemberTestSuite struct {
	suite.Suite
	conn *sql.DB
	mock sqlmock.Sqlmock
	h    *Handler
}

func TestMemberTestSuite(t *testing.T) {
	suite.Run(t, new(MemberTestSuite))
}

func (s *MemberTestSuite) SetupTest() {
	conn, mock, err := testhelper.OpenDB()
	if err != nil {
		s.FailNow("could not establish database connection: " + err.Error())
	}

	s.conn = conn
	s.mock = mock
	s.h = NewHandler(s.conn)
}

func (s *MemberTestSuite) TeardownTest() {
	s.conn.Close()
}

func (s *MemberTestSuite) TestNewHandler_OK() {
	h := NewHandler(s.conn)
	s.NotNil(h)
}

func (s *MemberTestSuite) TestCreate_OK() {
	s.mock.ExpectBegin()
	s.mock.ExpectPrepare("INSERT INTO").ExpectExec().
		WithArgs(sqlmock.AnyArg(), "tester@tester.com", sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(0, 1)) // no insert ID, 1 row affected
	s.mock.ExpectCommit()

	// get by ID
	s.mock.ExpectPrepare("SELECT .+ FROM member WHERE").
		ExpectQuery().
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "created_on"}).
			AddRow("test_uuid", "tester@tester.com", "hashed_asdf"),
		)

	newmem := &model.NewMember{
		Email:    "tester@tester.com",
		Password: "asdf",
	}

	mem, err := s.h.Create(newmem)

	s.NoError(err)
	s.NotNil(mem)
	s.Equal("tester@tester.com", mem.Email)
	s.NotEmpty(mem.ID)
	s.NotEmpty(mem.CreatedOn)

	testhelper.ExpectationsWereMet(s.mock)
}

func (s *MemberTestSuite) TestGetAll_OK() {
	cols := []string{"id", "email", "created_on"}
	rows := sqlmock.NewRows(cols).AddRow("test_uuid", "tester@tester.com", "test_created_on")

	s.mock.ExpectPrepare(sqlGetAll).
		ExpectQuery().
		WillReturnRows(rows)

	members, err := s.h.GetAll()

	s.NoError(err)
	s.Equal(1, len(members))
	s.Equal("tester@tester.com", members[0].Email)

	testhelper.ExpectationsWereMet(s.mock)
}

func (s *MemberTestSuite) TestGetByID_OK() {
	id := util.GenerateID()

	if testhelper.IsEnd2EndTest() {
		created, err := s.h.Create(&model.NewMember{Email: "getbyid@test.com", Password: "123abc"})
		s.NoError(err)
		id = created.ID
	}

	s.mock.ExpectPrepare("SELECT .+ FROM member WHERE id=").
		ExpectQuery().
		WithArgs(id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "email", "created_on"}).AddRow(id, "asdf@asdf.com", "getbyid_created_on"),
		)

	m, err := s.h.GetByID(id)
	s.NoError(err)
	s.Equal(id, m.ID)

	testhelper.ExpectationsWereMet(s.mock)
}

func (s *MemberTestSuite) TestLogin_OK() {
	email := "login@test.com"
	pass := "asdf"
	hashedPass, _ := auth.HashPassword(pass)

	id := util.GenerateID()
	if testhelper.IsEnd2EndTest() {
		m, _ := s.h.Create(&model.NewMember{Email: email, Password: pass})
		id = m.ID

	}

	s.mock.ExpectPrepare("SELECT .+ FROM member WHERE email=").
		ExpectQuery().
		WithArgs(email).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "email", "password"}).AddRow(id, email, hashedPass),
		)

	loginID, err := s.h.Login(&model.Login{Email: email, Password: pass})
	s.NoError(err)

	s.Equal(id, loginID)

	testhelper.ExpectationsWereMet(s.mock)
}

func (s *MemberTestSuite) TestLogin_ErrInvalidCredentials() {
	email := "login@test.com"
	pass := "asdf"
	hashedPass := "bad_hash"

	id := util.GenerateID()
	if testhelper.IsEnd2EndTest() {
		m, _ := s.h.Create(&model.NewMember{Email: email, Password: pass})
		id = m.ID

	}

	s.mock.ExpectPrepare("SELECT .+ FROM member WHERE email=").
		ExpectQuery().
		WithArgs(email).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "email", "password"}).AddRow(id, email, hashedPass),
		)

	loginID, err := s.h.Login(&model.Login{Email: email, Password: pass})
	s.Error(err)
	s.NotEqual(id, loginID)

	testhelper.ExpectationsWereMet(s.mock)
}
