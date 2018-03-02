package core

import (
	"encoding/json"
	"fmt"
)

type Profile struct {
	PrimaryStorage   Storage
	SecondaryStorage Storage
}

func (*Profile) UnmarshalJSON(b []byte) error {
	type storageConfig struct {
		Type    string
		Options interface{}
	}

	type profileConfig struct {
		PrimaryStorage   storageConfig
		SecondaryStorage storageConfig
	}

	var config profileConfig
	json.Unmarshal(b, &config)
	fmt.Println(config.PrimaryStorage)
	fmt.Println(config.SecondaryStorage)
	return nil
}
