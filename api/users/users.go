package users

import (
	"encoding/json"
	"fmt"

	"github.com/identixone/identixone-go/core"
)

type Users struct {
	request core.Requester
}

func NewUsers(request core.Requester) *Users {
	return &Users{request: request}
}

// Returns a username to which the token, specified in the request, belongs, as well as its group name.
func (u *Users) Me() (User, error) {
	var me User
	data, err := u.request.Get("/v1/users/me/", nil)
	if err != nil {
		return me, err
	}
	err = json.Unmarshal(data, &me)
	if err != nil {
		return me, err
	}
	return me, nil
}

// A list of all tokens, created by the user, is returned.
// Parameter: permanent (default nil)
//	will be displayed:
// 	false – only temporary tokens,
// 	true – only permanent tokens, no parameter set – all tokens
func (u *Users) ListTokens(permanent *bool) ([]Token, error) {
	var tokens []Token
	var query map[string]interface{}
	if permanent != nil {
		query = map[string]interface{}{"permanent": permanent}
	}
	data, err := u.request.Get("/v1/users/tokens/", query)
	if err != nil {
		return tokens, err
	}
	err = json.Unmarshal(data, &tokens)
	if err != nil {
		return tokens, err
	}
	return tokens, nil
}

// Returns information about a token with id or key, specified in url.
func (u *Users) GetToken(idOrToken interface{}) (Token, error) {
	var token Token
	data, err := u.request.Get(fmt.Sprintf("/v1/users/tokens/%v/", idOrToken), nil)
	if err != nil {
		return token, err
	}
	err = json.Unmarshal(data, &token)
	if err != nil {
		return token, err
	}
	return token, nil
}

// Deactivates / activates a token with id or key, specified in url.
func (u *Users) UpdateToken(idOrToken interface{}, isActive bool) (Token, error) {
	var token Token
	in, err := json.Marshal(map[string]bool{"is_active": isActive})
	if err != nil {
		return token, err
	}
	data, err := u.request.Patch(fmt.Sprintf("/v1/users/tokens/%v/", idOrToken), in, "application/json")
	if err != nil {
		return token, err
	}
	err = json.Unmarshal(data, &token)
	if err != nil {
		return token, err
	}
	return token, nil
}

// A token with id or key for deleting.
func (u *Users) DeleteToken(idOrToken interface{}) error {
	err := u.request.Delete(fmt.Sprintf("/v1/users/tokens/%v/", idOrToken), nil)
	if err != nil {
		return err
	}
	return nil
}

// All tokens of the given user are deleted.
func (u *Users) DeleteAllToken(permanent *bool) error {
	query := map[string]interface{}{}
	if permanent != nil {
		query["permanent"] = permanent
	}
	err := u.request.Delete("/v1/users/tokens/", query)
	if err != nil {
		return err
	}
	return nil
}

// Allows to change the user’s current password.
func (u *Users) ChangePassword(req ChangePasswordRequest) (ChangePasswordResponse, error) {
	var resp ChangePasswordResponse
	if err := req.Validate(); err != nil {
		return resp, err
	}
	in, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}
	data, err := u.request.Post("/v1/users/password/change/", in, "application/json")
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (u *Users) CreateToken(req CreateTokenRequest) (CreateTokenResponse, error) {
	var resp CreateTokenResponse
	url := "/v1/login/"
	if req.Permanent {
		url = "/v1/login/permanent"
	}

	in, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	data, err := u.request.Post(url, in, "application/json")
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
