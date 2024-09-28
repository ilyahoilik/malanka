package storage

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
	"testing"
)

var (
	st   *Storage
	logs *observer.ObservedLogs
)

func TestMain(m *testing.M) {
	var core zapcore.Core
	core, logs = observer.New(zap.InfoLevel)
	st = NewStorage(zap.New(core))
	m.Run()
}

func TestNewStorage(t *testing.T) {
	require.NotNil(t, st)
	assert.NotNil(t, st.engine)
	assert.NotNil(t, st.logger)
	assert.Equal(t, len(logs.TakeAll()), 1)
}

func TestStorage_Add(t *testing.T) {
	err := st.Set("key", "value")
	require.NoError(t, err)
	assert.Equal(t, len(logs.TakeAll()), 1)

	value, err := st.Get("key")
	require.NoError(t, err)
	assert.Equal(t, len(logs.TakeAll()), 1)

	assert.Equal(t, value, "value")
}

func TestStorage_Get(t *testing.T) {
	err := st.Set("key", "value")
	require.NoError(t, err)
	assert.Equal(t, len(logs.TakeAll()), 1)

	retrieved, err := st.Get("key")
	require.NoError(t, err)
	assert.Equal(t, len(logs.TakeAll()), 1)

	assert.Equal(t, retrieved, "value")
}

func TestStorage_Del(t *testing.T) {
	_ = st.Set("key", "value")
	assert.Equal(t, len(logs.TakeAll()), 1)

	err := st.Del("key")
	require.NoError(t, err)
	assert.Equal(t, len(logs.TakeAll()), 1)

	retrieved, err := st.Get("key")
	require.Error(t, err)
	assert.Equal(t, len(logs.TakeAll()), 1)

	assert.Empty(t, retrieved)
}
