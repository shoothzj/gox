package syncx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet_Add(t *testing.T) {
	set := &Set[int]{}

	// Add elements to the set
	set.Add(1)
	set.Add(2)

	// Test if the elements were added correctly
	assert.True(t, set.Contains(1), "Set should contain 1")
	assert.True(t, set.Contains(2), "Set should contain 2")
}

func TestSet_Remove(t *testing.T) {
	set := &Set[int]{}

	// Add elements to the set
	set.Add(1)
	set.Add(2)

	// Remove an element
	set.Remove(1)

	// Test if the element was removed correctly
	assert.False(t, set.Contains(1), "Set should not contain 1 after removal")
	assert.True(t, set.Contains(2), "Set should still contain 2")
}

func TestSet_Get(t *testing.T) {
	set := &Set[int]{}

	// Add elements to the set
	set.Add(1)

	// Test Get functionality
	value, ok := set.Get(1)
	assert.True(t, ok, "Get should return true for 1")
	assert.Equal(t, 1, value, "Get should return the value 1")

	// Test Get for a non-existing value
	_, ok = set.Get(2)
	assert.False(t, ok, "Get should return false for non-existing value 2")
}

func TestSet_Contains(t *testing.T) {
	set := &Set[int]{}

	// Add elements to the set
	set.Add(1)

	// Test Contains functionality
	assert.True(t, set.Contains(1), "Set should contain 1")
	assert.False(t, set.Contains(2), "Set should not contain 2")
}

func TestSet_Range(t *testing.T) {
	set := &Set[int]{}

	// Add elements to the set
	set.Add(1)
	set.Add(2)
	set.Add(3)

	// Create a map to track visited elements
	visited := make(map[int]bool)

	// Use Range to iterate over the elements
	set.Range(func(value int) bool {
		visited[value] = true
		return true
	})

	// Check if all elements were visited
	assert.True(t, visited[1], "Range should visit 1")
	assert.True(t, visited[2], "Range should visit 2")
	assert.True(t, visited[3], "Range should visit 3")
}
