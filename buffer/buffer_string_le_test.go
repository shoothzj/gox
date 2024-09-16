package buffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPutStringLe(t *testing.T) {
	b := NewBuffer(1024)
	str := "hello world"

	err := b.PutStringLe(str)
	assert.NoError(t, err, "expected no error during PutStringLe")
	assert.Equal(t, len(str)+4, b.ReadableSize(), "expected size to be string length + 4 for length prefix")
}

func TestReadStringLe(t *testing.T) {
	b := NewBuffer(1024)
	str := "hello world"

	// Write the string
	err := b.PutStringLe(str)
	assert.NoError(t, err, "expected no error during PutStringLe")

	// Read the string
	readStr, err := b.ReadStringLe()
	assert.NoError(t, err, "expected no error during ReadStringLe")
	assert.Equal(t, str, readStr, "expected read string to match written string")
}

func TestPutAndReadStringLe(t *testing.T) {
	b := NewBuffer(1024)
	str := "test string"

	// Write the string
	err := b.PutStringLe(str)
	assert.NoError(t, err, "expected no error during PutStringLe")

	// Reset read cursor for reading
	b.readCursor = 0

	// Read the string
	readStr, err := b.ReadStringLe()
	assert.NoError(t, err, "expected no error during ReadStringLe")
	assert.Equal(t, str, readStr, "expected read string to match written string")
}
