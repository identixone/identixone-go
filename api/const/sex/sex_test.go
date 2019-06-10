package sex

import (
	"reflect"
	"testing"
)

func TestSex_String(t *testing.T) {
	tests := []struct {
		name string
		s    Sex
		want string
	}{
		{name: "male", s: Sex(0), want: "male"},
		{name: "male", s: Sex(1), want: "female"},
		{name: "unknown", s: Sex(2), want: "unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("Sex.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAll(t *testing.T) {
	tests := []struct {
		name string
		want []Sex
	}{
		{name: "all", want: []Sex{Male, Female}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSexes_Validate(t *testing.T) {
	tests := []struct {
		name    string
		sx      Sexes
		wantErr bool
	}{
		{name: "valid", sx: Sexes{Male, Female}, wantErr: false},
		{name: "invalid", sx: Sexes{Sex(0), Sex(1), Sex(10)}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.sx.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Sexes.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSex_Validate(t *testing.T) {
	tests := []struct {
		name    string
		s       Sex
		wantErr bool
	}{
		{name: "valid", s: Sex(1), wantErr: false},
		{name: "invalid", s: Sex(10), wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Sex.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
