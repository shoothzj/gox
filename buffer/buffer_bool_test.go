package buffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBufferBoolTrue(t *testing.T) {
	b := NewBuffer(1)

	err := b.PutBool(true)
	assert.NoError(t, err, "PutBool true failed")

	readBoolVal, err := b.ReadBool()
	assert.NoError(t, err, "ReadBool failed")
	assert.True(t, readBoolVal, "ReadBool expected true, got false")
}

func TestBufferBoolFalse(t *testing.T) {
	b := NewBuffer(1)

	err := b.PutBool(false)
	assert.NoError(t, err, "PutBool false failed")

	readBoolVal, err := b.ReadBool()
	assert.NoError(t, err, "ReadBool failed")
	assert.False(t, readBoolVal, "ReadBool expected false, got true")
}

func TestBufferBoolBoundary(t *testing.T) {
	b := NewBuffer(0)

	// Attempt to read beyond buffer size should fail
	_, err := b.ReadBool()
	assert.Error(t, err, "Expected error when reading beyond buffer size")
}
