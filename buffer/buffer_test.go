package buffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBuffer(t *testing.T) {
	b := NewBuffer(1024)
	assert.Equal(t, 1024, b.Capacity(), "expected capacity to be 1024")
	assert.Equal(t, 0, b.ReadableSize(), "expected size to be 0")
}

func TestBufferWrite(t *testing.T) {
	b := NewBuffer(1024)
	data := []byte("hello")

	n, err := b.Write(data)
	assert.NoError(t, err, "expected no error during write")
	assert.Equal(t, len(data), n, "expected number of bytes written to match input")
	assert.Equal(t, len(data), b.ReadableSize(), "expected size to match written data")
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
	assert.Equal(t, 0, b.ReadableSize(), "expected size to be 0 after reading all data")
}

func TestBufferOverflow(t *testing.T) {
	b := NewBuffer(5)
	data := []byte("this is too long")

	_, err := b.Write(data)
	assert.Error(t, err, "expected overflow error during write")
}

func TestBufferPeek(t *testing.T) {
	b := NewBuffer(1024)
	data := []byte("hello")
	_, err := b.Write(data)
	assert.NoError(t, err, "expected no error during write")

	dst := make([]byte, 5)
	n, err := b.Peek(dst)
	assert.NoError(t, err, "expected no error during peek")
	assert.Equal(t, len(data), n, "expected number of bytes peeked to match input")
	assert.Equal(t, "hello", string(dst), "expected output to be 'hello'")

	// Ensure that the read cursor has not moved
	assert.Equal(t, 0, b.readCursor, "expected read cursor to remain unchanged after peek")
}

func TestBufferSkip(t *testing.T) {
	b := NewBuffer(1024)
	data := []byte("hello world")
	_, err := b.Write(data)
	assert.NoError(t, err, "expected no error during write")

	err = b.Skip(6) // Skipping "hello "
	assert.NoError(t, err, "expected no error during skip")
	assert.Equal(t, 6, b.readCursor, "expected read cursor to be at 6 after skip")
	assert.Equal(t, 5, b.ReadableSize(), "expected size to be 5 after skip")

	dst := make([]byte, 5)
	n, err := b.Read(dst)
	assert.NoError(t, err, "expected no error during read after skip")
	assert.Equal(t, 5, n, "expected to read 5 bytes after skip")
	assert.Equal(t, "world", string(dst), "expected to read 'world' after skip")
}

func TestAdjustWriteCursor(t *testing.T) {
	b := NewBuffer(1024)
	_, err := b.Write([]byte("test"))
	assert.NoError(t, err, "expected no error during write")
	assert.Equal(t, 4, b.ReadableSize(), "expected size to be 4 after initial write")

	err = b.AdjustWriteCursor(-2)
	assert.NoError(t, err, "expected no error during valid cursor adjustment")
	assert.Equal(t, 2, b.writeCursor, "expected write cursor to be adjusted to 2")
	assert.Equal(t, 2, b.ReadableSize(), "expected size to be adjusted to 2 after cursor adjustment")

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

	b.Compact()

	assert.Equal(t, 0, b.readCursor, "expected read cursor reset to 0")
	assert.Equal(t, 6, b.writeCursor, "expected write cursor at 6 after reading")
}

func TestBufferReadAll(t *testing.T) {
	b := NewBuffer(1024)
	data := []byte("hello world")
	_, err := b.Write(data)
	assert.NoError(t, err, "expected no error during write")

	readData := b.ReadAll()

	assert.Equal(t, data, readData, "expected data read to match written data")
	assert.Equal(t, b.writeCursor, b.readCursor, "expected read cursor to be equal to write cursor after ReadAll")
	assert.Equal(t, 0, b.ReadableSize(), "expected size to be 0 after ReadAll")
}
