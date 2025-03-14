package endianio

import (
	"bytes"
	"encoding/binary"
	"testing"
)

// Benchmark data - create once and reuse
var (
	// Data for big-endian tests
	bigEndianUint16Value uint16 = 0x1234
	bigEndianUint32Value uint32 = 0x12345678
	bigEndianUint64Value uint64 = 0x123456789ABCDEF0

	// Data for little-endian tests
	littleEndianUint16Value uint16 = 0x1234
	littleEndianUint32Value uint32 = 0x12345678
	littleEndianUint64Value uint64 = 0x123456789ABCDEF0
)

func BenchmarkWriteBigUint16(b *testing.B) {
	// Create a bytes buffer outside the loop
	buf := &bytes.Buffer{}
	w := NewWriter(buf)

	for b.Loop() {
		buf.Reset()

		_, err := w.WriteBigUint16(bigEndianUint16Value)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWriteBigUint32(b *testing.B) {
	buf := &bytes.Buffer{}
	w := NewWriter(buf)

	for b.Loop() {
		buf.Reset()

		_, err := w.WriteBigUint32(bigEndianUint32Value)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWriteBigUint64(b *testing.B) {
	buf := &bytes.Buffer{}
	w := NewWriter(buf)

	for b.Loop() {
		buf.Reset()

		_, err := w.WriteBigUint64(bigEndianUint64Value)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWriteLittleUint16(b *testing.B) {
	buf := &bytes.Buffer{}
	w := NewWriter(buf)

	for b.Loop() {
		buf.Reset()

		_, err := w.WriteLittleUint16(littleEndianUint16Value)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWriteLittleUint32(b *testing.B) {
	buf := &bytes.Buffer{}
	w := NewWriter(buf)

	for b.Loop() {
		buf.Reset()

		_, err := w.WriteLittleUint32(littleEndianUint32Value)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWriteLittleUint64(b *testing.B) {
	buf := &bytes.Buffer{}
	w := NewWriter(buf)

	for b.Loop() {
		buf.Reset()

		_, err := w.WriteLittleUint64(littleEndianUint64Value)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWriteBigUint16Stdlib(b *testing.B) {
	buf := &bytes.Buffer{}

	for b.Loop() {
		buf.Reset()
		err := binary.Write(buf, binary.BigEndian, bigEndianUint16Value)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWriteLittleUint32Stdlib(b *testing.B) {
	buf := &bytes.Buffer{}

	for b.Loop() {
		buf.Reset()
		err := binary.Write(buf, binary.LittleEndian, littleEndianUint32Value)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWriteLittleUint64Stdlib(b *testing.B) {
	buf := &bytes.Buffer{}

	for b.Loop() {
		buf.Reset()
		err := binary.Write(buf, binary.LittleEndian, littleEndianUint64Value)
		if err != nil {
			b.Fatal(err)
		}
	}
}
