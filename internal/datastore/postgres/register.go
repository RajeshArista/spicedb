package postgres

import (
	"github.com/authzed/spicedb/internal/datastore"
)

type drv struct{}

func (d *drv) Open(opts datastore.Config) (datastore.Datastore, error) {
	pgOpts := []Option{
		GCWindow(opts.GCWindow),
		RevisionFuzzingTimedelta(opts.RevisionQuantization),
		ConnMaxIdleTime(opts.MaxIdleTime),
		ConnMaxLifetime(opts.MaxLifetime),
		MaxOpenConns(opts.MaxOpenConns),
		MinOpenConns(opts.MinOpenConns),
		SplitAtUsersetCount(opts.SplitQueryCount),
		HealthCheckPeriod(opts.HealthCheckPeriod),
		GCInterval(opts.GCInterval),
		GCMaxOperationTime(opts.GCMaxOperationTime),
		EnableTracing(),
		WatchBufferLength(opts.WatchBufferLength),
	}
	if opts.EnableDatastoreMetrics {
		pgOpts = append(pgOpts, EnablePrometheusStats())
	}
	return NewPostgresDatastore(opts.URI, pgOpts...)
}

func init() {
	datastore.Register("postgres", &drv{})
}
