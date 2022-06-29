package acceptance

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RandomNumberSuite struct {
	suite.Suite
}

func (s *RandomNumberSuite) SetupTest() {
	resetMocks(s.T())
}

func (s *RandomNumberSuite) TearDownTest() {
	assertExpectations(s.T())
}

func (s *RandomNumberSuite) Test_RandomNumber() {
	// Mock dependencies
	mockRandomNumber.On("GetRandomNumber").
		Once().
		Return(1, nil)

	mockRandomNumber.On("GetRandomNumber").
		Once().
		Return(5, nil)

	url := "http://randomnumber:8080/random-sum"
	response, statusCode := makeRequest(s.T(), http.MethodGet, url)
	s.Require().Equal(http.StatusOK, statusCode)
	s.Require().Equal(6, response)
}

func TestRandomNumberSuite(t *testing.T) {
	suite.Run(t, new(RandomNumberSuite))
}
