package utility

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/identixone/identixone-go/api/common"
	"github.com/identixone/identixone-go/api/const/conf"
	"github.com/identixone/identixone-go/api/const/liveness"
	"github.com/identixone/identixone-go/api/const/moods"
	"github.com/identixone/identixone-go/api/const/sex"
)

type AsmRequest struct {
	common.Photo
}

func (asm *AsmRequest) MultipartWriterData() ([]byte, *multipart.Writer, error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	part, err := w.CreateFormFile("photo", asm.PhotoName)
	reader := bytes.NewReader(asm.PhotoData)
	_, err = io.Copy(part, reader)
	if err != nil {
		return nil, nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, nil, err
	}
	return buf.Bytes(), w, nil
}

func NewAsmRequest(photoPath string) (AsmRequest, error) {
	req := AsmRequest{}
	err := req.FromFile(photoPath)
	if err != nil {
		return req, err
	}
	return req, nil
}

func NewAsmRequestFromReader(reader io.Reader, name string) (*AsmRequest, error) {
	req := AsmRequest{}
	err := req.FromReader(reader, name)
	if err != nil {
		return &req, err
	}
	return &req, nil
}

type LivenessRequest struct {
	common.Photo
}

func NewLivenessRequest(photoPath string) (LivenessRequest, error) {
	req := LivenessRequest{}
	err := req.FromFile(photoPath)
	if err != nil {
		return req, err
	}
	return req, nil
}

func NewLivenessRequestFromReader(reader io.Reader, name string) (*LivenessRequest, error) {
	req := LivenessRequest{}
	err := req.FromReader(reader, name)
	if err != nil {
		return &req, err
	}
	return &req, nil
}

func (lr *LivenessRequest) MultipartWriterData() ([]byte, *multipart.Writer, error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	part, err := w.CreateFormFile("photo", lr.PhotoName)
	reader := bytes.NewReader(lr.PhotoData)
	_, err = io.Copy(part, reader)
	if err != nil {
		return nil, nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, nil, err
	}
	return buf.Bytes(), w, nil
}

type AsmResponse struct {
	Age  int        `json:"age"`
	Sex  sex.Sex    `json:"sex"`
	Mood moods.Mood `json:"mood"`
}

type LivenessResponse struct {
	Liveness liveness.Liveness `json:"liveness"`
}

type CompareRequest struct {
	Photo1         common.Photo `json:"-"`
	Photo2         common.Photo `json:"-"`
	LivenessPhoto1 bool         `json:"liveness_photo_1"`
	LivenessPhoto2 bool         `json:"liveness_photo_2"`
	Conf           *conf.Conf   `json:"conf,omitempty"`
}

func (compReq *CompareRequest) MultipartWriterData() ([]byte, *multipart.Writer, error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	part, err := w.CreateFormFile("photo1", compReq.Photo1.PhotoName)
	reader := bytes.NewReader(compReq.Photo1.PhotoData)

	_, err = io.Copy(part, reader)
	if err != nil {
		return nil, nil, err
	}

	part, err = w.CreateFormFile("photo2", compReq.Photo2.PhotoName)
	reader = bytes.NewReader(compReq.Photo2.PhotoData)

	_, err = io.Copy(part, reader)
	if err != nil {
		return nil, nil, err
	}

	err = w.WriteField("liveness_photo1", fmt.Sprintf("%v", compReq.LivenessPhoto1))
	if err != nil {
		return nil, nil, err
	}

	err = w.WriteField("liveness_photo2", fmt.Sprintf("%v", compReq.LivenessPhoto2))
	if err != nil {
		return nil, nil, err
	}

	if compReq.Conf != nil {
		err = w.WriteField("conf", string(*compReq.Conf))
		if err != nil {
			return nil, nil, err
		}
	}

	err = w.Close()
	if err != nil {
		return nil, nil, err
	}
	return buf.Bytes(), w, nil
}

func (compReq *CompareRequest) Validate() error {
	err := compReq.Photo1.Validate()
	if err != nil {
		return err
	}
	err = compReq.Photo2.Validate()
	if err != nil {
		return err
	}
	if compReq.Conf != nil {
		err = compReq.Conf.Validate()
	}
	return err
}

func NewCompareRequest(photo1Path, photo2Path string, c *conf.Conf) (CompareRequest, error) {
	req := CompareRequest{Conf: c}
	err := req.Photo1.FromFile(photo1Path)
	if err != nil {
		return req, err
	}
	err = req.Photo2.FromFile(photo2Path)
	if err != nil {
		return req, err
	}
	return req, nil
}

func NewCompareRequestFromReader(reader1, reader2 io.Reader, name1, name2 string, c *conf.Conf) (*CompareRequest, error) {
	req := CompareRequest{Conf: c}
	err := req.Photo1.FromReader(reader1, name1)
	if err != nil {
		return &req, err
	}
	err = req.Photo2.FromReader(reader2, name2)
	if err != nil {
		return &req, err
	}
	return &req, nil
}

type CompareResponse struct {
	Similar        bool              `json:"similar"`
	Conf           conf.Conf         `json:"conf"`
	LivenessPhoto1 liveness.Liveness `json:"liveness_photo1"`
	LivenessPhoto2 liveness.Liveness `json:"liveness_photo2"`
}

type CustomerRequest struct {
	Source string `json:"source"`
	Offset int    `json:"offset"`
}

func (custReq *CustomerRequest) Validate() error {
	if custReq.Source == "" {
		return fmt.Errorf("soure name is required")
	}
	return nil
}

type CustomerResponse struct {
	Idxid  string `json:"idxid"`
	Detail string `json:"detail"`
}
