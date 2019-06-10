package utility

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/identixone/identixone-go/api/const/liveness"
	"github.com/identixone/identixone-go/api/const/sex"
	"github.com/identixone/identixone-go/core"
	mock_core "github.com/identixone/identixone-go/core/mock"
	"github.com/identixone/identixone-go/utils"
)

func TestNewUtility(t *testing.T) {

}

func TestUtility_Asm(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_core.NewMockRequester(ctrl)
	u := NewUtility(m)

	req, err := NewAsmRequest("../../img/v2878.png")
	if err != nil {
		t.Fatal(err)
	}

	in, w, err := req.MultipartWriterData()
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().
		Post("/v1/utility/asm/", gomock.AssignableToTypeOf(in), gomock.AssignableToTypeOf(w.FormDataContentType())).
		Return([]byte(`{"age": 33, "sex": 1}`), nil).
		AnyTimes()

	resp, err := u.Asm(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Sex != sex.Female {
		t.Fatal("expect fail")
	}

	_, err = u.Asm(AsmRequest{})
	if err == nil {
		t.Fatal("expect fail")
	}
}

func TestUtility_Liveness(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_core.NewMockRequester(ctrl)
	u := NewUtility(m)

	req, err := NewLivenessRequest("../../img/v2878.png")
	if err != nil {
		t.Fatal(err)
	}

	in, w, err := req.MultipartWriterData()
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().
		Post("/v1/utility/liveness/", gomock.AssignableToTypeOf(in), gomock.AssignableToTypeOf(w.FormDataContentType())).
		Return([]byte(`{"liveness": "passed"}`), nil).
		AnyTimes()

	resp, err := u.Liveness(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Liveness != liveness.Passed {
		t.Fatal("expect fail")
	}

	_, err = u.Liveness(LivenessRequest{})
	if err == nil {
		t.Fatal("expect fail")
	}
}

func TestUtility_Compare(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_core.NewMockRequester(ctrl)
	u := NewUtility(m)

	req, err := NewCompareRequest("../../img/v2878.png", "../../img/v2885.png", nil)
	if err != nil {
		t.Fatal(err)
	}

	in, w, err := req.MultipartWriterData()
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().
		Post("/v1/utility/compare/", gomock.AssignableToTypeOf(in), gomock.AssignableToTypeOf(w.FormDataContentType())).
		Return([]byte(`{"similar": true}`), nil).
		AnyTimes()

	resp, err := u.Compare(req)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Similar {
		t.Fatal("expect fail")
	}

	_, err = u.Compare(CompareRequest{})
	if err == nil {
		t.Fatal("expect fail")
	}
}

func TestUtility_Customer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_core.NewMockRequester(ctrl)
	u := NewUtility(m)

	req := CustomerRequest{Source: "default", Offset: 10}

	in, err := utils.ToMap(req)
	if err != nil {
		t.Fatal(err)
	}
	req2 := CustomerRequest{Source: "asm", Offset: 10}

	in2, err := utils.ToMap(req2)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().
		Get("/v1/utility/customer/", in).
		Return([]byte(`{"idxid": "idxid"}`), nil).
		AnyTimes()

	m.EXPECT().
		Get("/v1/utility/customer/", in2).
		Return(nil, core.NewError(core.NotFound, fmt.Errorf("not found"), []byte(`{"detail": "not found"}`))).
		AnyTimes()

	resp, err := u.Customer(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Idxid != "idxid" {
		t.Fatal("expect fail")
	}

	resp, err = u.Customer(req2)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Idxid != "" {
		t.Fatal("expect fail")
	}

	if resp.Detail != "not found" {
		t.Fatal("expect fail")
	}
}
