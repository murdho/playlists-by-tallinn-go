package playlistsbytallinn

import (
	"cloud.google.com/go/firestore"
	"context"
	"crypto/md5"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewFirestoreTrackHandler(client *firestore.Client, projectID string, collectionPath string) *firestoreTrackHandler {
	return &firestoreTrackHandler{
		client:         client,
		projectID:      projectID,
		collectionPath: collectionPath,
	}
}

type firestoreTrackHandler struct {
	client         *firestore.Client
	projectID      string
	collectionPath string
}

func (f *firestoreTrackHandler) LoadTrack(ctx context.Context, trackName string) (*Track, error) {
	track := Track{
		Name: trackName,
	}

	snapshot, err := f.getDocument(trackName).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return &track, nil
		}

		return nil, errors.Wrap(err, "getting track data snapshot failed")
	}

	if err := snapshot.DataTo(&track); err != nil {
		return nil, errors.Wrap(err, "building track from snapshot failed")
	}

	return &track, nil
}

func (f *firestoreTrackHandler) SaveTrack(ctx context.Context, track *Track) error {
	if _, err := f.getDocument(track.Name).Set(ctx, track); err != nil {
		return errors.Wrap(err, "saving track failed")
	}

	return nil
}

func (f *firestoreTrackHandler) getDocument(trackName string) *firestore.DocumentRef {
	documentID := firestoreDocumentID(trackName)
	return f.client.Collection(f.collectionPath).Doc(documentID)
}

func firestoreDocumentID(trackName string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(trackName)))
}
