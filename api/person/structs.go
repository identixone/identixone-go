package person

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
	"github.com/identixone/identixone-go/utils"
)

type Person struct {
	Conf     conf.Conf         `json:"conf"`
	Idxid    string            `json:"idxid"`
	Age      *int              `json:"age"`
	Sex      *sex.Sex          `json:"sex"`
	Mood     moods.Mood        `json:"mood"`
	Liveness liveness.Liveness `json:"liveness"`
}

type PersonaCreateRequest struct {
	common.Photo
	Source             string `json:"source"`
	Facesize           int    `json:"facesize"`
	CreateOnHa         bool   `json:"create_on_ha"`
	CreateOnJunk       bool   `json:"create_on_junk"`
	Asm                bool   `json:"asm"`
	Liveness           bool   `json:"liveness"`
	CreateLivenessOnly bool   `json:"create_liveness_only"`
}

func (pc *PersonaCreateRequest) MultipartWriterData() ([]byte, *multipart.Writer, error) {

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	part, err := w.CreateFormFile("photo", pc.PhotoName)
	reader := bytes.NewReader(pc.PhotoData)
	_, err = io.Copy(part, reader)
	if err != nil {
		return nil, nil, err
	}

	perMap, err := utils.ToMap(pc)
	if err != nil {
		return nil, nil, err
	}
	for k, v := range perMap {
		err := w.WriteField(k, fmt.Sprintf("%v", v))
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

func (pc *PersonaCreateRequest) Validate() error {
	if pc.PhotoData == nil && pc.PhotoName == "" {
		return fmt.Errorf("photo data or name is empty")
	}
	if pc.Source == "" {
		return fmt.Errorf("source is required")
	}
	return nil
}

func (pc *PersonaCreateRequest) SetSource(name string) *PersonaCreateRequest {
	pc.Source = name
	return pc
}
func (pc *PersonaCreateRequest) SetFacesize(val int) *PersonaCreateRequest {
	pc.Facesize = val
	return pc
}
func (pc *PersonaCreateRequest) SetCreateOnHa(val bool) *PersonaCreateRequest {
	pc.CreateOnHa = val
	return pc
}
func (pc *PersonaCreateRequest) SetCreateOnJunk(val bool) *PersonaCreateRequest {
	pc.CreateOnJunk = val
	return pc
}
func (pc *PersonaCreateRequest) SetAsm(val bool) *PersonaCreateRequest {
	pc.Asm = val
	return pc
}
func (pc *PersonaCreateRequest) SetLiveness(val bool) *PersonaCreateRequest {
	pc.Liveness = val
	return pc
}
func (pc *PersonaCreateRequest) SetCreateLivenessOnly(val bool) *PersonaCreateRequest {
	pc.CreateLivenessOnly = val
	return pc
}

func NewPersonaCreateRequest(photoPath, sourceName string) (PersonaCreateRequest, error) {
	req := PersonaCreateRequest{Source: sourceName}
	err := req.FromFile(photoPath)
	if err != nil {
		return req, err
	}
	return req, nil
}

func NewPersonaCreateRequestFromReader(reader io.Reader, photoName, sourceName string) (PersonaCreateRequest, error) {
	req := PersonaCreateRequest{Source: sourceName}
	err := req.FromReader(reader, photoName)
	if err != nil {
		return req, err
	}
	return req, nil
}

type Search struct {
	common.Photo
	Asm      bool `json:"asm"`
	Liveness bool `json:"liveness"`
}

func (s *Search) MultipartWriterData() ([]byte, *multipart.Writer, error) {

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	part, err := w.CreateFormFile("photo", s.PhotoName)
	reader := bytes.NewReader(s.PhotoData)
	_, err = io.Copy(part, reader)
	if err != nil {
		return nil, nil, err
	}

	err = w.WriteField("asm", fmt.Sprintf("%v", s.Asm))
	if err != nil {
		return nil, nil, err
	}
	err = w.WriteField("liveness", fmt.Sprintf("%v", s.Liveness))
	if err != nil {
		return nil, nil, err
	}

	err = w.Close()
	if err != nil {
		return nil, nil, err
	}
	return buf.Bytes(), w, nil
}

func NewSearch(photoPath string, asm, liveness bool) (Search, error) {
	search := Search{Asm: asm, Liveness: liveness}
	err := search.FromFile(photoPath)
	if err != nil {
		return search, err
	}
	return search, nil
}

func NewSearchFromReader(reader io.Reader, photoName string, asm, liveness bool) (Search, error) {
	search := Search{Asm: asm, Liveness: liveness}
	err := search.FromReader(reader, photoName)
	if err != nil {
		return search, err
	}
	return search, nil
}

type SearchResult Person

func (sr SearchResult) Success() bool {
	switch sr.Conf {
	case conf.Exact:
		return true
	case conf.Ha:
		return true
	default:
		return false
	}
}

type ReinitImageRequest struct {
	common.Photo
	Conf               conf.Conf `json:"conf"`
	Source             string    `json:"source"`
	Facesize           int       `json:"facesize"`
	Liveness           bool      `json:"liveness"`
	ReinitLivenessOnly bool      `json:"reinit_liveness_only"`
	Idxid              string    `json:"-"`
}

func (ri *ReinitImageRequest) MultipartWriterData() ([]byte, *multipart.Writer, error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	part, err := w.CreateFormFile("photo", ri.PhotoName)
	reader := bytes.NewReader(ri.PhotoData)
	_, err = io.Copy(part, reader)
	if err != nil {
		return nil, nil, err
	}

	perMap, err := utils.ToMap(ri)
	if err != nil {
		return nil, nil, err
	}

	for k, v := range perMap {
		err := w.WriteField(k, fmt.Sprintf("%v", v))
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

func (ri *ReinitImageRequest) Validate() error {
	if ri.PhotoData == nil && ri.PhotoName == "" {
		return fmt.Errorf("photo data or name is empty")
	}
	if ri.Source == "" {
		return fmt.Errorf("source is required")
	}

	err := ri.Conf.Validate()
	return err
}

func NewReinitImageRequest(filePath string, idxid string, sourceName string) (ReinitImageRequest, error) {
	var request ReinitImageRequest
	request = ReinitImageRequest{Conf: conf.Ha, Idxid: idxid, Source: sourceName}
	err := request.FromFile(filePath)
	if err != nil {
		return request, err
	}
	return request, nil
}

func NewReinitImageRequestFromReader(reader io.Reader, photoName string, idxid string) (ReinitImageRequest, error) {
	var request ReinitImageRequest
	request = ReinitImageRequest{Conf: conf.Ha, Idxid: idxid}
	err := request.FromReader(reader, photoName)
	if err != nil {
		return request, err
	}
	return request, nil
}

func (ri *ReinitImageRequest) SetFacesize(val int) *ReinitImageRequest {
	ri.Facesize = val
	return ri
}

func (ri *ReinitImageRequest) SetLiveness(val bool) *ReinitImageRequest {
	ri.Liveness = val
	return ri
}

func (ri *ReinitImageRequest) SetReinitLivenessOnly(val bool) *ReinitImageRequest {
	ri.ReinitLivenessOnly = val
	return ri
}

func (ri *ReinitImageRequest) SetConf(val conf.Conf) *ReinitImageRequest {
	ri.Conf = val
	return ri
}

type ReinitIdRequest struct {
	ID       int `json:"id"`
	Facesize int `json:"facesize"`
}

func (re *ReinitIdRequest) Validate() error {
	if re.ID < 1 {
		return fmt.Errorf("id is required")
	}
	return nil
}
