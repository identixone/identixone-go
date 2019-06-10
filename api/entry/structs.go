package entry

import (
	"github.com/identixone/identixone-go/api/const/conf"
	"github.com/identixone/identixone-go/api/const/liveness"

	"time"
)

type Entry struct {
	ID      int       `json:"id"`
	Created time.Time `json:"created"`
	Photo   string    `json:"photo"`
	Source  struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"source"`
	Facesize     int       `json:"facesize"`
	Age          int       `json:"age"`
	Sex          int       `json:"sex"`
	Mood         string    `json:"mood"`
	Liveness     string    `json:"liveness"`
	Idxid        string    `json:"idxid"`
	Conf         conf.Conf `json:"conf"`
	IdxidCreated time.Time `json:"idxid_created"`
	InitialPhoto string    `json:"initial_photo"`
}

type StatSource struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Total int    `json:"total"`
	Conf  struct {
		Reinit int `json:"reinit"`
		Exact  int `json:"exact"`
		Ha     int `json:"ha"`
		Junk   int `json:"junk"`
		Nm     int `json:"nm"`
		Det    int `json:"det"`
		New    int `json:"new"`
	} `json:"conf"`
	Liveness struct {
		Failed       int `json:"failed"`
		Passed       int `json:"passed"`
		Undetermined int `json:"undetermined"`
	} `json:"liveness"`
}
type StatsSourcesRequest struct {
	Idxid       string            `json:"idxid,omitempty"`
	Conf        conf.Conf         `json:"conf,omitempty"`
	Liveness    liveness.Liveness `json:"liveness,omitempty"`
	Source      int               `json:"source,omitempty"`
	EntryIdFrom int               `json:"entry_id_from,omitempty"`
	DateFrom    time.Time         `json:"date_from,omitempty"`
	DateTo      time.Time         `json:"date_to,omitempty"`
}
type StatSourcesResponse struct {
	Count       int          `json:"count"`
	Next        *string      `json:"next"`
	Previous    *string      `json:"previous"`
	StatSources []StatSource `json:"results"`
}

type ListResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Entries  []Entry `json:"results"`
}

type StatsIdxid struct {
	Idxid        string    `json:"idxid"`
	IdxidCreated time.Time `json:"idxid_created"`
	IdxidSource  struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"idxid_source"`
	InitialPhoto    string `json:"initial_photo"`
	InitialFacesize int    `json:"initial_facesize"`
	InitialLiveness string `json:"initial_liveness"`
	Age             int    `json:"age"`
	Sex             int    `json:"sex"`
	Total           int    `json:"total"`
	Exact           int    `json:"exact"`
	Ha              int    `json:"ha"`
	Junk            int    `json:"junk"`
	Reinit          int    `json:"reinit"`
	Liveness        struct {
		Failed       int `json:"failed"`
		Passed       int `json:"passed"`
		Undetermined int `json:"undetermined"`
	} `json:"liveness"`
}
