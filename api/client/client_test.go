package client

import (
	"os"
	"reflect"
	"testing"

	"github.com/identixone/identixone-go/api/entry"
	"github.com/identixone/identixone-go/api/notification"
	"github.com/identixone/identixone-go/api/person"
	"github.com/identixone/identixone-go/api/source"
	"github.com/identixone/identixone-go/api/users"
	"github.com/identixone/identixone-go/api/utility"
)

var token = os.Getenv("IDENTIXONE_TOKEN")
var client = NewClientWithToken(token)

func TestNewClient(t *testing.T) {
	_, err := NewClient()
	if err != nil {
		t.Error(err)
	}
	_ = os.Setenv("IDENTIXONE_TOKEN", "")
	_, err = NewClient()
	if err == nil {
		t.Errorf("fail test without token")
	}
	_ = os.Setenv("IDENTIXONE_TOKEN", token)
}

func TestNewClientWithToken(t *testing.T) {
	c := NewClientWithToken(token)
	type args struct {
		token string
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{name: "valid", args: args{token: token}, want: c},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClientWithToken(tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClientWithToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Entries(t *testing.T) {
	type fields struct {
		users         *users.Users
		entries       *entry.Entries
		sources       *source.Sources
		persons       *person.Persons
		utility       *utility.Utility
		notifications *notification.Notifications
	}
	tests := []struct {
		name   string
		fields fields
		want   *entry.Entries
	}{
		{name: "entries", fields: fields{entries: client.entries}, want: client.Entries()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				users:         tt.fields.users,
				entries:       tt.fields.entries,
				sources:       tt.fields.sources,
				persons:       tt.fields.persons,
				utility:       tt.fields.utility,
				notifications: tt.fields.notifications,
			}
			if got := c.Entries(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Entries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Sources(t *testing.T) {
	type fields struct {
		users         *users.Users
		entries       *entry.Entries
		sources       *source.Sources
		persons       *person.Persons
		utility       *utility.Utility
		notifications *notification.Notifications
	}
	tests := []struct {
		name   string
		fields fields
		want   *source.Sources
	}{
		{name: "sources", fields: fields{sources: client.sources}, want: client.Sources()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				users:         tt.fields.users,
				entries:       tt.fields.entries,
				sources:       tt.fields.sources,
				persons:       tt.fields.persons,
				utility:       tt.fields.utility,
				notifications: tt.fields.notifications,
			}
			if got := c.Sources(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Sources() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Users(t *testing.T) {
	type fields struct {
		users         *users.Users
		entries       *entry.Entries
		sources       *source.Sources
		persons       *person.Persons
		utility       *utility.Utility
		notifications *notification.Notifications
	}
	tests := []struct {
		name   string
		fields fields
		want   *users.Users
	}{
		{name: "users", fields: fields{users: client.users}, want: client.Users()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				users:         tt.fields.users,
				entries:       tt.fields.entries,
				sources:       tt.fields.sources,
				persons:       tt.fields.persons,
				utility:       tt.fields.utility,
				notifications: tt.fields.notifications,
			}
			if got := c.Users(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Users() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Persons(t *testing.T) {
	type fields struct {
		users         *users.Users
		entries       *entry.Entries
		sources       *source.Sources
		persons       *person.Persons
		utility       *utility.Utility
		notifications *notification.Notifications
	}
	tests := []struct {
		name   string
		fields fields
		want   *person.Persons
	}{
		{name: "persons", fields: fields{persons: client.persons}, want: client.Persons()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				users:         tt.fields.users,
				entries:       tt.fields.entries,
				sources:       tt.fields.sources,
				persons:       tt.fields.persons,
				utility:       tt.fields.utility,
				notifications: tt.fields.notifications,
			}
			if got := c.Persons(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Persons() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Notifications(t *testing.T) {
	type fields struct {
		users         *users.Users
		entries       *entry.Entries
		sources       *source.Sources
		persons       *person.Persons
		utility       *utility.Utility
		notifications *notification.Notifications
	}
	tests := []struct {
		name   string
		fields fields
		want   *notification.Notifications
	}{
		{name: "notifications", fields: fields{notifications: client.notifications}, want: client.Notifications()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				users:         tt.fields.users,
				entries:       tt.fields.entries,
				sources:       tt.fields.sources,
				persons:       tt.fields.persons,
				utility:       tt.fields.utility,
				notifications: tt.fields.notifications,
			}
			if got := c.Notifications(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Notifications() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Utility(t *testing.T) {
	type fields struct {
		users         *users.Users
		entries       *entry.Entries
		sources       *source.Sources
		persons       *person.Persons
		utility       *utility.Utility
		notifications *notification.Notifications
	}
	tests := []struct {
		name   string
		fields fields
		want   *utility.Utility
	}{
		{name: "utility", fields: fields{utility: client.utility}, want: client.Utility()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				users:         tt.fields.users,
				entries:       tt.fields.entries,
				sources:       tt.fields.sources,
				persons:       tt.fields.persons,
				utility:       tt.fields.utility,
				notifications: tt.fields.notifications,
			}
			if got := c.Utility(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Utility() = %v, want %v", got, tt.want)
			}
		})
	}
}
