package endianio

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"testing"
)

type failingReader struct{}

func (fr *failingReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("read failed")
}
func TestLittleEndianReader(t *testing.T) {
	// Test ReadUint16
	t.Run("ReadUint16", func(t *testing.T) {
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
				r := NewLittleEndianReader(bytes.NewReader(tt.data))
				got, err := r.ReadUint16()
				if err != nil {
					t.Errorf("ReadUint16() error = %v", err)
					return
				}
				if got != tt.want {
					t.Errorf("ReadUint16() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test ReadUint32
	t.Run("ReadUint32", func(t *testing.T) {
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
				r := NewLittleEndianReader(bytes.NewReader(tt.data))
				got, err := r.ReadUint32()
				if err != nil {
					t.Errorf("ReadUint32() error = %v", err)
					return
				}
				if got != tt.want {
					t.Errorf("ReadUint32() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test ReadUint64
	t.Run("ReadUint64", func(t *testing.T) {
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
				r := NewLittleEndianReader(bytes.NewReader(tt.data))
				got, err := r.ReadUint64()
				if err != nil {
					t.Errorf("ReadUint64() error = %v", err)
					return
				}
				if got != tt.want {
					t.Errorf("ReadUint64() got = %v, want %v", got, tt.want)
				}
			})
		}
	})
	t.Run("ReadFloat32", func(t *testing.T) {
		var tests = []struct {
			name string
			data []byte
			want float32
		}{
			{"float32_1", []byte{0x00, 0x00, 0x00, 0x00}, 0.0},
			{"float32_2", []byte{0xcd, 0xcc, 0xcc, 0x3d}, 0.1},
			{"float32_3", []byte{0xcd, 0xcc, 0xcc, 0xbd}, -0.1},
			{"float32_4", []byte{0x00, 0x00, 0x80, 0xff}, float32(math.Inf(-1))},
			{"float32_5", []byte{0x00, 0x00, 0xc0, 0x7f}, float32(math.NaN())},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				r := NewLittleEndianReader(bytes.NewReader(tt.data))

				got, err := r.ReadFloat32()
				if err != nil {
					t.Errorf("ReadFloat32() error = %v", err)
					return
				}

				if got != tt.want {
					if math.IsNaN(float64(tt.want)) && !math.IsNaN(float64(got)) {
						t.Errorf("ReadFloat32() got = %v, want %v", got, tt.want)
					}
				}
			})
		}
	})
	// Test ReadFloat64
	t.Run("ReadFloat64", func(t *testing.T) {
		var tests = []struct {
			name string
			data []byte
			want float64
		}{
			{"float64_1", []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0.0},
			{"float64_2", []byte{0x9a, 0x99, 0x99, 0x99, 0x99, 0x99, 0xb9, 0x3f}, 0.1},
			{"float64_3", []byte{0x9a, 0x99, 0x99, 0x99, 0x99, 0x99, 0xb9, 0xbf}, -0.1},
			{"float64_4", []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0xff}, math.Inf(-1)},
			{"float64_5", []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x7f}, math.NaN()},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				r := NewLittleEndianReader(bytes.NewReader(tt.data))

				got, err := r.ReadFloat64()
				if err != nil {
					t.Errorf("ReadFloat64() error = %v", err)
					return
				}
				if got != tt.want {
					if math.IsNaN(float64(tt.want)) && !math.IsNaN(float64(got)) {
						t.Errorf("ReadFloat64() got = %v, want %v", got, tt.want)
					}
				}
			})
		}
	})
	// Test reading from a failing reader
	t.Run("FailingReader", func(t *testing.T) {
		failReader := &failingReader{}
		r := NewLittleEndianReader(failReader)

		_, err := r.ReadUint8()
		if err == nil {
			t.Errorf("ReadUint8() expected error for failing reader")
		}

		_, err = r.ReadUint16()
		if err == nil {
			t.Errorf("ReadUint16() expected error for failing reader")
		}

		_, err = r.ReadUint32()
		if err == nil {
			t.Errorf("ReadUint32() expected error for failing reader")
		}

		_, err = r.ReadUint64()
		if err == nil {
			t.Errorf("ReadUint64() expected error for failing reader")
		}
	})
}

