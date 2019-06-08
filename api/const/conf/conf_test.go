package conf

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestConf_MarshalJSON(t *testing.T) {
	c := Conf("exact")
	c2 := Conf("exactor")

	b, err := json.Marshal(c)
	if err != nil {
		t.Error(err)
	}
	b2, err := json.Marshal(c2)
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		name    string
		c       *Conf
		want    []byte
		wantErr bool
	}{
		{name: "valid marshal", c: &c, want: b, wantErr: false},
		{name: "unknown conf", c: &c2, want: b2, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Conf.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conf.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConf_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	n := Conf("new")
	nm := Conf("nm")
	exact := Conf("exact")
	junk := Conf("junk")
	ha := Conf("ha")
	det := Conf("det")
	reinit := Conf("reinit")
	reiniter := Conf("reiniter")

	tests := []struct {
		name    string
		c       *Conf
		args    args
		wantErr bool
	}{
		{name: "conf new", c: &n, args: args{b: []byte(n)}, wantErr: false},
		{name: "conf nm", c: &nm, args: args{b: []byte(nm)}, wantErr: false},
		{name: "conf exact", c: &exact, args: args{b: []byte(exact)}, wantErr: false},
		{name: "conf junk", c: &junk, args: args{b: []byte(junk)}, wantErr: false},
		{name: "conf ha", c: &ha, args: args{b: []byte(ha)}, wantErr: false},
		{name: "conf det", c: &det, args: args{b: []byte(det)}, wantErr: false},
		{name: "conf reinit", c: &reinit, args: args{b: []byte(reinit)}, wantErr: false},
		{name: "conf reiniter", c: &reiniter, args: args{b: []byte(reiniter)}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Conf.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAll(t *testing.T) {
	tests := []struct {
		name string
		want []Conf
	}{
		{name: "all", want: []Conf{Nm, New, Exact, Ha, Junk, Det, Reinit}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}
