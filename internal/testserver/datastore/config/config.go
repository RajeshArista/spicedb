package config

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/authzed/spicedb/internal/datastore"
	testdatastore "github.com/authzed/spicedb/internal/testserver/datastore"
	dsconfig "github.com/authzed/spicedb/pkg/cmd/datastore"
)

// DatastoreConfigInitFunc returns a InitFunc that constructs a ds
// with the top-level cmd/datastore machinery.
// It can't be used everywhere due to import cycles, but makes it easy to write
// an independent test with CLI-like config where possible.
func DatastoreConfigInitFunc(t testing.TB, options ...dsconfig.ConfigOption) testdatastore.InitFunc {
	return func(engine, uri string) datastore.Datastore {
		ds, err := dsconfig.NewDatastore(append(options, dsconfig.WithEngine(engine), dsconfig.WithURI(uri))...)
		require.NoError(t, err)
		return ds
	}
}
