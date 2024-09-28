package engine

type Engine interface {
	Set(k, v string)
	Get(k string) (string, bool)
	Del(k string) bool
}
