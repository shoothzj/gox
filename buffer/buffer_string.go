package buffer

// ReadString reads a string in big-endian format. It first reads the length (uint32), then the string data.
func (b *Buffer) ReadString() (string, error) {
	length, err := b.ReadUInt32()
	if err != nil {
		return "", err
	}

	data, err := b.ReadNBytes(int(length))
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// PutString writes a string in big-endian format. It writes the length (uint32), followed by the string data.
func (b *Buffer) PutString(s string) error {
	length := uint32(len(s))
	err := b.PutUInt32(length)
	if err != nil {
		return err
	}

	_, err = b.Write([]byte(s))
	return err
}
