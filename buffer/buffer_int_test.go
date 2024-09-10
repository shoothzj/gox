package buffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBufferInt32(t *testing.T) {
	b := NewBuffer(4)

	var int32Val int32 = -12345678
	err := b.PutInt32(int32Val)
	assert.NoError(t, err, "PutInt32 failed")

	readInt32Val, err := b.ReadInt32()
	assert.NoError(t, err, "ReadInt32 failed")
	assert.Equal(t, int32Val, readInt32Val, "ReadInt32 value mismatch")
}

func TestBufferUInt32(t *testing.T) {
	b := NewBuffer(4)

	var uint32Val uint32 = 12345678
	err := b.PutUInt32(uint32Val)
	assert.NoError(t, err, "PutUInt32 failed")

	readUInt32Val, err := b.ReadUInt32()
	assert.NoError(t, err, "ReadUInt32 failed")
	assert.Equal(t, uint32Val, readUInt32Val, "ReadUInt32 value mismatch")
}

func TestBufferInt64(t *testing.T) {
	b := NewBuffer(8)

	var int64Val int64 = -123456789012345
	err := b.PutInt64(int64Val)
	assert.NoError(t, err, "PutInt64 failed")

	readInt64Val, err := b.ReadInt64()
	assert.NoError(t, err, "ReadInt64 failed")
	assert.Equal(t, int64Val, readInt64Val, "ReadInt64 value mismatch")
}

func TestBufferUInt64(t *testing.T) {
	b := NewBuffer(8)

	var uint64Val uint64 = 123456789012345
	err := b.PutUInt64(uint64Val)
	assert.NoError(t, err, "PutUInt64 failed")

	readUInt64Val, err := b.ReadUInt64()
	assert.NoError(t, err, "ReadUInt64 failed")
	assert.Equal(t, uint64Val, readUInt64Val, "ReadUInt64 value mismatch")
}

func TestBufferBoundary(t *testing.T) {
	b := NewBuffer(4)

	var uint32Val uint32 = 12345678
	err := b.PutUInt32(uint32Val)
	assert.NoError(t, err, "PutUInt32 failed")

	// Attempt to read beyond buffer size should fail
	_, err = b.ReadUInt64()
	assert.Error(t, err, "Expected error when reading beyond buffer size")
}
