package entry

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/identixone/identixone-go/core"
)

type MockRequester struct{}

func (MockRequester) Get(string, map[string]interface{}) ([]byte, error) {
	return []byte(`{"hello": "world"}`), nil
}

func (MockRequester) Post(string, []byte, string) ([]byte, error) {
	return []byte(`{"hello": "world"}`), nil
}

func (MockRequester) Patch(string, []byte, string) ([]byte, error) {
	return []byte(`{"hello": "world"}`), nil
}

func (MockRequester) Delete(string, map[string]interface{}) error {
	return nil
}

type MockRequesterErr struct{}

func (MockRequesterErr) Get(string, map[string]interface{}) ([]byte, error) {
	return nil, fmt.Errorf("oops")
}

func (MockRequesterErr) Post(string, []byte, string) ([]byte, error) {
	return nil, fmt.Errorf("oops")
}

func (MockRequesterErr) Patch(string, []byte, string) ([]byte, error) {
	return nil, fmt.Errorf("oops")
}

func (MockRequesterErr) Delete(string, map[string]interface{}) error {
	return fmt.Errorf("oops")
}

func TestNewEntries(t *testing.T) {

	c, _ := core.NewRequest()
	e := NewEntries(c)
	type args struct {
		client *core.Request
	}
	tests := []struct {
		name string
		args args
		want *Entries
	}{
		{name: "valid", args: args{client: c}, want: e},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEntries(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEntries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntries_List(t *testing.T) {

	e := NewEntries(MockRequester{})
	type args struct {
		query map[string]interface{}
	}
	tests := []struct {
		name    string
		es      *Entries
		args    args
		want    ListResponse
		wantErr bool
	}{
		{name: "list", es: e, args: args{query: nil}, want: ListResponse{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.es.List(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entries.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entries.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntries_Delete(t *testing.T) {
	e := NewEntries(MockRequester{})
	e2 := NewEntries(MockRequesterErr{})
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		es      *Entries
		args    args
		wantErr bool
	}{
		{name: "not found", es: e2, args: args{id: 0}, wantErr: true},
		{name: "valid", es: e, args: args{id: 1}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.es.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Entries.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEntries_StatsIdxid(t *testing.T) {
	e := NewEntries(MockRequester{})
	type args struct {
		idxid string
	}
	tests := []struct {
		name    string
		es      *Entries
		args    args
		want    StatsIdxid
		wantErr bool
	}{
		{name: "stats idxid", es: e, args: args{idxid: ""}, want: StatsIdxid{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.es.StatsIdxid(tt.args.idxid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entries.StatsIdxid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entries.StatsIdxid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntries_StatsSources(t *testing.T) {
	e := NewEntries(MockRequester{})
	type args struct {
		query map[string]interface{}
	}
	tests := []struct {
		name    string
		es      *Entries
		args    args
		want    StatSourceResponse
		wantErr bool
	}{
		{name: "stats sources", es: e, args: args{query: nil}, want: StatSourceResponse{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.es.StatsSources(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entries.StatsSources() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entries.StatsSources() = %v, want %v", got, tt.want)
			}
		})
	}
}
