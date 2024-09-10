package buffer

import "encoding/binary"

func (b *Buffer) ReadInt16Le() (int16, error) {
	bytes, err := b.ReadNBytes(2)
	if err != nil {
		return 0, err
	}
	return int16(binary.LittleEndian.Uint16(bytes)), nil
}

func (b *Buffer) PutInt16Le(x int16) error {
	return b.PutUInt16Le(uint16(x))
}

func (b *Buffer) ReadUInt16Le() (uint16, error) {
	bytes, err := b.ReadNBytes(2)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(bytes), nil
}

func (b *Buffer) PutUInt16Le(x uint16) error {
	binary.LittleEndian.PutUint16(b.WritableSlice(), x)
	return b.AdjustWriteCursor(2)
}

func (b *Buffer) ReadInt32Le() (int32, error) {
	bytes, err := b.ReadNBytes(4)
	if err != nil {
		return 0, err
	}
	return int32(binary.LittleEndian.Uint32(bytes)), nil
}

func (b *Buffer) PutInt32Le(x int32) error {
	return b.PutUInt32Le(uint32(x))
}

func (b *Buffer) ReadUInt32Le() (uint32, error) {
	bytes, err := b.ReadNBytes(4)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(bytes), nil
}

func (b *Buffer) PutUInt32Le(x uint32) error {
	binary.LittleEndian.PutUint32(b.WritableSlice(), x)
	return b.AdjustWriteCursor(4)
}

func (b *Buffer) ReadInt64Le() (int64, error) {
	bytes, err := b.ReadNBytes(8)
	if err != nil {
		return 0, err
	}
	return int64(binary.LittleEndian.Uint64(bytes)), nil
}

func (b *Buffer) PutInt64Le(x int64) error {
	return b.PutUInt64Le(uint64(x))
}

func (b *Buffer) ReadUInt64Le() (uint64, error) {
	bytes, err := b.ReadNBytes(8)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(bytes), nil
}

func (b *Buffer) PutUInt64Le(x uint64) error {
	binary.LittleEndian.PutUint64(b.WritableSlice(), x)
	return b.AdjustWriteCursor(8)
}
