package compute

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidCommand(t *testing.T) {
	commands := []string{"SET", "GET", "DEL"}

	for _, cmd := range commands {
		command, ok := NewCommand(cmd, logger)

		require.NoError(t, ok)
		assert.Equal(t, command, Command(cmd))
		assert.Equal(t, 0, len(logs.TakeAll()))
	}
}

func TestInvalidCommand(t *testing.T) {
	commands := []string{"Set", "set", "Get", "get", "Del", "del"}

	for _, cmd := range commands {
		command, err := NewCommand(cmd, logger)

		require.Error(t, err)
		assert.Empty(t, command)
	}
}
