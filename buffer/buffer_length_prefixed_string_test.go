package buffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteLengthPrefixedString(t *testing.T) {
	b := NewBuffer(1024)
	str := "hello world"

	err := b.WriteLengthPrefixedString(str)
	assert.NoError(t, err, "expected no error during WriteLengthPrefixedString")
	assert.Equal(t, len(str)+4, b.ReadableSize(), "expected size to be string length + 4 for length prefix")
}

func TestReadLengthPrefixedString(t *testing.T) {
	b := NewBuffer(1024)
	str := "hello world"

	// Write the string
	err := b.WriteLengthPrefixedString(str)
	assert.NoError(t, err, "expected no error during WriteLengthPrefixedString")

	// Reset read cursor for reading
	b.readCursor = 0

	// Read the string
	readStr, err := b.ReadLengthPrefixedString()
	assert.NoError(t, err, "expected no error during ReadLengthPrefixedString")
	assert.Equal(t, str, readStr, "expected read string to match written string")
}

func TestWriteAndReadLengthPrefixedString(t *testing.T) {
	b := NewBuffer(1024)
	str := "test string"

	// Write the string
	err := b.WriteLengthPrefixedString(str)
	assert.NoError(t, err, "expected no error during WriteLengthPrefixedString")

	// Reset read cursor for reading
	b.readCursor = 0

	// Read the string
	readStr, err := b.ReadLengthPrefixedString()
	assert.NoError(t, err, "expected no error during ReadLengthPrefixedString")
	assert.Equal(t, str, readStr, "expected read string to match written string")
}

func TestWriteLengthPrefixedStringLe(t *testing.T) {
	b := NewBuffer(1024)
	str := "hello world"

	err := b.WriteLengthPrefixedStringLe(str)
	assert.NoError(t, err, "expected no error during WriteLengthPrefixedStringLe")
	assert.Equal(t, len(str)+4, b.ReadableSize(), "expected size to be string length + 4 for length prefix")
}

func TestReadLengthPrefixedStringLe(t *testing.T) {
	b := NewBuffer(1024)
	str := "hello world"

	// Write the string
	err := b.WriteLengthPrefixedStringLe(str)
	assert.NoError(t, err, "expected no error during WriteLengthPrefixedStringLe")

	// Read the string
	readStr, err := b.ReadLengthPrefixedStringLe()
	assert.NoError(t, err, "expected no error during ReadLengthPrefixedStringLe")
	assert.Equal(t, str, readStr, "expected read string to match written string")
}

func TestWriteAndReadLengthPrefixedStringLe(t *testing.T) {
	b := NewBuffer(1024)
	str := "test string"

	// Write the string
	err := b.WriteLengthPrefixedStringLe(str)
	assert.NoError(t, err, "expected no error during WriteLengthPrefixedStringLe")

	// Reset read cursor for reading
	b.readCursor = 0

	// Read the string
	readStr, err := b.ReadLengthPrefixedStringLe()
	assert.NoError(t, err, "expected no error during ReadLengthPrefixedStringLe")
	assert.Equal(t, str, readStr, "expected read string to match written string")
}
