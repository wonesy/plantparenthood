package plant

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PlantTestSuite struct {
	suite.Suite
}

func (s *PlantTestSuite) TestPlantTestSuite(t *testing.T) {
	suite.Run(t, new(PlantTestSuite))
}

func (s *PlantTestSuite) SetupTest() {
}
