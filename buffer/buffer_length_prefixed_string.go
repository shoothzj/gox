package buffer

// ReadLengthPrefixedString reads a string in big-endian format. It first reads the length (uint32), then the string data.
func (b *Buffer) ReadLengthPrefixedString() (string, error) {
	length, err := b.ReadUInt32()
	if err != nil {
		return "", err
	}

	return b.ReadString(int(length))
}

// WriteLengthPrefixedString writes a string in big-endian format. It writes the length (uint32), followed by the string data.
func (b *Buffer) WriteLengthPrefixedString(s string) error {
	length := uint32(len(s))
	err := b.WriteUInt32(length)
	if err != nil {
		return err
	}

	err = b.WriteString(s)
	return err
}

// ReadLengthPrefixedStringLe reads a string in little-endian format. First, it reads the length (uint32), then the string data.
func (b *Buffer) ReadLengthPrefixedStringLe() (string, error) {
	length, err := b.ReadUInt32Le()
	if err != nil {
		return "", err
	}

	data, err := b.ReadNBytes(int(length))
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// WriteLengthPrefixedStringLe writes a string in little-endian format. It writes the length (uint32), followed by the string data.
func (b *Buffer) WriteLengthPrefixedStringLe(s string) error {
	length := uint32(len(s))
	err := b.WriteUInt32Le(length)
	if err != nil {
		return err
	}

	err = b.WriteExactly([]byte(s))
	return err
}
