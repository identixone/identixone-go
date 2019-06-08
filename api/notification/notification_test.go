package notification

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/identixone/identixone-go/api/const/transport"
	"github.com/identixone/identixone-go/core"
)

type MockRequesterList struct{}

func (MockRequesterList) Get(string, map[string]interface{}) ([]byte, error) {
	return []byte(`[{"hello": "world"}]`), nil
}

func (MockRequesterList) Post(string, []byte, string) ([]byte, error) {
	return []byte(`{"hello": "world"}`), nil
}

func (MockRequesterList) Patch(string, []byte, string) ([]byte, error) {
	return []byte(`{"hello": "world"}`), nil
}

func (MockRequesterList) Delete(string, map[string]interface{}) error {
	return nil
}

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

func TestNotifications_List(t *testing.T) {
	type args struct {
		query map[string]interface{}
	}
	tests := []struct {
		name    string
		n       *Notifications
		args    args
		want    []Notification
		wantErr bool
	}{
		{name: "list", n: NewNotifications(MockRequesterList{}), args: args{query: map[string]interface{}{}}, wantErr: false, want: []Notification{{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.List(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("Notifications.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Notifications.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotifications_Get(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		n       *Notifications
		args    args
		want    Notification
		wantErr bool
	}{
		{name: "get", n: NewNotifications(MockRequester{}), args: args{id: 1}, want: Notification{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Notifications.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Notifications.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotifications_Create(t *testing.T) {
	type args struct {
		req CreateRequest
	}
	tests := []struct {
		name    string
		n       *Notifications
		args    args
		want    Notification
		wantErr bool
	}{
		{name: "create", n: NewNotifications(MockRequester{}), args: args{req: CreateRequest{}}, want: Notification{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.Create(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Notifications.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Notifications.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotifications_Update(t *testing.T) {
	r := MockRequester{}

	type args struct {
		id  int
		req map[string]interface{}
	}
	tests := []struct {
		name    string
		n       *Notifications
		args    args
		want    Notification
		wantErr bool
	}{
		{name: "valid", n: NewNotifications(r), args: args{id: 1, req: map[string]interface{}{}}, want: Notification{}, wantErr: false},
		{name: "oops", n: NewNotifications(MockRequesterErr{}), args: args{id: 2, req: map[string]interface{}{"transport": transport.Webhook}}, want: Notification{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.Update(tt.args.id, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Notifications.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Notifications.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotifications_Delete(t *testing.T) {
	r := MockRequester{}
	r2 := MockRequesterErr{}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		n       *Notifications
		args    args
		wantErr bool
	}{
		{name: "not found", n: NewNotifications(r2), args: args{0}, wantErr: true},
		{name: "valid", n: NewNotifications(r), args: args{1}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Notifications.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewNotifications(t *testing.T) {
	r := MockRequester{}
	n := NewNotifications(r)
	type args struct {
		request core.Requester
	}
	tests := []struct {
		name string
		args args
		want *Notifications
	}{
		{name: "new", args: args{request: r}, want: n},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNotifications(tt.args.request); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNotifications() = %v, want %v", got, tt.want)
			}
		})
	}
}
