package buffer

// ReadString reads a string of the specified length.
func (b *Buffer) ReadString(length int) (string, error) {
	data, err := b.ReadNBytes(length)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// WriteString writes the string data to the buffer.
func (b *Buffer) WriteString(s string) error {
	return b.WriteExactly([]byte(s))
}
