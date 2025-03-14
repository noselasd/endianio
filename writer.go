// Package endianio provides utilities for reading and writing binary data in both big-endian and little-endian formats.
package endianio

import (
	"encoding/binary"
	"io"
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
func (w *Writer) WriteUint8(v uint8) error {
	var b [1]byte
	b[0] = v
	_, err := w.Write(b[:])
	return err
}

// WriteBigUint16 writes a 16-bit unsigned integer in big-endian format.
func (w *Writer) WriteBigUint16(v uint16) error {
	var b [2]byte
	binary.BigEndian.PutUint16(b[:], v)
	_, err := w.Write(b[:])
	return err
}

// WriteBigUint32 writes a 32-bit unsigned integer in big-endian format.
func (w *Writer) WriteBigUint32(v uint32) error {
	var b [4]byte
	binary.BigEndian.PutUint32(b[:], v)
	_, err := w.Write(b[:])
	return err
}

// WriteBigUint64 writes a 64-bit unsigned integer in big-endian format.
func (w *Writer) WriteBigUint64(v uint64) error {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], v)
	_, err := w.Write(b[:])
	return err
}

// WriteLittleUint16 writes a 16-bit unsigned integer in little-endian format.
func (w *Writer) WriteLittleUint16(v uint16) error {
	var b [2]byte
	binary.LittleEndian.PutUint16(b[:], v)
	_, err := w.Write(b[:])
	return err
}

// WriteLittleUint32 writes a 32-bit unsigned integer in little-endian format.
func (w *Writer) WriteLittleUint32(v uint32) error {
	var b [4]byte
	binary.LittleEndian.PutUint32(b[:], v)
	_, err := w.Write(b[:])
	return err
}

// WriteLittleUint64 writes a 64-bit unsigned integer in little-endian format.
func (w *Writer) WriteLittleUint64(v uint64) error {
	var b [8]byte
	binary.LittleEndian.PutUint64(b[:], v)
	_, err := w.Write(b[:])
	return err
}
