package generate

import (
	"github.com/tayalone/go-ess-package/router"
	gin "github.com/tayalone/go-ess-package/router/gin"
)

/*
New Return Http Router varaint with Routet Type (routerType)
*/
func New(routerType string) router.Route {
	switch routerType {
	case "GIN":
		return gin.NewHTTPRouter()
	default:
		return gin.NewHTTPRouter()
	}
}
