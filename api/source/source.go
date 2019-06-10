package source

import (
	"encoding/json"
	"fmt"

	"github.com/identixone/identixone-go/core"
)

type Sources struct {
	request core.Requester
}

func NewSource(request core.Requester) *Sources {
	return &Sources{request: request}
}

// List user sources
func (s *Sources) List(query map[string]interface{}) (ListResponse, error) {
	response := ListResponse{}

	data, err := s.request.Get("/v1/sources/", query)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// Create new source
func (s *Sources) Create(source Source) (Source, error) {
	var resp Source
	if err := source.Validate(); err != nil {
		return resp, err
	}
	in, err := json.Marshal(source)
	if err != nil {
		return resp, err
	}
	data, err := s.request.Post("/v1/sources/", in, "application/json")
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// Get source
func (s *Sources) Get(id int) (Source, error) {
	var resp Source
	data, err := s.request.Get(fmt.Sprintf("/v1/sources/%d/", id), nil)
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return resp, err
	}
	return resp, err
}

// Delete source
func (s *Sources) Delete(id int) error {
	return s.request.Delete(fmt.Sprintf("/v1/sources/%d/", id), nil)
}

// Update source
func (s *Sources) Update(req UpdateRequest) (Source, error) {
	var source Source
	if err := req.Validate(); err != nil {
		return source, err
	}

	id := req.ID
	req.ID = 0

	in, err := json.Marshal(req)
	if err != nil {
		return source, err
	}
	out, err := s.request.Patch(fmt.Sprintf("/v1/sources/%d/", id), in, "application/json")
	if err != nil {
		return source, err
	}
	err = json.Unmarshal(out, &source)
	if err != nil {
		return source, err
	}
	return source, nil
}
