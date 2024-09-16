package buffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteAndReadFloat32(t *testing.T) {
	b := NewBuffer(1024)
	f := float32(3.14159)

	// Write float32
	err := b.WriteFloat32(f)
	assert.NoError(t, err, "expected no error during WriteFloat32")

	// Reset read cursor for reading
	b.readCursor = 0

	// Read float32
	readF, err := b.ReadFloat32()
	assert.NoError(t, err, "expected no error during ReadFloat32")
	assert.Equal(t, f, readF, "expected read float32 to match written float32")
}

func TestWriteAndReadFloat32Le(t *testing.T) {
	b := NewBuffer(1024)
	f := float32(3.14159)

	// Write float32 in little-endian format
	err := b.WriteFloat32Le(f)
	assert.NoError(t, err, "expected no error during WriteFloat32Le")

	// Reset read cursor for reading
	b.readCursor = 0

	// Read float32 in little-endian format
	readF, err := b.ReadFloat32Le()
	assert.NoError(t, err, "expected no error during ReadFloat32Le")
	assert.Equal(t, f, readF, "expected read float32 to match written float32 in little-endian")
}

func TestWriteAndReadFloat64(t *testing.T) {
	b := NewBuffer(1024)
	f := float64(3.141592653589793)

	// Write float64
	err := b.WriteFloat64(f)
	assert.NoError(t, err, "expected no error during WriteFloat64")

	// Reset read cursor for reading
	b.readCursor = 0

	// Read float64
	readF, err := b.ReadFloat64()
	assert.NoError(t, err, "expected no error during ReadFloat64")
	assert.Equal(t, f, readF, "expected read float64 to match written float64")
}

func TestWriteAndReadFloat64Le(t *testing.T) {
	b := NewBuffer(1024)
	f := float64(3.141592653589793)

	// Write float64 in little-endian format
	err := b.WriteFloat64Le(f)
	assert.NoError(t, err, "expected no error during WriteFloat64Le")

	// Reset read cursor for reading
	b.readCursor = 0

	// Read float64 in little-endian format
	readF, err := b.ReadFloat64Le()
	assert.NoError(t, err, "expected no error during ReadFloat64Le")
	assert.Equal(t, f, readF, "expected read float64 to match written float64 in little-endian")
}
