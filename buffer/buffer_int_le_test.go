package buffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBufferInt16Le(t *testing.T) {
	b := NewBuffer(2)

	var int16Val int16 = -12345
	err := b.WriteInt16Le(int16Val)
	assert.NoError(t, err, "WriteInt16Le failed")

	readInt16Val, err := b.ReadInt16Le()
	assert.NoError(t, err, "ReadInt16Le failed")
	assert.Equal(t, int16Val, readInt16Val, "ReadInt16Le value mismatch")
}

func TestBufferUInt16Le(t *testing.T) {
	b := NewBuffer(2)

	var uint16Val uint16 = 12345
	err := b.WriteUInt16Le(uint16Val)
	assert.NoError(t, err, "WriteUInt16Le failed")

	readUInt16Val, err := b.ReadUInt16Le()
	assert.NoError(t, err, "ReadUInt16Le failed")
	assert.Equal(t, uint16Val, readUInt16Val, "ReadUInt16Le value mismatch")
}

func TestBufferInt32Le(t *testing.T) {
	b := NewBuffer(4)

	var int32Val int32 = -12345678
	err := b.WriteInt32Le(int32Val)
	assert.NoError(t, err, "WriteInt32Le failed")

	readInt32Val, err := b.ReadInt32Le()
	assert.NoError(t, err, "ReadInt32Le failed")
	assert.Equal(t, int32Val, readInt32Val, "ReadInt32Le value mismatch")
}

func TestBufferUInt32Le(t *testing.T) {
	b := NewBuffer(4)

	var uint32Val uint32 = 12345678
	err := b.WriteUInt32Le(uint32Val)
	assert.NoError(t, err, "WriteUInt32Le failed")

	readUInt32Val, err := b.ReadUInt32Le()
	assert.NoError(t, err, "ReadUInt32Le failed")
	assert.Equal(t, uint32Val, readUInt32Val, "ReadUInt32Le value mismatch")
}

func TestBufferInt64Le(t *testing.T) {
	b := NewBuffer(8)

	var int64Val int64 = -123456789012345
	err := b.WriteInt64Le(int64Val)
	assert.NoError(t, err, "WriteInt64Le failed")

	readInt64Val, err := b.ReadInt64Le()
	assert.NoError(t, err, "ReadInt64Le failed")
	assert.Equal(t, int64Val, readInt64Val, "ReadInt64Le value mismatch")
}

func TestBufferUInt64Le(t *testing.T) {
	b := NewBuffer(8)

	var uint64Val uint64 = 123456789012345
	err := b.WriteUInt64Le(uint64Val)
	assert.NoError(t, err, "WriteUInt64Le failed")

	readUInt64Val, err := b.ReadUInt64Le()
	assert.NoError(t, err, "ReadUInt64Le failed")
	assert.Equal(t, uint64Val, readUInt64Val, "ReadUInt64Le value mismatch")
}
