package engine

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var ht *HashTable

func TestMain(m *testing.M) {
	ht = NewHashTable()
	m.Run()
}

func TestNewHashTable(t *testing.T) {
	require.NotNil(t, ht)
	assert.NotNil(t, ht.data)
}

func TestHashTable_Add(t *testing.T) {
	ht.Set("key", "value")

	assert.Equal(t, ht.data["key"], "value")
}

func TestHashTable_Get(t *testing.T) {
	ht.Set("key", "value")
	retrieved, found := ht.Get("key")

	assert.Equal(t, retrieved, "value")
	assert.True(t, found)
}

func TestHashTable_Del(t *testing.T) {
	ht.Set("key", "value")
	deleted := ht.Del("key")
	retrieved, found := ht.Get("key")

	assert.True(t, deleted)
	assert.Empty(t, retrieved)
	assert.False(t, found)
}
