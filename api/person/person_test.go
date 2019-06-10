package person

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/identixone/identixone-go/core"
	mock_core "github.com/identixone/identixone-go/core/mock"
)

func TestNewPersons(t *testing.T) {
	c, _ := core.NewRequest()
	p := NewPersons(c)
	type args struct {
		client *core.Request
	}
	tests := []struct {
		name string
		args args
		want *Persons
	}{
		{name: "new", args: args{client: c}, want: p},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPersons(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPersons() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPersons_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req, err := NewPersonaCreateRequest("../../img/v2878.png", "default")
	if err != nil {
		t.Fatal(err)
	}
	in, w, err := req.MultipartWriterData()
	if err != nil {
		t.Fatal(err)
	}

	m := mock_core.NewMockRequester(ctrl)
	m.EXPECT().
		Post("/v1/persons/", gomock.AssignableToTypeOf(in), gomock.AssignableToTypeOf(w.FormDataContentType())).
		Return([]byte(`{"idxid": "idxid"}`), nil).
		AnyTimes()

	p := NewPersons(m)

	resp, err := p.Create(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Idxid != "idxid" {
		t.Fatal("expect fail")
	}
}

func TestPersons_Search(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req, err := NewSearch("../../img/v2878.png", true, true)
	if err != nil {
		t.Fatal(err)
	}
	in, w, err := req.MultipartWriterData()
	if err != nil {
		t.Fatal(err)
	}
	m := mock_core.NewMockRequester(ctrl)
	m.EXPECT().
		Post("/v1/persons/search/", gomock.AssignableToTypeOf(in), gomock.AssignableToTypeOf(w.FormDataContentType())).
		Return([]byte(`{"idxid": "idxid"}`), nil).
		AnyTimes()

	p := NewPersons(m)

	resp, err := p.Search(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Idxid != "idxid" {
		t.Fatal("expect fail")
	}
}

func TestPersons_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_core.NewMockRequester(ctrl)
	m.EXPECT().
		Delete("/v1/persons/idxid/", nil).
		Return(nil).
		AnyTimes()

	p := NewPersons(m)

	err := p.Delete("idxid")
	if err != nil {
		t.Fatal(err)
	}
}

func TestPersons_ReinitImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req, err := NewReinitImageRequest("../../img/v2878.png", "idxid", "default")
	if err != nil {
		t.Fatal(err)
	}
	in, w, err := req.MultipartWriterData()
	if err != nil {
		t.Fatal(err)
	}
	m := mock_core.NewMockRequester(ctrl)
	m.EXPECT().
		Post("/v1/persons/reinit/idxid/", gomock.AssignableToTypeOf(in), gomock.AssignableToTypeOf(w.FormDataContentType())).
		Return(nil, nil).
		AnyTimes()

	p := NewPersons(m)

	err = p.ReinitImage(req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPersons_ReinitId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	req := ReinitIdRequest{ID: 10, Facesize: 100}
	in, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}

	m := mock_core.NewMockRequester(ctrl)

	m.EXPECT().
		Post("/v1/persons/reinit/", in, "application/json").
		Return(nil, nil).
		AnyTimes()

	p := NewPersons(m)

	err = p.ReinitId(req)
	if err != nil {
		t.Fatal(err)
	}
}
