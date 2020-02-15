package testhelpers

import (
	"testing"

	"github.com/Flaque/filet"
)

// WithTmpDir creates a tmp directory for a test
func WithTmpDir(t *testing.T, test func(string)) {
	dir := filet.TmpDir(t, "")
	defer filet.CleanUp(t)
	test(dir)
}
