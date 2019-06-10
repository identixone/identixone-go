package common

import (
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
)

type LimitOffset struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Photo struct {
	PhotoData []byte `json:"-"`
	PhotoName string `json:"-"`
}

func (p *Photo) Validate() error {
	if p.PhotoName == "" || len(p.PhotoData) == 0 || p.PhotoData == nil {
		return fmt.Errorf("photo data or name is empty")
	}
	return nil
}

func (p *Photo) FromFile(photoPath string) error {
	data, err := ioutil.ReadFile(photoPath)
	if err != nil {
		return err
	}
	_, fn := filepath.Split(photoPath)
	p.PhotoData = data
	p.PhotoName = fn
	return nil
}
func (p *Photo) FromBytes(data []byte, name string) error {
	if data == nil || len(data) == 0 {
		return fmt.Errorf("data is empty")
	}
	p.PhotoData = data
	p.PhotoName = name
	return nil
}

func (p *Photo) FromReader(r io.Reader, name string) error {
	var buf []byte
	_, err := r.Read(buf)
	if err != nil {
		return err
	}
	p.PhotoData = buf
	p.PhotoName = name
	return nil
}

func NewPhotoFromFile(photoPath string) (Photo, error) {
	p := Photo{}
	err := p.FromFile(photoPath)
	if err != nil {
		return p, err
	}
	return p, nil
}

func NewPhotoFromReader(r io.Reader, name string) (Photo, error) {
	p := Photo{}
	err := p.FromReader(r, name)
	if err != nil {
		return p, err
	}
	return p, nil
}

func NewSearchQuery(q string) map[string]interface{} {
	query := map[string]interface{}{}
	if q != "" {
		query["q"] = q
	}
	return query
}

func NewPaginationQuery(limit, offset int) map[string]interface{} {
	if limit == 0 {
		limit = 50
	}
	query := map[string]interface{}{
		"limit":  limit,
		"offset": offset,
	}
	return query
}

func NewSearchPaginationQuery(q string, limit, offset int) map[string]interface{} {
	query := NewSearchQuery(q)
	for k, v := range NewPaginationQuery(limit, offset) {
		query[k] = v
	}
	return query
}
