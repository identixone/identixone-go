package entry

import (
	"encoding/json"
	"fmt"

	"github.com/identixone/identixone-go/core"
	"github.com/identixone/identixone-go/utils"
)

type Entries struct {
	request core.Requester
}

func NewEntries(request core.Requester) *Entries {
	return &Entries{request: request}
}

func (es *Entries) List(query map[string]interface{}) (ListResponse, error) {
	var resp ListResponse
	data, err := es.request.Get("/v1/entries/", query)
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (es *Entries) Delete(id int) error {
	return es.request.Delete(fmt.Sprintf("/v1/entries/%d/", id), nil)
}

func (es *Entries) StatsIdxid(idxid string) (StatsIdxid, error) {
	var resp StatsIdxid
	data, err := es.request.Get(fmt.Sprintf("/v1/entries/stats/idxid/%s/", idxid), nil)
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (es *Entries) StatsSources(req StatsSourcesRequest) (StatSourcesResponse, error) {
	var resp StatSourcesResponse

	query, err := utils.ToMap(req)
	if err != nil {
		return resp, err
	}
	data, err := es.request.Get("/v1/entries/stats/sources/", query)
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
