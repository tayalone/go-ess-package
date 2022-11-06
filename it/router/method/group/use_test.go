package group

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tayalone/go-ess-package/it/router/mock"
	"github.com/tayalone/go-ess-package/router"
)

type GetUseSuite struct {
	suite.Suite
	router router.Route
}

/*SetupSuite init setup for Router*/
func (suite *GetUseSuite) SetupSuite() {
	/* Do Not Thing */
}

// BeforeTest run before each test
func (suite *GetUseSuite) BeforeTest(suiteName, testName string) {
	var routeType string

	switch testName {
	case "TestFiber":
		routeType = "FIBER"
	default:
		routeType = "GIN"
	}
	suite.router = mock.MakeRoute(routeType)
}

func (suite *GetUseSuite) runTest() {
	statusCode, actual := suite.router.Testing(http.MethodGet, "/v1/test-added-use", nil)

	wantMap := map[string]interface{}{
		"message":  "OK",
		"group-v1": "v1",
	}

	want, _ := json.Marshal(wantMap)

	suite.Equal(http.StatusOK, statusCode)
	suite.JSONEq(string(want), actual)
}

func (suite *GetUseSuite) TestGin() {
	suite.runTest()
}

/*TestGinRouteSuite is trigger run it test*/
func TestRouteUseSuite(t *testing.T) {
	suite.Run(t, new(GetTestSuite))
}
