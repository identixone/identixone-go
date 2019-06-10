package sex

import (
	"fmt"
)

type Sexes []Sex

func (sx Sexes) Validate() error {
	for _, s := range sx {
		if err := s.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Sex uint8

func (s Sex) Validate() error {
	switch s {
	case Male, Female:
		return nil
	default:
		return fmt.Errorf("unknown Sex %s`", s)
	}
}

func (s Sex) String() string {
	switch s {
	case Male:
		return "male"
	case Female:
		return "female"
	default:
		return "unknown"
	}
}

const (
	Male Sex = iota
	Female
)

func All() []Sex {
	return []Sex{Male, Female}
}
