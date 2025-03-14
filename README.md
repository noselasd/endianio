# endianio

A Go package that provides utilities for reading and writing binary data in both big-endian and little-endian formats.

## Features

- Efficient reading and writing of binary data
- Support for both big-endian and little-endian formats
- Simple API that wraps standard io.Reader and io.Writer interfaces
- Methods for reading and writing uint8, uint16, uint32, and uint64 values

## Installation

```bash
go get github.com/yourusername/endianio
```

## Usage

### Reading binary data

```go
package main

import (
    "bytes"
    "fmt"
    "log"

    "github.com/yourusername/endianio"
)

func main() {
    // Sample binary data in big-endian format
    data := []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xF0}

    // Create a reader
    buf := bytes.NewReader(data)
    reader := endianio.NewReader(buf)

    // Read a 16-bit unsigned integer in big-endian format
    val16, err := reader.ReadBigUint16()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("16-bit value: 0x%04X\n", val16) // Output: 0x1234

    // Read a 32-bit unsigned integer in big-endian format
    val32, err := reader.ReadBigUint32()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("32-bit value: 0x%08X\n", val32) // Output: 0x56789ABC

    // Read a 16-bit unsigned integer in little-endian format
    // (Note: we're creating a new reader with little-endian data)
    leData := []byte{0x34, 0x12, 0x78, 0x56}
    leBuf := bytes.NewReader(leData)
    leReader := endianio.NewReader(leBuf)

    leVal16, err := leReader.ReadLittleUint16()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("16-bit LE value: 0x%04X\n", leVal16) // Output: 0x1234
}
```

### Writing binary data

```go
package main

import (
    "bytes"
    "fmt"

    "github.com/yourusername/endianio"
)

func main() {
    // Create a buffer to write to
    buf := &bytes.Buffer{}

    // Create a writer
    writer := endianio.NewWriter(buf)

    // Write a 16-bit unsigned integer in big-endian format
    err := writer.WriteBigUint16(0x1234)
    if err != nil {
        panic(err)
    }

    // Write a 32-bit unsigned integer in big-endian format
    err = writer.WriteBigUint32(0x56789ABC)
    if err != nil {
        panic(err)
    }

    // Get the written bytes
    data := buf.Bytes()
    fmt.Printf("Written bytes: % X\n", data) // Output: 12 34 56 78 9A BC

    // Create a new buffer for little-endian writing
    leBuf := &bytes.Buffer{}
    leWriter := endianio.NewWriter(leBuf)

    // Write a 16-bit unsigned integer in little-endian format
    err = leWriter.WriteLittleUint16(0x1234)
    if err != nil {
        panic(err)
    }

    // Get the written bytes
    leData := leBuf.Bytes()
    fmt.Printf("Written LE bytes: % X\n", leData) // Output: 34 12
}
```

## Available Methods

### Reader Methods

- `ReadUint8() (uint8, error)` - Read a single byte
- `ReadBigUint16() (uint16, error)` - Read a 16-bit unsigned integer in big-endian format
- `ReadBigUint32() (uint32, error)` - Read a 32-bit unsigned integer in big-endian format
- `ReadBigUint64() (uint64, error)` - Read a 64-bit unsigned integer in big-endian format
- `ReadLittleUint16() (uint16, error)` - Read a 16-bit unsigned integer in little-endian format
- `ReadLittleUint32() (uint32, error)` - Read a 32-bit unsigned integer in little-endian format
- `ReadLittleUint64() (uint64, error)` - Read a 64-bit unsigned integer in little-endian format

### Writer Methods

- `WriteUint8(v uint8) error` - Write a single byte
- `WriteBigUint16(v uint16) error` - Write a 16-bit unsigned integer in big-endian format
- `WriteBigUint32(v uint32) error` - Write a 32-bit unsigned integer in big-endian format
- `WriteBigUint64(v uint64) error` - Write a 64-bit unsigned integer in big-endian format
- `WriteLittleUint16(v uint16) error` - Write a 16-bit unsigned integer in little-endian format
- `WriteLittleUint32(v uint32) error` - Write a 32-bit unsigned integer in little-endian format
- `WriteLittleUint64(v uint64) error` - Write a 64-bit unsigned integer in little-endian format

## License

See the [LICENSE](LICENSE) file for details.
