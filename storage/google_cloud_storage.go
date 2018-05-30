package storage

import (
	"github.com/katsuya94/memo/util"
)

type GoogleCloudStorage struct{}

func (GoogleCloudStorage) Get(d util.Date) (util.Memo, error) {
	return util.Memo{}, nil
}

func (GoogleCloudStorage) Put(d util.Date, memo util.Memo) error {
	return nil
}

func (GoogleCloudStorage) List() ([]util.Date, error) {
	return []util.Date{}, nil
}
