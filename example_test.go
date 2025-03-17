package endianio

import (
	"bytes"
	"fmt"
)

// This example demonstrates how to use the BigEndianReader and LittleEndianReader
func Example_readers() {
	// Create a buffer with some binary data
	data := []byte{
		// Big-endian values
		0x12, 0x34, // uint16: 0x1234
		0x56, 0x78, 0x9A, 0xBC, // uint32: 0x56789ABC

		// Little-endian values
		0x34, 0x12, // uint16: 0x1234
		0xBC, 0x9A, 0x78, 0x56, // uint32: 0x56789ABC
	}

	// Create a reader for the data
	r := bytes.NewReader(data)

	// Create a big-endian reader
	bigEndian := NewBigEndianReader(r)

	// Read big-endian values
	beUint16, _ := bigEndian.ReadUint16()
	beUint32, _ := bigEndian.ReadUint32()

	// Reset the reader to the position for little-endian values
	r.Reset(data)
	r.Seek(6, 0) // Skip the first 6 bytes (big-endian values)

	// Create a little-endian reader
	littleEndian := NewLittleEndianReader(r)

	// Read little-endian values
	leUint16, _ := littleEndian.ReadUint16()
	leUint32, _ := littleEndian.ReadUint32()

	// Print the values
	fmt.Printf("Big-endian uint16: 0x%04X\n", beUint16)
	fmt.Printf("Big-endian uint32: 0x%08X\n", beUint32)
	fmt.Printf("Little-endian uint16: 0x%04X\n", leUint16)
	fmt.Printf("Little-endian uint32: 0x%08X\n", leUint32)

	// Output:
	// Big-endian uint16: 0x1234
	// Big-endian uint32: 0x56789ABC
	// Little-endian uint16: 0x1234
	// Little-endian uint32: 0x56789ABC
}

// This example demonstrates how to use the BigEndianWriter and LittleEndianWriter
func Example_writers() {
	// Create buffers to write to
	bigEndianBuf := &bytes.Buffer{}
	littleEndianBuf := &bytes.Buffer{}

	// Create writers
	bigEndian := NewBigEndianWriter(bigEndianBuf)
	littleEndian := NewLittleEndianWriter(littleEndianBuf)

	// Write values
	bigEndian.WriteUint16(0x1234)
	bigEndian.WriteUint32(0x56789ABC)

	littleEndian.WriteUint16(0x1234)
	littleEndian.WriteUint32(0x56789ABC)

	// Print the written bytes
	fmt.Printf("Big-endian bytes: % X\n", bigEndianBuf.Bytes())
	fmt.Printf("Little-endian bytes: % X\n", littleEndianBuf.Bytes())

	// Output:
	// Big-endian bytes: 12 34 56 78 9A BC
	// Little-endian bytes: 34 12 BC 9A 78 56
}
