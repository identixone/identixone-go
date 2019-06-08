package common

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestPhoto_FromFile(t *testing.T) {
	type args struct {
		photoPath string
	}
	tests := []struct {
		name    string
		p       *Photo
		args    args
		wantErr bool
	}{
		{name: "valid file", p: &Photo{}, args: args{"../../img/v2878.png"}, wantErr: false},
		{name: "invalid file", p: &Photo{}, args: args{"img/0000.png"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.FromFile(tt.args.photoPath); (err != nil) != tt.wantErr {
				t.Errorf("Photo.FromFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPhoto_FromBytes(t *testing.T) {
	type args struct {
		data []byte
		name string
	}

	buf, err := ioutil.ReadFile("../../img/v2878.png")
	if err != nil {
		t.Error(err)
	}
	var empty []byte
	tests := []struct {
		name    string
		p       *Photo
		args    args
		wantErr bool
	}{
		{name: "valid data", p: &Photo{}, args: args{data: buf, name: "v2878.png"}, wantErr: false},
		{name: "empty data", p: &Photo{}, args: args{data: empty, name: ""}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.FromBytes(tt.args.data, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Photo.FromBytes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewPhotoFromFile(t *testing.T) {
	type args struct {
		photoPath string
	}

	p, err := NewPhotoFromFile("../../img/v2878.png")
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		name    string
		args    args
		want    Photo
		wantErr bool
	}{
		{name: "valid file", args: args{"../../img/v2878.png"}, wantErr: false, want: p},
		{name: "invalid file", args: args{"img/0000.png"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPhotoFromFile(tt.args.photoPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPhotoFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPhotoFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSearchQuery(t *testing.T) {
	type args struct {
		q string
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{name: "empty query", args: args{q: ""}, want: map[string]interface{}{}},
		{name: "empty query", args: args{q: "name"}, want: map[string]interface{}{"q": "name"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSearchQuery(tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSearchQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPaginationQuery(t *testing.T) {
	type args struct {
		limit  int
		offset int
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{name: "default query", args: args{limit: 0}, want: map[string]interface{}{"limit": 50, "offset": 0}},
		{name: "user query", args: args{limit: 100, offset: 100}, want: map[string]interface{}{"limit": 100, "offset": 100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPaginationQuery(tt.args.limit, tt.args.offset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPaginationQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSearchPaginationQuery(t *testing.T) {
	type args struct {
		q      string
		limit  int
		offset int
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{name: "default query", args: args{limit: 0}, want: map[string]interface{}{"limit": 50, "offset": 0}},
		{name: "default query", args: args{limit: 0, q: "name"}, want: map[string]interface{}{"limit": 50, "offset": 0, "q": "name"}},
		{name: "user query", args: args{limit: 100, offset: 100}, want: map[string]interface{}{"limit": 100, "offset": 100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSearchPaginationQuery(tt.args.q, tt.args.limit, tt.args.offset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSearchPaginationQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhoto_IsValid(t *testing.T) {
	type fields struct {
		PhotoData []byte
		PhotoName string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "invalid", fields: fields{PhotoData: nil, PhotoName: ""}, wantErr: true},
		{name: "invalid", fields: fields{PhotoData: []byte("hello world"), PhotoName: "some.name"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Photo{
				PhotoData: tt.fields.PhotoData,
				PhotoName: tt.fields.PhotoName,
			}
			if err := p.IsValid(); (err != nil) != tt.wantErr {
				t.Errorf("Photo.IsValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPhoto_FromReader(t *testing.T) {
	buf := []byte("hello world")
	r := bytes.NewReader(buf)

	r2, err := os.Open("../../img/v2878.png")
	if err != nil {
		t.Fatal(err)
	}
	type fields struct {
		PhotoData []byte
		PhotoName string
	}
	type args struct {
		r    io.Reader
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "valid", fields: fields{PhotoName: "", PhotoData: nil}, args: args{r: r, name: "some.name"}, wantErr: false},
		{name: "valid", fields: fields{PhotoName: "", PhotoData: nil}, args: args{r: r2, name: "some.name"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Photo{
				PhotoData: tt.fields.PhotoData,
				PhotoName: tt.fields.PhotoName,
			}
			if err := p.FromReader(tt.args.r, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Photo.FromReader() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewPhotoFromReader(t *testing.T) {
	r, err := os.Open("../../img/v2878.png")
	r2, _ := os.Open("../../img/v.png")
	if err != nil {
		t.Fatal(err)
	}
	p, err := NewPhotoFromReader(r, "some.name")
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		r    io.Reader
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    Photo
		wantErr bool
	}{
		{name: "valid", args: args{r: r, name: "some.name"}, want: p, wantErr: false},
		{name: "invalid", args: args{r: r2, name: "some.name"}, want: Photo{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPhotoFromReader(tt.args.r, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPhotoFromReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPhotoFromReader() = %v, want %v", got, tt.want)
			}
		})
	}
}
