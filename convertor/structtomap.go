package convertor

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/iancoleman/strcase"
)

/*StructToMap covert every struct to  map[string]interface{} */
func StructToMap(s interface{}) (map[string]interface{}, error) {
	v := reflect.ValueOf(s)

	if v.Kind() != reflect.Struct {
		return nil, errors.New("s is not struct")
	}

	r := make(map[string]interface{})

	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		// r[typeOfS.Field(i).Name] = v.Field(i).Interface()
		fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
		r[strcase.ToLowerCamel(typeOfS.Field(i).Name)] = v.Field(i).Interface()
	}
	return r, nil
}
