package moods

import (
	"reflect"
	"testing"
)

func TestMood_String(t *testing.T) {
	tests := []struct {
		name string
		m    Mood
		want string
	}{
		{name: "string", m: Mood("neutral"), want: "neutral"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.String(); got != tt.want {
				t.Errorf("Mood.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMood_Validate(t *testing.T) {
	tests := []struct {
		name    string
		m       Mood
		wantErr bool
	}{
		{name: "valid", m: Mood("neutral"), wantErr: false},
		{name: "valid", m: Mood("packman"), wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Mood.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAll(t *testing.T) {
	all := []Mood{
		Neutral,
		Anger,
		Contempt,
		Disgust,
		Fear,
		Happiness,
		Sadness,
		Surprise,
	}
	tests := []struct {
		name string
		want []Mood
	}{
		{name: "valid", want: all},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoods_Validate(t *testing.T) {
	tests := []struct {
		name    string
		ms      Moods
		wantErr bool
	}{
		{name: "valid", ms: Moods{
			Mood("neutral"),
			Mood("anger"),
			Mood("contempt"),
			Mood("disgust"),
			Mood("fear"),
			Mood("happiness"),
			Mood("sadness"),
			Mood("surprise"),
		}, wantErr: false},
		{name: "invalid", ms: Moods{
			Mood("neutral"),
			Mood("anger"),
			Mood("contempt"),
			Mood("disgust"),
			Mood("fear"),
			Mood("happiness"),
			Mood("sadness"),
			Mood("surprise"),
			Mood("teapot"),
		}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ms.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Moods.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
