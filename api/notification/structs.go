package notification

import (
	"github.com/identixone/identixone-go/api/const/conf"
	"github.com/identixone/identixone-go/api/const/liveness"
	"github.com/identixone/identixone-go/api/const/method"
	"github.com/identixone/identixone-go/api/const/moods"
	"github.com/identixone/identixone-go/api/const/sex"
	"github.com/identixone/identixone-go/api/const/transport"
)

type Notification struct {
	ID             int                  `json:"id"`
	Name           string               `json:"name"`
	IsActive       bool                 `json:"is_active"`
	Transport      transport.Transport  `json:"transport"`
	DestinationURL string               `json:"destination_url"`
	ConfThresholds *[]conf.Conf         `json:"conf_thresholds"`
	AgeFrom        *int                 `json:"age_from"`
	AgeTo          *int                 `json:"age_to"`
	Sex            *[]sex.Sex           `json:"sex"`
	Moods          *[]moods.Mood        `json:"moods"`
	Liveness       *[]liveness.Liveness `json:"liveness"`
	Sources        *string              `json:"sources"`
	HTTPMethod     *method.Method       `json:"http_method"`
}

type ListResponse struct {
	Count         int             `json:"count"`
	Next          *string         `json:"next"`
	Previous      *string         `json:"previous"`
	Notifications []Notifications `json:"results"`
}

type CreateRequest struct {
}
