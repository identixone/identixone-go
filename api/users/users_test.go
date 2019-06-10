package users

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"

	mock_core "github.com/identixone/identixone-go/core/mock"
)

func TestNewUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_core.NewMockRequester(ctrl)
	u := NewUsers(m)
	if u == nil {
		t.Fatal("new users fail")
	}
}

func TestUsers_Me(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_core.NewMockRequester(ctrl)
	u := NewUsers(m)

	m.EXPECT().
		Get("/v1/users/me/", nil).
		Return([]byte(`{"username": "username", "group": "group"}`), nil).
		AnyTimes()

	resp, err := u.Me()
	if err != nil {
		t.Fatal(err)
	}

	if resp.Username != "username" || resp.Group != "group" {
		t.Fatal("expect fail")
	}
}

func TestUsers_ListTokens(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_core.NewMockRequester(ctrl)
	u := NewUsers(m)
	per := true

	m.EXPECT().
		Get("/v1/users/tokens/", nil).
		Return([]byte(`[{"id": 1}]`), nil).AnyTimes()

	m.EXPECT().
		Get("/v1/users/tokens/", map[string]interface{}{"permanent": &per}).
		Return([]byte(`[{"id": 2}]`), nil).AnyTimes()

	resp, err := u.ListTokens(nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(resp) != 1 || resp[0].ID != 1 {
		t.Fatal("expect fail")
	}
	resp, err = u.ListTokens(&per)
	if err != nil {
		t.Fatal(err)
	}

	if len(resp) != 1 || resp[0].ID != 2 {
		t.Fatal("expect fail")
	}
}

func TestUsers_GetToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_core.NewMockRequester(ctrl)
	u := NewUsers(m)

	m.EXPECT().
		Get("/v1/users/tokens/secret/", nil).
		Return([]byte(`{"id": 3}`), nil).AnyTimes()
	m.EXPECT().
		Get("/v1/users/tokens/3/", nil).
		Return([]byte(`{"id": 3}`), nil).AnyTimes()

	reps, err := u.GetToken("secret")
	if err != nil {
		t.Fatal(err)
	}

	if reps.ID != 3 {
		t.Fatal("expect fail")
	}

	reps, err = u.GetToken(3)
	if err != nil {
		t.Fatal(err)
	}

	if reps.ID != 3 {
		t.Fatal("expect fail")
	}
}

func TestUsers_UpdateToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_core.NewMockRequester(ctrl)
	u := NewUsers(m)

	in, err := json.Marshal(map[string]bool{"is_active": true})
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().
		Patch("/v1/users/tokens/secret/", in, "application/json").
		Return([]byte(`{"id": 3}`), nil).AnyTimes()
	m.EXPECT().
		Patch("/v1/users/tokens/3/", in, "application/json").
		Return([]byte(`{"id": 3}`), nil).AnyTimes()

	reps, err := u.UpdateToken("secret", true)
	if err != nil {
		t.Fatal(err)
	}

	if reps.ID != 3 {
		t.Fatal("expect fail")
	}

	reps, err = u.UpdateToken(3, true)
	if err != nil {
		t.Fatal(err)
	}

	if reps.ID != 3 {
		t.Fatal("expect fail")
	}

}

func TestUsers_DeleteToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_core.NewMockRequester(ctrl)
	u := NewUsers(m)

	m.EXPECT().
		Delete("/v1/users/tokens/secret/", nil).
		Return(nil).AnyTimes()
	m.EXPECT().
		Delete("/v1/users/tokens/3/", nil).
		Return(nil).AnyTimes()

	err := u.DeleteToken("secret")
	if err != nil {
		t.Fatal(err)
	}

	err = u.DeleteToken(3)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUsers_DeleteAllToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_core.NewMockRequester(ctrl)
	u := NewUsers(m)
	per := true

	in := map[string]interface{}{"permanent": &per}

	m.EXPECT().
		Delete("/v1/users/tokens/", in).
		Return(nil).AnyTimes()

	err := u.DeleteAllToken(&per)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUsers_ChangePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_core.NewMockRequester(ctrl)
	u := NewUsers(m)

	req := ChangePasswordRequest{Password1: "somepassword", Password2: "somepassword"}
	req2 := ChangePasswordRequest{Password1: "somepassword", Password2: "password"}
	in, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().
		Post("/v1/users/password/change/", in, "application/json").
		Return([]byte(`{"success": true}`), nil).AnyTimes()

	resp, err := u.ChangePassword(req)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Success {
		t.Fatal("expect fail")
	}

	_, err = u.ChangePassword(req2)
	if err == nil {
		t.Fatal(err)
	}
}

func TestUsers_CreateToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_core.NewMockRequester(ctrl)
	u := NewUsers(m)

	req := CreateTokenRequest{Password: "password", Username: "username", Permanent: true}
	req2 := CreateTokenRequest{Password: "password", Username: "username"}
	in, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	in2, err := json.Marshal(req2)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().
		Post("/v1/login/permanent", in, "application/json").
		Return([]byte(`{"user": {"username": "username"}, "token": {"key": "key"}}`), nil).
		AnyTimes()

	m.EXPECT().
		Post("/v1/login/", in2, "application/json").
		Return([]byte(`{"user": {"username": "username"}, "token": {"key": "key"}}`), nil).
		AnyTimes()

	resp, err := u.CreateToken(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.User.Username != "username" {
		t.Fatal("expect fail")
	}

	resp, err = u.CreateToken(req2)
	if err != nil {
		t.Fatal(err)
	}

	if resp.User.Username != "username" {
		t.Fatal("expect fail")
	}

}
