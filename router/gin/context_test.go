package gin

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/tayalone/go-ess-package/router/response"
)

func TestMyContext_Next(t *testing.T) {
	type fields struct {
		Context *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &MyContext{
				Context: tt.fields.Context,
			}
			mc.Next()
		})
	}
}

func TestMyContext_JSON(t *testing.T) {
	type fields struct {
		Context *gin.Context
	}
	type args struct {
		statusCode int
		v          interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &MyContext{
				Context: tt.fields.Context,
			}
			mc.JSON(tt.args.statusCode, tt.args.v)
		})
	}
}

func TestMyContext_Set(t *testing.T) {
	type fields struct {
		Context *gin.Context
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &MyContext{
				Context: tt.fields.Context,
			}
			mc.Set(tt.args.key, tt.args.value)
		})
	}
}

func TestMyContext_BindURI(t *testing.T) {
	type fields struct {
		Context *gin.Context
	}
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    response.BadReqResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &MyContext{
				Context: tt.fields.Context,
			}
			got, err := mc.BindURI(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("MyContext.BindURI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MyContext.BindURI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyContext_Get(t *testing.T) {
	type fields struct {
		Context *gin.Context
	}
	type args struct {
		key string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantValue   interface{}
		wantIsExist bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &MyContext{
				Context: tt.fields.Context,
			}
			gotValue, gotIsExist := mc.Get(tt.args.key)
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("MyContext.Get() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotIsExist != tt.wantIsExist {
				t.Errorf("MyContext.Get() gotIsExist = %v, want %v", gotIsExist, tt.wantIsExist)
			}
		})
	}
}

func TestNewMyContext(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
		want *MyContext
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMyContext(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMyContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
