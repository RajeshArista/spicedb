package memdb

import (
	"github.com/authzed/spicedb/internal/datastore"
	"github.com/authzed/spicedb/internal/datastore/memdb"
	ds "github.com/authzed/spicedb/pkg/cmd/datastore"
	"github.com/rs/zerolog/log"
)

type drv struct{}

func (d *drv) Open(opts ds.Config) (datastore.Datastore, error) {
	log.Warn().Msg("in-memory datastore is not persistent and not feasible to run in a high availability fashion")
	return memdb.NewMemdbDatastore(opts.WatchBufferLength, opts.RevisionQuantization, opts.GCWindow, 0)
}

func init() {
	ds.Register("memory", &drv{})
}
