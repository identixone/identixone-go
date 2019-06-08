package conf

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Conf string

func (c *Conf) String() string {
	return string(*c)
}

func (c *Conf) IsValid() error {
	for _, x := range All() {
		if x.String() == c.String() {
			return nil
		}
	}
	return fmt.Errorf("unknown type of Conf %s", c.String())
}

func (c *Conf) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(*c))
}

func (c *Conf) UnmarshalJSON(b []byte) error {
	source := string(b)
	source = strings.Replace(source, `"`, "", -1)
	conf := Conf(source)
	if err := conf.IsValid(); err != nil {
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
