package testhelpers

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithTmpDirCreatesATmpDirectoryForTheTestAndCleansItUp(t *testing.T) {
	var dirName string
	WithTmpDir(t, func(tmpDir string) {
		dirName = tmpDir
		checkDirExists(t, dirName)
	})
	checkDirDoesNotExist(t, dirName)
}

func checkDirExists(t *testing.T, dirName string) {
	dir, err := os.Stat(dirName)
	assert.NoError(t, err)
	assert.True(t, dir.IsDir())
}

func checkDirDoesNotExist(t *testing.T, dirName string) {
	_, err := os.Stat(dirName)
	assert.True(t, os.IsNotExist(err))
}
