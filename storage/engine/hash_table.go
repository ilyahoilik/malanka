package engine

type HashTable struct {
	data map[string]string
}

func NewHashTable() *HashTable {
	return &HashTable{
		data: make(map[string]string),
	}
}

func (h *HashTable) Set(key, value string) {
	h.data[key] = value
}

func (h *HashTable) Get(key string) (string, bool) {
	value, found := h.data[key]
	return value, found
}

func (h *HashTable) Del(key string) bool {
	_, found := h.data[key]
	delete(h.data, key)
	return found
}
