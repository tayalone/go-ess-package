package response

import (
	"net/http"

	"github.com/tayalone/go-ess-package/convertor"
)

/*BadReqResponse is Struct which use in JSON Method */
type BadReqResponse struct {
	StatusCode int `default:"400"`
	Payload    map[string]interface{}
}

/*BadReqPayload is Struct which convert to body of http response*/
type BadReqPayload struct {
	Message   string `default:"Validation Failed"`
	Parameter string
	Details   []convertor.ErrorObj
}

/*GenBadReqRes Gen Bad Req Payload */
func GenBadReqRes(err error, p string) (BadReqResponse, error) {
	details, errDetail := convertor.ErrorValidate(err)
	if errDetail != nil {
		return BadReqResponse{}, errDetail
	}
	brp := BadReqPayload{
		Message:   "Validation Failed",
		Parameter: p,
		Details:   details,
	}
	payload, _ := convertor.StructToMap(brp)

	return BadReqResponse{StatusCode: http.StatusBadRequest, Payload: payload}, nil
}
