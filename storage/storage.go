package storage

import (
	"context"
)

//go:generate moq -out storage_moq_test.go . Firestore

type Firestore interface {
	Get(ctx context.Context, dataTo interface{}, documentID string) error
	Set(ctx context.Context, documentID string, data interface{}) error
}
