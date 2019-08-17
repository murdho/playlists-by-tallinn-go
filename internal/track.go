package internal

import "fmt"

type Track struct {
	Name     string
	Persists bool
}

func NewTrack(name string, persists bool) *Track {
	return &Track{
		Name:     name,
		Persists: persists,
	}
}

func (t *Track) String() string {
	return fmt.Sprintf("<Track name=%#v persists=%t>", t.Name, t.Persists)
}
