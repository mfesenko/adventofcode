package input

import (
	"math/rand"
	"path"
	"strings"
	"testing"

	"github.com/Flaque/filet"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestWhenFileDoesNotExistThenLoadFromFileReturnsAnError(t *testing.T) {
	withTmpDir(t, func(dir string) {
		filePath := path.Join(dir, randomString())
		assert.False(t, filet.Exists(t, filePath))

		content, err := LoadFromFile(filePath)

		assert.Error(t, err)
		assert.Nil(t, content)
	})
}

func TestWhenPathPointsToADirectoryThenLoadFromFileReturnsAnError(t *testing.T) {
	withTmpDir(t, func(dir string) {
		content, err := LoadFromFile(dir)

		assert.Error(t, err)
		assert.Nil(t, content)
	})
}

func TestWhenFileExistsThenLoadFromFileReturnsContentOfTheFile(t *testing.T) {
	withTmpDir(t, func(dir string) {
		testContent := randomContent()
		file := filet.TmpFile(t, dir, strings.Join(testContent, "\n"))

		content, err := LoadFromFile(file.Name())

		assert.NoError(t, err)
		assert.Equal(t, testContent, content)
	})
}

func withTmpDir(t *testing.T, test func(string)) {
	dir := filet.TmpDir(t, "")
	defer filet.CleanUp(t)
	test(dir)
}

func randomContent() []string {
	lineCount := rand.Intn(5) + 1
	lines := make([]string, lineCount)
	for i := 0; i < lineCount; i++ {
		lines[i] = randomString()
	}
	return lines
}

func randomString() string {
	return uuid.New().String()
}
