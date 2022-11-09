package convertor

import (
	"reflect"
	"testing"
)

// type BadReqResponse struct {
// 	StatusCode int `default:"400"`
// 	Payloads   map[string]interface{}
// }

type Book struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type ReqPayLoads struct {
	Message string
	Books   []Book
}

func TestStructToMap(t *testing.T) {
	books := []Book{
		{ID: 1, Title: "book1"},
		{ID: 2, Title: "book2"},
	}
	payload := ReqPayLoads{
		Message: "OK",
		Books:   books,
	}

	type args struct {
		s interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Case: Non-Struct",
			args: args{
				s: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Case: Struct",
			args: args{
				s: payload,
			},
			want: map[string]interface{}{
				"message": "OK",
				"books":   books,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StructToMap(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
