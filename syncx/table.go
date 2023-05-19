package syncx

import (
	"github.com/Shoothzj/gox/table"
	"sync"
)

type SyncTable[K1 comparable, K2 comparable, V any] struct {
	mu   *sync.RWMutex
	data table.Table[K1, K2, V]
}

// Set sets the value for the given keys
func (t *SyncTable[K1, K2, V]) Set(key1 K1, key2 K2, value V) {
	t.mu.Lock()
	if t.data == nil {
		t.data = make(table.Table[K1, K2, V])
	}
	_, ok := t.data[key1]
	if !ok {
		t.data[key1] = make(map[K2]V)
		t.data[key1][key2] = value
	} else {
		t.data[key1][key2] = value
	}
	t.mu.Unlock()
}

// Get gets the value for the given keys
func (t *SyncTable[K1, K2, V]) Get(key1 K1, key2 K2) (V, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	var value V
	if t.data == nil {
		return value, false
	}
	m1, ok := t.data[key1]
	if !ok {
		return value, false
	}
	value, ok = m1[key2]
	return value, ok
}

// Delete deletes the value for the given keys
func (t *SyncTable[K1, K2, V]) Delete(key1 K1, key2 K2) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.data == nil {
		return
	}
	m1, ok := t.data[key1]
	if !ok {
		return
	}
	delete(m1, key2)

	if len(m1) == 0 {
		delete(t.data, key1)
	}
}

// LoadAndDelete deletes the value for the given keys when loaded.
func (t *SyncTable[K1, K2, V]) LoadAndDelete(key1 K1) ([]V, bool) {
	t.mu.Lock()
	defer t.mu.Unlock()
	var value = make([]V, 0)
	if t.data == nil {
		return value, false
	}
	m1, ok := t.data[key1]
	if !ok {
		return value, false
	}
	for _, v := range m1 {
		value = append(value, v)
	}
	delete(t.data, key1)
	return value, ok
}

func (t *SyncTable[K1, K2, V]) Range(f func(key1 K1, key2 K2, v V) bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	for k1, v := range t.data {
		for k2, vv := range v {
			if !f(k1, k2, vv) {
				return
			}
		}
	}
}

// Clear deletes all data
func (t *SyncTable[K1, K2, V]) Clear() {
	t.mu.Lock()
	t.data = make(table.Table[K1, K2, V])
	t.mu.Unlock()
}
