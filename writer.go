// Package endianio provides utilities for reading and writing binary data in both big-endian and little-endian formats.
package endianio

import (
	"encoding/binary"
	"io"
	"math"
)

// Writer wraps an io.Writer to provide methods for writing binary data in different endian formats.
type Writer struct {
	io.Writer
}

// NewWriter creates a new Writer writing to the provided io.Writer.
func NewWriter(w io.Writer) *Writer {
	return &Writer{w}
}

// WriteUint8 writes a uint8 (byte)
func (w *Writer) WriteUint8(v uint8) (n int, err error) {
	var b [1]byte
	b[0] = v
	return w.Write(b[:])
}

// WriteBigUint16 writes a 16-bit unsigned integer in big-endian format.
func (w *Writer) WriteBigUint16(v uint16) (n int, err error) {
	var b [2]byte
	binary.BigEndian.PutUint16(b[:], v)
	return w.Write(b[:])
}

// WriteBigUint32 writes a 32-bit unsigned integer in big-endian format.
func (w *Writer) WriteBigUint32(v uint32) (n int, err error) {
	var b [4]byte
	binary.BigEndian.PutUint32(b[:], v)
	return w.Write(b[:])
}

// WriteBigUint64 writes a 64-bit unsigned integer in big-endian format.
func (w *Writer) WriteBigUint64(v uint64) (n int, err error) {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], v)
	return w.Write(b[:])
}

// WriteBigFloat32 writes a 32-bit float encoded as a 32-bit unsigned integer in big-endian format.
func (w *Writer) WriteBigFloat32(v float32) (n int, err error) {
	var b [4]byte
	binary.BigEndian.PutUint32(b[:], math.Float32bits(v))
	return w.Write(b[:])
}

// WriteBigFloat64 writes a 64-bit float encoded as a 64-bit unsigned integer in big-endian format.
func (w *Writer) WriteBigFloat64(v float64) (n int, err error) {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], math.Float64bits(v))
	return w.Write(b[:])
}

// WriteLittleUint16 writes a 16-bit unsigned integer in little-endian format.
func (w *Writer) WriteLittleUint16(v uint16) (n int, err error) {
	var b [2]byte
	binary.LittleEndian.PutUint16(b[:], v)
	return w.Write(b[:])
}

// WriteLittleUint32 writes a 32-bit unsigned integer in little-endian format.
func (w *Writer) WriteLittleUint32(v uint32) (n int, err error) {
	var b [4]byte
	binary.LittleEndian.PutUint32(b[:], v)
	return w.Write(b[:])
}

// WriteLittleUint64 writes a 64-bit unsigned integer in little-endian format.
func (w *Writer) WriteLittleUint64(v uint64) (n int, err error) {
	var b [8]byte
	binary.LittleEndian.PutUint64(b[:], v)
	return w.Write(b[:])
}

// WriteBigFloat32 writes a 32-bit float encoded as a 32-bit unsigned integer in little-endian format.
func (w *Writer) WriteLittleFloat32(v float32) (n int, err error) {
	var b [4]byte
	binary.LittleEndian.PutUint32(b[:], math.Float32bits(v))
	return w.Write(b[:])
}

// WriteBigFloat64 writes a 64-bit float encoded as a 64-bit unsigned integer in little-endian format.
func (w *Writer) WriteLittleFloat64(v float64) (n int, err error) {
	var b [8]byte
	binary.LittleEndian.PutUint64(b[:], math.Float64bits(v))
	return w.Write(b[:])
}
