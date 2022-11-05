package mock

import (
	"net/http"

	"github.com/tayalone/go-ess-package/router"
	routerGenerator "github.com/tayalone/go-ess-package/router/generator"
)

// MakeRoute create router for integration test
func MakeRoute(routeType string) router.Route {
	myRouter := routerGenerator.New(routeType)

	myRouter.GET("/test-get", func(c router.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Test Route 'GET' OK!!",
		})
	})

	return myRouter
}
