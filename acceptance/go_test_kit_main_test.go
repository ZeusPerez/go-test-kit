package acceptance

import (
	"net/http"
	"os"
	"strconv"
	"testing"
	"testkit/acceptance/mocks/randomnumber"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
)

var (
	mockRandomNumber = &randomnumber.MockInterface{}
	mocksCollection  = []*mock.Mock{
		&mockRandomNumber.Mock,
	}
)

func TestMain(m *testing.M) {
	setUpMockedServices()
	code := m.Run()

	os.Exit(code)
}

func setUpMockedServices() {
	log.Info("setUpMockedServices")

	randomNumberHandler := randomnumber.NewHandler(mockRandomNumber)
	randomNumberHandler = newWithMethod(randomNumberHandler, http.MethodGet, "/").(*mux.Router)

	startHTTPServer(randomNumberHandler, 5028)
}

func resetMocks(t *testing.T) {
	for _, m := range mocksCollection {
		m.ExpectedCalls = []*mock.Call{}
		m.Calls = []mock.Call{}
		m.Test(t)
	}
}

func assertExpectations(t *testing.T) {
	mock.AssertExpectationsForObjects(t,
		mockRandomNumber,
	)
}

func startHTTPServer(handler http.Handler, port int) {
	go http.ListenAndServe(":"+strconv.FormatInt(int64(port), 10), handler)
}

func newWithMethod(handler http.Handler, method string, path string) http.Handler {
	r := mux.NewRouter()
	r.Handle(path, handler).Methods(method)
	return r
}
