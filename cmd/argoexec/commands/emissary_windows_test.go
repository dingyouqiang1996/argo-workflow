//go:build windows

package commands

import (
	"os"
	"testing"

	"github.com/argoproj/argo-workflows/v3/util/errors"
	"github.com/stretchr/testify/require"
)

func TestEmissary(t *testing.T) {
	tmp := t.TempDir()

	varRunArgo = tmp
	includeScriptOutput = true

	err := os.WriteFile(varRunArgo+"/template", []byte(`{}`), 0o600)
	require.NoError(t, err)

	t.Run("Exit0", func(t *testing.T) {
		err := run("exit")
		require.NoError(t, err)
		data, err := os.ReadFile(varRunArgo + "/ctr/main/exitcode")
		require.NoError(t, err)
		require.Equal(t, "0", string(data))
	})

	t.Run("Exit1", func(t *testing.T) {
		err := run("exit 1")
		require.Equal(t, 1, err.(errors.Exited).ExitCode())
		data, err := os.ReadFile(varRunArgo + "/ctr/main/exitcode")
		require.NoError(t, err)
		require.Equal(t, "1", string(data))
	})
	t.Run("Exit13", func(t *testing.T) {
		err := run("exit 13")
		require.Equal(t, 13, err.(errors.Exited).ExitCode())
		require.EqualError(t, err, "exit status 13")
	})
}

func run(script string) error {
	cmd := NewEmissaryCommand()
	containerName = "main"
	return cmd.RunE(cmd, append([]string{"powershell", "-c"}, script))
}
