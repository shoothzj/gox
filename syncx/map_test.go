package syncx

import (
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	m := &Map[string, int]{}

	// Define some test data
	testData := []struct {
		key   string
		value int
	}{
		{"key1", 10},
		{"key2", 20},
		{"key3", 30},
	}

	var wg sync.WaitGroup
	wg.Add(len(testData))

	for i := 0; i < len(testData); i++ {
		go func(i int) {
			m.Store(testData[i].key, testData[i].value)
			wg.Done()
		}(i)
	}

	wg.Wait()

	for i := 0; i < len(testData); i++ {
		if v, ok := m.Load(testData[i].key); !ok || v != testData[i].value {
			t.Fatalf("Got %v, %v, expected %v, %v", v, ok, testData[i].value, true)
		}
	}
}
