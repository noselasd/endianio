package endianio

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestReader_ReadLittle(t *testing.T) {
	// Test ReadLittleUint16
	t.Run("ReadLittleUint16", func(t *testing.T) {
		var tests = []struct {
			name string
			data []byte
			want uint16
		}{
			{"Uint16_1", []byte{0x00, 0x00}, 0x0000},
			{"Uint16_2", []byte{0xff, 0xff}, 0xffff},
			{"Uint16_3", []byte{0x01, 0x00}, 0x0001},
			{"Uint16_4", []byte{0x00, 0x10}, 0x1000},
			{"Uint16_5", []byte{0x34, 0x12}, 0x1234},
			{"Uint16_6", []byte{0x12, 0x34}, 0x3412},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				r := NewReader(bytes.NewReader(tt.data))
				got, err := r.ReadLittleUint16()
				if err != nil {
					t.Errorf("ReadLittleUint16() error = %v", err)
					return
				}
				if got != tt.want {
					t.Errorf("ReadLittleUint16() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test ReadLittleUint32
	t.Run("ReadLittleUint32", func(t *testing.T) {
		var tests = []struct {
			name string
			data []byte
			want uint32
		}{
			{"Uint32_1", []byte{0x00, 0x00, 0x00, 0x00}, 0x00000000},
			{"Uint32_2", []byte{0xff, 0xff, 0xff, 0xff}, 0xffffffff},
			{"Uint32_3", []byte{0x01, 0x00, 0x00, 0x00}, 0x00000001},
			{"Uint32_4", []byte{0x00, 0x00, 0x00, 0x10}, 0x10000000},
			{"Uint32_5", []byte{0x78, 0x56, 0x34, 0x12}, 0x12345678},
			{"Uint32_6", []byte{0x12, 0x34, 0x56, 0x78}, 0x78563412},
			{"Uint32_7", []byte{0xD0, 0xC0, 0xB0, 0xA0}, 0xA0B0C0D0},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				r := NewReader(bytes.NewReader(tt.data))
				got, err := r.ReadLittleUint32()
				if err != nil {
					t.Errorf("ReadLittleUint32() error = %v", err)
					return
				}
				if got != tt.want {
					t.Errorf("ReadLittleUint32() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test ReadLittleUint64
	t.Run("ReadLittleUint64", func(t *testing.T) {
		var tests = []struct {
			name string
			data []byte
			want uint64
		}{
			{"Uint64_1", []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0x0000000000000000},
			{"Uint64_2", []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, 0xffffffffffffffff},
			{"Uint64_3", []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0x0000000000000001},
			{"Uint64_4", []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10}, 0x1000000000000000},
			{"Uint64_5", []byte{0xF0, 0xDE, 0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12}, 0x123456789ABCDEF0},
			{"Uint64_6", []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xF0}, 0xF0DEBC9A78563412},
			{"Uint64_7", []byte{0xB1, 0xA1, 0xF0, 0xE0, 0xD0, 0xC0, 0xB0, 0xA0}, 0xA0B0C0D0E0F0A1B1},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				r := NewReader(bytes.NewReader(tt.data))
				got, err := r.ReadLittleUint64()
				if err != nil {
					t.Errorf("ReadLittleUint64() error = %v", err)
					return
				}
				if got != tt.want {
					t.Errorf("ReadLittleUint64() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test error cases
	t.Run("ErrorCases", func(t *testing.T) {
		// Test reading with insufficient data
		t.Run("InsufficientData", func(t *testing.T) {
			// For uint16 (need 2 bytes)
			r := NewReader(bytes.NewReader([]byte{0x12}))
			_, err := r.ReadLittleUint16()
			if err == nil {
				t.Errorf("ReadLittleUint16() expected error for insufficient data")
			}

			// For uint32 (need 4 bytes)
			r = NewReader(bytes.NewReader([]byte{0x12, 0x34, 0x56}))
			_, err = r.ReadLittleUint32()
			if err == nil {
				t.Errorf("ReadLittleUint32() expected error for insufficient data")
			}

			// For uint64 (need 8 bytes)
			r = NewReader(bytes.NewReader([]byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE}))
			_, err = r.ReadLittleUint64()
			if err == nil {
				t.Errorf("ReadLittleUint64() expected error for insufficient data")
			}
		})

		// Test reading from empty reader
		t.Run("EmptyReader", func(t *testing.T) {
			r := NewReader(bytes.NewReader([]byte{}))

			_, err := r.ReadLittleUint16()
			if err == nil {
				t.Errorf("ReadLittleUint16() expected error for empty reader")
			}

			_, err = r.ReadLittleUint32()
			if err == nil {
				t.Errorf("ReadLittleUint32() expected error for empty reader")
			}

			_, err = r.ReadLittleUint64()
			if err == nil {
				t.Errorf("ReadLittleUint64() expected error for empty reader")
			}
		})
	})
}

func TestReader_ReadBig(t *testing.T) {
	// Test ReadBigUint16
	t.Run("ReadBigUint16", func(t *testing.T) {
		var tests = []struct {
			name string
			data []byte
			want uint16
		}{
			{"Uint16_1", []byte{0x00, 0x00}, 0x0000},
			{"Uint16_2", []byte{0xff, 0xff}, 0xffff},
			{"Uint16_3", []byte{0x00, 0x01}, 0x0001},
			{"Uint16_4", []byte{0x10, 0x00}, 0x1000},
			{"Uint16_5", []byte{0x12, 0x34}, 0x1234},
			{"Uint16_6", []byte{0x34, 0x12}, 0x3412},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				r := NewReader(bytes.NewReader(tt.data))
				got, err := r.ReadBigUint16()
				if err != nil {
					t.Errorf("ReadBigUint16() error = %v", err)
					return
				}
				if got != tt.want {
					t.Errorf("ReadBigUint16() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test ReadBigUint32
	t.Run("ReadBigUint32", func(t *testing.T) {
		var tests = []struct {
			name string
			data []byte
			want uint32
		}{
			{"Uint32_1", []byte{0x00, 0x00, 0x00, 0x00}, 0x00000000},
			{"Uint32_2", []byte{0xff, 0xff, 0xff, 0xff}, 0xffffffff},
			{"Uint32_3", []byte{0x00, 0x00, 0x00, 0x01}, 0x00000001},
			{"Uint32_4", []byte{0x10, 0x00, 0x00, 0x00}, 0x10000000},
			{"Uint32_5", []byte{0x12, 0x34, 0x56, 0x78}, 0x12345678},
			{"Uint32_6", []byte{0x78, 0x56, 0x34, 0x12}, 0x78563412},
			{"Uint32_7", []byte{0xA0, 0xB0, 0xC0, 0xD0}, 0xA0B0C0D0},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				r := NewReader(bytes.NewReader(tt.data))
				got, err := r.ReadBigUint32()
				if err != nil {
					t.Errorf("ReadBigUint32() error = %v", err)
					return
				}
				if got != tt.want {
					t.Errorf("ReadBigUint32() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test ReadBigUint64
	t.Run("ReadBigUint64", func(t *testing.T) {
		var tests = []struct {
			name string
			data []byte
			want uint64
		}{
			{"Uint64_1", []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0x0000000000000000},
			{"Uint64_2", []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, 0xffffffffffffffff},
			{"Uint64_3", []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, 0x0000000000000001},
			{"Uint64_4", []byte{0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0x1000000000000000},
			{"Uint64_5", []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xF0}, 0x123456789ABCDEF0},
			{"Uint64_6", []byte{0xF0, 0xDE, 0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12}, 0xF0DEBC9A78563412},
			{"Uint64_7", []byte{0xA0, 0xB0, 0xC0, 0xD0, 0xE0, 0xF0, 0xA1, 0xB1}, 0xA0B0C0D0E0F0A1B1},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				r := NewReader(bytes.NewReader(tt.data))
				got, err := r.ReadBigUint64()
				if err != nil {
					t.Errorf("ReadBigUint64() error = %v", err)
					return
				}
				if got != tt.want {
					t.Errorf("ReadBigUint64() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test ReadUint8
	t.Run("ReadUint8", func(t *testing.T) {
		var tests = []struct {
			name string
			data []byte
			want uint8
		}{
			{"Uint8_1", []byte{0x00}, 0x00},
			{"Uint8_2", []byte{0xff}, 0xff},
			{"Uint8_3", []byte{0x01}, 0x01},
			{"Uint8_4", []byte{0x7F}, 0x7F},
			{"Uint8_5", []byte{0x80}, 0x80},
			{"Uint8_6", []byte{0xA5}, 0xA5},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				r := NewReader(bytes.NewReader(tt.data))
				got, err := r.ReadUint8()
				if err != nil {
					t.Errorf("ReadUint8() error = %v", err)
					return
				}
				if got != tt.want {
					t.Errorf("ReadUint8() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test error cases
	t.Run("ErrorCases", func(t *testing.T) {
		// Test reading with insufficient data
		t.Run("InsufficientData", func(t *testing.T) {
			// For uint16 (need 2 bytes)
			r := NewReader(bytes.NewReader([]byte{0x12}))
			_, err := r.ReadBigUint16()
			if err == nil {
				t.Errorf("ReadBigUint16() expected error for insufficient data")
			}

			// For uint32 (need 4 bytes)
			r = NewReader(bytes.NewReader([]byte{0x12, 0x34, 0x56}))
			_, err = r.ReadBigUint32()
			if err == nil {
				t.Errorf("ReadBigUint32() expected error for insufficient data")
			}

			// For uint64 (need 8 bytes)
			r = NewReader(bytes.NewReader([]byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE}))
			_, err = r.ReadBigUint64()
			if err == nil {
				t.Errorf("ReadBigUint64() expected error for insufficient data")
			}
		})

		// Test reading from empty reader
		t.Run("EmptyReader", func(t *testing.T) {
			r := NewReader(bytes.NewReader([]byte{}))

			_, err := r.ReadUint8()
			if err == nil {
				t.Errorf("ReadUint8() expected error for empty reader")
			}

			_, err = r.ReadBigUint16()
			if err == nil {
				t.Errorf("ReadBigUint16() expected error for empty reader")
			}

			_, err = r.ReadBigUint32()
			if err == nil {
				t.Errorf("ReadBigUint32() expected error for empty reader")
			}

			_, err = r.ReadBigUint64()
			if err == nil {
				t.Errorf("ReadBigUint64() expected error for empty reader")
			}
		})
	})
}

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
