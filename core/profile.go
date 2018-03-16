package core

import (
	"fmt"

	"github.com/katsuya94/memo/util"
)

type Profile struct {
	PrimaryStorage   Storage
	SecondaryStorage Storage
}

func (p Profile) Open(d util.Date) error {
	fmt.Println(d)
	return nil
}
