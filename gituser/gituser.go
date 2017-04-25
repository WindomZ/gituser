package gituser

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

var (
	_CONFIG_DIR       string = ".gituser"
	_CONFIG_FILE             = "config"
	_CONFIG_FILE_PATH        = "~/.gituser/config"
)

func init() {
	home, err := user.Current()
	if err != nil {
		panic(err)
	}
	_CONFIG_DIR = path.Join(home.HomeDir, _CONFIG_DIR)
	_CONFIG_FILE_PATH = path.Join(_CONFIG_DIR, _CONFIG_FILE)

	if existFile(_CONFIG_FILE_PATH) {
		return
	}

	if existDir(_CONFIG_DIR) {
	} else if err := os.MkdirAll(_CONFIG_DIR, os.ModePerm); err != nil {
		panic(err)
	}

	if err := writeEmptyConfig(); err != nil {
		panic(err)
	}
}

func writeEmptyConfig() error {
	fOut, err := os.Create(_CONFIG_FILE_PATH)
	if err != nil {
		return err
	}
	fOut.WriteString("")
	if err := fOut.Close(); err != nil {
		return err
	}
	return nil
}

func writeConfig(gitUser *_GitUser) error {
	if gitUser == nil {
		return errors.New("gitUser must not be empty!")
	}

	gitUsers, err := readConfig()
	if err != nil {
		return err
	}

	return writeConfigs(gitUsers.Add(gitUser))
}

func writeConfigs(gitUsers *_GitUsers) error {
	fOut, err := os.Create(_CONFIG_FILE_PATH)
	if err != nil {
		return err
	}
	defer fOut.Close()

	data, err := json.Marshal(gitUsers)
	if err != nil {
		return err
	}

	if _, err := fOut.Write(data); err != nil {
		return err
	}

	return nil
}

func readConfig() (*_GitUsers, error) {
	data, err := ioutil.ReadFile(_CONFIG_FILE_PATH)
	if err != nil {
		return nil, err
	} else if len(data) == 0 {
		return &_GitUsers{}, nil
	}

	var gitUsers _GitUsers
	if err := json.Unmarshal(data, &gitUsers); err != nil {
		return nil, err
	}

	return &gitUsers, nil
}
