// Package endianio provides utilities for reading binary data in both big-endian and little-endian formats.
package endianio

import (
	"encoding/binary"
	"io"
	"math"
)

// EndianReader is an interface that defines methods for reading binary data.
type EndianReader interface {
	// ReadUint8 reads a uint8 (byte)
	ReadUint8() (uint8, error)
	// ReadUint16 reads a 16-bit unsigned integer
	ReadUint16() (uint16, error)
	// ReadUint32 reads a 32-bit unsigned integer
	ReadUint32() (uint32, error)
	// ReadUint64 reads a 64-bit unsigned integer
	ReadUint64() (uint64, error)
	// ReadFloat32 reads a 32-bit float
	ReadFloat32() (float32, error)
	// ReadFloat64 reads a 64-bit float
	ReadFloat64() (float64, error)
}

// baseReader provides common functionality for both big-endian and little-endian readers.
type baseReader struct {
	io.Reader
}

// ReadUint8 reads a uint8 (byte)
func (r *baseReader) ReadUint8() (uint8, error) {
	var b [1]byte
	_, err := r.Read(b[:])
	return b[0], err
}

// BigEndianReader reads binary data in big-endian format.
type BigEndianReader struct {
	baseReader
}

// NewBigEndianReader creates a new BigEndianReader reading from the provided io.Reader.
func NewBigEndianReader(r io.Reader) *BigEndianReader {
	return &BigEndianReader{baseReader{r}}
}

// ReadUint16 reads a 16-bit unsigned integer in big-endian format.
func (r *BigEndianReader) ReadUint16() (uint16, error) {
	var b [2]byte
	_, err := io.ReadFull(r, b[:])
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(b[:]), nil
}

// ReadUint32 reads a 32-bit unsigned integer in big-endian format.
func (r *BigEndianReader) ReadUint32() (uint32, error) {
	var b [4]byte
	_, err := io.ReadFull(r, b[:])
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(b[:]), nil
}

// ReadUint64 reads a 64-bit unsigned integer in big-endian format.
func (r *BigEndianReader) ReadUint64() (uint64, error) {
	var b [8]byte
	_, err := io.ReadFull(r, b[:])
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(b[:]), nil
}

// ReadFloat32 reads a 32-bit float encoded as a 32-bit unsigned integer in big-endian format.
func (r *BigEndianReader) ReadFloat32() (float32, error) {
	var b [4]byte
	_, err := io.ReadFull(r, b[:])
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(binary.BigEndian.Uint32(b[:])), nil
}

// ReadFloat64 reads a 64-bit float encoded as a 64-bit unsigned integer in big-endian format.
func (r *BigEndianReader) ReadFloat64() (float64, error) {
	var b [8]byte
	_, err := io.ReadFull(r, b[:])
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(binary.BigEndian.Uint64(b[:])), nil
}

// LittleEndianReader reads binary data in little-endian format.
type LittleEndianReader struct {
	baseReader
}

// NewLittleEndianReader creates a new LittleEndianReader reading from the provided io.Reader.
func NewLittleEndianReader(r io.Reader) *LittleEndianReader {
	return &LittleEndianReader{baseReader{r}}
}

// ReadUint16 reads a 16-bit unsigned integer in little-endian format.
func (r *LittleEndianReader) ReadUint16() (uint16, error) {
	var b [2]byte
	_, err := io.ReadFull(r, b[:])
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(b[:]), nil
}

// ReadUint32 reads a 32-bit unsigned integer in little-endian format.
func (r *LittleEndianReader) ReadUint32() (uint32, error) {
	var b [4]byte
	_, err := io.ReadFull(r, b[:])
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(b[:]), nil
}

// ReadUint64 reads a 64-bit unsigned integer in little-endian format.
func (r *LittleEndianReader) ReadUint64() (uint64, error) {
	var b [8]byte
	_, err := io.ReadFull(r, b[:])
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(b[:]), nil
}

// ReadFloat32 reads a 32-bit float encoded as a 32-bit unsigned integer in little-endian format.
func (r *LittleEndianReader) ReadFloat32() (float32, error) {
	var b [4]byte
	_, err := io.ReadFull(r, b[:])
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(binary.LittleEndian.Uint32(b[:])), nil
}

// ReadFloat64 reads a 64-bit float encoded as a 64-bit unsigned integer in little-endian format.
func (r *LittleEndianReader) ReadFloat64() (float64, error) {
	var b [8]byte
	_, err := io.ReadFull(r, b[:])
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(binary.LittleEndian.Uint64(b[:])), nil
}
