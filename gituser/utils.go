package gituser

import (
	"os"
	"path"
)

func existFile(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}

func existDir(fileDir string) bool {
	if f, err := os.Stat(fileDir); f != nil &&
		err == nil && os.IsExist(err) {
		return f.IsDir()
	}
	return false
}

func isGitRepoDir(filePath string) bool {
	f, err := os.Stat(filePath)
	if f == nil {
	} else if err != nil && !os.IsExist(err) {
	} else if f.IsDir() {
		return existFile(path.Join(filePath, ".git"))
	}
	return false
}
