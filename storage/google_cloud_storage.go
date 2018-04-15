package storage

import (
	"github.com/katsuya94/memo/core"
	"github.com/katsuya94/memo/util"
)

type GoogleCloudStorage struct{}

func (GoogleCloudStorage) Get(d util.Date) (core.Memo, error) {
	return core.Memo{}, nil
}

func (GoogleCloudStorage) Put(d util.Date, memo core.Memo) error {
	return nil
}
