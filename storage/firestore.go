package storage

import (
	"cloud.google.com/go/firestore"
	"context"
	"crypto/md5"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewFirestoreClient(gcpProject, collectionName string) Client {
	return &firestoreClient{
		firestoreClientFunc: initLazyFirestoreClientFunc(gcpProject),
		collectionName:      collectionName,
	}
}

type firestoreClient struct {
	firestoreClientFunc firestoreClientFunc
	collectionName      string
}

func (f *firestoreClient) LoadTrack(ctx context.Context, trackName string) (*Track, error) {
	track := Track{
		Name:     trackName,
		Persists: false,
	}

	client, err := f.firestoreClientFunc()
	if err != nil {
		return nil, errors.Wrap(err, "getting Firestore client failed")
	}

	snapshot, err := f.getDocument(client, trackName).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return &track, nil
		}

		return nil, errors.Wrap(err, "getting track data snapshot failed")
	}

	if err := snapshot.DataTo(&track); err != nil {
		return nil, errors.Wrap(err, "building track from snapshot failed")
	}

	track.Persists = true

	return &track, nil
}

func (f *firestoreClient) SaveTrack(ctx context.Context, track *Track) error {
	client, err := f.firestoreClientFunc()
	if err != nil {
		return errors.Wrap(err, "getting Firestore client failed")
	}

	if _, err := f.getDocument(client, track.Name).Set(ctx, track); err != nil {
		return errors.Wrap(err, "saving track failed")
	}

	return nil

}

func (f *firestoreClient) getDocument(client *firestore.Client, trackName string) *firestore.DocumentRef {
	documentID := firestoreDocumentID(trackName)
	return client.Collection(f.collectionName).Doc(documentID)
}

func firestoreDocumentID(trackName string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(trackName)))
}

type firestoreClientFunc func() (*firestore.Client, error)

var lazyFirestoreClient *firestore.Client

func initLazyFirestoreClientFunc(gcpProject string) firestoreClientFunc {
	return func() (*firestore.Client, error) {
		if lazyFirestoreClient == nil {
			var err error
			lazyFirestoreClient, err = firestore.NewClient(context.Background(), gcpProject)
			if err != nil {
				return nil, errors.Wrap(err, "initializing Firestore client failed")
			}
		}

		return lazyFirestoreClient, nil
	}
}
