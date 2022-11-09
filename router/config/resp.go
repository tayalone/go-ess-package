package config

type BadReqResponse struct {
	StatusCode int `default:"400"`
	Payloads   map[string]interface{}
}

// type BadReqPayLoads struct {
// 	Message
// }
