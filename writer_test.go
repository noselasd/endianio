package endianio

import (
	"bytes"
	"fmt"
	"testing"
)

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

type failingWriter struct{}

func (fw *failingWriter) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("write failed")
}
func TestBigEndianWriter(t *testing.T) {
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
				w := NewBigEndianWriter(buf)
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

	// Test WriteUint16
	t.Run("WriteUint16", func(t *testing.T) {
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
				w := NewBigEndianWriter(buf)
				_, err := w.WriteUint16(tt.value)
				if err != nil {
					t.Errorf("WriteUint16() error = %v", err)
					return
				}
				got := buf.Bytes()
				if !bytes.Equal(got, tt.want) {
					t.Errorf("WriteUint16() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test WriteUint32
	t.Run("WriteUint32", func(t *testing.T) {
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
				w := NewBigEndianWriter(buf)
				_, err := w.WriteUint32(tt.value)
				if err != nil {
					t.Errorf("WriteUint32() error = %v", err)
					return
				}
				got := buf.Bytes()
				if !bytes.Equal(got, tt.want) {
					t.Errorf("WriteUint32() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test WriteUint64
	t.Run("WriteUint64", func(t *testing.T) {
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
				w := NewBigEndianWriter(buf)
				_, err := w.WriteUint64(tt.value)
				if err != nil {
					t.Errorf("WriteUint64() error = %v", err)
					return
				}
				got := buf.Bytes()
				if !bytes.Equal(got, tt.want) {
					t.Errorf("WriteUint64() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test WriteFloat32
	t.Run("WriteFloat32", func(t *testing.T) {
		var tests = []struct {
			name  string
			value float32
			want  []byte
		}{
			{"float32_1", 0.0, []byte{0x00, 0x00, 0x00, 0x00}},
			{"float32_2", 0.1, []byte{0x3d, 0xcc, 0xcc, 0xcd}},
			{"float32_3", -0.1, []byte{0xbd, 0xcc, 0xcc, 0xcd}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				buf := &bytes.Buffer{}
				w := NewBigEndianWriter(buf)
				_, err := w.WriteFloat32(tt.value)
				if err != nil {
					t.Errorf("WriteFloat32() error = %v", err)
					return
				}
				got := buf.Bytes()
				if !bytes.Equal(got, tt.want) {
					t.Errorf("WriteFloat32() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test WriteFloat64
	t.Run("WriteFloat64", func(t *testing.T) {
		var tests = []struct {
			name  string
			value float64
			want  []byte
		}{
			{"float64_1", 0.0, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
			{"float64_2", 0.1, []byte{0x3f, 0xb9, 0x99, 0x99, 0x99, 0x99, 0x99, 0x9a}},
			{"float64_3", -0.1, []byte{0xbf, 0xb9, 0x99, 0x99, 0x99, 0x99, 0x99, 0x9a}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				buf := &bytes.Buffer{}
				w := NewBigEndianWriter(buf)
				_, err := w.WriteFloat64(tt.value)
				if err != nil {
					t.Errorf("WriteFloat64() error = %v", err)
					return
				}
				got := buf.Bytes()
				if !bytes.Equal(got, tt.want) {
					t.Errorf("WriteFloat64() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test error cases
	t.Run("ErrorCases", func(t *testing.T) {
		// Test writing to a failing writer
		t.Run("FailingWriter", func(t *testing.T) {
			failWriter := &failingWriter{}
			w := NewBigEndianWriter(failWriter)

			_, err := w.WriteUint8(0x12)
			if err == nil {
				t.Errorf("WriteUint8() expected error for failing writer")
			}

			_, err = w.WriteUint16(0x1234)
			if err == nil {
				t.Errorf("WriteUint16() expected error for failing writer")
			}

			_, err = w.WriteUint32(0x12345678)
			if err == nil {
				t.Errorf("WriteUint32() expected error for failing writer")
			}

			_, err = w.WriteUint64(0x123456789ABCDEF0)
			if err == nil {
				t.Errorf("WriteUint64() expected error for failing writer")
			}
		})
	})
}

func BenchmarkBigEndianWriter_WriteUint16(b *testing.B) {
	buf := &bytes.Buffer{}
	w := NewBigEndianWriter(buf)

	for b.Loop() {
		buf.Reset()

		_, err := w.WriteUint16(bigEndianUint16Value)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBigEndianWriter_WriteUint32(b *testing.B) {
	buf := &bytes.Buffer{}
	w := NewBigEndianWriter(buf)

	for b.Loop() {
		buf.Reset()

		_, err := w.WriteUint32(bigEndianUint32Value)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBigEndianWriter_WriteUint64(b *testing.B) {
	buf := &bytes.Buffer{}
	w := NewBigEndianWriter(buf)

	for b.Loop() {
		buf.Reset()

		_, err := w.WriteUint64(bigEndianUint64Value)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestLittleEndianWriter(t *testing.T) {
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
				w := NewLittleEndianWriter(buf)
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

	// Test WriteUint16
	t.Run("WriteUint16", func(t *testing.T) {
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
				w := NewLittleEndianWriter(buf)
				n, err := w.WriteUint16(tt.value)
				if err != nil {
					t.Errorf("WriteUint16() error = %v", err)
					return
				}
				if n != 2 {
					t.Errorf("WriteUint16() wrote %d bytes, want 2", n)
				}
				if !bytes.Equal(buf.Bytes(), tt.want) {
					t.Errorf("WriteUint16() got = %v, want %v", buf.Bytes(), tt.want)
				}
			})
		}
	})

	// Test WriteUint32
	t.Run("WriteUint32", func(t *testing.T) {
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
				w := NewLittleEndianWriter(buf)
				n, err := w.WriteUint32(tt.value)
				if err != nil {
					t.Errorf("WriteUint32() error = %v", err)
					return
				}
				if n != 4 {
					t.Errorf("WriteUint32() wrote %d bytes, want 4", n)
				}
				if !bytes.Equal(buf.Bytes(), tt.want) {
					t.Errorf("WriteUint32() got = %v, want %v", buf.Bytes(), tt.want)
				}
			})
		}
	})

	// Test WriteUint64
	t.Run("WriteUint64", func(t *testing.T) {
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
				w := NewLittleEndianWriter(buf)
				n, err := w.WriteUint64(tt.value)
				if err != nil {
					t.Errorf("WriteUint64() error = %v", err)
					return
				}
				if n != 8 {
					t.Errorf("WriteUint64() wrote %d bytes, want 8", n)
				}
				if !bytes.Equal(buf.Bytes(), tt.want) {
					t.Errorf("WriteUint64() got = %v, want %v", buf.Bytes(), tt.want)
				}
			})
		}
	})

	// Test WriteFloat32
	t.Run("WriteFloat32", func(t *testing.T) {
		var tests = []struct {
			name  string
			value float32
			want  []byte
		}{
			{"float32_1", 0.0, []byte{0x00, 0x00, 0x00, 0x00}},
			{"float32_2", 0.1, []byte{0xcd, 0xcc, 0xcc, 0x3d}},
			{"float32_3", -0.1, []byte{0xcd, 0xcc, 0xcc, 0xbd}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				buf := &bytes.Buffer{}
				w := NewLittleEndianWriter(buf)
				_, err := w.WriteFloat32(tt.value)
				if err != nil {
					t.Errorf("WriteFloat32() error = %v", err)
					return
				}
				got := buf.Bytes()
				if !bytes.Equal(got, tt.want) {
					t.Errorf("WriteFloat32() got = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test WriteFloat64
	t.Run("WriteFloat64", func(t *testing.T) {
		var tests = []struct {
			name  string
			value float64
			want  []byte
		}{
			{"float64_1", 0.0, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
			{"float64_2", 0.1, []byte{0x9a, 0x99, 0x99, 0x99, 0x99, 0x99, 0xb9, 0x3f}},
			{"float64_3", -0.1, []byte{0x9a, 0x99, 0x99, 0x99, 0x99, 0x99, 0xb9, 0xbf}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				buf := &bytes.Buffer{}
				w := NewLittleEndianWriter(buf)
				_, err := w.WriteFloat64(tt.value)
				if err != nil {
					t.Errorf("WriteFloat64() error = %v", err)
					return
				}
				got := buf.Bytes()
				if !bytes.Equal(got, tt.want) {
					t.Errorf("WriteFloat64() got = %v, want %v", got, tt.want)
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
			w := NewLittleEndianWriter(failWriter)

			_, err := w.WriteUint16(0x1234)
			if err == nil {
				t.Errorf("WriteUint16() expected error for failing writer")
			}

			_, err = w.WriteUint32(0x12345678)
			if err == nil {
				t.Errorf("WriteUint32() expected error for failing writer")
			}

			_, err = w.WriteUint64(0x123456789ABCDEF0)
			if err == nil {
				t.Errorf("WriteUint64() expected error for failing writer")
			}
		})
	})
}

func BenchmarkLittleEndianWriter_WriteUint16(b *testing.B) {
	buf := &bytes.Buffer{}
	w := NewLittleEndianWriter(buf)

	for b.Loop() {
		buf.Reset()

		_, err := w.WriteUint16(littleEndianUint16Value)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLittleEndianWriter_WriteUint32(b *testing.B) {
	buf := &bytes.Buffer{}
	w := NewLittleEndianWriter(buf)

	for b.Loop() {
		buf.Reset()

		_, err := w.WriteUint32(littleEndianUint32Value)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLittleEndianWriter_WriteUint64(b *testing.B) {
	buf := &bytes.Buffer{}
	w := NewLittleEndianWriter(buf)

	for b.Loop() {
		buf.Reset()

		_, err := w.WriteUint64(littleEndianUint64Value)
		if err != nil {
			b.Fatal(err)
		}
	}
}
