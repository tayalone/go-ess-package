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

	myRouter.POST("/test-post", func(c router.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Test Route 'POST' OK!!",
		})
	})

	myRouter.PATCH("/test-patch", func(c router.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Test Route 'PATCH' OK!!",
		})
	})

	myRouter.PUT("/test-put", func(c router.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Test Route 'PUT' OK!!",
		})
	})

	myRouter.DELETE("/test-delete", func(c router.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Test Route 'DELETE' OK!!",
		})
	})

	v1 := myRouter.Group("/v1")
	{
		sub := v1.Group("/sub")
		{
			sub.GET("/test-get", func(c router.Context) {
				c.JSON(http.StatusOK, map[string]interface{}{
					"message": "Test Sub Route Grouper 'GET' OK!!",
				})
			})
			sub.POST("/test-post", func(c router.Context) {
				c.JSON(http.StatusOK, map[string]interface{}{
					"message": "Test Sub Route Grouper 'POST' OK!!",
				})
			})
			sub.PATCH("/test-patch", func(c router.Context) {
				c.JSON(http.StatusOK, map[string]interface{}{
					"message": "Test Sub Route Grouper 'PATCH' OK!!",
				})
			})
			sub.PUT("/test-put", func(c router.Context) {
				c.JSON(http.StatusOK, map[string]interface{}{
					"message": "Test Sub Route Grouper 'PUT' OK!!",
				})
			})
			sub.DELETE("/test-delete", func(c router.Context) {
				c.JSON(http.StatusOK, map[string]interface{}{
					"message": "Test Sub Route Grouper 'DELETE' OK!!",
				})
			})
		}

		v1.GET("/test-get", func(c router.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"message": "Test Route Grouper 'GET' OK!!",
			})
		})
		v1.POST("/test-post", func(c router.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"message": "Test Route Grouper 'POST' OK!!",
			})
		})
		v1.PATCH("/test-patch", func(c router.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"message": "Test Route Grouper 'PATCH' OK!!",
			})
		})
		v1.PUT("/test-put", func(c router.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"message": "Test Route Grouper 'PUT' OK!!",
			})
		})
		v1.DELETE("/test-delete", func(c router.Context) {
			c.JSON(http.StatusOK, map[string]interface{}{
				"message": "Test Route Grouper 'DELETE' OK!!",
			})
		})

	}

	myRouter.GET("/test-not-add-use", GetGlobalFromCtx)

	myRouter.Use(UseGlobal)

	myRouter.GET("/test-added-use", GetGlobalFromCtx)

	return myRouter
}
