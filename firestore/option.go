package firestore

type Option func(*firestore)

func Project(project string) Option {
	return func(fs *firestore) {
		fs.projectID = project
	}
}

func Collection(collection string) Option {
	return func(fs *firestore) {
		fs.collection = collection
	}
}

