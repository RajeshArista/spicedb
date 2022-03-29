package datastore

import (
	"fmt"
	"sync"

	"github.com/authzed/spicedb/internal/datastore"
)

var (
	driversMu sync.RWMutex
	drivers   = make(map[string]Driver)
)

type Driver interface {
	Open(opts Config) (datastore.Datastore, error)
}

func Register(name string, driver Driver) error {
	driversMu.Lock()
	defer driversMu.Unlock()
	if driver == nil {
		return fmt.Errorf("driver is nil")
	}
	if _, dup := drivers[name]; dup {
		return fmt.Errorf("duplicate driver: %s", name)
	}
	drivers[name] = driver
	return nil
}

// Open opens a database specified by its database driver name and options
func Open(driverName string, opts Config) (datastore.Datastore,
	error) {
	driversMu.RLock()
	d, ok := drivers[driverName]
	driversMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf(
			"unknown driver %s (forgotten import)?", driverName)
	}
	return d.Open(opts)
}
