package storage

import (
	"github.com/katsuya94/memo/util"
)

type Storage interface {
	Get(util.Date) (Memo, error)
	Store(util.Date, Memo) error
}
