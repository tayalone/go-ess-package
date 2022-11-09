package config

import "github.com/tayalone/go-ess-package/convertor"

/*BadReqResponse is Struct which use in JSON Method */
type BadReqResponse struct {
	StatusCode int `default:"400"`
	Payloads   map[string]interface{}
}

/*BadReqPayload is Struct which convert to body of http response*/
type BadReqPayload struct {
	Message   string `default:"Validation Failed"`
	Parameter string
	Details   []convertor.ErrorObj
}
