package spanner

import "github.com/authzed/spicedb/internal/datastore"

type drv struct{}

func (d *drv) Open(opts datastore.Config) (datastore.Datastore, error) {
	return NewSpannerDatastore(
		opts.URI,
		FollowerReadDelay(opts.FollowerReadDelay),
		GCInterval(opts.GCInterval),
		GCWindow(opts.GCWindow),
		CredentialsFile(opts.SpannerCredentialsFile),
		WatchBufferLength(opts.WatchBufferLength),
	)
}

func init() {
	datastore.Register("spanner", &drv{})
}
