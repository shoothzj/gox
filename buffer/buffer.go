package buffer

import "fmt"

type Buffer struct {
	cap         int
	bytes       []byte
	size        int
	readCursor  int
	writeCursor int
}

func NewBuffer(cap int) *Buffer {
	var initCap int
	switch {
	case cap > 4*1024*1024:
		initCap = cap / 64
	case cap > 1*1024*1024:
		initCap = cap / 32
	case cap > 512*1024:
		initCap = cap / 16
	case cap > 128*1024:
		initCap = cap / 4
	default:
		initCap = cap
	}
	return &Buffer{
		cap:   cap,
		bytes: make([]byte, initCap),
	}
}

func (b *Buffer) Read(dst []byte) (int, error) {
	n := len(dst)
	if n > b.size {
		n = b.size
	}

	copy(dst, b.bytes[b.readCursor:b.readCursor+n])
	b.readCursor += n
	b.size -= n

	b.compact()

	return n, nil
}

func (b *Buffer) ReadExactly(dst []byte) error {
	n := len(dst)
	if n > b.size {
		return fmt.Errorf("not enough data to read")
	}

	copy(dst, b.bytes[b.readCursor:b.readCursor+n])
	b.readCursor += n
	b.size -= n

	b.compact()

	return nil
}

func (b *Buffer) compact() {
	if b.readCursor > 0 {
		copy(b.bytes, b.bytes[b.readCursor:b.writeCursor])
		b.writeCursor -= b.readCursor
		b.readCursor = 0
	}
}

func (b *Buffer) Write(p []byte) (int, error) {
	n := len(p)
	if n > (b.cap - b.size) {
		b.expand()
		if n > (b.cap - b.size) {
			return 0, fmt.Errorf("buffer overflow")
		}
	}

	copy(b.bytes[b.writeCursor:], p)
	b.writeCursor += n
	b.size += n

	return n, nil
}

func (b *Buffer) expand() {
	newCap := b.cap * 2
	if newCap > b.cap {
		buf := make([]byte, newCap)
		copy(buf, b.bytes)
		b.bytes = buf
		b.cap = newCap
	}
}

func (b *Buffer) Peek(dst []byte) (int, error) {
	n := len(dst)
	if n > b.size {
		n = b.size
	}

	copy(dst, b.bytes[b.readCursor:b.readCursor+n])
	return n, nil
}

func (b *Buffer) WritableSlice() []byte {
	return b.bytes[b.writeCursor:]
}

func (b *Buffer) AdjustWriteCursor(offset int) error {
	newCursor := b.writeCursor + offset
	if newCursor < 0 || newCursor > b.cap {
		return fmt.Errorf("write cursor adjustment out of bounds")
	}
	b.writeCursor = newCursor
	b.size += offset
	return nil
}

func (b *Buffer) Capacity() int {
	return b.cap
}

func (b *Buffer) Size() int {
	return b.size
}
