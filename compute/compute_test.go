package compute

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
	"strings"
	"testing"
)

var (
	logger *zap.Logger
	logs   *observer.ObservedLogs
)

func TestMain(m *testing.M) {
	var core zapcore.Core
	core, logs = observer.New(zap.InfoLevel)
	logger = zap.New(core)
	m.Run()
}

func TestNewCompute(t *testing.T) {
	compute := NewCompute(zap.NewExample())
	require.NotNil(t, compute)
}

func TestParser(t *testing.T) {
	cases := [][]string{
		{"SET", "key value"},
		{"SET", "/etc/nginx/config some-information"},
		{"SET", "!,#*(&$%@^ (*)&($^(&#!"},
		{"GET", "key"},
		{"GET", "/etc/nginx/config"},
		{"GET", "!,#*(&$%@^"},
		{"DEL", "key"},
		{"DEL", "/etc/nginx/config"},
		{"DEL", "!,#*(&$%@^"},
	}

	compute := NewCompute(zap.NewExample())

	for _, c := range cases {
		query, err := compute.Parse(strings.Join(c, " "))
		require.NoError(t, err)
		assert.Equal(t, strings.ToUpper(c[0]), string(query.Command))
		assert.Equal(t, c[1], strings.Join(query.Args, " "))
	}
}

func TestParserWithIncorrectNumberOfArguments(t *testing.T) {
	cases := [][]string{
		{"SET", ""},
		{"SET", "one"},
		{"SET", "one two three"},
		{"GET", ""},
		{"GET", "one two"},
		{"GET", "one two three"},
		{"DEL", ""},
		{"DEL", "one two"},
		{"DEL", "one two three"},
	}

	compute := NewCompute(zap.NewExample())

	for _, c := range cases {
		_, err := compute.Parse(strings.Join(c, " "))
		require.Error(t, err)
	}
}
