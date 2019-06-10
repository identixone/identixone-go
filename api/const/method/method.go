package method

import (
	"fmt"
)

type Method uint8

func (m Method) String() string {
	switch m {
	case Post:
		return "POST"
	case Get:
		return "GET"
	default:
		return "unknown"
	}
}

func (m Method) Validate() error {
	switch m {
	case Post, Get:
		return nil
	default:
		return fmt.Errorf("unknown Method %s", m)
	}
}

const (
	Post Method = iota
	Get
)

func Point(m Method) *Method {
	return &m
}
