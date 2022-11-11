package gin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/tayalone/go-ess-package/router"
	"github.com/tayalone/go-ess-package/router/config"
	rotuerConfig "github.com/tayalone/go-ess-package/router/config"
)

/*NewRouterHandler covert  MyContext (OverComposiont Gin) -> Gin Context */
func NewRouterHandler(handler func(c router.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(NewMyContext(c))
	}
}

// HTTPRouter is Overided Gin Engine
type HTTPRouter struct {
	*gin.Engine
	config config.Config
}

// NewHTTPRouter retun my engin
func NewHTTPRouter() router.Route {
	config := rotuerConfig.Read()

	if config.Mode != "DEBUG" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": rotuerConfig.RespMsg["NOT_FOUND"],
		})
	})

	return &HTTPRouter{r, config}
}

func handlerConvertor(h []func(router.Context)) []gin.HandlerFunc {
	ginHandlers := []gin.HandlerFunc{}
	for _, handler := range h {
		ginHandlers = append(ginHandlers, NewRouterHandler(handler))
	}
	return ginHandlers
}

/*Start Http-Router*/
func (r *HTTPRouter) Start() {
	r.GET("/status", func(ctx router.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "OK",
		})
	})

	port := fmt.Sprintf(":%d", r.config.Port)

	r.Run(port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

/*Testing make Gin Testing Call API and return result and statuscode*/
func (r *HTTPRouter) Testing(method string, path string, body map[string]interface{}) (int, string) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)

	req, _ := http.NewRequest(method, path, b)
	w := httptest.NewRecorder()
	r.Engine.ServeHTTP(w, req)

	return w.Code, w.Body.String()
}

/*GET is HTTP Method Get */
func (r *HTTPRouter) GET(path string, handlers ...func(router.Context)) {
	ginHandlers := handlerConvertor(handlers)
	r.Engine.GET(path, ginHandlers...)
}

/*POST is HTTP Method Post */
func (r *HTTPRouter) POST(path string, handlers ...func(router.Context)) {
	ginHandlers := handlerConvertor(handlers)
	r.Engine.POST(path, ginHandlers...)
}

/*PATCH is HTTP Method Patch */
func (r *HTTPRouter) PATCH(path string, handlers ...func(router.Context)) {
	ginHandlers := handlerConvertor(handlers)
	r.Engine.PATCH(path, ginHandlers...)
}

/*PUT is HTTP Method Put */
func (r *HTTPRouter) PUT(path string, handlers ...func(router.Context)) {
	ginHandlers := handlerConvertor(handlers)
	r.Engine.PUT(path, ginHandlers...)
}

/*DELETE is HTTP Method Put */
func (r *HTTPRouter) DELETE(path string, handlers ...func(router.Context)) {
	ginHandlers := handlerConvertor(handlers)
	r.Engine.DELETE(path, ginHandlers...)
}

/*Use is inject Middleware To Http Router */
func (r *HTTPRouter) Use(middleware func(router.Context)) {
	// ginHandlers := handlerConvertor(middlewares)
	r.Engine.Use(NewRouterHandler(middleware))
}

/*Group  Routing*/
func (r *HTTPRouter) Group(path string, handlers ...func(router.Context)) router.RoterGrouper {
	ginHandlers := handlerConvertor(handlers)
	return &HTTPRouterGroup{RouterGroup: r.Engine.Group(path, ginHandlers...)}
}

/*HTTPRouterGroup .... */
type HTTPRouterGroup struct {
	*gin.RouterGroup
}

/*GET is Router Grouper HTTP Method Get */
func (g *HTTPRouterGroup) GET(path string, handlers ...func(router.Context)) {
	ginHandlers := handlerConvertor(handlers)
	g.RouterGroup.GET(path, ginHandlers...)
}

/*POST is Router Grouper HTTP Method Get */
func (g *HTTPRouterGroup) POST(path string, handlers ...func(router.Context)) {
	ginHandlers := handlerConvertor(handlers)
	g.RouterGroup.POST(path, ginHandlers...)
}

/*PATCH is Router Grouper HTTP Method Get */
func (g *HTTPRouterGroup) PATCH(path string, handlers ...func(router.Context)) {
	ginHandlers := handlerConvertor(handlers)
	g.RouterGroup.PATCH(path, ginHandlers...)
}

/*PUT is Router Grouper HTTP Method Get */
func (g *HTTPRouterGroup) PUT(path string, handlers ...func(router.Context)) {
	ginHandlers := handlerConvertor(handlers)
	g.RouterGroup.PUT(path, ginHandlers...)
}

/*DELETE is Router Grouper HTTP Method Get */
func (g *HTTPRouterGroup) DELETE(path string, handlers ...func(router.Context)) {
	ginHandlers := handlerConvertor(handlers)
	g.RouterGroup.DELETE(path, ginHandlers...)
}

/*Use is inject Middleware To Http Router */
func (g *HTTPRouterGroup) Use(middleware func(router.Context)) {
	g.RouterGroup.Use(NewRouterHandler(middleware))
}

/*Group  Routing*/
func (g *HTTPRouterGroup) Group(path string, handlers ...func(router.Context)) router.RoterGrouper {
	ginHandlers := handlerConvertor(handlers)
	return &HTTPRouterGroup{RouterGroup: g.RouterGroup.Group(path, ginHandlers...)}
}
