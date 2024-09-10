package buffer

import "encoding/binary"

func (b *Buffer) ReadInt16() (int16, error) {
	bytes, err := b.ReadNBytes(2)
	if err != nil {
		return 0, err
	}
	return int16(binary.BigEndian.Uint16(bytes)), nil
}

func (b *Buffer) PutInt16(x int16) error {
	return b.PutUInt16(uint16(x))
}

func (b *Buffer) ReadUInt16() (uint16, error) {
	bytes, err := b.ReadNBytes(2)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(bytes), nil
}

func (b *Buffer) PutUInt16(x uint16) error {
	binary.BigEndian.PutUint16(b.WritableSlice(), x)
	return b.AdjustWriteCursor(2)
}

func (b *Buffer) ReadInt32() (int32, error) {
	bytes, err := b.ReadNBytes(4)
	if err != nil {
		return 0, err
	}
	return int32(binary.BigEndian.Uint32(bytes)), nil
}

func (b *Buffer) PutInt32(x int32) error {
	return b.PutUInt32(uint32(x))
}

func (b *Buffer) ReadUInt32() (uint32, error) {
	bytes, err := b.ReadNBytes(4)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(bytes), nil
}

func (b *Buffer) PutUInt32(x uint32) error {
	binary.BigEndian.PutUint32(b.WritableSlice(), x)
	return b.AdjustWriteCursor(4)
}

func (b *Buffer) ReadInt64() (int64, error) {
	bytes, err := b.ReadNBytes(8)
	if err != nil {
		return 0, err
	}
	return int64(binary.BigEndian.Uint64(bytes)), nil
}

func (b *Buffer) PutInt64(x int64) error {
	return b.PutUInt64(uint64(x))
}

func (b *Buffer) ReadUInt64() (uint64, error) {
	bytes, err := b.ReadNBytes(8)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(bytes), nil
}

func (b *Buffer) PutUInt64(x uint64) error {
	binary.BigEndian.PutUint64(b.WritableSlice(), x)
	return b.AdjustWriteCursor(8)
}
