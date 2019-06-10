package moods

import (
	"fmt"
)

type Moods []Mood

func (ms Moods) Validate() error {
	for _, m := range ms {
		if err := m.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Mood string

func (m Mood) String() string {
	return string(m)
}

func (m Mood) Validate() error {
	switch m {
	case Neutral, Anger, Contempt, Disgust, Fear, Happiness, Sadness, Surprise:
		return nil
	default:
		return fmt.Errorf("unknown Mood %s", m.String())
	}
}

const (
	Neutral   Mood = "neutral"
	Anger          = "anger"
	Contempt       = "contempt"
	Disgust        = "disgust"
	Fear           = "fear"
	Happiness      = "happiness"
	Sadness        = "sadness"
	Surprise       = "surprise"
)

func All() []Mood {
	return []Mood{Neutral, Anger, Contempt, Disgust, Fear, Happiness, Sadness, Surprise}
}
