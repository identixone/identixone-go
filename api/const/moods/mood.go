package moods

import (
	"fmt"
)

type Mood string

func (m Mood) String() string {
	return string(m)
}

func (m Mood) IsValid() error {
	for _, x := range All() {
		if x.String() == m.String() {
			return nil
		}
	}
	return fmt.Errorf("unknown Mood %s", m.String())
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
	return []Mood{
		Neutral,
		Anger,
		Contempt,
		Disgust,
		Fear,
		Happiness,
		Sadness,
		Surprise,
	}
}
