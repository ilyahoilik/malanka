package logger

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger, err := NewLogger()
	require.NoError(t, err)
	require.NotNil(t, logger)
}
