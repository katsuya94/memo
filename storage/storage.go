package storage

import (
	"github.com/katsuya94/memo/util"
)

type Storage interface {
	Get(util.Date) (util.Memo, error)
	Put(util.Date, util.Memo) error
}
