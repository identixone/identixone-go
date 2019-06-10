package utility

import (
	"encoding/json"

	"github.com/identixone/identixone-go/core"
	"github.com/identixone/identixone-go/utils"
)

type Utility struct {
	request core.Requester
}

func NewUtility(request core.Requester) *Utility {
	return &Utility{request: request}
}

func (u *Utility) Asm(req AsmRequest) (AsmResponse, error) {
	var resp AsmResponse

	if err := req.Validate(); err != nil {
		return resp, err
	}
	in, w, err := req.MultipartWriterData()
	if err != nil {
		return resp, err
	}
	data, err := u.request.Post("/v1/utility/asm/", in, w.FormDataContentType())
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (u *Utility) Liveness(req LivenessRequest) (LivenessResponse, error) {
	var resp LivenessResponse
	if err := req.Validate(); err != nil {
		return resp, err
	}
	in, w, err := req.MultipartWriterData()
	if err != nil {
		return resp, err
	}
	data, err := u.request.Post("/v1/utility/liveness/", in, w.FormDataContentType())
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (u *Utility) Compare(req CompareRequest) (CompareResponse, error) {
	var resp CompareResponse
	if err := req.Validate(); err != nil {
		return resp, err
	}
	in, w, err := req.MultipartWriterData()
	if err != nil {
		return resp, err
	}
	data, err := u.request.Post("/v1/utility/compare/", in, w.FormDataContentType())
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
	if err := req.Validate(); err != nil {
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
