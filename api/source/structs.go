package source

import (
	"fmt"

	"github.com/identixone/identixone-go/api/const/conf"
)

type Source struct {
	ID                            int        `json:"id"`
	PpsTimestamp                  bool       `json:"pps_timestamp"`
	StoreImagesForConfs           conf.Confs `json:"store_images_for_confs"`
	Name                          string     `json:"name"`
	IdentifyFacesizeThreshold     int        `json:"identify_facesize_threshold"`
	AutoCreatePersons             bool       `json:"auto_create_persons"`
	AutoCreateFacesizeThreshold   int        `json:"auto_create_facesize_threshold"`
	AutoCreateOnHa                bool       `json:"auto_create_on_ha"`
	AutoCreateOnJunk              bool       `json:"auto_create_on_junk"`
	AutoCheckFaceAngle            bool       `json:"auto_check_face_angle"`
	AutoCheckAngleThreshold       int        `json:"auto_check_angle_threshold"`
	AutoCheckAsm                  bool       `json:"auto_check_asm"`
	AutoCreateCheckBlur           bool       `json:"auto_create_check_blur"`
	AutoCreateCheckExp            bool       `json:"auto_create_check_exp"`
	AutoCheckLiveness             bool       `json:"auto_check_liveness"`
	AutoCreateLivenessOnly        bool       `json:"auto_create_liveness_only"`
	ManualCreateFacesizeThreshold int        `json:"manual_create_facesize_threshold"`
	ManualCreateOnHa              bool       `json:"manual_create_on_ha"`
	ManualCreateOnJunk            bool       `json:"manual_create_on_junk"`
	ManualCheckAsm                bool       `json:"manual_check_asm"`
	ManualCreateLivenessOnly      bool       `json:"manual_create_liveness_only"`
	ManualCheckLiveness           bool       `json:"manual_check_liveness"`
}

func (s *Source) Validate() error {
	if s.Name == "" {
		return fmt.Errorf("source name is required")
	}
	return nil
}

func (s *Source) SetPpsTimestamp(val bool) *Source {
	s.PpsTimestamp = val
	return s
}

func (s *Source) SetStoreImagesForConfs(val conf.Confs) *Source {
	s.StoreImagesForConfs = val
	return s
}

func (s *Source) SetName(val string) *Source {
	s.Name = val
	return s
}

func (s *Source) SetIdentifyFacesizeThreshold(val int) *Source {
	s.IdentifyFacesizeThreshold = val
	return s
}

func (s *Source) SetAutoCreatePersons(val bool) *Source {
	s.AutoCreatePersons = val
	return s
}

func (s *Source) SetAutoCreateFacesizeThreshold(val int) *Source {
	s.AutoCreateFacesizeThreshold = val
	return s
}

func (s *Source) SetAutoCreateOnHa(val bool) *Source {
	s.AutoCreateOnHa = val
	return s
}

func (s *Source) SetAutoCreateOnJunk(val bool) *Source {
	s.AutoCreateOnJunk = val
	return s
}

func (s *Source) SetAutoCheckFaceAngle(val bool) *Source {
	s.AutoCheckFaceAngle = val
	return s
}

func (s *Source) SetAutoCheckAngleThreshold(val int) *Source {
	s.AutoCheckAngleThreshold = val
	return s
}

func (s *Source) SetAutoCheckAsm(val bool) *Source {
	s.AutoCheckAsm = val
	return s
}

func (s *Source) SetAutoCreateCheckBlur(val bool) *Source {
	s.AutoCreateCheckBlur = val
	return s
}

func (s *Source) SetAutoCreateCheckExp(val bool) *Source {
	s.AutoCreateCheckExp = val
	return s
}

func (s *Source) SetAutoCheckLiveness(val bool) *Source {
	s.AutoCheckLiveness = val
	return s
}

func (s *Source) SetAutoCreateLivenessOnly(val bool) *Source {
	s.AutoCreateLivenessOnly = val
	return s
}

func (s *Source) SetManualCreateFacesizeThreshold(val int) *Source {
	s.ManualCreateFacesizeThreshold = val
	return s
}

func (s *Source) SetManualCreateOnHa(val bool) *Source {
	s.ManualCreateOnHa = val
	return s
}

func (s *Source) SetManualCreateOnJunk(val bool) *Source {
	s.ManualCreateOnJunk = val
	return s
}

func (s *Source) SetManualCheckAsm(val bool) *Source {
	s.ManualCheckAsm = val
	return s
}

func (s *Source) SetManualCreateLivenessOnly(val bool) *Source {
	s.ManualCreateLivenessOnly = val
	return s
}

func (s *Source) SetManualCheckLiveness(val bool) *Source {
	s.ManualCheckLiveness = val
	return s
}

func DefaultSource() Source {
	return Source{
		PpsTimestamp:                  false,
		StoreImagesForConfs:           conf.All(),
		IdentifyFacesizeThreshold:     7000,
		AutoCreatePersons:             false,
		AutoCreateFacesizeThreshold:   25000,
		AutoCreateOnHa:                false,
		AutoCreateOnJunk:              false,
		AutoCheckFaceAngle:            true,
		AutoCheckAngleThreshold:       9,
		AutoCheckAsm:                  true,
		AutoCreateCheckBlur:           true,
		AutoCreateCheckExp:            true,
		AutoCheckLiveness:             false,
		AutoCreateLivenessOnly:        true,
		ManualCreateFacesizeThreshold: 25000,
		ManualCreateOnHa:              false,
		ManualCreateOnJunk:            false,
		ManualCheckAsm:                true,
		ManualCreateLivenessOnly:      true,
		ManualCheckLiveness:           false,
	}
}

func DefaultSourceWithName(name string) Source {
	s := DefaultSource()
	s.Name = name
	return s
}

