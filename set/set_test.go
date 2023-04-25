package set

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSet(t *testing.T) {
	set := Set[int]{}

	set.Add(1)
	set.Add(2)

	require.Equal(t, 2, set.Len())

	require.True(t, set.Contains(1))

	set.Remove(1)

	require.Equal(t, 1, set.Len())

	elements := set.Elements()
	require.Equal(t, 1, len(elements))
	require.Equal(t, 2, elements[0])
}
