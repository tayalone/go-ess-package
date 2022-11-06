package post

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tayalone/go-ess-package/it/router/mock"
	"github.com/tayalone/go-ess-package/router"
)

type PostTestSuite struct {
	suite.Suite
	router router.Route
}

/*SetupSuite init setup for Router*/
func (suite *PostTestSuite) SetupSuite() {
	/* Do Not Thing */
}

// BeforeTest run before each test
func (suite *PostTestSuite) BeforeTest(suiteName, testName string) {
	var routeType string

	switch testName {
	case "TestFiber":
		routeType = "FIBER"
	default:
		routeType = "GIN"
	}
	suite.router = mock.MakeRoute(routeType)
}

func (suite *PostTestSuite) runTest() {
	statusCode, actual := suite.router.Testing(http.MethodPost, "/v1/test-post", nil)

	wantMap := map[string]interface{}{
		"message": "Test Route Grouper 'POST' OK!!",
	}

	want, _ := json.Marshal(wantMap)

	suite.Equal(http.StatusOK, statusCode)
	suite.JSONEq(string(want), actual)
}

func (suite *PostTestSuite) TestGin() {
	suite.runTest()
}

/*TestGinRouteSuite is trigger run it test*/
func TestRoutePostSuite(t *testing.T) {
	suite.Run(t, new(PostTestSuite))
}
