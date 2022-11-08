package method

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tayalone/go-ess-package/it/router/mock"
	"github.com/tayalone/go-ess-package/router"
)

type NotFoundTestSuite struct {
	suite.Suite
	router router.Route
}

/*SetupSuite init setup for Router*/
func (suite *NotFoundTestSuite) SetupSuite() {
	/* Do Not Thing */
}

// BeforeTest run before each test
func (suite *NotFoundTestSuite) BeforeTest(suiteName, testName string) {
	var routeType string

	switch testName {
	case "TestFiber":
		routeType = "FIBER"
	default:
		routeType = "GIN"
	}
	suite.router = mock.MakeRoute(routeType)
}

func (suite *NotFoundTestSuite) runTest() {
	statusCode, actual := suite.router.Testing(http.MethodGet, "/test-get", nil)

	wantMap := map[string]interface{}{
		"message": "Not Found",
	}

	want, _ := json.Marshal(wantMap)

	suite.Equal(http.StatusNotFound, statusCode)
	suite.JSONEq(string(want), actual)
}

func (suite *NotFoundTestSuite) TestGin() {
	suite.runTest()
}

/*TestGinRouteSuite is trigger run it test*/
func TestNotFoundRouteSuite(t *testing.T) {
	suite.Run(t, new(GetTestSuite))
}
