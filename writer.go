// Package endianio provides utilities for reading and writing binary data in both big-endian and little-endian formats.
package endianio

import (
	"encoding/binary"
	"io"
	"math"
)

// EndianWriter is an interface that defines methods for writing binary data.
type EndianWriter interface {
	// WriteUint8 writes a uint8 (byte)
	WriteUint8(v uint8) (n int, err error)
	// WriteUint16 writes a 16-bit unsigned integer
	WriteUint16(v uint16) (n int, err error)
	// WriteUint32 writes a 32-bit unsigned integer
	WriteUint32(v uint32) (n int, err error)
	// WriteUint64 writes a 64-bit unsigned integer
	WriteUint64(v uint64) (n int, err error)
	// WriteFloat32 writes a 32-bit float
	WriteFloat32(v float32) (n int, err error)
	// WriteFloat64 writes a 64-bit float
	WriteFloat64(v float64) (n int, err error)
}

// baseWriter provides common functionality for both big-endian and little-endian writers.
type baseWriter struct {
	io.Writer
}

// WriteUint8 writes a uint8 (byte)
func (w *baseWriter) WriteUint8(v uint8) (n int, err error) {
	var b [1]byte
	b[0] = v
	return w.Write(b[:])
}

// BigEndianWriter writes binary data in big-endian format.
type BigEndianWriter struct {
	baseWriter
}

// NewBigEndianWriter creates a new BigEndianWriter writing to the provided io.Writer.
func NewBigEndianWriter(w io.Writer) *BigEndianWriter {
	return &BigEndianWriter{baseWriter{w}}
}

// WriteUint16 writes a 16-bit unsigned integer in big-endian format.
func (w *BigEndianWriter) WriteUint16(v uint16) (n int, err error) {
	var b [2]byte
	binary.BigEndian.PutUint16(b[:], v)
	return w.Write(b[:])
}

// WriteUint32 writes a 32-bit unsigned integer in big-endian format.
func (w *BigEndianWriter) WriteUint32(v uint32) (n int, err error) {
	var b [4]byte
	binary.BigEndian.PutUint32(b[:], v)
	return w.Write(b[:])
}

// WriteUint64 writes a 64-bit unsigned integer in big-endian format.
func (w *BigEndianWriter) WriteUint64(v uint64) (n int, err error) {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], v)
	return w.Write(b[:])
}

// WriteFloat32 writes a 32-bit float encoded as a 32-bit unsigned integer in big-endian format.
func (w *BigEndianWriter) WriteFloat32(v float32) (n int, err error) {
	var b [4]byte
	binary.BigEndian.PutUint32(b[:], math.Float32bits(v))
	return w.Write(b[:])
}

// WriteFloat64 writes a 64-bit float encoded as a 64-bit unsigned integer in big-endian format.
func (w *BigEndianWriter) WriteFloat64(v float64) (n int, err error) {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], math.Float64bits(v))
	return w.Write(b[:])
}

// LittleEndianWriter writes binary data in little-endian format.
type LittleEndianWriter struct {
	baseWriter
}

// NewLittleEndianWriter creates a new LittleEndianWriter writing to the provided io.Writer.
func NewLittleEndianWriter(w io.Writer) *LittleEndianWriter {
	return &LittleEndianWriter{baseWriter{w}}
}

// WriteUint16 writes a 16-bit unsigned integer in little-endian format.
func (w *LittleEndianWriter) WriteUint16(v uint16) (n int, err error) {
	var b [2]byte
	binary.LittleEndian.PutUint16(b[:], v)
	return w.Write(b[:])
}

// WriteUint32 writes a 32-bit unsigned integer in little-endian format.
func (w *LittleEndianWriter) WriteUint32(v uint32) (n int, err error) {
	var b [4]byte
	binary.LittleEndian.PutUint32(b[:], v)
	return w.Write(b[:])
}

// WriteUint64 writes a 64-bit unsigned integer in little-endian format.
func (w *LittleEndianWriter) WriteUint64(v uint64) (n int, err error) {
	var b [8]byte
	binary.LittleEndian.PutUint64(b[:], v)
	return w.Write(b[:])
}

// WriteFloat32 writes a 32-bit float encoded as a 32-bit unsigned integer in little-endian format.
func (w *LittleEndianWriter) WriteFloat32(v float32) (n int, err error) {
	var b [4]byte
	binary.LittleEndian.PutUint32(b[:], math.Float32bits(v))
	return w.Write(b[:])
}

// WriteFloat64 writes a 64-bit float encoded as a 64-bit unsigned integer in little-endian format.
func (w *LittleEndianWriter) WriteFloat64(v float64) (n int, err error) {
	var b [8]byte
	binary.LittleEndian.PutUint64(b[:], math.Float64bits(v))
	return w.Write(b[:])
}
