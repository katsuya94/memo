package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path"

	"github.com/katsuya94/memo/core"
	"github.com/katsuya94/memo/storage"
	"github.com/spf13/cobra"
)

func Setup(cmd *cobra.Command, args []string) {
	if err := setup(); err != nil {
		fmt.Fprintf(os.Stderr, "memo: %v\n", err)
		os.Exit(1)
	}
}

func setup() error {
	if err := setHome(); err != nil {
		return err
	}

	if err := setProfile(); err != nil {
		return err
	}

	return nil
}

func setHome() error {
	if home := os.Getenv("MEMO_HOME"); home != "" {
		Home = home
		return nil
	}

	usr, err := user.Current()
	if err != nil {
		return err
	}

	Home = path.Join(usr.HomeDir, ".memo")

	return nil
}

type storageConfig struct {
	Kind    string
	Options interface{}
}

type profileConfig struct {
	PrimaryStorage   *storageConfig
	SecondaryStorage *storageConfig
}

type memoConfig struct {
	DefaultProfile string
	Profiles       map[string]profileConfig
}

func setProfile() error {
	var config = memoConfig{
		DefaultProfile: "default",
		Profiles: map[string]profileConfig{
			"default": profileConfig{
				PrimaryStorage: &storageConfig{
					Kind: "LocalStorage",
				},
			},
		},
	}

	configPath := path.Join(Home, "config.json")

	if f, err := os.Open(configPath); err == nil {
		decoder := json.NewDecoder(f)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&config); err != nil {
			return fmt.Errorf("%v: %v", configPath, err)
		}
	} else if !os.IsNotExist(err) {
		return err
	}

	if ProfileName == "" {
		ProfileName = config.DefaultProfile
	}

	electedProfileConfig, ok := config.Profiles[ProfileName]
	if !ok {
		return fmt.Errorf("%v: profile `%v` not found", configPath, ProfileName)
	}

	return configureProfile(electedProfileConfig)
}

func configureProfile(config profileConfig) error {
	if config.PrimaryStorage != nil {
		storage, err := createStorage(*config.PrimaryStorage)
		if err != nil {
			return fmt.Errorf("primaryStorage: %v", err)
		}
		Profile.PrimaryStorage = storage
	}

	if config.SecondaryStorage != nil {
		storage, err := createStorage(*config.SecondaryStorage)
		if err != nil {
			return fmt.Errorf("secondaryStorage: %v", err)
		}
		Profile.SecondaryStorage = storage
	}

	return nil
}

func createStorage(config storageConfig) (core.Storage, error) {
	b, err := json.Marshal(config.Options)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(b)
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()

	switch config.Kind {
	case "LocalStorage":
		var options struct {
			Path string
		}
		decoder.Decode(&options)
		return createLocalStorage(options.Path), nil
	case "EncryptedLocalStorage":
		var options struct {
			Path string
		}
		decoder.Decode(&options)
		return createEncryptedLocalStorage(options.Path), nil
	case "GoogleCloudStorage":
		var options struct{}
		decoder.Decode(&options)
		return createGoogleCloudStorage(), nil
	default:
		return nil, fmt.Errorf("kind `%v` not found", config.Kind)
	}
}

func normalizePath(p string) string {
	if p == "" {
		return path.Join(Home, ProfileName)
	} else if !path.IsAbs(p) {
		return path.Join(Home, p)
	} else {
		return p
	}
}

func createLocalStorage(p string) *storage.LocalStorage {
	return &storage.LocalStorage{
		Path: normalizePath(p),
	}
}

func createEncryptedLocalStorage(p string) *storage.EncryptedLocalStorage {
	return &storage.EncryptedLocalStorage{
		Path: normalizePath(p),
	}
}

func createGoogleCloudStorage() *storage.GoogleCloudStorage {
	return &storage.GoogleCloudStorage{}
}
