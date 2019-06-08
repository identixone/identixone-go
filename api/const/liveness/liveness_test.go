package liveness

import (
	"reflect"
	"testing"
)

func TestLiveness_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}

	passed := Liveness("passed")
	failed := Liveness("failed")
	undetermined := Liveness("undetermined")
	u := Liveness("u")

	tests := []struct {
		name    string
		l       *Liveness
		args    args
		wantErr bool
	}{
		{name: "passed liveness", l: &passed, args: args{b: []byte(passed)}, wantErr: false},
		{name: "failed liveness", l: &failed, args: args{b: []byte(failed)}, wantErr: false},
		{name: "undetermined liveness", l: &undetermined, args: args{b: []byte(undetermined)}, wantErr: false},
		{name: "u liveness", l: &u, args: args{b: []byte(u)}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.l.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Liveness.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLiveness_IsLive(t *testing.T) {
	tests := []struct {
		name string
		l    Liveness
		want bool
	}{
		{name: "live", l: Passed, want: true},
		{name: "not live", l: Failed, want: false},
		{name: "not live 2", l: Undetermined, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.IsLive(); got != tt.want {
				t.Errorf("Liveness.IsLive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAll(t *testing.T) {
	tests := []struct {
		name string
		want []Liveness
	}{
		{name: "all", want: []Liveness{Passed, Failed, Undetermined}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}
