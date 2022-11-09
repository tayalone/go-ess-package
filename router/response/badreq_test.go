package response

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/tayalone/go-ess-package/convertor"
)

type User struct {
	FirstName      string `validate:"required"`
	LastName       string `validate:"required"`
	Age            uint8  `validate:"gte=0,lte=130"`
	Email          string `validate:"required,email"`
	FavouriteColor string `validate:"iscolor"` // alias for 'hexcolor|rgb|rgba|hsl|hsla'
}

func TestGenBadReqRes(t *testing.T) {
	validate := validator.New()

	user := &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
	}

	errPGV := validate.Struct(user)

	got, err := convertor.ErrorValidate(errPGV)

	fmt.Println("got", got)
	fmt.Println("err", err)

	type args struct {
		err error
		p   string
	}
	tests := []struct {
		name    string
		args    args
		want    BadReqResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Case: Wrong Error Format",
			args: args{
				err: errors.New("wrong format"),
				p:   "params",
			},
			want:    BadReqResponse{},
			wantErr: true,
		},
		{
			name: "Case: Correct Error Format",
			args: args{
				err: errPGV,
				p:   "params",
			},
			want: BadReqResponse{
				StatusCode: http.StatusBadRequest,
				Payload: map[string]interface{}{
					"message":   "Validation Failed",
					"parameter": "params",
					"details": []convertor.ErrorObj{
						{
							Field:   "Age",
							Type:    "uint8.lte_130",
							Value:   uint8(135),
							Message: "Key: 'User.Age' Error:Field validation for 'Age' failed on the 'lte' tag",
						},
						{
							Field:   "FavouriteColor",
							Type:    "string.iscolor",
							Value:   "#000-",
							Message: "Key: 'User.FavouriteColor' Error:Field validation for 'FavouriteColor' failed on the 'iscolor' tag",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenBadReqRes(tt.args.err, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenBadReqRes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			fmt.Println("got", got)
			fmt.Println("tt.want", tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenBadReqRes() = %v, want %v", got, tt.want)
			}
		})
	}
}
