package person

import (
	"encoding/json"
	"fmt"

	"github.com/identixone/identixone-go/core"
)

type Persons struct {
	request core.Requester
}

func NewPersons(request core.Requester) *Persons {
	return &Persons{request: request}
}

func (p *Persons) Create(req PersonaCreateRequest) (Person, error) {
	var person Person
	if err := req.Validate(); err != nil {
		return person, err
	}

	in, w, err := req.MultipartWriterData()
	if err != nil {
		return person, err
	}
	out, err := p.request.Post("/v1/persons/", in, w.FormDataContentType())
	if err != nil {
		return person, err
	}

	err = json.Unmarshal(out, &person)
	if err != nil {
		return person, err
	}

	return person, nil
}

func (p *Persons) Search(search Search) (SearchResult, error) {
	var result SearchResult
	if err := search.Validate(); err != nil {
		return result, err
	}
	in, w, err := search.MultipartWriterData()
	if err != nil {
		return result, err
	}
	data, err := p.request.Post("/v1/persons/search/", in, w.FormDataContentType())
	if err != nil {
		if idxErr, ok := err.(core.IdentixOneError); ok {
			if idxErr.Code == core.NotFound {
				data = *idxErr.Data
			} else {
				return result, err
			}
		} else {
			return result, err
		}
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (p *Persons) Delete(idxid string) error {
	return p.request.Delete(fmt.Sprintf("/v1/persons/%s/", idxid), nil)
}

func (p *Persons) ReinitImage(req ReinitImageRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	in, w, err := req.MultipartWriterData()
	if err != nil {
		return err
	}
	_, err = p.request.Post(fmt.Sprintf("/v1/persons/reinit/%s/", req.Idxid), in, w.FormDataContentType())
	if err != nil {
		return err
	}
	return nil
}

func (p *Persons) ReinitId(req ReinitIdRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	in, err := json.Marshal(req)
	if err != nil {
		return err
	}
	_, err = p.request.Post("/v1/persons/reinit/", in, "application/json")
	if err != nil {
		return err
	}
	return nil
}
