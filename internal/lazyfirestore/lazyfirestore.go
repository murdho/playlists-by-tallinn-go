package lazyfirestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/pkg/errors"
)

var client *firestore.Client

func NewClientFunc(projectID string) func() (*firestore.Client, error) {
	return func() (*firestore.Client, error) {
		return NewClient(projectID)
	}
}

func NewClient(projectID string) (*firestore.Client, error) {
	if client == nil {
		var err error

		client, err = firestore.NewClient(context.Background(), projectID)
		if err != nil {
			return nil, errors.Wrap(err, "initializing Firestore client failed")
		}
	}

	return client, nil
}
