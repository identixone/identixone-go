package sex

type Sex uint8

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
