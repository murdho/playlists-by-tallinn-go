package internal

type Track struct {
	Name     string `firestore:"name"`
	Persists bool   `firestore:"persists"`
}

func NewTrack(name string, persists bool) *Track {
	return &Track{
		Name:     name,
		Persists: persists,
	}
}
