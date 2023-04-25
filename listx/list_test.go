package listx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestList(t *testing.T) {
	// Create a new list
	l := New[string]()

	// Check that the list is empty
	require.Equal(t, 0, l.Len())

	// Add some items to the list
	l.PushBack("hello")
	l.PushBack("world")

	// Check the length of the list
	require.Equal(t, 2, l.Len())

	// Check the front of the list
	require.Equal(t, "hello", l.Front().Value)

	// Add another item to the list
	l.PushBack("foo")

	// Check the length of the list again
	require.Equal(t, 3, l.Len())

	// Check the front of the list again
	require.Equal(t, "hello", l.Front().Value)
}
