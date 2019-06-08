package person

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/identixone/identixone-go/core"
	"github.com/identixone/identixone-go/utils"
)

type Persons struct {
	request core.Requester
}

func NewPersons(request core.Requester) *Persons {
	return &Persons{request: request}
}

func (p *Persons) Create(req PersonaCreateRequest) (Person, error) {
	var person Person
	if err := req.IsValid(); err != nil {
		return person, err
	}

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	part, err := w.CreateFormFile("photo", req.PhotoName)
	reader := bytes.NewReader(req.PhotoData)
	_, err = io.Copy(part, reader)
	if err != nil {
		return person, err
	}

	perMap, err := utils.ToMap(req)
	if err != nil {
		return person, err
	}
	for k, v := range perMap {
		err := w.WriteField(k, fmt.Sprintf("%v", v))
		if err != nil {
			return person, err
		}
	}
	err = w.Close()
	if err != nil {
		return person, err
	}

	out, err := p.request.Post("/v1/persons/", buf.Bytes(), w.FormDataContentType())
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
	if err := search.IsValid(); err != nil {
		return result, err
	}
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	part, err := w.CreateFormFile("photo", search.PhotoName)
	reader := bytes.NewReader(search.PhotoData)
	_, err = io.Copy(part, reader)
	if err != nil {
		return result, err
	}

	err = w.WriteField("asm", fmt.Sprintf("%v", search.Asm))
	if err != nil {
		return result, err
	}
	err = w.WriteField("liveness", fmt.Sprintf("%v", search.Liveness))
	if err != nil {
		return result, err
	}

	err = w.Close()
	if err != nil {
		return result, err
	}
	data, err := p.request.Post("/v1/persons/search/", buf.Bytes(), w.FormDataContentType())
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
	if err := req.IsValid(); err != nil {
		return err
	}
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	part, err := w.CreateFormFile("photo", req.PhotoName)
	reader := bytes.NewReader(req.PhotoData)
	_, err = io.Copy(part, reader)
	if err != nil {
		return err
	}

	perMap, err := utils.ToMap(req)
	if err != nil {
		return err
	}
	for k, v := range perMap {
		err := w.WriteField(k, fmt.Sprintf("%v", v))
		if err != nil {
			return err
		}
	}
	err = w.Close()
	if err != nil {
		return err
	}
	_, err = p.request.Post(fmt.Sprintf("/v1/persons/reinit/%s/", req.Idxid), buf.Bytes(), w.FormDataContentType())
	if err != nil {
		return err
	}
	return nil
}

func (p *Persons) ReinitId(req ReinitIdRequest) error {
	if err := req.IsValid(); err != nil {
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
