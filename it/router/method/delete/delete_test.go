package delete

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tayalone/go-ess-package/it/router/mock"
	"github.com/tayalone/go-ess-package/router"
)

type DeleteTestSuite struct {
	suite.Suite
	router router.Route
}

/*SetupSuite init setup for Router*/
func (suite *DeleteTestSuite) SetupSuite() {
	/* Do Not Thing */
}

// BeforeTest run before each test
func (suite *DeleteTestSuite) BeforeTest(suiteName, testName string) {
	var routeType string

	switch testName {
	case "TestFiber":
		routeType = "FIBER"
	default:
		routeType = "GIN"
	}
	suite.router = mock.MakeRoute(routeType)
}

func (suite *DeleteTestSuite) runTest() {
	statusCode, actual := suite.router.Testing(http.MethodDelete, "/test-delete", nil)

	wantMap := map[string]interface{}{
		"message": "Test Route 'DELETE' OK!!",
	}

	want, _ := json.Marshal(wantMap)

	suite.Equal(http.StatusOK, statusCode)
	suite.JSONEq(string(want), actual)
}

func (suite *DeleteTestSuite) TestGin() {
	suite.runTest()
}

/*TestGinRouteSuite is trigger run it test*/
func TestRouteDeleteSuite(t *testing.T) {
	suite.Run(t, new(DeleteTestSuite))
}
