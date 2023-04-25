package syncx

import "sync"

type Map[K comparable, V any] struct {
	sync.Map
}

func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.Map.Load(key)
	if ok {
		value = v.(V)
	}
	return value, ok
}

func (m *Map[K, V]) Store(key K, value V) {
	m.Map.Store(key, value)
}

func (m *Map[K, V]) Delete(key K) {
	m.Map.Delete(key)
}

func (m *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	v, loaded := m.Map.LoadOrStore(key, value)
	if loaded {
		actual = v.(V)
	} else {
		actual = value
	}
	return actual, loaded
}

func (m *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, loaded := m.Map.LoadAndDelete(key)
	if loaded {
		value = v.(V)
	}
	return value, loaded
}

func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.Map.Range(func(k, v interface{}) bool {
		return f(k.(K), v.(V))
	})
}
