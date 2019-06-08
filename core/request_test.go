package core

import (
	"reflect"
	"testing"
)

func TestNewRequest(t *testing.T) {
	r, err := NewRequest()
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name    string
		want    *Request
		wantErr bool
	}{
		{name: "new", want: r, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_SetToken(t *testing.T) {
	r := Request{}
	type args struct {
		token string
	}
	tests := []struct {
		name string
		c    *Request
		args args
	}{
		{name: "set token", c: &r, args: args{token: "sometoken"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.SetToken(tt.args.token)
		})
	}
}
