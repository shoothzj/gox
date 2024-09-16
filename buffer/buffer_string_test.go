package buffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPutString(t *testing.T) {
	b := NewBuffer(1024)
	str := "hello world"

	err := b.PutString(str)
	assert.NoError(t, err, "expected no error during PutString")
	assert.Equal(t, len(str)+4, b.ReadableSize(), "expected size to be string length + 4 for length prefix")
}

func TestReadString(t *testing.T) {
	b := NewBuffer(1024)
	str := "hello world"

	// Write the string
	err := b.PutString(str)
	assert.NoError(t, err, "expected no error during PutString")

	// Reset read cursor for reading
	b.readCursor = 0

	// Read the string
	readStr, err := b.ReadString()
	assert.NoError(t, err, "expected no error during ReadString")
	assert.Equal(t, str, readStr, "expected read string to match written string")
}

func TestPutAndReadString(t *testing.T) {
	b := NewBuffer(1024)
	str := "test string"

	// Write the string
	err := b.PutString(str)
	assert.NoError(t, err, "expected no error during PutString")

	// Reset read cursor for reading
	b.readCursor = 0

	// Read the string
	readStr, err := b.ReadString()
	assert.NoError(t, err, "expected no error during ReadString")
	assert.Equal(t, str, readStr, "expected read string to match written string")
}
