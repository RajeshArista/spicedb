package crdb

import "github.com/authzed/spicedb/internal/datastore"

type drv struct{}

func (d *drv) Open(opts datastore.Config) (datastore.Datastore, error) {
	return NewCRDBDatastore(
		opts.URI,
		GCWindow(opts.GCWindow),
		RevisionQuantization(opts.RevisionQuantization),
		ConnMaxIdleTime(opts.MaxIdleTime),
		ConnMaxLifetime(opts.MaxLifetime),
		MaxOpenConns(opts.MaxOpenConns),
		MinOpenConns(opts.MinOpenConns),
		SplitAtUsersetCount(opts.SplitQueryCount),
		FollowerReadDelay(opts.FollowerReadDelay),
		MaxRetries(opts.MaxRetries),
		OverlapKey(opts.OverlapKey),
		OverlapStrategy(opts.OverlapStrategy),
		WatchBufferLength(opts.WatchBufferLength),
	)
}

func init() {
	datastore.Register("cockroachdb", &drv{})
}
