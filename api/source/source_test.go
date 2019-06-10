package source

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	mock_core "github.com/identixone/identixone-go/core/mock"
)

func TestNewSource(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_core.NewMockRequester(ctrl)
	s := NewSource(m)
	if s == nil {
		t.Fatal("new sources error")
	}
}

func TestSources_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_core.NewMockRequester(ctrl)
	s := NewSource(m)

	m.EXPECT().
		Get("/v1/sources/", nil).
		Return([]byte(`{"count": 1, "results": [{"id": 1}]}`), nil).
		AnyTimes()

	offsetQuery := map[string]interface{}{"offset": 10}
	m.EXPECT().
		Get("/v1/sources/", offsetQuery).
		Return([]byte(`{"count": 0, "results": []}`), nil).
		AnyTimes()

	limitQuery := map[string]interface{}{"limitQuery": -1}
	m.EXPECT().
		Get("/v1/sources/", limitQuery).
		Return(nil, fmt.Errorf("invalid params")).
		AnyTimes()

	m.EXPECT().
		Get("/v1/sources/", map[string]interface{}{"teapot": 1}).
		Return([]byte(`teapot`), nil).
		AnyTimes()

	resp, err := s.List(nil)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Count != 1 {
		t.Fatal("expect fail")
	}

	if len(resp.Sources) != 1 {
		t.Fatal("sources count fail")
	}

	resp, err = s.List(offsetQuery)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Count != 0 {
		t.Fatal("expect fail")
	}

	if len(resp.Sources) != 0 {
		t.Fatal("sources count fail")
	}
	_, err = s.List(limitQuery)
	if err == nil {
		t.Fatal("expect fail")
	}
	_, err = s.List(map[string]interface{}{"teapot": 1})
	if err == nil {
		t.Fatal("expect fail")
	}

}

func TestSources_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_core.NewMockRequester(ctrl)

	source := DefaultSourceWithName("default")
	in, err := json.Marshal(source)
	if err != nil {
		t.Fatal(err)
	}
	source.ID = 1
	resData, err := json.Marshal(source)

	m.EXPECT().
		Post("/v1/sources/", gomock.AssignableToTypeOf(in), "application/json").
		Return(resData, nil).
		AnyTimes()

	s := NewSource(m)

	resp, err := s.Create(source)
	if err != nil {
		t.Fatal(err)
	}

	if resp.ID != 1 {
		t.Fatal("expect fail")
	}
}

func TestSources_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_core.NewMockRequester(ctrl)
	s := NewSource(m)

	m.EXPECT().
		Get("/v1/sources/1/", nil).
		Return([]byte(`{"id": 1}`), nil)

	m.EXPECT().
		Get("/v1/sources/0/", nil).
		Return(nil, fmt.Errorf("not found"))

	resp, err := s.Get(1)
	if err != nil {
		t.Fatal(err)
	}

	if resp.ID != 1 {
		t.Fatal("expect fail")
	}

	_, err = s.Get(0)
	if err == nil {
		t.Fatal("expect fail")
	}

}

func TestSources_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_core.NewMockRequester(ctrl)
	s := NewSource(m)

	m.EXPECT().Delete("/v1/sources/1/", nil).Return(nil).AnyTimes()
	m.EXPECT().Delete("/v1/sources/0/", nil).Return(fmt.Errorf("not found")).AnyTimes()
	err := s.Delete(1)
	if err != nil {
		t.Fatal(err)
	}

	err = s.Delete(0)
	if err == nil {
		t.Fatal("expect fail")
	}
}

func TestSources_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_core.NewMockRequester(ctrl)
	s := NewSource(m)

	req := UpdateRequest{ID: 1}

	in, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().Patch("/v1/sources/1/", gomock.AssignableToTypeOf(in), "application/json").
		Return(in, nil).
		AnyTimes()

	resp, err := s.Update(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.ID != 1 {
		t.Fatal("expect fail")
	}

}
