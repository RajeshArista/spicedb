package postgres

import (
	"github.com/authzed/spicedb/internal/datastore"
	ds "github.com/authzed/spicedb/pkg/cmd/datastore"

	"github.com/authzed/spicedb/internal/datastore/postgres"
)

type drv struct{}

func (d *drv) Open(opts ds.Config) (datastore.Datastore, error) {
	pgOpts := []postgres.Option{
		postgres.GCWindow(opts.GCWindow),
		postgres.RevisionFuzzingTimedelta(opts.RevisionQuantization),
		postgres.ConnMaxIdleTime(opts.MaxIdleTime),
		postgres.ConnMaxLifetime(opts.MaxLifetime),
		postgres.MaxOpenConns(opts.MaxOpenConns),
		postgres.MinOpenConns(opts.MinOpenConns),
		postgres.SplitAtUsersetCount(opts.SplitQueryCount),
		postgres.HealthCheckPeriod(opts.HealthCheckPeriod),
		postgres.GCInterval(opts.GCInterval),
		postgres.GCMaxOperationTime(opts.GCMaxOperationTime),
		postgres.EnableTracing(),
		postgres.WatchBufferLength(opts.WatchBufferLength),
	}
	if opts.EnableDatastoreMetrics {
		pgOpts = append(pgOpts, postgres.EnablePrometheusStats())
	}
	return postgres.NewPostgresDatastore(opts.URI, pgOpts...)
}

func init() {
	ds.Register("postgres", &drv{})
}
