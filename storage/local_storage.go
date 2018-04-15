package storage

import (
	"github.com/katsuya94/memo/util"
)

type LocalStorage struct {
	Path string
}

func (LocalStorage) Get(d util.Date) (util.Memo, error) {
	return util.Memo{}, nil
}

func (LocalStorage) Put(d util.Date, memo util.Memo) error {
	return nil
}
