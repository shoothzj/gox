package buffer

// ReadBool reads a boolean value from the buffer.
// It assumes that a boolean is stored as a single byte (0 for false, 1 for true).
func (b *Buffer) ReadBool() (bool, error) {
	bytes, err := b.ReadNBytes(1)
	if err != nil {
		return false, err
	}
	return bytes[0] == 1, nil
}

// WriteBool writes a boolean value to the buffer.
// It stores the boolean as a single byte (0 for false, 1 for true).
func (b *Buffer) WriteBool(x bool) error {
	var val byte
	if x {
		val = 1
	} else {
		val = 0
	}
	b.WritableSlice()[0] = val
	return b.AdjustWriteCursor(1)
}
