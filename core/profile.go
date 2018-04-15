package core

import (
	"github.com/katsuya94/memo/storage"
	"github.com/katsuya94/memo/util"
)

type Profile struct {
	PrimaryStorage   storage.Storage
	SecondaryStorage storage.Storage
}

func (p Profile) Get(d util.Date) (util.Memo, error) {
	return p.PrimaryStorage.Get(d)
}

func (p Profile) Put(d util.Date, memo util.Memo) error {
	return p.PrimaryStorage.Put(d, memo)
}
