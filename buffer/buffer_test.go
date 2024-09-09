package buffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBuffer(t *testing.T) {
	b := NewBuffer(1024)
	assert.Equal(t, 1024, b.Capacity(), "expected capacity to be 1024")
	assert.Equal(t, 0, b.Size(), "expected size to be 0")
}

func TestBufferWrite(t *testing.T) {
	b := NewBuffer(1024)
	data := []byte("hello")

	n, err := b.Write(data)
	assert.NoError(t, err, "expected no error during write")
	assert.Equal(t, len(data), n, "expected number of bytes written to match input")
	assert.Equal(t, len(data), b.Size(), "expected size to match written data")
}

func TestBufferRead(t *testing.T) {
	b := NewBuffer(1024)
	data := []byte("hello")
	_, err := b.Write(data)
	assert.NoError(t, err, "expected no error during write")

	dst := make([]byte, 5)
	n, err := b.Read(dst)
	assert.NoError(t, err, "expected no error during read")
	assert.Equal(t, len(data), n, "expected number of bytes read to match input")
	assert.Equal(t, "hello", string(dst), "expected output to be 'hello'")
	assert.Equal(t, 0, b.Size(), "expected size to be 0 after reading all data")
}

func TestBufferOverflow(t *testing.T) {
	b := NewBuffer(5)
	data := []byte("this is too long")

	_, err := b.Write(data)
	assert.Error(t, err, "expected overflow error during write")
}

func TestAdjustWriteCursor(t *testing.T) {
	b := NewBuffer(1024)
	_, err := b.Write([]byte("test"))
	assert.NoError(t, err, "expected no error during write")
	assert.Equal(t, 4, b.Size(), "expected size to be 4 after initial write")

	err = b.AdjustWriteCursor(-2)
	assert.NoError(t, err, "expected no error during valid cursor adjustment")
	assert.Equal(t, 2, b.writeCursor, "expected write cursor to be adjusted to 2")
	assert.Equal(t, 2, b.Size(), "expected size to be adjusted to 2 after cursor adjustment")

	err = b.AdjustWriteCursor(2000)
	assert.Error(t, err, "expected out of bounds error for invalid cursor adjustment")
}

func TestBufferCompact(t *testing.T) {
	b := NewBuffer(1024)
	data := []byte("hello world")
	_, err := b.Write(data)
	assert.NoError(t, err, "expected no error during write")

	dst := make([]byte, 5)
	_, err = b.Read(dst) // read 'hello'
	assert.NoError(t, err, "expected no error during read")

	assert.Equal(t, 0, b.readCursor, "expected read cursor reset to 0")
	assert.Equal(t, 6, b.writeCursor, "expected write cursor at 6 after reading")
}
