package member

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/wonesy/plantparenthood/internal/pkg/testhelper"
)

type MemberTestSuite struct {
	suite.Suite
	conn *sql.DB
}

func (s *MemberTestSuite) TestMemberTestSuite(t *testing.T) {
	suite.Run(t, new(MemberTestSuite))
}

func (s *MemberTestSuite) SetupTest() {
	conn, err := testhelper.OpenDB()
	if err != nil {
		s.FailNow("could not establish database connection")
	}
	s.conn = conn
}
