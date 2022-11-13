package gin

import (
	"fmt"

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

/*GetHeader by Key*/
func (mc *MyContext) GetHeader(key string) (string, bool) {
	h := mc.Context.GetHeader(key)
	fmt.Println("GetHeader", key, h)
	return h, h != ""
}

/*SetHeader by Key*/
func (mc *MyContext) SetHeader(key string, value string) {
	fmt.Println("SetHeader", key, value)
	mc.Context.Writer.Header().Set(key, value)
}

/*GetQuery by name*/
func (mc *MyContext) GetQuery(key string) string {
	return mc.Context.Query(key)
}

/*BindJSON must retun erro when query params invalidate*/
func (mc *MyContext) BindJSON(i interface{}) (response.BadReqResponse, error) {
	err := mc.Context.ShouldBindJSON(i)
	if err != nil {
		resp, _ := response.GenBadReqRes(err, "body")
		return resp, err
	}
	return response.BadReqResponse{}, nil
}

/*BindFormData must retun erro when query params invalidate*/
func (mc *MyContext) BindFormData(i interface{}) (response.BadReqResponse, error) {
	err := mc.Context.ShouldBind(i)
	if err != nil {
		resp, _ := response.GenBadReqRes(err, "form-data")
		return resp, err
	}
	return response.BadReqResponse{}, nil
}

/*NewMyContext return MyContext whince overcomposition GIN Context*/
func NewMyContext(c *gin.Context) *MyContext {
	return &MyContext{Context: c}
}
