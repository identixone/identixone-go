package entry

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/identixone/identixone-go/core"
	mock_core "github.com/identixone/identixone-go/core/mock"
	"github.com/identixone/identixone-go/utils"
)

func TestNewEntries(t *testing.T) {

	c, _ := core.NewRequest()
	e := NewEntries(c)
	type args struct {
		client *core.Request
	}
	tests := []struct {
		name string
		args args
		want *Entries
	}{
		{name: "valid", args: args{client: c}, want: e},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEntries(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEntries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntries_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_core.NewMockRequester(ctrl)

	m.EXPECT().
		Get(gomock.Eq("/v1/entries/"), nil).
		Return([]byte(`{"results": [{"id": 1}, {"id": 2}]}`), nil).
		AnyTimes()

	e := NewEntries(m)

	resp, err := e.List(nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Entries) != 2 {
		t.Fatal(fmt.Errorf("list response fail"))
	}

	m.EXPECT().
		Get(gomock.Eq("/v1/entries/"), gomock.Eq(map[string]interface{}{"offset": 20})).
		Return([]byte(`{"results": []}`), nil).
		AnyTimes()

	resp, err = e.List(map[string]interface{}{"offset": 20})
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Entries) != 0 {
		t.Fatal(fmt.Errorf("list response fail"))
	}

	m.EXPECT().
		Get(gomock.Eq("/v1/entries/"), gomock.Eq(map[string]interface{}{"bar": 20})).
		Return(nil, fmt.Errorf("interanl error")).
		AnyTimes()

	_, err = e.List(map[string]interface{}{"bar": 20})
	if err == nil {
		t.Fatal(err)
	}

	m.EXPECT().
		Get(gomock.Eq("/v1/entries/"), gomock.Eq(map[string]interface{}{"limit": -1})).
		Return([]byte("teapot"), nil).
		AnyTimes()

	_, err = e.List(map[string]interface{}{"limit": -1})
	if err == nil {
		t.Fatal(err)
	}
}

func TestEntries_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_core.NewMockRequester(ctrl)

	m.EXPECT().Delete(gomock.Eq("/v1/entries/1/"), nil).
		Return(nil).AnyTimes()
	m.EXPECT().Delete(gomock.Eq("/v1/entries/0/"), nil).
		Return(fmt.Errorf("not found")).AnyTimes()
	e := NewEntries(m)

	err := e.Delete(1)
	if err != nil {
		t.Fatal(err)
	}
	err = e.Delete(0)
	if err == nil {
		t.Fatal(err)
	}

}

func TestEntries_StatsIdxid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_core.NewMockRequester(ctrl)

	m.EXPECT().
		Get(gomock.Eq("/v1/entries/stats/idxid/idxid1/"), nil).
		Return([]byte(`{"idxid": "idxid1"}`), nil).
		AnyTimes()

	m.EXPECT().
		Get(gomock.Eq("/v1/entries/stats/idxid/idxid0/"), nil).
		Return(nil, fmt.Errorf("not found")).
		AnyTimes()

	e := NewEntries(m)

	data, err := e.StatsIdxid("idxid1")
	if err != nil {
		t.Fatal(err)
	}
	if data.Idxid != "idxid1" {
		t.Fatal("idxid not found")
	}

	data, err = e.StatsIdxid("idxid0")
	if err == nil {
		t.Fatal("expected error fail")
	}
}

func TestEntries_StatsSources(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_core.NewMockRequester(ctrl)
	req1, err := utils.ToMap(StatsSourcesRequest{Idxid: "idxid"})
	if err != nil {
		t.Fatal(err)
	}
	req2, err := utils.ToMap(StatsSourcesRequest{})
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().
		Get(gomock.Eq("/v1/entries/stats/sources/"), req1).
		Return([]byte(`{"count": 1}`), nil).
		AnyTimes()

	m.EXPECT().
		Get(gomock.Eq("/v1/entries/stats/sources/"), req2).
		Return(nil, fmt.Errorf("not found")).
		AnyTimes()

	e := NewEntries(m)

	data, err := e.StatsSources(StatsSourcesRequest{Idxid: "idxid"})
	if err != nil {
		t.Fatal(err)
	}
	if data.Count != 1 {
		t.Fatal("idxid not found")
	}

	data, err = e.StatsSources(StatsSourcesRequest{})
	if err == nil {
		t.Fatal("expected error fail")
	}
}
