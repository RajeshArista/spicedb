package crdb

import (
	"github.com/authzed/spicedb/internal/datastore"
	ds "github.com/authzed/spicedb/pkg/cmd/datastore"

	"github.com/authzed/spicedb/internal/datastore/crdb"
)

type drv struct{}

func (d *drv) Open(opts ds.Config) (datastore.Datastore, error) {
	return crdb.NewCRDBDatastore(
		opts.URI,
		crdb.GCWindow(opts.GCWindow),
		crdb.RevisionQuantization(opts.RevisionQuantization),
		crdb.ConnMaxIdleTime(opts.MaxIdleTime),
		crdb.ConnMaxLifetime(opts.MaxLifetime),
		crdb.MaxOpenConns(opts.MaxOpenConns),
		crdb.MinOpenConns(opts.MinOpenConns),
		crdb.SplitAtUsersetCount(opts.SplitQueryCount),
		crdb.FollowerReadDelay(opts.FollowerReadDelay),
		crdb.MaxRetries(opts.MaxRetries),
		crdb.OverlapKey(opts.OverlapKey),
		crdb.OverlapStrategy(opts.OverlapStrategy),
		crdb.WatchBufferLength(opts.WatchBufferLength),
	)
}

func init() {
	ds.Register("cockroachdb", &drv{})
}
