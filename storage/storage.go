package storage

import (
	"context"

	"github.com/murdho/playlists-by-tallinn/firestore"
)

type Firestore interface {
	Get(ctx context.Context, dataTo interface{}, documentID string, opts ...firestore.Option) error
	Set(ctx context.Context, documentID string, data interface{}, opts ...firestore.Option) error
}
