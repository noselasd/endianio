# endianio

A Go package that provides utilities for reading and writing binary data in both big-endian and little-endian formats.

## Features

- Efficient reading and writing of binary data
- Support for both big-endian and little-endian formats
- Simple API that wraps standard io.Reader and io.Writer interfaces
- Methods for reading and writing uint8, uint16, uint32, uint64, float32, and float64 values
- Specialized readers for big-endian and little-endian formats

## Installation

```bash
go get github.com/noselasd/endianio
```

## Usage

## Available Methods

### Interfaces

The package provides an `EndianReader` and `EndianWriter interface that defines methods for reading
and writingbinary data:

```go
type EndianReader interface {
    ReadUint8() (uint8, error)
    ReadUint16() (uint16, error)
    ReadUint32() (uint32, error)
    ReadUint64() (uint64, error)
    ReadFloat32() (float32, error)
    ReadFloat64() (float64, error)
}

type EndianWriter interface {
	WriteUint8(v uint8) (n int, err error)
	WriteUint16(v uint16) (n int, err error)
	WriteUint32(v uint32) (n int, err error)
	WriteUint64(v uint64) (n int, err error)
	WriteFloat32(v float32) (n int, err error)
	WriteFloat64(v float64) (n int, err error)
}
```
The Writer methods return the number of bytes written. If everything is written ok, the `err` return is
nil. If an error occured, the `err` return is non-nil.

### BigEndianReader/LittleEndianReader Methods

- `ReadUint8() (uint8, error)` - Read a single byte
- `ReadUint16() (uint16, error)` - Read a 16-bit unsigned integer in the relevant endian format
- `ReadUint32() (uint32, error)` - Read a 32-bit unsigned integer in the relevant endian format
- `ReadUint64() (uint64, error)` - Read a 64-bit unsigned integer in the relevant endian format
- `ReadFloat32() (float32, error)` - Read a 32-bit float in the relevant endian format
- `ReadFloat64() (float64, error)` - Read a 64-bit float in the relevant endian format

float32/float64 are written with the bit pattern of the float converted to an uint32/uint64

### BigEndianWriter/LittleEndianWriter Methods

- `WriteUint8(v uint8) (n int, err error)` - Write a single byte
- `WriteUint16(v uint16) (n int, err error)` - Write a 16-bit unsigned integer the relevant endian format
- `WriteUint32(v uint32) (n int, err error)` - Write a 32-bit unsigned integer the relevant endian format
- `WriteUint64(v uint64) (n int, err error)` - Write a 64-bit unsigned integer the relevant endian format
- `WriteFloat32(v float32) (n int, err error)` - Write a 32-bit float the relevant endian format
- `WriteFloat64(v float64) (n int, err error)` - Write a 64-bit float the relevant endian format

float32/float64 are read with the bit pattern of an uint32/uint64 then converted to a float

### Reading binary data

```go
package main

import (
    "bytes"
    "fmt"
    "log"

    "github.com/noselasd/endianio"
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

### Using specialized endian readers

```go
package main

import (
    "bytes"
    "fmt"
    "log"

    "github.com/noselasd/endianio"
)

func main() {
    // Sample binary data
    data := []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xF0}

    // Create a big-endian reader
    buf := bytes.NewReader(data)
    reader := endianio.NewBigEndianReader(buf)

    // Read values in big-endian format
    val16, err := reader.ReadUint16()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("16-bit value: 0x%04X\n", val16) // Output: 0x1234

    val32, err := reader.ReadUint32()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("32-bit value: 0x%08X\n", val32) // Output: 0x56789ABC

    // Create a little-endian reader for the same data
    buf.Reset(data)
    leReader := endianio.NewLittleEndianReader(buf)

    // Read values in little-endian format
    leVal16, err := leReader.ReadUint16()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("16-bit LE value: 0x%04X\n", leVal16) // Output: 0x3412

    leVal32, err := leReader.ReadUint32()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("32-bit LE value: 0x%08X\n", leVal32) // Output: 0x78563412
}
```

### Writing binary data

```go
package main

import (
    "bytes"
    "fmt"

    "github.com/noselasd/endianio"
)

func main() {
    // Create a buffer to write to
    buf := &bytes.Buffer{}

    // Create a writer
    writer := endianio.NewWriter(buf)

    // Write a 16-bit unsigned integer in big-endian format
    _,err := writer.WriteBigUint16(0x1234)
    if err != nil {
        panic(err)
    }

    // Write a 32-bit unsigned integer in big-endian format
    _,err = writer.WriteBigUint32(0x56789ABC)
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
    _,err = leWriter.WriteLittleUint16(0x1234)
    if err != nil {
        panic(err)
    }

    // Get the written bytes
    leData := leBuf.Bytes()
    fmt.Printf("Written LE bytes: % X\n", leData) // Output: 34 12
}
```



## License

MIT License

See the [LICENSE](LICENSE) file for details.
