package client

import (
	"github.com/identixone/identixone-go/api/entry"
	"github.com/identixone/identixone-go/api/notification"
	"github.com/identixone/identixone-go/api/person"
	"github.com/identixone/identixone-go/api/source"
	"github.com/identixone/identixone-go/api/users"
	"github.com/identixone/identixone-go/api/utility"
	"github.com/identixone/identixone-go/core"
)

type Client struct {
	users         *users.Users
	entries       *entry.Entries
	sources       *source.Sources
	persons       *person.Persons
	utility       *utility.Utility
	notifications *notification.Notifications
}

func (c *Client) Entries() *entry.Entries {
	return c.entries
}

func (c *Client) Sources() *source.Sources {
	return c.sources
}

func (c *Client) Users() *users.Users {
	return c.users
}

func (c *Client) Persons() *person.Persons {
	return c.persons
}

func (c *Client) Notifications() *notification.Notifications {
	return c.notifications
}

func (c *Client) Utility() *utility.Utility {
	return c.utility
}

func NewClient() (*Client, error) {
	requester, err := core.NewRequest()
	if err != nil {
		return nil, err
	}
	client := &Client{
		users:         users.NewUsers(requester),
		entries:       entry.NewEntries(requester),
		sources:       source.NewSource(requester),
		persons:       person.NewPersons(requester),
		utility:       utility.NewUtility(requester),
		notifications: notification.NewNotifications(requester),
	}
	return client, nil
}

func NewClientWithToken(token string) *Client {
	requester := &core.Request{}
	requester.SetToken(token)
	client := &Client{
		users:         users.NewUsers(requester),
		entries:       entry.NewEntries(requester),
		sources:       source.NewSource(requester),
		persons:       person.NewPersons(requester),
		utility:       utility.NewUtility(requester),
		notifications: notification.NewNotifications(requester),
	}
	return client
}
