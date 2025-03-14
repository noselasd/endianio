package endianio

import (
	"bytes"
	"encoding/binary"
	"testing"
)

// Benchmark data - create once and reuse
var (
	// Data for big-endian tests
	bigEndianUint16Data = []byte{0x12, 0x34}                                     // 0x1234
	bigEndianUint32Data = []byte{0x12, 0x34, 0x56, 0x78}                         // 0x12345678
	bigEndianUint64Data = []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xF0} // 0x123456789ABCDEF0

	// Data for little-endian tests
	littleEndianUint16Data = []byte{0x34, 0x12}                                     // 0x1234
	littleEndianUint32Data = []byte{0x78, 0x56, 0x34, 0x12}                         // 0x12345678
	littleEndianUint64Data = []byte{0xF0, 0xDE, 0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12} // 0x123456789ABCDEF0
)

func BenchmarkReadBigUint16(b *testing.B) {
	// Create a bytes reader outside the loop
	br := bytes.NewReader(bigEndianUint16Data)
	r := NewReader(br)

	for b.Loop() {
		br.Reset(bigEndianUint16Data)

		_, err := r.ReadBigUint16()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkReadBigUint32(b *testing.B) {
	br := bytes.NewReader(bigEndianUint32Data)
	r := NewReader(br)

	for b.Loop() {
		br.Reset(bigEndianUint32Data)

		_, err := r.ReadBigUint32()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkReadBigUint64(b *testing.B) {
	br := bytes.NewReader(bigEndianUint64Data)
	r := NewReader(br)

	for b.Loop() {
		br.Reset(bigEndianUint64Data)

		_, err := r.ReadBigUint64()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkReadLittleUint16(b *testing.B) {
	// Create a bytes reader outside the loop
	br := bytes.NewReader(littleEndianUint16Data)
	r := NewReader(br)

	for b.Loop() {
		br.Reset(littleEndianUint16Data)

		_, err := r.ReadLittleUint16()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkReadLittleUint32(b *testing.B) {
	// Create a bytes reader outside the loop
	br := bytes.NewReader(littleEndianUint32Data)
	r := NewReader(br)

	for b.Loop() {
		br.Reset(littleEndianUint32Data)

		_, err := r.ReadLittleUint32()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkReadLittleUint64(b *testing.B) {
	// Create a bytes reader outside the loop
	br := bytes.NewReader(littleEndianUint64Data)
	r := NewReader(br)

	for b.Loop() {
		br.Reset(littleEndianUint64Data)

		_, err := r.ReadLittleUint64()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkReadBigUint16Stdlib(b *testing.B) {
	br := bytes.NewReader(bigEndianUint32Data)

	for b.Loop() {
		br.Reset(bigEndianUint32Data)
		var v uint16
		err := binary.Read(br, binary.LittleEndian, &v)
		if err != nil {
			b.Fatal(err)
		}
	}
}
func BenchmarkReadLittleUint32Stdlib(b *testing.B) {
	br := bytes.NewReader(littleEndianUint32Data)

	for b.Loop() {
		br.Reset(littleEndianUint32Data)
		var v uint32
		err := binary.Read(br, binary.LittleEndian, &v)
		if err != nil {
			b.Fatal(err)
		}
	}
}
func BenchmarkReadLittleUint64Stdlib(b *testing.B) {
	// Create a bytes reader outside the loop
	br := bytes.NewReader(littleEndianUint64Data)

	for b.Loop() {
		br.Reset(littleEndianUint64Data)
		var v uint64
		err := binary.Read(br, binary.LittleEndian, &v)
		if err != nil {
			b.Fatal(err)
		}
	}
}
