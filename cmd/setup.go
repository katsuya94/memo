package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path"

	"github.com/katsuya94/memo/core"
	"github.com/spf13/cobra"
)

func Setup(cmd *cobra.Command, args []string) {
	if err := setup(); err != nil {
		fmt.Fprintf(os.Stderr, "memo: %v\n", err)
		os.Exit(1)
	}
}

func setup() error {
	if err := SetHome(); err != nil {
		return err
	}

	if err := SetProfile(); err != nil {
		return err
	}

	return nil
}

func SetHome() error {
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

func SetProfile() error {
	configPath := path.Join(Home, "config.json")

	f, err := os.Open(configPath)
	if err != nil {
		return err
	}

	var config map[string]core.Profile

	decoder := json.NewDecoder(f)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&config); err != nil {
		return fmt.Errorf("%v: %v", configPath, err)
	}
	f.Close()

	profile, ok := config[ProfileName]
	if !ok {
		return fmt.Errorf("profile `%v` not found", ProfileName)
	}

	Profile = profile

	return nil
}
