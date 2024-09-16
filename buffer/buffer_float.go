package buffer

import (
	"encoding/binary"
	"math"
)

// ReadFloat32 reads a float32 value in big-endian format from the buffer.
func (b *Buffer) ReadFloat32() (float32, error) {
	bytes, err := b.ReadNBytes(4)
	if err != nil {
		return 0, err
	}
	bits := binary.BigEndian.Uint32(bytes)
	return math.Float32frombits(bits), nil
}

// WriteFloat32 writes a float32 value in big-endian format to the buffer.
func (b *Buffer) WriteFloat32(f float32) error {
	bits := math.Float32bits(f)
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, bits)
	return b.WriteExactly(bytes)
}

// ReadFloat32Le reads a float32 value in little-endian format from the buffer.
func (b *Buffer) ReadFloat32Le() (float32, error) {
	bytes, err := b.ReadNBytes(4)
	if err != nil {
		return 0, err
	}
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits), nil
}

// WriteFloat32Le writes a float32 value in little-endian format to the buffer.
func (b *Buffer) WriteFloat32Le(f float32) error {
	bits := math.Float32bits(f)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return b.WriteExactly(bytes)
}

// ReadFloat64 reads a float64 value in big-endian format from the buffer.
func (b *Buffer) ReadFloat64() (float64, error) {
	bytes, err := b.ReadNBytes(8)
	if err != nil {
		return 0, err
	}
	bits := binary.BigEndian.Uint64(bytes)
	return math.Float64frombits(bits), nil
}

// WriteFloat64 writes a float64 value in big-endian format to the buffer.
func (b *Buffer) WriteFloat64(f float64) error {
	bits := math.Float64bits(f)
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, bits)
	return b.WriteExactly(bytes)
}

// ReadFloat64Le reads a float64 value in little-endian format from the buffer.
func (b *Buffer) ReadFloat64Le() (float64, error) {
	bytes, err := b.ReadNBytes(8)
	if err != nil {
		return 0, err
	}
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits), nil
}

// WriteFloat64Le writes a float64 value in little-endian format to the buffer.
func (b *Buffer) WriteFloat64Le(f float64) error {
	bits := math.Float64bits(f)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return b.WriteExactly(bytes)
}
