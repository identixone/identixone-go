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
