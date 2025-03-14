// Package endianio provides utilities for reading binary data in both big-endian and little-endian formats.
package endianio

import (
	"encoding/binary"
	"io"
	"math"
)

// Reader wraps an io.Reader to provide methods for reading binary data in different endian formats.
type Reader struct {
	io.Reader
}

// NewReader creates a new Reader reading from the provided io.Reader.
func NewReader(r io.Reader) *Reader {
	return &Reader{r}
}

// ReadUint8 reads a uint8 (byte)
func (r *Reader) ReadUint8() (uint8, error) {
	var b [1]byte
	_, err := r.Read(b[:])
	return b[0], err

}

// ReadBigUint16 reads a 16-bit unsigned integer in big-endian format.
func (r *Reader) ReadBigUint16() (uint16, error) {
	var b [2]byte
	_, err := io.ReadFull(r, b[:])
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(b[:]), nil
}

// ReadBigUint32 reads a 32-bit unsigned integer in big-endian format.
func (r *Reader) ReadBigUint32() (uint32, error) {
	var b [4]byte
	_, err := io.ReadFull(r, b[:])
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(b[:]), nil
}

// ReadBigUint64 reads a 64-bit unsigned integer in big-endian format.
func (r *Reader) ReadBigUint64() (uint64, error) {
	var b [8]byte
	_, err := io.ReadFull(r, b[:])
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(b[:]), nil
}

// ReadBigFloat32 reads a 32-bit float encoded as a 32-bit unsigned integer in big-endian format.
func (r *Reader) ReadBigFloat32() (float32, error) {
	var b [4]byte
	_, err := io.ReadFull(r, b[:])
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(binary.BigEndian.Uint32(b[:])), nil
}

// ReadBigFloat64 reads a 64-bit float encoded as a 64-bit unsigned integer in big-endian format.
func (r *Reader) ReadBigFloat64() (float64, error) {
	var b [8]byte
	_, err := io.ReadFull(r, b[:])
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(binary.BigEndian.Uint64(b[:])), nil
}

// ReadLittleUint16 reads a 16-bit unsigned integer in little-endian format.
func (r *Reader) ReadLittleUint16() (uint16, error) {
	var b [2]byte
	_, err := io.ReadFull(r, b[:])
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(b[:]), nil
}

// ReadLittleUint32 reads a 32-bit unsigned integer in little-endian format.
func (r *Reader) ReadLittleUint32() (uint32, error) {
	var b [4]byte
	_, err := io.ReadFull(r, b[:])
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(b[:]), nil
}

// ReadLittleUint64 reads a 64-bit unsigned integer in little-endian format.
func (r *Reader) ReadLittleUint64() (uint64, error) {
	var b [8]byte
	_, err := io.ReadFull(r, b[:])
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(b[:]), nil
}

// ReadLittleFloat32 reads a 32-bit float encoded as a 32-bit unsigned integer in little-endian format.
func (r *Reader) ReadLittleFloat32() (float32, error) {
	var b [4]byte
	_, err := io.ReadFull(r, b[:])
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(binary.LittleEndian.Uint32(b[:])), nil
}

// ReadLittleFloat32 reads a 32-bit float encoded as a 32-bit unsigned integer in little-endian format.
func (r *Reader) ReadLittleFloat64() (float64, error) {
	var b [8]byte
	_, err := io.ReadFull(r, b[:])
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(binary.LittleEndian.Uint64(b[:])), nil
}