func TestBigEndianReader(t *testing.T) {
	// Test ReadUint16
	t.Run("ReadUint16", func(t *testing.T) {
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
				r := NewBigEndianReader(bytes.NewReader(tt.data))
				got, err := r.ReadUint16()
				if err != nil {
					t.Errorf("ReadUint16() error = %v", err)
					return
				}
				if got != tt.want {
					t.Errorf("ReadUint16() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test ReadUint32
	t.Run("ReadUint32", func(t *testing.T) {
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
				r := NewBigEndianReader(bytes.NewReader(tt.data))
				got, err := r.ReadUint32()
				if err != nil {
					t.Errorf("ReadUint32() error = %v", err)
					return
				}
				if got != tt.want {
					t.Errorf("ReadUint32() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test ReadUint64
	t.Run("ReadUint64", func(t *testing.T) {
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
				r := NewBigEndianReader(bytes.NewReader(tt.data))
				got, err := r.ReadUint64()
				if err != nil {
					t.Errorf("ReadUint64() error = %v", err)
					return
				}
				if got != tt.want {
					t.Errorf("ReadUint64() got = %v, want %v", got, tt.want)
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
				r := NewBigEndianReader(bytes.NewReader(tt.data))
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
	t.Run("ReadFloat32", func(t *testing.T) {
		var tests = []struct {
			name string
			data []byte
			want float32
		}{
			{"float32_1", []byte{0x00, 0x00, 0x00, 0x00}, 0.0},
			{"float32_2", []byte{0x3d, 0xcc, 0xcc, 0xcd}, 0.1},
			{"float32_3", []byte{0xbd, 0xcc, 0xcc, 0xcd}, -0.1},
			{"float32_4", []byte{0xff, 0x80, 0x00, 0x00}, float32(math.Inf(-1))},
			{"float32_5", []byte{0x7f, 0xc0, 0x00, 0x00}, float32(math.NaN())},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				r := NewBigEndianReader(bytes.NewReader(tt.data))

				got, err := r.ReadFloat32()
				if err != nil {
					t.Errorf("ReadFloat32() error = %v", err)
					return
				}

				if got != tt.want {
					if math.IsNaN(float64(tt.want)) && !math.IsNaN(float64(got)) {
						t.Errorf("ReadFloat32() got = %v, want %v", got, tt.want)
					}
				}
			})
		}
	})
	// Test ReadFloat64
	t.Run("ReadFloat64", func(t *testing.T) {
		var tests = []struct {
			name string
			data []byte
			want float64
		}{
			{"float64_1", []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0.0},
			{"float64_2", []byte{0x3f, 0xb9, 0x99, 0x99, 0x99, 0x99, 0x99, 0x9a}, 0.1},
			{"float64_3", []byte{0xbf, 0xb9, 0x99, 0x99, 0x99, 0x99, 0x99, 0x9a}, -0.1},
			{"float64_4", []byte{0xff, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, math.Inf(-1)},
			{"float64_5", []byte{0x7f, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, math.NaN()},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				r := NewBigEndianReader(bytes.NewReader(tt.data))

				got, err := r.ReadFloat64()
				if err != nil {
					t.Errorf("ReadFloat64() error = %v", err)
					return
				}
				if got != tt.want {
					if math.IsNaN(tt.want) && !math.IsNaN(got) {
						t.Errorf("ReadFloat64() got = %v, want %v", got, tt.want)
					}
				}
			})
		}
	})
	// Test reading from a failing reader
	t.Run("FailingReader", func(t *testing.T) {
		failReader := &failingReader{}
		r := NewBigEndianReader(failReader)

		_, err := r.ReadUint8()
		if err == nil {
			t.Errorf("ReadUint8() expected error for failing reader")
		}

		_, err = r.ReadUint16()
		if err == nil {
			t.Errorf("ReadUint16() expected error for failing reader")
		}

		_, err = r.ReadUint32()
		if err == nil {
			t.Errorf("ReadUint32() expected error for failing reader")
		}

		_, err = r.ReadUint64()
		if err == nil {
			t.Errorf("ReadUint64() expected error for failing reader")
		}
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

func BenchmarkBigEndianReader_ReadUint16(b *testing.B) {
	br := bytes.NewReader(bigEndianUint16Data)
	r := NewBigEndianReader(br)

	for b.Loop() {
		br.Reset(bigEndianUint16Data)

		_, err := r.ReadUint16()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBigEndianReader_ReadUint32(b *testing.B) {
	br := bytes.NewReader(bigEndianUint32Data)
	r := NewBigEndianReader(br)

	for b.Loop() {
		br.Reset(bigEndianUint32Data)

		_, err := r.ReadUint32()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBigEndianReader_ReadUint64(b *testing.B) {
	br := bytes.NewReader(bigEndianUint64Data)
	r := NewBigEndianReader(br)

	for b.Loop() {
		br.Reset(bigEndianUint64Data)

		_, err := r.ReadUint64()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLittleEndianReader_ReadUint16(b *testing.B) {
	br := bytes.NewReader(littleEndianUint16Data)
	r := NewLittleEndianReader(br)

	for b.Loop() {
		br.Reset(littleEndianUint16Data)

		_, err := r.ReadUint16()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLittleEndianReader_ReadUint32(b *testing.B) {
	br := bytes.NewReader(littleEndianUint32Data)
	r := NewLittleEndianReader(br)

	for b.Loop() {
		br.Reset(littleEndianUint32Data)

		_, err := r.ReadUint32()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLittleEndianReader_ReadUint64(b *testing.B) {
	br := bytes.NewReader(littleEndianUint64Data)
	r := NewLittleEndianReader(br)

	for b.Loop() {
		br.Reset(littleEndianUint64Data)

		_, err := r.ReadUint64()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkReadUint16Stdlib(b *testing.B) {
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

func BenchmarkReadUint32Stdlib(b *testing.B) {
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

func BenchmarkReadUint64Stdlib(b *testing.B) {
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
