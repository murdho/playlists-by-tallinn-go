package track

func New(name string) Track {
	return Track{
		Name: name,
	}
}

type Track struct {
	Name string `firestore:"name"`
}
