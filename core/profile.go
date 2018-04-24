package core

import (
	"github.com/katsuya94/memo/storage"
	"github.com/katsuya94/memo/util"
)

type Profile struct {
	PrimaryStorage   storage.Storage
	SecondaryStorage storage.Storage
}

func (p Profile) GetLast() (util.Memo, error) {
	dates, err := p.PrimaryStorage.List()
	if err != nil {
		return nil, err
	}
	return p.Get(dates[len(dates)-1])
}

func (p Profile) Get(d util.Date) (util.Memo, error) {
	return p.PrimaryStorage.Get(d)
}

func (p Profile) Put(d util.Date, memo util.Memo) error {
	return p.PrimaryStorage.Put(d, memo)
}
