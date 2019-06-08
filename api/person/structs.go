package person

import (
	"fmt"
	"io"

	"github.com/identixone/identixone-go/api/common"
	"github.com/identixone/identixone-go/api/const/conf"
	"github.com/identixone/identixone-go/api/const/liveness"
	"github.com/identixone/identixone-go/api/const/moods"
	"github.com/identixone/identixone-go/api/const/sex"
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

func (pc *PersonaCreateRequest) IsValid() error {
	if pc.PhotoData != nil && pc.PhotoName != "" {
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

func (ri *ReinitImageRequest) IsValid() error {
	if ri.PhotoData != nil && ri.PhotoName != "" {
		return fmt.Errorf("photo data or name is empty")
	}
	if ri.Source == "" {
		return fmt.Errorf("source is required")
	}

	err := ri.Conf.IsValid()
	return err
}

func NewReinitImageRequest(filePath string, idxid string) (ReinitImageRequest, error) {
	var request ReinitImageRequest
	request = ReinitImageRequest{Conf: conf.Ha, Idxid: idxid}
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
	Id       int `json:"id"`
	Facesize int `json:"facesize"`
}

func (re *ReinitIdRequest) IsValid() error {
	if re.Id < 1 {
		return fmt.Errorf("id is required")
	}
	return nil
}
