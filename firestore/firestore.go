package firestore

import (
	"context"
	"fmt"

	cloudfirestore "cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func New(ctx context.Context, opts ...Option) (*firestore, error) {
	fs := &firestore{}
	fs.applyOptions(opts)

	if err := fs.connect(ctx); err != nil {
		return nil, fmt.Errorf("connect: %w", err)
	}

	return fs, nil
}

type firestore struct {
	client     *cloudfirestore.Client
	projectID  string
	collection string
}

func (fs *firestore) applyOptions(opts []Option) {
	for _, opt := range opts {
		opt(fs)
	}
}

func (fs *firestore) connect(ctx context.Context) error {
	if fs.projectID == "" {
		return fmt.Errorf("projectID required")
	}

	client, err := cloudfirestore.NewClient(ctx, fs.projectID)
	if err != nil {
		return fmt.Errorf("new cloud firestore client: %w", err)
	}

	fs.client = client
	return nil
}

func (fs firestore) Get(ctx context.Context, dataTo interface{}, documentID string, opts ...Option) error {
	fs.applyOptions(opts)

	if fs.collection == "" {
		return fmt.Errorf("collection required")
	}

	snapshot, err := fs.client.Collection(fs.collection).Doc(documentID).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return err
		}

		return fmt.Errorf("get document: %w", err)
	}

	if err := snapshot.DataTo(dataTo); err != nil {
		return fmt.Errorf("load document data: %w", err)
	}

	return nil
}

func (fs firestore) Set(ctx context.Context, documentID string, data interface{}, opts ...Option) error {
	fs.applyOptions(opts)
	if fs.collection == "" {
		return fmt.Errorf("collection required")
	}

	_, err := fs.client.Collection(fs.collection).Doc(documentID).Set(ctx, data)
	if err != nil {
		return fmt.Errorf("save document data: %w", err)
	}

	return nil
}

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
