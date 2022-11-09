package convertor

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

/*ErrorObj is Error Detail which extract from go-playground/validator error */
type ErrorObj struct {
	Field   string      `json:"field"`
	Type    string      `json:"type"`
	Value   interface{} `json:"value"`
	Message string      `json:"message"`
}

/*ErrorValidate convert  go-playground/validator error to []ErrorObj*/
func ErrorValidate(err error) (errObjs []ErrorObj, panicErr error) {
	errObjs = []ErrorObj{}

	defer func() {
		if r := recover(); r != nil {
			panicErr = errors.New("err not go-playground/validator error")
		}
	}()

	for _, err := range err.(validator.ValidationErrors) {
		param := ""
		if err.Param() != "" {
			param = "_" + err.Param()
		}
		errObjs = append(errObjs, ErrorObj{
			Field:   err.StructField(),
			Type:    fmt.Sprintf("%s.%s%s", err.Type(), err.Tag(), param),
			Value:   err.Value(),
			Message: err.Error(),
		})
	}

	return errObjs, nil
}
