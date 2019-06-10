package notification

import (
	"fmt"

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
	Count         int            `json:"count"`
	Next          *string        `json:"next"`
	Previous      *string        `json:"previous"`
	Notifications []Notification `json:"results"`
}

type CreateRequest struct {
	Name           string               `json:"name,omitempty"`
	IsActive       *bool                `json:"is_active,omitempty"`
	Transport      transport.Transport  `json:"transport,omitempty"`
	HTTPMethod     *method.Method       `json:"http_method,omitempty"`
	DestinationURL string               `json:"destination_url,omitempty"`
	Moods          *moods.Moods         `json:"moods,omitempty"`
	ConfThresholds *conf.Confs          `json:"conf_thresholds,omitempty"`
	Liveness       *liveness.Livenesses `json:"liveness,omitempty"`
	AgeFrom        *int                 `json:"age_from,omitempty"`
	AgeTo          *int                 `json:"age_to,omitempty"`
	Sex            *sex.Sexes           `json:"sex,omitempty"`
	Sources        *[]string            `json:"sources,omitempty"`
}

func (cr CreateRequest) Validate() error {
	if cr.Name == "" {
		return fmt.Errorf("name is required")
	}

	if err := cr.Transport.Validate(); err != nil {
		return err
	}

	if cr.HTTPMethod == nil {
		return fmt.Errorf("HTTPMethod is required")
	} else {
		if err := cr.HTTPMethod.Validate(); err != nil {
			return err
		}
	}

	if (cr.Transport == transport.Webhook || cr.Transport == transport.WebsocketClient) && cr.DestinationURL == "" {
		return fmt.Errorf("DestinationUrl is required for transports Webhook or WebsocketClient")
	}

	if cr.Moods != nil {
		if err := cr.Moods.Validate(); err != nil {
			return err
		}
	}

	if cr.ConfThresholds != nil {
		if err := cr.ConfThresholds.Validate(); err != nil {
			return err
		}
	}

	if cr.Liveness != nil {
		if err := cr.Liveness.Validate(); err != nil {
			return err
		}
	}

	if cr.Sex != nil {
		if err := cr.Sex.Validate(); err != nil {
			return err
		}
	}

	if cr.Sources != nil {
		for _, s := range *cr.Sources {
			if s == "" {
				return fmt.Errorf("source name canot be empty")
			}
		}
	}

	return nil
}

type UpdateRequest struct {
	ID             int                  `json:"id"`
	Name           string               `json:"name,omitempty"`
	IsActive       *bool                `json:"is_active,omitempty"`
	Transport      transport.Transport  `json:"transport,omitempty"`
	HTTPMethod     *method.Method       `json:"http_method,omitempty"`
	DestinationURL string               `json:"destination_url,omitempty"`
	Moods          *moods.Moods         `json:"moods,omitempty"`
	ConfThresholds *conf.Confs          `json:"conf_thresholds,omitempty"`
	Liveness       *liveness.Livenesses `json:"liveness,omitempty"`
	AgeFrom        *int                 `json:"age_from,omitempty"`
	AgeTo          *int                 `json:"age_to,omitempty"`
	Sex            *sex.Sexes           `json:"sex,omitempty"`
	Sources        *[]string            `json:"sources,omitempty"`
}

func (ur UpdateRequest) Validate() error {
	if ur.ID == 0 {
		return fmt.Errorf("id is reequired")
	}
	if ur.Name == "" {
		return fmt.Errorf("name is required")
	}

	if err := ur.Transport.Validate(); err != nil {
		return err
	}

	if ur.HTTPMethod == nil {
		return fmt.Errorf("HTTPMethod is required")
	} else {
		if err := ur.HTTPMethod.Validate(); err != nil {
			return err
		}
	}

	if (ur.Transport == transport.Webhook || ur.Transport == transport.WebsocketClient) && ur.DestinationURL == "" {
		return fmt.Errorf("DestinationUrl is required for transports Webhook or WebsocketClient")
	}

	if ur.Moods != nil {
		if err := ur.Moods.Validate(); err != nil {
			return err
		}
	}

	if ur.ConfThresholds != nil {
		if err := ur.ConfThresholds.Validate(); err != nil {
			return err
		}
	}

	if ur.Liveness != nil {
		if err := ur.Liveness.Validate(); err != nil {
			return err
		}
	}

	if ur.Sex != nil {
		if err := ur.Sex.Validate(); err != nil {
			return err
		}
	}

	if ur.Sources != nil {
		for _, s := range *ur.Sources {
			if s == "" {
				return fmt.Errorf("source name canot be empty")
			}
		}
	}

	return nil
}
