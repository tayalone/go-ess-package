package router

import "github.com/tayalone/go-ess-package/router/response"

/*Context is Behavior of Route Context In Application*/
type Context interface {
	Next()
	JSON(int, interface{})
	Set(string, interface{})
	Get(string) (interface{}, bool)
	BindURI(interface{}) (response.BadReqResponse, error)
}
