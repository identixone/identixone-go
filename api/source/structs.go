package source

import (
	"fmt"

	"github.com/identixone/identixone-go/api/const/conf"
)

type Source struct {
	ID                            int         `json:"id"`
	PpsTimestamp                  bool        `json:"pps_timestamp"`
	StoreImagesForConfs           []conf.Conf `json:"store_images_for_confs"`
	Name                          string      `json:"name"`
	IdentifyFacesizeThreshold     int         `json:"identify_facesize_threshold"`
	AutoCreatePersons             bool        `json:"auto_create_persons"`
	AutoCreateFacesizeThreshold   int         `json:"auto_create_facesize_threshold"`
	AutoCreateOnHa                bool        `json:"auto_create_on_ha"`
	AutoCreateOnJunk              bool        `json:"auto_create_on_junk"`
	AutoCheckFaceAngle            bool        `json:"auto_check_face_angle"`
	AutoCheckAngleThreshold       int         `json:"auto_check_angle_threshold"`
	AutoCheckAsm                  bool        `json:"auto_check_asm"`
	AutoCreateCheckBlur           bool        `json:"auto_create_check_blur"`
	AutoCreateCheckExp            bool        `json:"auto_create_check_exp"`
	AutoCheckLiveness             bool        `json:"auto_check_liveness"`
	AutoCreateLivenessOnly        bool        `json:"auto_create_liveness_only"`
	ManualCreateFacesizeThreshold int         `json:"manual_create_facesize_threshold"`
	ManualCreateOnHa              bool        `json:"manual_create_on_ha"`
	ManualCreateOnJunk            bool        `json:"manual_create_on_junk"`
	ManualCheckAsm                bool        `json:"manual_check_asm"`
	ManualCreateLivenessOnly      bool        `json:"manual_create_liveness_only"`
	ManualCheckLiveness           bool        `json:"manual_check_liveness"`
}

func (s *Source) IsValid() error {
	if s.Name == "" {
		return fmt.Errorf("source name is required")
	}
	return nil
}

func (s *Source) SetPpsTimestamp(val bool) *Source {
	s.PpsTimestamp = val
	return s
}

func (s *Source) SetStoreImagesForConfs(val []conf.Conf) *Source {
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
