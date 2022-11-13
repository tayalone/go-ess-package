package router

import (
	"net/http"

	"github.com/tayalone/go-ess-package/router/response"
)

/*Context is Behavior of Route Context In Application*/
type Context interface {
	Next()
	JSON(int, interface{})
	Set(string, interface{})
	Get(string) (interface{}, bool)
	GetHeader(string) (string, bool)
	GetQuery(string) string
	GetHTTPRequest() *http.Request
	SetHeader(string, string)
	BindJSON(interface{}) (response.BadReqResponse, error)
	BindFormData(interface{}) (response.BadReqResponse, error)
	// BindURI(interface{}) (response.BadReqResponse, error)
}
