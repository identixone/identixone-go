package person

import (
	"reflect"
	"testing"

	"github.com/identixone/identixone-go/core"
)

func TestNewPersons(t *testing.T) {
	c, _ := core.NewRequest()
	p := NewPersons(c)
	type args struct {
		client *core.Request
	}
	tests := []struct {
		name string
		args args
		want *Persons
	}{
		{name: "new", args: args{client: c}, want: p},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPersons(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPersons() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPersons_Create(t *testing.T) {
	type args struct {
		personCreate PersonaCreateRequest
	}
	tests := []struct {
		name    string
		p       *Persons
		args    args
		want    Person
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.Create(tt.args.personCreate)
			if (err != nil) != tt.wantErr {
				t.Errorf("Persons.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Persons.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPersons_Search(t *testing.T) {
	type args struct {
		search Search
	}
	tests := []struct {
		name    string
		p       *Persons
		args    args
		want    SearchResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.Search(tt.args.search)
			if (err != nil) != tt.wantErr {
				t.Errorf("Persons.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Persons.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPersons_Delete(t *testing.T) {
	c, _ := core.NewRequest()
	p := NewPersons(c)

	type args struct {
		idxid string
	}
	tests := []struct {
		name    string
		p       *Persons
		args    args
		wantErr bool
	}{
		{name: "not found", p: p, args: args{idxid: "p"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.Delete(tt.args.idxid); (err != nil) != tt.wantErr {
				t.Errorf("Persons.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPersons_ReinitImage(t *testing.T) {
	type args struct {
		reinitRequest ReinitImageRequest
	}
	tests := []struct {
		name    string
		p       *Persons
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.ReinitImage(tt.args.reinitRequest); (err != nil) != tt.wantErr {
				t.Errorf("Persons.ReinitImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPersons_ReinitId(t *testing.T) {
	type args struct {
		reinitIdRequest ReinitIdRequest
	}
	tests := []struct {
		name    string
		p       *Persons
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.ReinitId(tt.args.reinitIdRequest); (err != nil) != tt.wantErr {
				t.Errorf("Persons.ReinitId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
