package core

import (
	"encoding/json"
	"fmt"
)

type Profile struct {
	PrimaryStorage   Storage
	SecondaryStorage Storage
}

type storageConfig struct {
	Name    string
	Options interface{}
}

func (*Profile) UnmarshalJSON(b []byte) error {
	var config struct {
		PrimaryStorage   storageConfig
		SecondaryStorage storageConfig
	}
	json.Unmarshal(b, &config)
	fmt.Println(config.PrimaryStorage)
	fmt.Println(config.SecondaryStorage)
	return nil
}
