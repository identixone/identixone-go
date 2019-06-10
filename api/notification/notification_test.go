package notification

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/identixone/identixone-go/api/const/method"
	mock_core "github.com/identixone/identixone-go/core/mock"
)

func TestNotifications_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_core.NewMockRequester(ctrl)
	m.EXPECT().
		Get(gomock.Eq("/v1/settings/notifications/"), nil).
		Return([]byte(`{"count": 1, "results": [{"id": 1}]}`), nil).
		AnyTimes()

	query := map[string]interface{}{"offset": 10}
	m.EXPECT().
		Get(gomock.Eq("/v1/settings/notifications/"), query).
		Return([]byte(`{"count": 0, "results": []}`), nil).
		AnyTimes()

	n := NewNotifications(m)

	resp, err := n.List(nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Count != 1 {
		t.Fatal("expect fail")
	}

	if resp.Notifications[0].ID != 1 {
		t.Fatal("expect fail")
	}

	resp, err = n.List(query)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Count != 0 {
		t.Fatal("expect fail")
	}

	if len(resp.Notifications) != 0 {
		t.Fatal("expect fail")
	}
}

func TestNotifications_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_core.NewMockRequester(ctrl)
	m.EXPECT().
		Get(gomock.Eq("/v1/settings/notifications/1/"), nil).
		Return([]byte(`{"id": 1}`), nil).
		AnyTimes()

	m.EXPECT().
		Get(gomock.Eq("/v1/settings/notifications/0/"), nil).
		Return([]byte(``), fmt.Errorf("not found")).
		AnyTimes()

	n := NewNotifications(m)

	resp, err := n.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	if resp.ID != 1 {
		t.Fatal("expect fail")
	}

	resp, err = n.Get(0)
	if err == nil {
		t.Fatal(err)
	}
}

func TestNotifications_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := CreateRequest{Name: "notify", HTTPMethod: method.Point(method.Post), DestinationURL: "some/path"}
	req2 := CreateRequest{}
	data, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	m := mock_core.NewMockRequester(ctrl)
	m.EXPECT().
		Post(gomock.Eq("/v1/settings/notifications/"), data, "application/json").
		Return([]byte(`{"id": 1}`), nil).
		AnyTimes()

	n := NewNotifications(m)
	resp, err := n.Create(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.ID != 1 {
		t.Fatal("expect fail")
	}

	_, err = n.Create(req2)
	if err == nil {
		t.Fatal("expect validation error")
	}
}

func TestNotifications_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := UpdateRequest{ID: 1, Name: "notify", HTTPMethod: method.Point(method.Post), DestinationURL: "some/path"}
	req2 := UpdateRequest{}

	data, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	m := mock_core.NewMockRequester(ctrl)
	m.EXPECT().
		Patch(gomock.Eq("/v1/settings/notifications/1/"), data, "application/json").
		Return([]byte(`{"id": 1}`), nil).
		AnyTimes()

	n := NewNotifications(m)
	resp, err := n.Update(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.ID != 1 {
		t.Fatal("expect fail")
	}

	_, err = n.Update(req2)
	if err == nil {
		t.Fatal("expect validation error")
	}
}

func TestNotifications_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_core.NewMockRequester(ctrl)
	m.EXPECT().
		Delete(gomock.Eq("/v1/settings/notifications/1/"), nil).
		Return(nil).
		AnyTimes()

	m.EXPECT().
		Delete(gomock.Eq("/v1/settings/notifications/0/"), nil).
		Return(fmt.Errorf("not found")).
		AnyTimes()

	n := NewNotifications(m)
	err := n.Delete(1)
	if err != nil {
		t.Fatal(err)
	}

	err = n.Delete(0)
	if err == nil {
		t.Fatal("expect validation error")
	}
}

func TestNewNotifications(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_core.NewMockRequester(ctrl)
	_ = NewNotifications(m)

}
