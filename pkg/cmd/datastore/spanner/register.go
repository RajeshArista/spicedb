package spanner

import (
	"github.com/authzed/spicedb/internal/datastore"
	ds "github.com/authzed/spicedb/pkg/cmd/datastore"

	"github.com/authzed/spicedb/internal/datastore/spanner"
)

type drv struct{}

func (d *drv) Open(opts ds.Config) (datastore.Datastore, error) {
	return spanner.NewSpannerDatastore(
		opts.URI,
		spanner.FollowerReadDelay(opts.FollowerReadDelay),
		spanner.GCInterval(opts.GCInterval),
		spanner.GCWindow(opts.GCWindow),
		spanner.CredentialsFile(opts.SpannerCredentialsFile),
		spanner.WatchBufferLength(opts.WatchBufferLength),
	)
}

func init() {
	ds.Register("spanner", &drv{})
}
