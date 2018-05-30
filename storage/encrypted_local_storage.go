package storage

import (
	"github.com/katsuya94/memo/util"
)

type EncryptedLocalStorage struct {
	Path string
}

func (EncryptedLocalStorage) Get(d util.Date) (util.Memo, error) {
	return util.Memo{}, nil
}

func (EncryptedLocalStorage) Put(d util.Date, memo util.Memo) error {
	return nil
}

func (EncryptedLocalStorage) List() ([]util.Date, error) {
	return []util.Date{}, nil
}
