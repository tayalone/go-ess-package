package patch

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tayalone/go-ess-package/it/router/mock"
	"github.com/tayalone/go-ess-package/router"
)

type PatchTestSuite struct {
	suite.Suite
	router router.Route
}

/*SetupSuite init setup for Router*/
func (suite *PatchTestSuite) SetupSuite() {
	/* Do Not Thing */
}

// BeforeTest run before each test
func (suite *PatchTestSuite) BeforeTest(suiteName, testName string) {
	var routeType string

	switch testName {
	case "TestFiber":
		routeType = "FIBER"
	default:
		routeType = "GIN"
	}
	suite.router = mock.MakeRoute(routeType)
}

func (suite *PatchTestSuite) runTest() {
	statusCode, actual := suite.router.Testing(http.MethodPatch, "/test-patch", nil)

	wantMap := map[string]interface{}{
		"message": "Test Route 'PATCH' OK!!",
	}

	want, _ := json.Marshal(wantMap)

	suite.Equal(http.StatusOK, statusCode)
	suite.JSONEq(string(want), actual)
}

func (suite *PatchTestSuite) TestGin() {
	suite.runTest()
}

/*TestGinRouteSuite is trigger run it test*/
func TestRoutePatchSuite(t *testing.T) {
	suite.Run(t, new(PatchTestSuite))
}
