package storage

import (
	"github.com/katsuya94/memo/core"
	"github.com/katsuya94/memo/util"
)

type LocalStorage struct {
	Path string
}

func (LocalStorage) Get(d util.Date) (core.Memo, error) {
	return core.Memo{}, nil
}

func (LocalStorage) Put(d util.Date, memo core.Memo) error {
	return nil
}
