package gin

import (
	"fmt"
	"net/http"

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
func NewHTTPRouter() *HTTPRouter {
	r := gin.Default()
	// c := router.
	config := config.Read()
	return &HTTPRouter{r, config}
}

func handlerConvertor(h []func(router.Context)) []gin.HandlerFunc {
	ginHandlers := []gin.HandlerFunc{}
	for _, handler := range h {
		ginHandlers = append(ginHandlers, NewRouterHandler(handler))
	}
	return ginHandlers
}

/*GET is HTTP Method Get */
func (r *HTTPRouter) GET(path string, handlers ...func(router.Context)) {
	ginHandlers := handlerConvertor(handlers)

	r.Engine.GET(path, ginHandlers...)
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
