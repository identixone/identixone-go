package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/identixone/identixone-go/utils"
)

type Requester interface {
	Get(string, map[string]interface{}) ([]byte, error)
	Post(string, []byte, string) ([]byte, error)
	Patch(string, []byte, string) ([]byte, error)
	Delete(string, map[string]interface{}) error
}

type Request struct {
	token   string
	headers map[string]string
}

func NewRequest() (*Request, error) {
	token := os.Getenv("IDENTIXONE_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("IDENTIXONE_TOKEN environment not found")
	}
	headers := makeHeaders(token)
	return &Request{token: token, headers: headers}, nil
}

func (c *Request) SetToken(token string) {
	c.token = token
	c.headers = makeHeaders(token)
}

func (c *Request) Get(path string, query map[string]interface{}) ([]byte, error) {
	return c.request("GET", path, query, nil, "")
}

func (c *Request) Post(path string, data []byte, contentType string) ([]byte, error) {
	return c.request("POST", path, nil, data, contentType)
}

func (c *Request) Patch(path string, data []byte, contentType string) ([]byte, error) {
	return c.request("PATCH", path, nil, data, contentType)
}

func (c *Request) Delete(path string, query map[string]interface{}) error {
	_, err := c.request("DELETE", path, query, nil, "")
	return err
}

func (c *Request) request(method, path string, query map[string]interface{}, body []byte, contentType string) ([]byte, error) {
	url, err := makeUrl(path, query)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(body)
	req, err := http.NewRequest(method, url, buffer)

	if err != nil {
		return nil, err
	}

	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	u := uuid.New()
	req.Header.Set("Request-ID", u.String())

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, NewError(Internal, err, nil)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, NewError(Internal, err, nil)
	}

	switch resp.StatusCode {
	case http.StatusOK:
		break
	case http.StatusCreated:
		break
	case http.StatusNoContent:
		break
	case http.StatusNotFound:
		idxError := NewError(NotFound, nil, data)
		utils.Debug().
			Int("status", resp.StatusCode).
			Str("method", req.Method).
			Str("url", req.URL.String()).
			Msg("response")
		return nil, idxError
	case http.StatusBadRequest:
		var detail map[string]interface{}
		_ = json.Unmarshal(data, &detail)
		det, _ := utils.GetPretty(detail)
		err := fmt.Errorf("Request to: %s Status: %s\nDetail: %s", req.URL, resp.Status, string(det))
		utils.Debug().
			Int("status", resp.StatusCode).
			Str("method", req.Method).
			Str("url", req.URL.String()).
			Interface("body", string(data)).
			Msg("response")
		return nil, NewError(BadRequest, err, det)
	default:
		err := fmt.Errorf("Request to: %s Status: %s\n", req.URL, resp.Status)
		utils.Debug().
			Int("status", resp.StatusCode).
			Str("method", req.Method).
			Str("url", req.URL.String()).
			Interface("body", string(data)).
			Msg("response")
		return nil, NewError(ApiInternal, err, data)
	}

	utils.Debug().
		Int("status", resp.StatusCode).
		Str("method", req.Method).
		Msg(req.URL.String())

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Error().Err(err).Msg("close body")
		}
	}()
	return data, nil
}
