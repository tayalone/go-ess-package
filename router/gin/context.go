package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/tayalone/go-ess-package/router/response"
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

/*Set is assingning key n' value to router ctx*/
func (mc *MyContext) Set(key string, value interface{}) {
	mc.Context.Set(key, value)
}

/*BindURI must retun erro when query params invalidate*/
func (mc *MyContext) BindURI(i interface{}) (response.BadReqResponse, error) {
	err := mc.Context.ShouldBindUri(i)
	if err != nil {
		resp, _ := response.GenBadReqRes(err, "params")
		return resp, err
	}
	return response.BadReqResponse{}, nil
}

/*Get is assingning key n' value to router ctx*/
func (mc *MyContext) Get(key string) (value interface{}, isExist bool) {
	value, isExist = mc.Context.Get(key)
	if !isExist {
		return nil, false
	}
	return value, isExist
}

/*NewMyContext return MyContext whince overcomposition GIN Context*/
func NewMyContext(c *gin.Context) *MyContext {
	return &MyContext{Context: c}
}
