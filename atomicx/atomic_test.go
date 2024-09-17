package atomx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenericAtomicInt(t *testing.T) {
	// Create a new atomic value for int
	atomicInt := NewAtomic(42)

	// Test loading the initial value
	assert.Equal(t, 42, atomicInt.Load())

	// Test storing a new value
	atomicInt.Store(100)
	assert.Equal(t, 100, atomicInt.Load())

	// Test CompareAndSwap success
	casResult := atomicInt.CompareAndSwap(100, 200)
	assert.True(t, casResult)
	assert.Equal(t, 200, atomicInt.Load())

	// Test CompareAndSwap failure
	casResult = atomicInt.CompareAndSwap(100, 300)
	assert.False(t, casResult)
	assert.Equal(t, 200, atomicInt.Load())
}

func TestGenericAtomicString(t *testing.T) {
	// Create a new atomic value for string
	atomicStr := NewAtomic("hello")

	// Test loading the initial value
	assert.Equal(t, "hello", atomicStr.Load())

	// Test storing a new value
	atomicStr.Store("world")
	assert.Equal(t, "world", atomicStr.Load())

	// Test CompareAndSwap success
	casResult := atomicStr.CompareAndSwap("world", "gophers")
	assert.True(t, casResult)
	assert.Equal(t, "gophers", atomicStr.Load())

	// Test CompareAndSwap failure
	casResult = atomicStr.CompareAndSwap("world", "goroutines")
	assert.False(t, casResult)
	assert.Equal(t, "gophers", atomicStr.Load())
}

func TestGenericAtomicStruct(t *testing.T) {
	type MyStruct struct {
		ID   int
		Name string
	}

	// Create a new atomic value for a struct
	initialStruct := MyStruct{ID: 1, Name: "First"}
	atomicStruct := NewAtomic(initialStruct)

	// Test loading the initial value
	loadedStruct := atomicStruct.Load()
	assert.Equal(t, 1, loadedStruct.ID)
	assert.Equal(t, "First", loadedStruct.Name)

	// Test storing a new value
	newStruct := MyStruct{ID: 2, Name: "Second"}
	atomicStruct.Store(newStruct)
	loadedStruct = atomicStruct.Load()
	assert.Equal(t, 2, loadedStruct.ID)
	assert.Equal(t, "Second", loadedStruct.Name)

	// Test CompareAndSwap success
	casResult := atomicStruct.CompareAndSwap(newStruct, MyStruct{ID: 3, Name: "Third"})
	assert.True(t, casResult)
	loadedStruct = atomicStruct.Load()
	assert.Equal(t, 3, loadedStruct.ID)
	assert.Equal(t, "Third", loadedStruct.Name)

	// Test CompareAndSwap failure
	casResult = atomicStruct.CompareAndSwap(newStruct, MyStruct{ID: 4, Name: "Fourth"})
	assert.False(t, casResult)
	loadedStruct = atomicStruct.Load()
	assert.Equal(t, 3, loadedStruct.ID)
	assert.Equal(t, "Third", loadedStruct.Name)
}
