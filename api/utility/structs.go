package utility

import (
	"fmt"

	"github.com/identixone/identixone-go/api/common"
	"github.com/identixone/identixone-go/api/const/conf"
	"github.com/identixone/identixone-go/api/const/liveness"
)

type AsmResponse struct {
	Age  int    `json:"age"`
	Sex  int    `json:"sex"`
	Mood string `json:"mood"`
}

type LivenessResponse struct {
	Liveness liveness.Liveness `json:"liveness"`
}

type CompareRequest struct {
	Photo1         common.Photo `json:"-"`
	Photo2         common.Photo `json:"-"`
	LivenessPhoto1 bool         `json:"liveness_photo_1"`
	LivenessPhoto2 bool         `json:"liveness_photo_2"`
	Conf           conf.Conf    `json:"conf"`
}

func (compReq *CompareRequest) IsValid() error {
	err := compReq.Photo1.IsValid()
	if err != nil {
		return err
	}
	err = compReq.Photo2.IsValid()
	if err != nil {
		return err
	}
	err = compReq.Conf.IsValid()
	return err
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

func (custReq *CustomerRequest) IsValid() error {
	if custReq.Source == "" {
		return fmt.Errorf("soure name is required")
	}
	return nil
}

type CustomerResponse struct {
	Idxid string `json:"idxid"`
}
