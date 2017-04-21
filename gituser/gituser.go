package gituser

import (
	"encoding/json"
	"github.com/btcsuite/goleveldb/leveldb/errors"
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
	fOut.WriteString("[]")
	if err := fOut.Close(); err != nil {
		return err
	}
	return nil
}

func writeConfig(gitUser *_GitUser) error {
	if gitUser == nil {
		return errors.New("gitUser must not be empty!")
	}
	var gitUsers []*_GitUser
	if existFile(_CONFIG_FILE_PATH) {
		var err error
		if gitUsers, err = readConfig(); err != nil {
			return err
		}
	}
	fOut, err := os.Create(_CONFIG_FILE_PATH)
	if err != nil {
		return err
	}
	defer fOut.Close()

	for i, u := range gitUsers {
		if u.user == gitUser.user {
			gitUsers[i] = gitUser
			goto Write
		}
	}
	gitUsers = append(gitUsers, gitUser)

Write:
	data, err := json.Marshal(gitUsers)
	if err != nil {
		return err
	}

	if _, err := fOut.Write(data); err != nil {
		return err
	}

	return nil
}

func readConfig() ([]*_GitUser, error) {
	fIn, err := os.Open(_CONFIG_FILE_PATH)
	if err != nil {
		return nil, err
	}
	defer fIn.Close()

	data := make([]byte, 1024)
	for true {
		if n, _ := fIn.Read(data); 0 <= n {
			break
		}
	}

	gitUsers := make([]*_GitUser, 0, 10)
	if err := json.Unmarshal(data, &gitUsers); err != nil {
		return nil, err
	}
	return gitUsers, nil
}
