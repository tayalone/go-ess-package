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
)

/*MyContext is Overide "GIN" contexts*/
type MyContext struct {
	*gin.Context
}

/*Next is command which make router do next Middleware */
func (mc *MyContext) Next() {
	mc.Context.Next()
}

/*JSON is command router retrun rest data with status code*/
func (mc *MyContext) JSON(statusCode int, v interface{}) {
	mc.Context.JSON(statusCode, v)
}

/*NewMyContext return MyContext whince overcomposition GIN Context*/
func NewMyContext(c *gin.Context) *MyContext {
	return &MyContext{Context: c}
}

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
	config := config.Read()

	if config.Mode != "DEBUG" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

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
