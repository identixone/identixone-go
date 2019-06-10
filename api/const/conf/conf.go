package conf

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Confs []Conf

func (cs Confs) Validate() error {
	for _, c := range cs {
		if err := c.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Conf string

func (c Conf) String() string {
	return string(c)
}

func (c Conf) Validate() error {
	switch c {
	case Nm, New, Exact, Ha, Junk, Det, Reinit:
		return nil
	default:
		return fmt.Errorf("unknown type of Conf %s", c.String())
	}

}

func (c *Conf) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(*c))
}

func (c *Conf) UnmarshalJSON(b []byte) error {
	source := string(b)
	source = strings.Replace(source, `"`, "", -1)
	conf := Conf(source)
	if err := conf.Validate(); err != nil {
		return err
	}
	*c = conf
	return nil
}

const (
	New    Conf = "new"
	Nm          = "nm"
	Exact       = "exact"
	Junk        = "junk"
	Ha          = "ha"
	Det         = "det"
	Reinit      = "reinit"
)

func All() []Conf {
	return []Conf{Nm, New, Exact, Ha, Junk, Det, Reinit}
}

func Point(c Conf) *Conf {
	return &c
}
