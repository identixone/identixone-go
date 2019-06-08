package utility

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/identixone/identixone-go/api/common"
	"github.com/identixone/identixone-go/core"
	"github.com/identixone/identixone-go/utils"
)

type Utility struct {
	request *core.Request
}

func NewUtility(request *core.Request) *Utility {
	return &Utility{request: request}
}

func (u *Utility) Asm(photo common.Photo) (AsmResponse, error) {
	var resp AsmResponse

	if err := photo.IsValid(); err != nil {
		return resp, err
	}

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	part, err := w.CreateFormFile("photo", photo.PhotoName)
	reader := bytes.NewReader(photo.PhotoData)
	_, err = io.Copy(part, reader)
	if err != nil {
		return resp, err
	}
	err = w.Close()
	if err != nil {
		return resp, err
	}
	data, err := u.request.Post("/v1/utility/asm/", buf.Bytes(), w.FormDataContentType())
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (u *Utility) Liveness(photo common.Photo) (LivenessResponse, error) {
	var resp LivenessResponse
	if err := photo.IsValid(); err != nil {
		return resp, err
	}
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	part, err := w.CreateFormFile("photo", photo.PhotoName)
	reader := bytes.NewReader(photo.PhotoData)
	_, err = io.Copy(part, reader)
	if err != nil {
		return resp, err
	}
	err = w.Close()
	if err != nil {
		return resp, err
	}
	data, err := u.request.Post("/v1/utility/liveness/", buf.Bytes(), w.FormDataContentType())
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (u *Utility) Compare(req CompareRequest) (LivenessResponse, error) {
	var resp LivenessResponse
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	part, err := w.CreateFormFile("photo1", req.Photo1.PhotoName)
	reader := bytes.NewReader(req.Photo1.PhotoData)

	_, err = io.Copy(part, reader)
	if err != nil {
		return resp, err
	}

	part, err = w.CreateFormFile("photo2", req.Photo2.PhotoName)
	reader = bytes.NewReader(req.Photo2.PhotoData)

	_, err = io.Copy(part, reader)
	if err != nil {
		return resp, err
	}

	err = w.WriteField("liveness_photo1", fmt.Sprintf("%v", req.LivenessPhoto1))
	if err != nil {
		return resp, err
	}

	err = w.WriteField("liveness_photo2", fmt.Sprintf("%v", req.LivenessPhoto2))
	if err != nil {
		return resp, err
	}

	err = w.WriteField("conf", string(req.Conf))
	if err != nil {
		return resp, err
	}

	err = w.Close()
	if err != nil {
		return resp, err
	}
	data, err := u.request.Post("/v1/utility/compare/", buf.Bytes(), w.FormDataContentType())
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (u *Utility) Customer(req CustomerRequest) (CustomerResponse, error) {
	var resp CustomerResponse
	if err := req.IsValid(); err != nil {
		return resp, err
	}
	if req.Offset == 0 {
		req.Offset = 10
	}

	in, err := utils.ToMap(req)
	if err != nil {
		return resp, err
	}
	data, err := u.request.Get("/v1/utility/customer/", in)
	if err != nil {
		if idxErr, ok := err.(core.IdentixOneError); ok {
			data = *idxErr.Data
		} else {
			return resp, nil
		}
	}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
