package buffer

// ReadStringLe reads a string in little-endian format. First, it reads the length (uint32), then the string data.
func (b *Buffer) ReadStringLe() (string, error) {
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

// PutStringLe writes a string in little-endian format. It writes the length (uint32), followed by the string data.
func (b *Buffer) PutStringLe(s string) error {
	length := uint32(len(s))
	err := b.PutUInt32Le(length)
	if err != nil {
		return err
	}

	_, err = b.Write([]byte(s))
	return err
}
