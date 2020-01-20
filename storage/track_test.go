package storage

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/murdho/playlists-by-tallinn/firestore"
	"github.com/murdho/playlists-by-tallinn/track"
)

func TestNewTrack(t *testing.T) {
	fs := &FirestoreMock{}
	trackStorage := NewTrack(fs)

	if trackStorage.firestore != fs {
		t.Errorf("track firestore:\ngot  %+v\nwant %+v", trackStorage.firestore, fs)
	}
}

func TestTrackStorage_Load(t *testing.T) {
	var gotDocumentID string
	fs := &FirestoreMock{
		GetFunc: func(_ context.Context, dataTo interface{}, documentID string, _ ...firestore.Option) error {
			dataTo.(*track.Track).Name = "b"
			gotDocumentID = documentID
			return nil
		},
	}

	trackStorage := NewTrack(fs)
	trk, err := trackStorage.Load(context.Background(), "a")
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
	}

	expectedDocumentID := documentID("a")
	if gotDocumentID != expectedDocumentID {
		t.Errorf("track load document ID:\ngot  %+v\nwant %+v", gotDocumentID, expectedDocumentID)
	}

	expectedTrack := &track.Track{Name: "b"}
	if !reflect.DeepEqual(trk, expectedTrack) {
		t.Errorf("track load:\ngot  %+v\nwant %+v", trk, expectedTrack)
	}
}

func TestTrackStorage_LoadErrNotFound(t *testing.T) {
	fs := &FirestoreMock{
		GetFunc: func(_ context.Context, _ interface{}, _ string, _ ...firestore.Option) error {
			return status.Error(codes.NotFound, "")
		},
	}

	trackStorage := NewTrack(fs)
	trk, err := trackStorage.Load(context.Background(), "a")
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
	}

	if trk != nil {
		t.Errorf("track load:\ngot  %+v\nwant %+v", trk, nil)
	}
}

func TestTrackStorage_LoadErr(t *testing.T) {
	expectedErr := errors.New("x")
	fs := &FirestoreMock{
		GetFunc: func(_ context.Context, _ interface{}, _ string, _ ...firestore.Option) error {
			return expectedErr
		},
	}

	trackStorage := NewTrack(fs)
	trk, err := trackStorage.Load(context.Background(), "a")

	if !errors.Is(err, expectedErr) {
		t.Errorf("track load err:\ngot  %+v\nwant %+v", err, expectedErr)
	}

	if trk != nil {
		t.Errorf("track load:\ngot  %+v\nwant %+v", trk, nil)
	}
}

func TestTrackStorage_Save(t *testing.T) {
	var gotDocumentID string
	var gotTrack track.Track
	fs := &FirestoreMock{
		SetFunc: func(_ context.Context, documentID string, data interface{}, _ ...firestore.Option) error {
			gotDocumentID = documentID
			gotTrack = data.(track.Track)
			return nil
		},
	}

	trk := track.New("a")

	trackStorage := NewTrack(fs)
	err := trackStorage.Save(context.Background(), trk)
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
	}

	expectedDocumentID := documentID("a")
	if gotDocumentID != expectedDocumentID {
		t.Errorf("track save document ID:\ngot  %+v\nwant %+v", gotDocumentID, expectedDocumentID)
	}

	if gotTrack != trk {
		t.Errorf("track save track:\ngot  %+v\nwant %+v", gotTrack, trk)
	}
}

func TestTrackStorage_SaveErr(t *testing.T) {
	expectedErr := errors.New("x")
	fs := &FirestoreMock{
		SetFunc: func(_ context.Context, _ string, _ interface{}, _ ...firestore.Option) error {
			return expectedErr
		},
	}

	trackStorage := NewTrack(fs)
	err := trackStorage.Save(context.Background(), track.New("a"))

	if !errors.Is(err, expectedErr) {
		t.Errorf("track save err:\ngot  %+v\nwant %+v", err, expectedErr)
	}
}

func TestDocumentID(t *testing.T) {
	name := "a"
	expectedDocumentID := "0cc175b9c0f1b6a831c399e269772661"

	docID := documentID(name)
	if docID != expectedDocumentID {
		t.Errorf("documentID:\ngot  %+v\nwant %+v", docID, expectedDocumentID)
	}
}
