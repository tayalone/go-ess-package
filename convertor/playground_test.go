package convertor

import (
	"errors"
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"
)

type User struct {
	FirstName      string `validate:"required"`
	LastName       string `validate:"required"`
	Age            uint8  `validate:"gte=0,lte=130"`
	Email          string `validate:"required,email"`
	FavouriteColor string `validate:"iscolor"` // alias for 'hexcolor|rgb|rgba|hsl|hsla'
}

func TestErrorValidate(t *testing.T) {
	validate := validator.New()

	user := &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
	}

	errPGV := validate.Struct(user)

	type args struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		want    []ErrorObj
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Case: Not go-playground/validator Error",
			args: args{
				err: errors.New("self made error"),
			},
			want:    []ErrorObj{},
			wantErr: true,
		},
		{
			name: "Case: Not go-playground/validator Error",
			args: args{
				err: errPGV,
			},
			want: []ErrorObj{
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ErrorValidate(tt.args.err)
			if (err != nil) != tt.wantErr {
				t.Errorf("ErrorValidate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrorValidate() = %v, want %v", got, tt.want)
			}
		})
	}
}
