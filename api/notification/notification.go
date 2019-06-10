package notification

import (
	"encoding/json"
	"fmt"

	"github.com/identixone/identixone-go/core"
)

type Notifications struct {
	request core.Requester
}

func NewNotifications(request core.Requester) *Notifications {
	return &Notifications{request: request}
}

func (n *Notifications) List(query map[string]interface{}) (ListResponse, error) {
	var resp ListResponse
	data, err := n.request.Get("/v1/settings/notifications/", query)
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (n *Notifications) Get(id int) (Notification, error) {
	var resp Notification
	data, err := n.request.Get(fmt.Sprintf("/v1/settings/notifications/%d/", id), nil)
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (n *Notifications) Create(req CreateRequest) (Notification, error) {
	var resp Notification

	if err := req.Validate(); err != nil {
		return resp, err
	}

	in, err := json.Marshal(req)
	data, err := n.request.Post("/v1/settings/notifications/", in, "application/json")
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (n *Notifications) Update(req UpdateRequest) (Notification, error) {
	var resp Notification
	if err := req.Validate(); err != nil {
		return resp, err
	}
	in, err := json.Marshal(req)
	data, err := n.request.Patch(fmt.Sprintf("/v1/settings/notifications/%d/", req.ID), in, "application/json")
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (n *Notifications) Delete(id int) error {
	return n.request.Delete(fmt.Sprintf("/v1/settings/notifications/%d/", id), nil)
}