type ListResponse struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Sources  []Source `json:"results"`
}

type UpdateRequest struct {
	ID                            int         `json:"id"`
	PpsTimestamp                  *bool       `json:"pps_timestamp,omitempty"`
	StoreImagesForConfs           *conf.Confs `json:"store_images_for_confs,omitempty"`
	Name                          *string     `json:"name,omitempty"`
	IdentifyFacesizeThreshold     *int        `json:"identify_facesize_threshold,omitempty"`
	AutoCreatePersons             *bool       `json:"auto_create_persons,omitempty"`
	AutoCreateFacesizeThreshold   *int        `json:"auto_create_facesize_threshold,omitempty"`
	AutoCreateOnHa                *bool       `json:"auto_create_on_ha,omitempty"`
	AutoCreateOnJunk              *bool       `json:"auto_create_on_junk,omitempty"`
	AutoCheckFaceAngle            *bool       `json:"auto_check_face_angle,omitempty"`
	AutoCheckAngleThreshold       *int        `json:"auto_check_angle_threshold,omitempty"`
	AutoCheckAsm                  *bool       `json:"auto_check_asm,omitempty"`
	AutoCreateCheckBlur           *bool       `json:"auto_create_check_blur,omitempty"`
	AutoCreateCheckExp            *bool       `json:"auto_create_check_exp,omitempty"`
	AutoCheckLiveness             *bool       `json:"auto_check_liveness,omitempty"`
	AutoCreateLivenessOnly        *bool       `json:"auto_create_liveness_only,omitempty"`
	ManualCreateFacesizeThreshold *int        `json:"manual_create_facesize_threshold,omitempty"`
	ManualCreateOnHa              *bool       `json:"manual_create_on_ha,omitempty"`
	ManualCreateOnJunk            *bool       `json:"manual_create_on_junk,omitempty"`
	ManualCheckAsm                *bool       `json:"manual_check_asm,omitempty"`
	ManualCreateLivenessOnly      *bool       `json:"manual_create_liveness_only,omitempty"`
	ManualCheckLiveness           *bool       `json:"manual_check_liveness,omitempty"`
}

func (u UpdateRequest) Validate() error {
	if u.ID == 0 {
		return fmt.Errorf("id is required")
	}
	if u.Name != nil && *u.Name == "" {
		return fmt.Errorf("name canot bbe empty")
	}
	return nil
}

func (u *UpdateRequest) SetPpsTimestamp(val bool) *UpdateRequest {
	u.PpsTimestamp = &val
	return u
}

func (u *UpdateRequest) SetStoreImagesForConfs(val conf.Confs) *UpdateRequest {
	u.StoreImagesForConfs = &val
	return u
}

func (u *UpdateRequest) SetName(val string) *UpdateRequest {
	u.Name = &val
	return u
}

func (u *UpdateRequest) SetIdentifyFacesizeThreshold(val int) *UpdateRequest {
	u.IdentifyFacesizeThreshold = &val
	return u
}

func (u *UpdateRequest) SetAutoCreatePersons(val bool) *UpdateRequest {
	u.AutoCreatePersons = &val
	return u
}

func (u *UpdateRequest) SetAutoCreateFacesizeThreshold(val int) *UpdateRequest {
	u.AutoCreateFacesizeThreshold = &val
	return u
}

func (u *UpdateRequest) SetAutoCreateOnHa(val bool) *UpdateRequest {
	u.AutoCreateOnHa = &val
	return u
}

func (u *UpdateRequest) SetAutoCreateOnJunk(val bool) *UpdateRequest {
	u.AutoCreateOnJunk = &val
	return u
}

func (u *UpdateRequest) SetAutoCheckFaceAngle(val bool) *UpdateRequest {
	u.AutoCheckFaceAngle = &val
	return u
}

func (u *UpdateRequest) SetAutoCheckAngleThreshold(val int) *UpdateRequest {
	u.AutoCheckAngleThreshold = &val
	return u
}

func (u *UpdateRequest) SetAutoCheckAsm(val bool) *UpdateRequest {
	u.AutoCheckAsm = &val
	return u
}

func (u *UpdateRequest) SetAutoCreateCheckBlur(val bool) *UpdateRequest {
	u.AutoCreateCheckBlur = &val
	return u
}

func (u *UpdateRequest) SetAutoCreateCheckExp(val bool) *UpdateRequest {
	u.AutoCreateCheckExp = &val
	return u
}

func (u *UpdateRequest) SetAutoCheckLiveness(val bool) *UpdateRequest {
	u.AutoCheckLiveness = &val
	return u
}

func (u *UpdateRequest) SetAutoCreateLivenessOnly(val bool) *UpdateRequest {
	u.AutoCreateLivenessOnly = &val
	return u
}

func (u *UpdateRequest) SetManualCreateFacesizeThreshold(val int) *UpdateRequest {
	u.ManualCreateFacesizeThreshold = &val
	return u
}

func (u *UpdateRequest) SetManualCreateOnHa(val bool) *UpdateRequest {
	u.ManualCreateOnHa = &val
	return u
}

func (u *UpdateRequest) SetManualCreateOnJunk(val bool) *UpdateRequest {
	u.ManualCreateOnJunk = &val
	return u
}

func (u *UpdateRequest) SetManualCheckAsm(val bool) *UpdateRequest {
	u.ManualCheckAsm = &val
	return u
}

func (u *UpdateRequest) SetManualCreateLivenessOnly(val bool) *UpdateRequest {
	u.ManualCreateLivenessOnly = &val
	return u
}

func (u *UpdateRequest) SetManualCheckLiveness(val bool) *UpdateRequest {
	u.ManualCheckLiveness = &val
	return u
}
