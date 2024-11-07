package database

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
	"testing"
)

var (
	logger *zap.Logger
	logs   *observer.ObservedLogs
)

func TestMain(m *testing.M) {
	var core zapcore.Core
	core, logs = observer.New(zap.DebugLevel)
	logger = zap.New(core)
	m.Run()
}

func TestNewDatabase(t *testing.T) {
	db := NewDatabase(logger)

	require.NotNil(t, db)
	assert.NotNil(t, db.storage)
	assert.NotNil(t, db.logger)
	assert.NotZero(t, len(logs.TakeAll()))
}
