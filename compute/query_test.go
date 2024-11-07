package compute

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestNewQuery(t *testing.T) {
	cases := map[string]string{
		"SET": "key value",
		"GET": "key",
		"DEL": "key",
	}

	for cmd, args := range cases {
		query, err := NewQuery(Command(cmd), strings.Split(args, " "), logger)

		require.NoError(t, err, args)
		require.NotNil(t, query)
		assert.Equal(t, Command(cmd), query.Command)
		assert.Equal(t, args, strings.Join(query.Args, " "))
	}
}

func TestNewQueryWithLessArguments(t *testing.T) {
	cases := map[string][]string{
		"set": {"key"},
		"get": {""},
		"del": {""},
	}

	for cmd, arguments := range cases {
		for _, args := range arguments {
			query, err := NewQuery(Command(cmd), strings.Split(args, " "), logger)

			require.Error(t, err)
			require.Nil(t, query)
		}
	}
}

func TestNewQueryWithMoreArguments(t *testing.T) {
	cases := map[string][]string{
		"set": {"key value oops"},
		"get": {"key value"},
		"del": {"key value"},
	}

	for cmd, arguments := range cases {
		for _, args := range arguments {
			query, err := NewQuery(Command(cmd), strings.Split(args, " "), logger)

			require.Error(t, err)
			require.Nil(t, query)
		}
	}
}
