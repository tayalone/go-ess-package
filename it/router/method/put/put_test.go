package put

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tayalone/go-ess-package/it/router/mock"
	"github.com/tayalone/go-ess-package/router"
)

type PutTestSuite struct {
	suite.Suite
	router router.Route
}

/*SetupSuite init setup for Router*/
func (suite *PutTestSuite) SetupSuite() {
	/* Do Not Thing */
}

// BeforeTest run before each test
func (suite *PutTestSuite) BeforeTest(suiteName, testName string) {
	var routeType string

	switch testName {
	case "TestFiber":
		routeType = "FIBER"
	default:
		routeType = "GIN"
	}
	suite.router = mock.MakeRoute(routeType)
}

func (suite *PutTestSuite) runTest() {
	statusCode, actual := suite.router.Testing(http.MethodPut, "/test-put", nil)

	wantMap := map[string]interface{}{
		"message": "Test Route 'PUT' OK!!",
	}

	want, _ := json.Marshal(wantMap)

	suite.Equal(http.StatusOK, statusCode)
	suite.JSONEq(string(want), actual)
}

func (suite *PutTestSuite) TestGin() {
	suite.runTest()
}

/*TestGinRouteSuite is trigger run it test*/
func TestRoutePutSuite(t *testing.T) {
	suite.Run(t, new(PutTestSuite))
}
