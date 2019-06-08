package liveness

import (
	"fmt"
	"strings"
)

type Liveness string

const (
	Passed       Liveness = "passed"
	Failed                = "failed"
	Undetermined          = "undetermined"
)

func (l *Liveness) UnmarshalJSON(b []byte) error {
	source := string(b)
	source = strings.Replace(source, `"`, "", -1)
	live := Liveness(source)
	hasMatch := false
	for _, l := range All() {
		if l == live {
			hasMatch = true
		}
	}
	if !hasMatch && source != "" {
		return fmt.Errorf("unknown liveness %s", string(b))
	}
	*l = live
	return nil
}

func (l Liveness) IsLive() bool {
	switch l {
	case Passed:
		return true
	default:
		return false
	}
}

func All() []Liveness {
	return []Liveness{Passed, Failed, Undetermined}
}
