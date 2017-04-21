package gituser

import (
	"github.com/WindomZ/testify/assert"
	"path"
	"testing"
)

var testUser *_GitUser = &_GitUser{
	user:  "user",
	name:  "name",
	email: "email@email.com",
}

func Test_gituser_init(t *testing.T) {
	_CONFIG_FILE_PATH = path.Join(_CONFIG_DIR, _CONFIG_FILE+".test")
}

func Test_gituser_writeEmptyConfig(t *testing.T) {
	err := writeEmptyConfig()
	assert.NoError(t, err)
}

func Test_gituser_writeConfig(t *testing.T) {
	err := writeConfig(testUser)
	assert.NoError(t, err)
}

func Test_gituser_readConfig(t *testing.T) {
	users, err := readConfig()
	assert.NoError(t, err)
	assert.Equal(t, users, []*_GitUser{testUser})
}