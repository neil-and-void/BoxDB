package src

import (
	"github.com/neil-and-void/boxdb/src/memtable"
	"github.com/neil-and-void/boxdb/src/wal"
)

type BoxDB struct {
	options    Options
	memtable   memtable.Memtable
	walManager *wal.WalManager
}

func NewBoxDB(options Options) *BoxDB {
	skipList := memtable.NewMemTable(memtable.StringComparer{}, options.MaxLSMHeight)
	walManager := wal.NewWalManager(options.Path)

	return &BoxDB{options, skipList, walManager}
}

func (s BoxDB) Put(k, v string) error {
	// write to wal
	s.walManager.Write(k, v)
	return nil
}

func (s BoxDB) Get(k string) *string {
	val := s.memtable.Get(k)
	if val != nil {
		return val
	}

	// sst's

	return nil
}
