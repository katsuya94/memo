package storage

import (
	"github.com/katsuya94/memo/core"
	"github.com/katsuya94/memo/util"
)

type EncryptedLocalStorage struct {
	Path string
}

func (EncryptedLocalStorage) Get(d util.Date) (core.Memo, error) {
	return core.Memo{}, nil
}

func (EncryptedLocalStorage) Put(d util.Date, memo core.Memo) error {
	return nil
}
