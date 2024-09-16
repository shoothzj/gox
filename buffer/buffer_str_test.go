package buffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteString(t *testing.T) {
	b := NewBuffer(1024)
	str := "hello world"

	err := b.WriteString(str)
	assert.NoError(t, err, "expected no error during WriteString")
	assert.Equal(t, len(str), b.ReadableSize(), "expected size to be the string length")
}

func TestReadString(t *testing.T) {
	b := NewBuffer(1024)
	str := "hello world"

	// Write the string
	err := b.WriteString(str)
	assert.NoError(t, err, "expected no error during WriteString")

	// Reset read cursor for reading
	b.readCursor = 0

	// Read the string
	readStr, err := b.ReadString(len(str))
	assert.NoError(t, err, "expected no error during ReadString")
	assert.Equal(t, str, readStr, "expected read string to match written string")
}

func TestWriteAndReadString(t *testing.T) {
	b := NewBuffer(1024)
	str := "test string"

	// Write the string
	err := b.WriteString(str)
	assert.NoError(t, err, "expected no error during WriteString")

	// Reset read cursor for reading
	b.readCursor = 0

	// Read the string
	readStr, err := b.ReadString(len(str))
	assert.NoError(t, err, "expected no error during ReadString")
	assert.Equal(t, str, readStr, "expected read string to match written string")
}
