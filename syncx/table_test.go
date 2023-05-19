package syncx

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestTableNew(t *testing.T) {
	table := NewSyncTable[string, string, int]()
	table.Set("a1", "a2", 1)
	value, ok := table.Get("a1", "a2")
	assert.Equal(t, 1, value)
	assert.Equal(t, true, ok)
	table.Clear()
	value, ok = table.Get("a1", "a2")
	assert.Equal(t, 0, value)
	assert.Equal(t, false, ok)

	table.Set("a3", "a4", 2)
	value, ok = table.Get("a3", "a4")
	assert.Equal(t, 2, value)
	assert.Equal(t, true, ok)
	table.Delete("a3", "a4")
	value, ok = table.Get("a1", "a2")
	assert.Equal(t, 0, value)
	assert.Equal(t, false, ok)
}

func TestTableBasic(t *testing.T) {
	table := &SyncTable[string, string, int]{
		mu: new(sync.RWMutex),
	}
	table.Set("a1", "a2", 1)
	value, ok := table.Get("a1", "a2")
	assert.Equal(t, 1, value)
	assert.Equal(t, true, ok)
	table.Clear()
	value, ok = table.Get("a1", "a2")
	assert.Equal(t, 0, value)
	assert.Equal(t, false, ok)

	table.Set("a3", "a4", 2)
	value, ok = table.Get("a3", "a4")
	assert.Equal(t, 2, value)
	assert.Equal(t, true, ok)
	table.Delete("a3", "a4")
	value, ok = table.Get("a1", "a2")
	assert.Equal(t, 0, value)
	assert.Equal(t, false, ok)
}

func TestTableIterator(t *testing.T) {
	table := &SyncTable[string, string, int]{
		mu: new(sync.RWMutex),
	}
	table.Set("a1", "a2", 1)
	table.Set("b1", "b2", 2)
	table.Set("c1", "c2", 3)
	table.Set("d1", "d2", 4)
	table.Set("e1", "e2", 5)
	table.Set("f1", "f2", 6)
	table.Set("a1", "a3", 1)
	table.Set("b1", "b3", 2)
	table.Set("c1", "c3", 3)
	table.Set("d1", "d3", 4)
	table.Set("e1", "e3", 5)
	table.Set("f1", "f3", 6)

	i := 0
	table.Range(func(key1 string, key2 string, v int) bool {
		i++
		return true
	})
	assert.Equal(t, 12, i)
	i = 0
	table.Range(func(key1 string, key2 string, v int) bool {
		i++
		return false
	})
	assert.Equal(t, 1, i)
}

func TestTableLoadAndDelete(t *testing.T) {
	table := &SyncTable[string, string, int]{
		mu: new(sync.RWMutex),
	}
	table.Set("a1", "a2", 1)
	value, ok := table.Get("a1", "a2")
	assert.Equal(t, 1, value)
	assert.Equal(t, true, ok)

	valueList, ok := table.LoadAndDelete("a0")
	assert.Equal(t, 0, len(valueList))
	assert.Equal(t, false, ok)

	valueList, ok = table.LoadAndDelete("a1")
	assert.Equal(t, 1, valueList[0])
	assert.Equal(t, true, ok)

	value, ok = table.Get("a1", "a2")
	assert.Equal(t, 0, value)
	assert.Equal(t, false, ok)
}
