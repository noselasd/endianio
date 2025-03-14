package endianio

import (
	"bytes"
	"encoding/binary"
	"math"
	"testing"
)

// failingWriter is a writer that always fails
type failingWriter struct{}

func (w *failingWriter) Write(p []byte) (n int, err error) {
	return 0, bytes.ErrTooLarge // Using an existing error for simplicity
}
func TestWriter_WriteBig(t *testing.T) {
	// Test WriteUint8
	t.Run("WriteUint8", func(t *testing.T) {
		var tests = []struct {
			name  string
			value uint8
			want  []byte
		}{
			{"Uint8_1", 0x00, []byte{0x00}},
			{"Uint8_2", 0xff, []byte{0xff}},
			{"Uint8_3", 0x01, []byte{0x01}},
			{"Uint8_4", 0x7F, []byte{0x7F}},
			{"Uint8_5", 0x80, []byte{0x80}},
			{"Uint8_6", 0xA5, []byte{0xA5}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				buf := &bytes.Buffer{}
				w := NewWriter(buf)
				_, err := w.WriteUint8(tt.value)
				if err != nil {
					t.Errorf("WriteUint8() error = %v", err)
					return
				}
				got := buf.Bytes()
				if !bytes.Equal(got, tt.want) {
					t.Errorf("WriteUint8() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test WriteBigUint16
	t.Run("WriteBigUint16", func(t *testing.T) {
		var tests = []struct {
			name  string
			value uint16
			want  []byte
		}{
			{"Uint16_1", 0x0000, []byte{0x00, 0x00}},
			{"Uint16_2", 0xffff, []byte{0xff, 0xff}},
			{"Uint16_3", 0x0001, []byte{0x00, 0x01}},
			{"Uint16_4", 0x1000, []byte{0x10, 0x00}},
			{"Uint16_5", 0x1234, []byte{0x12, 0x34}},
			{"Uint16_6", 0x3412, []byte{0x34, 0x12}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				buf := &bytes.Buffer{}
				w := NewWriter(buf)
				_, err := w.WriteBigUint16(tt.value)
				if err != nil {
					t.Errorf("WriteBigUint16() error = %v", err)
					return
				}
				got := buf.Bytes()
				if !bytes.Equal(got, tt.want) {
					t.Errorf("WriteBigUint16() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test WriteBigUint32
	t.Run("WriteBigUint32", func(t *testing.T) {
		var tests = []struct {
			name  string
			value uint32
			want  []byte
		}{
			{"Uint32_1", 0x00000000, []byte{0x00, 0x00, 0x00, 0x00}},
			{"Uint32_2", 0xffffffff, []byte{0xff, 0xff, 0xff, 0xff}},
			{"Uint32_3", 0x00000001, []byte{0x00, 0x00, 0x00, 0x01}},
			{"Uint32_4", 0x10000000, []byte{0x10, 0x00, 0x00, 0x00}},
			{"Uint32_5", 0x12345678, []byte{0x12, 0x34, 0x56, 0x78}},
			{"Uint32_6", 0x78563412, []byte{0x78, 0x56, 0x34, 0x12}},
			{"Uint32_7", 0xA0B0C0D0, []byte{0xA0, 0xB0, 0xC0, 0xD0}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				buf := &bytes.Buffer{}
				w := NewWriter(buf)
				_, err := w.WriteBigUint32(tt.value)
				if err != nil {
					t.Errorf("WriteLittleUint32() error = %v", err)
					return
				}
				got := buf.Bytes()
				if !bytes.Equal(got, tt.want) {
					t.Errorf("WriteLittleUint32() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test WriteBigUint64
	t.Run("WriteBigUint64", func(t *testing.T) {
		var tests = []struct {
			name  string
			value uint64
			want  []byte
		}{
			{"Uint64_1", 0x0000000000000000, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
			{"Uint64_2", 0xffffffffffffffff, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
			{"Uint64_3", 0x0000000000000001, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}},
			{"Uint64_4", 0x1000000000000000, []byte{0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
			{"Uint64_5", 0x123456789ABCDEF0, []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xF0}},
			{"Uint64_6", 0xF0DEBC9A78563412, []byte{0xF0, 0xDE, 0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12}},
			{"Uint64_7", 0xA0B0C0D0E0F0A1B1, []byte{0xA0, 0xB0, 0xC0, 0xD0, 0xE0, 0xF0, 0xA1, 0xB1}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				buf := &bytes.Buffer{}
				w := NewWriter(buf)
				_, err := w.WriteBigUint64(tt.value)
				if err != nil {
					t.Errorf("WriteBigUint64() error = %v", err)
					return
				}
				got := buf.Bytes()
				if !bytes.Equal(got, tt.want) {
					t.Errorf("WriteBigUint64() got = %v, want %v", got, tt.want)
				}
			})
		}
	})
	t.Run("WriteBigFloat32", func(t *testing.T) {
		var tests = []struct {
			name  string
			value float32
			want  []byte
		}{
			{"float32_1", 0.0, []byte{0x00, 0x00, 0x00, 0x00}},
			{"float32_2", 0.1, []byte{0x3d, 0xcc, 0xcc, 0xcd}},
			{"float32_3", -0.1, []byte{0xbd, 0xcc, 0xcc, 0xcd}},
			{"float32_4", float32(math.Inf(-1)), []byte{0xff, 0x80, 0x00, 0x00}},
			{"float32_5", float32(math.NaN()), []byte{0x7f, 0xc0, 0x00, 0x00}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				buf := &bytes.Buffer{}
				w := NewWriter(buf)
				_, err := w.WriteBigFloat32(tt.value)
				if err != nil {
					t.Errorf("WriteBigUint64() error = %v", err)
					return
				}
				got := buf.Bytes()
				if !bytes.Equal(got, tt.want) {
					t.Errorf("WriteBigUint64() got = %v, want %v", got, tt.want)
				}
			})
		}
	})
	// Test WriteBigUint64
	t.Run("WriteBigFloat64", func(t *testing.T) {
		var tests = []struct {
			name  string
			value float64
			want  []byte
		}{
			{"float64_1", 0.0, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
			{"float64_2", 0.1, []byte{0x3f, 0xb9, 0x99, 0x99, 0x99, 0x99, 0x99, 0x9a}},
			{"float64_3", -0.1, []byte{0xbf, 0xb9, 0x99, 0x99, 0x99, 0x99, 0x99, 0x9a}},
			{"float64_4", math.Inf(-1), []byte{0xff, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
			{"float64_5", math.NaN(), []byte{0x7f, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				buf := &bytes.Buffer{}
				w := NewWriter(buf)
				_, err := w.WriteBigFloat64(tt.value)
				if err != nil {
					t.Errorf("WriteBigUint64() error = %v", err)
					return
				}
				got := buf.Bytes()
				if !bytes.Equal(got, tt.want) {
					t.Errorf("WriteBigUint64() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test error cases
	t.Run("ErrorCases", func(t *testing.T) {
		// Test writing to a failing writer
		t.Run("FailingWriter", func(t *testing.T) {
			failWriter := &failingWriter{}
			w := NewWriter(failWriter)

			_, err := w.WriteUint8(0x12)
			if err == nil {
				t.Errorf("WriteUint8() expected error for failing writer")
			}

			_, err = w.WriteBigUint16(0x1234)
			if err == nil {
				t.Errorf("WriteBigUint16() expected error for failing writer")
			}

			_, err = w.WriteBigUint32(0x12345678)
			if err == nil {
				t.Errorf("WriteBigUint32() expected error for failing writer")
			}

			_, err = w.WriteBigUint64(0x123456789ABCDEF0)
			if err == nil {
				t.Errorf("WriteBigUint64() expected error for failing writer")
			}
		})
	})
}

func TestWriter_WriteLittle(t *testing.T) {
	// Test WriteLittleUint16
	t.Run("WriteLittleUint16", func(t *testing.T) {
		var tests = []struct {
			name  string
			value uint16
			want  []byte
		}{
			{"Uint16_1", 0x0000, []byte{0x00, 0x00}},
			{"Uint16_2", 0xffff, []byte{0xff, 0xff}},
			{"Uint16_3", 0x0001, []byte{0x01, 0x00}},
			{"Uint16_4", 0x1000, []byte{0x00, 0x10}},
			{"Uint16_5", 0x1234, []byte{0x34, 0x12}},
			{"Uint16_6", 0x3412, []byte{0x12, 0x34}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				buf := &bytes.Buffer{}
				w := NewWriter(buf)
				n, err := w.WriteLittleUint16(tt.value)
				if err != nil {
					t.Errorf("WriteLittleUint16() error = %v", err)
					return
				}
				if n != 2 {
					t.Errorf("WriteLittleUint16() wrote %d bytes, want 2", n)
				}
				if !bytes.Equal(buf.Bytes(), tt.want) {
					t.Errorf("WriteLittleUint16() got = %v, want %v", buf.Bytes(), tt.want)
				}
			})
		}
	})

	// Test WriteLittleUint32
	t.Run("WriteLittleUint32", func(t *testing.T) {
		var tests = []struct {
			name  string
			value uint32
			want  []byte
		}{
			{"Uint32_1", 0x00000000, []byte{0x00, 0x00, 0x00, 0x00}},
			{"Uint32_2", 0xffffffff, []byte{0xff, 0xff, 0xff, 0xff}},
			{"Uint32_3", 0x00000001, []byte{0x01, 0x00, 0x00, 0x00}},
			{"Uint32_4", 0x10000000, []byte{0x00, 0x00, 0x00, 0x10}},
			{"Uint32_5", 0x12345678, []byte{0x78, 0x56, 0x34, 0x12}},
			{"Uint32_6", 0x78563412, []byte{0x12, 0x34, 0x56, 0x78}},
			{"Uint32_7", 0xA0B0C0D0, []byte{0xD0, 0xC0, 0xB0, 0xA0}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				buf := &bytes.Buffer{}
				w := NewWriter(buf)
				n, err := w.WriteLittleUint32(tt.value)
				if err != nil {
					t.Errorf("WriteLittleUint32() error = %v", err)
					return
				}
				if n != 4 {
					t.Errorf("WriteLittleUint32() wrote %d bytes, want 4", n)
				}
				if !bytes.Equal(buf.Bytes(), tt.want) {
					t.Errorf("WriteLittleUint32() got = %v, want %v", buf.Bytes(), tt.want)
				}
			})
		}
	})

	// Test WriteLittleUint64
	t.Run("WriteLittleUint64", func(t *testing.T) {
		var tests = []struct {
			name  string
			value uint64
			want  []byte
		}{
			{"Uint64_1", 0x0000000000000000, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
			{"Uint64_2", 0xffffffffffffffff, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
			{"Uint64_3", 0x0000000000000001, []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
			{"Uint64_4", 0x1000000000000000, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10}},
			{"Uint64_5", 0x123456789ABCDEF0, []byte{0xF0, 0xDE, 0xBC, 0x9A, 0x78, 0x56, 0x34, 0x12}},
			{"Uint64_6", 0xF0DEBC9A78563412, []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xF0}},
			{"Uint64_7", 0xA0B0C0D0E0F0A1B1, []byte{0xB1, 0xA1, 0xF0, 0xE0, 0xD0, 0xC0, 0xB0, 0xA0}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				buf := &bytes.Buffer{}
				w := NewWriter(buf)
				n, err := w.WriteLittleUint64(tt.value)
				if err != nil {
					t.Errorf("WriteLittleUint64() error = %v", err)
					return
				}
				if n != 8 {
					t.Errorf("WriteLittleUint64() wrote %d bytes, want 8", n)
				}
				if !bytes.Equal(buf.Bytes(), tt.want) {
					t.Errorf("WriteLittleUint64() got = %v, want %v", buf.Bytes(), tt.want)
				}
			})
		}
	})
	t.Run("WriteLittleFloat32", func(t *testing.T) {
		var tests = []struct {
			name  string
			value float32
			want  []byte
		}{
			{"float32_1", 0.0, []byte{0x00, 0x00, 0x00, 0x00}},
			{"float32_2", 0.1, []byte{0xcd, 0xcc, 0xcc, 0x3d}},
			{"float32_3", -0.1, []byte{0xcd, 0xcc, 0xcc, 0xbd}},
			{"float32_4", float32(math.Inf(-1)), []byte{0x00, 0x00, 0x80, 0xff}},
			{"float32_5", float32(math.NaN()), []byte{0x00, 0x00, 0xc0, 0x7f}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				buf := &bytes.Buffer{}
				w := NewWriter(buf)
				_, err := w.WriteLittleFloat32(tt.value)
				if err != nil {
					t.Errorf("WriteLittleUint32() error = %v", err)
					return
				}
				got := buf.Bytes()
				if !bytes.Equal(got, tt.want) {
					t.Errorf("WriteLittleUint32() got = %v, want %v", got, tt.want)
				}
			})
		}
	})
	// Test WriteBigUint64
	t.Run("WriteLittleFloat64", func(t *testing.T) {
		var tests = []struct {
			name  string
			value float64
			want  []byte
		}{
			{"float64_1", 0.0, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
			{"float64_2", 0.1, []byte{0x9a, 0x99, 0x99, 0x99, 0x99, 0x99, 0xb9, 0x3f}},
			{"float64_3", -0.1, []byte{0x9a, 0x99, 0x99, 0x99, 0x99, 0x99, 0xb9, 0xbf}},
			{"float64_4", math.Inf(-1), []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0xff}},
			{"float64_5", math.NaN(), []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x7f}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				buf := &bytes.Buffer{}
				w := NewWriter(buf)
				_, err := w.WriteLittleFloat64(tt.value)
				if err != nil {
					t.Errorf("WriteLittleUint64() error = %v", err)
					return
				}
				got := buf.Bytes()
				if !bytes.Equal(got, tt.want) {
					t.Errorf("WriteLittleUint64() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test error cases
	t.Run("ErrorCases", func(t *testing.T) {
		// Test writing to a writer that fails
		t.Run("FailingWriter", func(t *testing.T) {
			// Use the failingWriter from writer_big_test.go
			failWriter := &failingWriter{}
			w := NewWriter(failWriter)

			_, err := w.WriteLittleUint16(0x1234)
			if err == nil {
				t.Errorf("WriteLittleUint16() expected error for failing writer")
			}

			_, err = w.WriteLittleUint32(0x12345678)
			if err == nil {
				t.Errorf("WriteLittleUint32() expected error for failing writer")
			}

			_, err = w.WriteLittleUint64(0x123456789ABCDEF0)
			if err == nil {
				t.Errorf("WriteLittleUint64() expected error for failing writer")
			}
		})
	})
}

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
