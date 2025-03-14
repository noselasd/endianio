package endianio

import (
	"bytes"
	"fmt"
)

// This example demonstrates how to use the Writer and Reader together
func Example() {
	// Create a buffer to write to
	buf := &bytes.Buffer{}

	// Create a writer
	writer := NewWriter(buf)

	// Write some values in different formats
	writer.WriteUint8(0x12)
	writer.WriteBigUint16(0x3456)
	writer.WriteBigUint32(0x789ABCDE)
	writer.WriteLittleUint16(0x1234)
	writer.WriteLittleUint32(0x56789ABC)

	// Now read the values back
	// First, create a reader from the buffer
	reader := NewReader(bytes.NewReader(buf.Bytes()))

	// Read the values in the same order they were written
	uint8Val, _ := reader.ReadUint8()
	bigUint16Val, _ := reader.ReadBigUint16()
	bigUint32Val, _ := reader.ReadBigUint32()
	littleUint16Val, _ := reader.ReadLittleUint16()
	littleUint32Val, _ := reader.ReadLittleUint32()

	// Print the values
	fmt.Printf("uint8: 0x%02X\n", uint8Val)
	fmt.Printf("big uint16: 0x%04X\n", bigUint16Val)
	fmt.Printf("big uint32: 0x%08X\n", bigUint32Val)
	fmt.Printf("little uint16: 0x%04X\n", littleUint16Val)
	fmt.Printf("little uint32: 0x%08X\n", littleUint32Val)

	// Output:
	// uint8: 0x12
	// big uint16: 0x3456
	// big uint32: 0x789ABCDE
	// little uint16: 0x1234
	// little uint32: 0x56789ABC
}
