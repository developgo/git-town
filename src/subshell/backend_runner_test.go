package subshell_test

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/git-town/git-town/v12/src/gohacks"
	"github.com/git-town/git-town/v12/src/subshell"
	"github.com/shoenig/test/must"
)

func TestBackendRunner(t *testing.T) {
	t.Parallel()

	t.Run("Query", func(t *testing.T) {
		t.Parallel()
		t.Run("happy path", func(t *testing.T) {
			tmpDir := t.TempDir()
			runner := subshell.BackendRunner{Dir: &tmpDir, Verbose: false, CommandsCounter: &gohacks.Counter{}}
			output, err := runner.Query("echo", "hello", "world  ")
			must.NoError(t, err)
			must.EqOp(t, "hello world  \n", output)
		})

		t.Run("unknown executable", func(t *testing.T) {
			t.Parallel()
			tmpDir := t.TempDir()
			runner := subshell.BackendRunner{Dir: &tmpDir, Verbose: false, CommandsCounter: &gohacks.Counter{}}
			err := runner.Run("zonk")
			must.Error(t, err)
			var execError *exec.Error
			must.True(t, errors.As(err, &execError))
		})

		t.Run("non-zero exit code", func(t *testing.T) {
			t.Parallel()
			tmpDir := t.TempDir()
			runner := subshell.BackendRunner{Dir: &tmpDir, Verbose: false, CommandsCounter: &gohacks.Counter{}}
			err := runner.Run("bash", "-c", "echo hi && exit 2")
			expectedError := `
----------------------------------------
Diagnostic information of failed command

COMMAND: bash -c echo hi && exit 2
ERROR: exit status 2
OUTPUT START
hi

OUTPUT END
----------------------------------------`
			must.EqOp(t, expectedError, err.Error())
		})
	})

	t.Run("QueryTrim", func(t *testing.T) {
		t.Parallel()
		t.Run("trims whitespace", func(t *testing.T) {
			tmpDir := t.TempDir()
			runner := subshell.BackendRunner{Dir: &tmpDir, Verbose: false, CommandsCounter: &gohacks.Counter{}}
			output, err := runner.QueryTrim("echo", "hello", "world  ")
			must.NoError(t, err)
			must.EqOp(t, "hello world", output)
		})
	})

	t.Run("RunMany", func(t *testing.T) {
		t.Parallel()
		tmpDir := t.TempDir()
		runner := subshell.BackendRunner{Dir: &tmpDir, Verbose: false, CommandsCounter: &gohacks.Counter{}}
		err := runner.RunMany([][]string{
			{"mkdir", "tmp"},
			{"touch", "tmp/first"},
			{"touch", "tmp/second"},
		})
		must.NoError(t, err)
		entries, err := os.ReadDir(filepath.Join(tmpDir, "tmp"))
		must.NoError(t, err)
		must.EqOp(t, "first", entries[0].Name())
		must.EqOp(t, "second", entries[1].Name())
	})
}
