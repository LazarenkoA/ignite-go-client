package ignite

import (
	"bytes"
	"io"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

func Test_response_ReadByte(t *testing.T) {
	r := &response{message: bytes.NewBuffer([]byte{123})}

	tests := []struct {
		name    string
		r       *response
		want    byte
		wantErr bool
	}{
		{
			name: "1",
			r:    r,
			want: 123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadByte()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadShort(t *testing.T) {
	r := &response{message: bytes.NewBuffer([]byte{0x39, 0x30})}

	tests := []struct {
		name    string
		r       *response
		want    int16
		wantErr bool
	}{
		{
			name: "1",
			r:    r,
			want: 12345,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadShort()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadShort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadShort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadInt(t *testing.T) {
	r := &response{message: bytes.NewBuffer([]byte{0xD2, 0x02, 0x96, 0x49})}

	tests := []struct {
		name    string
		r       *response
		want    int32
		wantErr bool
	}{
		{
			name: "1",
			r:    r,
			want: 1234567890,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadInt()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadLong(t *testing.T) {
	r := &response{message: bytes.NewBuffer([]byte{0x15, 0x81, 0xE9, 0x7D, 0xF4, 0x10, 0x22, 0x11})}

	tests := []struct {
		name    string
		r       *response
		want    int64
		wantErr bool
	}{
		{
			name: "1",
			r:    r,
			want: 1234567890123456789,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadLong()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadLong() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadLong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadFloat(t *testing.T) {
	r := &response{message: bytes.NewBuffer([]byte{0x65, 0x20, 0xf1, 0x47})}

	tests := []struct {
		name    string
		r       *response
		want    float32
		wantErr bool
	}{
		{
			name: "1",
			r:    r,
			want: 123456.789,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadFloat()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadDouble(t *testing.T) {
	r := &response{message: bytes.NewBuffer([]byte{0xad, 0x69, 0x7e, 0x54, 0x34, 0x6f, 0x9d, 0x41})}

	tests := []struct {
		name    string
		r       *response
		want    float64
		wantErr bool
	}{
		{
			name: "1",
			r:    r,
			want: 123456789.12345,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadDouble()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadDouble() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadDouble() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadChar(t *testing.T) {
	r := &response{message: bytes.NewBuffer([]byte{0x41, 0x0})}

	tests := []struct {
		name    string
		r       *response
		want    Char
		wantErr bool
	}{
		{
			name: "1",
			r:    r,
			want: Char('A'),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadChar()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadChar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadBool(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer([]byte{1})}
	r2 := &response{message: bytes.NewBuffer([]byte{0})}
	r3 := &response{message: bytes.NewBuffer([]byte{2})}

	tests := []struct {
		name    string
		r       *response
		want    bool
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: true,
		},
		{
			name: "2",
			r:    r2,
			want: false,
		},
		{
			name:    "3",
			r:       r3,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadBool()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadOString(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer(
		[]byte{9, 0x0B, 0, 0, 0, 0x74, 0x65, 0x73, 0x74, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67})}
	r2 := &response{message: bytes.NewBuffer(
		[]byte{9, 0, 0, 0, 0})}
	r3 := &response{message: bytes.NewBuffer(
		[]byte{101})}
	r4 := &response{message: bytes.NewBuffer(
		[]byte{0, 0x0B, 0, 0, 0, 0x74, 0x65, 0x73, 0x74, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67})}

	tests := []struct {
		name    string
		r       *response
		want    string
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: "test string",
		},
		{
			name: "2",
			r:    r2,
			want: "",
		},
		{
			name: "3",
			r:    r3,
		},
		{
			name:    "4",
			r:       r4,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadOString()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadOString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadOString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadUUID(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer(
		[]byte{0xd6, 0x58, 0x9d, 0xa7, 0xf8, 0xb1, 0x46, 0x87, 0xb5,
			0xbd, 0x2d, 0xdc, 0x73, 0x62, 0xa4, 0xa4}[:])}
	v, _ := uuid.Parse("d6589da7-f8b1-4687-b5bd-2ddc7362a4a4")

	tests := []struct {
		name    string
		r       *response
		want    uuid.UUID
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: v,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadUUID()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response.ReadUUID() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadDate(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer([]byte{0x0, 0xa0, 0xcd, 0x88, 0x62, 0x1, 0x0, 0x0})}
	dm := time.Date(2018, 4, 3, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name    string
		r       *response
		want    time.Time
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: dm,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadDate()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response.ReadDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadArrayBytes(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer([]byte{3, 0, 0, 0, 1, 2, 3})}

	tests := []struct {
		name    string
		r       *response
		want    []byte
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: []byte{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadArrayBytes()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadArrayBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response.ReadArrayBytes() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadArrayShorts(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer([]byte{3, 0, 0, 0, 1, 0, 2, 0, 3, 0})}

	tests := []struct {
		name    string
		r       *response
		want    []int16
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: []int16{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadArrayShorts()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadArrayShorts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response.ReadArrayShorts() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadArrayInts(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer(
		[]byte{3, 0, 0, 0, 1, 0, 0, 0, 2, 0, 0, 0, 3, 0, 0, 0})}

	tests := []struct {
		name    string
		r       *response
		want    []int32
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: []int32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadArrayInts()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadArrayInts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response.ReadArrayInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadArrayLongs(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer(
		[]byte{3, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0})}

	tests := []struct {
		name    string
		r       *response
		want    []int64
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: []int64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadArrayLongs()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadArrayLongs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response.ReadArrayLongs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadArrayFloats(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer(
		[]byte{3, 0x0, 0x0, 0x0, 0xcd, 0xcc, 0x8c, 0x3f, 0xcd, 0xcc, 0xc, 0x40, 0x33, 0x33, 0x53, 0x40})}

	tests := []struct {
		name    string
		r       *response
		want    []float32
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: []float32{1.1, 2.2, 3.3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadArrayFloats()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadArrayFloats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response.ReadArrayFloats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadArrayDoubles(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer(
		[]byte{3, 0, 0, 0, 0x9a, 0x99, 0x99, 0x99, 0x99, 0x99, 0xf1, 0x3f, 0x9a, 0x99,
			0x99, 0x99, 0x99, 0x99, 0x1, 0x40, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0xa, 0x40})}

	tests := []struct {
		name    string
		r       *response
		want    []float64
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: []float64{1.1, 2.2, 3.3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadArrayDoubles()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadArrayDoubles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response.ReadArrayDoubles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadArrayChars(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer([]byte{3, 0, 0, 0, 0x41, 0x0, 0x42, 0x0, 0x2f, 0x4})}

	tests := []struct {
		name    string
		r       *response
		want    []Char
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: []Char{'A', 'B', 'Я'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadArrayChars()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadArrayChars() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response.ReadArrayChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadArrayBools(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer([]byte{3, 0, 0, 0, 1, 0, 1})}

	tests := []struct {
		name    string
		r       *response
		want    []bool
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: []bool{true, false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadArrayBools()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadArrayBools() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response.ReadArrayBools() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadArrayOStrings(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer([]byte{3, 0, 0, 0,
		0x9, 3, 0, 0, 0, 0x6f, 0x6e, 0x65,
		0x9, 3, 0, 0, 0, 0x74, 0x77, 0x6f,
		0x9, 6, 0, 0, 0, 0xd1, 0x82, 0xd1, 0x80, 0xd0, 0xb8})}

	tests := []struct {
		name    string
		r       *response
		want    []string
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: []string{"one", "two", "три"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadArrayOStrings()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadArrayOStrings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response.ReadArrayOStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadArrayOUUIDs(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer([]byte{3, 0, 0, 0,
		10, 0xa0, 0xc0, 0x7c, 0x4c, 0x7e, 0x2e, 0x43, 0xd3, 0x8e, 0xda, 0x17, 0x68, 0x81, 0x47, 0x7c, 0x81,
		10, 0x40, 0x15, 0xb5, 0x5f, 0x72, 0xf0, 0x48, 0xa4, 0x8d, 0x1, 0x64, 0x16, 0x8d, 0x50, 0xf6, 0x27,
		10, 0x82, 0x7d, 0x1b, 0xf0, 0xc5, 0xd4, 0x44, 0x43, 0x87, 0x8, 0xd8, 0xb5, 0xde, 0x31, 0xfe, 0x74})}
	uid1, _ := uuid.Parse("a0c07c4c-7e2e-43d3-8eda-176881477c81")
	uid2, _ := uuid.Parse("4015b55f-72f0-48a4-8d01-64168d50f627")
	uid3, _ := uuid.Parse("827d1bf0-c5d4-4443-8708-d8b5de31fe74")

	tests := []struct {
		name    string
		r       *response
		want    []uuid.UUID
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: []uuid.UUID{uid1, uid2, uid3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadArrayOUUIDs()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadArrayOUUIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response.ReadArrayOUUIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadArrayODates(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer([]byte{3, 0, 0, 0,
		11, 0x0, 0xa0, 0xcd, 0x88, 0x62, 0x1, 0x0, 0x0,
		11, 0x0, 0xf0, 0x23, 0x80, 0x6a, 0x1, 0x0, 0x0,
		11, 0x0, 0xf8, 0xc6, 0x81, 0x72, 0x1, 0x0, 0x0})}
	dm1 := time.Date(2018, 4, 3, 0, 0, 0, 0, time.UTC)
	dm2 := time.Date(2019, 5, 4, 0, 0, 0, 0, time.UTC)
	dm3 := time.Date(2020, 6, 5, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name    string
		r       *response
		want    []time.Time
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: []time.Time{dm1, dm2, dm3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadArrayODates()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadArrayODates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response.ReadArrayODates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadTimestamp(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer(
		[]byte{0xdb, 0xb, 0xe6, 0x8b, 0x62, 0x1, 0x0, 0x0, 0x55, 0xf8, 0x6, 0x0})}
	tm := time.Date(2018, 4, 3, 14, 25, 32, int(time.Millisecond*123+time.Microsecond*456+789), time.UTC)

	tests := []struct {
		name    string
		r       *response
		want    time.Time
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: tm,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadTimestamp()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadTimestamp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response.ReadTimestamp() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadArrayOTimestamps(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer([]byte{3, 0, 0, 0,
		33, 0xdb, 0xb, 0xe6, 0x8b, 0x62, 0x1, 0x0, 0x0, 0x55, 0xf8, 0x6, 0x0,
		33, 0xa3, 0x38, 0x74, 0x83, 0x6a, 0x1, 0x0, 0x0, 0x55, 0xf8, 0x6, 0x0,
		33, 0x6b, 0x1d, 0x4f, 0x85, 0x72, 0x1, 0x0, 0x0, 0x55, 0xf8, 0x6, 0x0})}
	tm1 := time.Date(2018, 4, 3, 14, 25, 32, int(time.Millisecond*123+time.Microsecond*456+789), time.UTC)
	tm2 := time.Date(2019, 5, 4, 15, 26, 33, int(time.Millisecond*123+time.Microsecond*456+789), time.UTC)
	tm3 := time.Date(2020, 6, 5, 16, 27, 34, int(time.Millisecond*123+time.Microsecond*456+789), time.UTC)

	tests := []struct {
		name    string
		r       *response
		want    []time.Time
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: []time.Time{tm1, tm2, tm3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadArrayOTimestamps()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadArrayOTimestamps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response.ReadArrayOTimestamps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadTime(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer([]byte{0xdb, 0x6b, 0x18, 0x3, 0x0, 0x0, 0x0, 0x0})}
	tm := time.Date(1, 1, 1, 14, 25, 32, int(time.Millisecond*123), time.UTC)

	tests := []struct {
		name    string
		r       *response
		want    time.Time
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: tm,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadTime()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response.ReadTime() = %s, want %s", got.String(), tt.want.String())
				// t.Errorf("response.ReadTime() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadArrayOTimes(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer([]byte{3, 0, 0, 0,
		36, 0xdb, 0x6b, 0x18, 0x3, 0x0, 0x0, 0x0, 0x0,
		36, 0xa3, 0x48, 0x50, 0x3, 0x0, 0x0, 0x0, 0x0,
		36, 0x6b, 0x25, 0x88, 0x3, 0x0, 0x0, 0x0, 0x0})}
	tm4 := time.Date(1, 1, 1, 14, 25, 32, int(time.Millisecond*123), time.UTC)
	tm5 := time.Date(1, 1, 1, 15, 26, 33, int(time.Millisecond*123), time.UTC)
	tm6 := time.Date(1, 1, 1, 16, 27, 34, int(time.Millisecond*123), time.UTC)

	tests := []struct {
		name    string
		r       *response
		want    []time.Time
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: []time.Time{tm4, tm5, tm6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadArrayOTimes()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadArrayOTimes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response.ReadArrayOTimes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadObject(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer([]byte{1, 123})}
	r2 := &response{message: bytes.NewBuffer([]byte{2, 0x39, 0x30})}
	r3 := &response{message: bytes.NewBuffer([]byte{3, 0xD2, 0x02, 0x96, 0x49})}
	r4 := &response{message: bytes.NewBuffer([]byte{4, 0x15, 0x81, 0xE9, 0x7D, 0xF4, 0x10, 0x22, 0x11})}
	r5 := &response{message: bytes.NewBuffer([]byte{5, 0x65, 0x20, 0xf1, 0x47})}
	r6 := &response{message: bytes.NewBuffer([]byte{6, 0xad, 0x69, 0x7e, 0x54, 0x34, 0x6f, 0x9d, 0x41})}
	r7 := &response{message: bytes.NewBuffer([]byte{7, 0x41, 0x0})}
	r8 := &response{message: bytes.NewBuffer([]byte{8, 1})}
	r9 := &response{message: bytes.NewBuffer(
		[]byte{9, 0x0B, 0, 0, 0, 0x74, 0x65, 0x73, 0x74, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67})}
	r10 := &response{message: bytes.NewBuffer([]byte{10, 0xd6, 0x58, 0x9d, 0xa7, 0xf8, 0xb1, 0x46, 0x87, 0xb5,
		0xbd, 0x2d, 0xdc, 0x73, 0x62, 0xa4, 0xa4})}
	uid, _ := uuid.Parse("d6589da7-f8b1-4687-b5bd-2ddc7362a4a4")
	r11 := &response{message: bytes.NewBuffer([]byte{11, 0x0, 0xa0, 0xcd, 0x88, 0x62, 0x1, 0x0, 0x0})}
	dm := time.Date(2018, 4, 3, 0, 0, 0, 0, time.UTC)
	r12 := &response{message: bytes.NewBuffer([]byte{12, 3, 0, 0, 0, 1, 2, 3})}
	r13 := &response{message: bytes.NewBuffer([]byte{13, 3, 0, 0, 0, 1, 0, 2, 0, 3, 0})}
	r14 := &response{message: bytes.NewBuffer(
		[]byte{14, 3, 0, 0, 0, 1, 0, 0, 0, 2, 0, 0, 0, 3, 0, 0, 0})}
	r15 := &response{message: bytes.NewBuffer(
		[]byte{15, 3, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0})}
	r16 := &response{message: bytes.NewBuffer(
		[]byte{16, 0x3, 0x0, 0x0, 0x0, 0xcd, 0xcc, 0x8c, 0x3f, 0xcd, 0xcc, 0xc, 0x40, 0x33, 0x33, 0x53, 0x40})}
	r17 := &response{message: bytes.NewBuffer(
		[]byte{17, 3, 0, 0, 0, 0x9a, 0x99, 0x99, 0x99, 0x99, 0x99, 0xf1, 0x3f, 0x9a, 0x99,
			0x99, 0x99, 0x99, 0x99, 0x1, 0x40, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0xa, 0x40})}
	r18 := &response{message: bytes.NewBuffer([]byte{18, 3, 0, 0, 0, 0x41, 0x0, 0x42, 0x0, 0x2f, 0x4})}
	r19 := &response{message: bytes.NewBuffer([]byte{19, 3, 0, 0, 0, 1, 0, 1})}
	r20 := &response{message: bytes.NewBuffer([]byte{20, 3, 0, 0, 0,
		0x9, 3, 0, 0, 0, 0x6f, 0x6e, 0x65,
		0x9, 3, 0, 0, 0, 0x74, 0x77, 0x6f,
		0x9, 6, 0, 0, 0, 0xd1, 0x82, 0xd1, 0x80, 0xd0, 0xb8})}
	r21 := &response{message: bytes.NewBuffer([]byte{21, 3, 0, 0, 0,
		10, 0xa0, 0xc0, 0x7c, 0x4c, 0x7e, 0x2e, 0x43, 0xd3, 0x8e, 0xda, 0x17, 0x68, 0x81, 0x47, 0x7c, 0x81,
		10, 0x40, 0x15, 0xb5, 0x5f, 0x72, 0xf0, 0x48, 0xa4, 0x8d, 0x1, 0x64, 0x16, 0x8d, 0x50, 0xf6, 0x27,
		10, 0x82, 0x7d, 0x1b, 0xf0, 0xc5, 0xd4, 0x44, 0x43, 0x87, 0x8, 0xd8, 0xb5, 0xde, 0x31, 0xfe, 0x74})}
	uid1, _ := uuid.Parse("a0c07c4c-7e2e-43d3-8eda-176881477c81")
	uid2, _ := uuid.Parse("4015b55f-72f0-48a4-8d01-64168d50f627")
	uid3, _ := uuid.Parse("827d1bf0-c5d4-4443-8708-d8b5de31fe74")
	r22 := &response{message: bytes.NewBuffer([]byte{22, 3, 0, 0, 0,
		11, 0x0, 0xa0, 0xcd, 0x88, 0x62, 0x1, 0x0, 0x0,
		11, 0x0, 0xf0, 0x23, 0x80, 0x6a, 0x1, 0x0, 0x0,
		11, 0x0, 0xf8, 0xc6, 0x81, 0x72, 0x1, 0x0, 0x0})}
	dm1 := time.Date(2018, 4, 3, 0, 0, 0, 0, time.UTC)
	dm2 := time.Date(2019, 5, 4, 0, 0, 0, 0, time.UTC)
	dm3 := time.Date(2020, 6, 5, 0, 0, 0, 0, time.UTC)
	r33 := &response{message: bytes.NewBuffer([]byte{33, 0xdb, 0xb, 0xe6, 0x8b, 0x62, 0x1, 0x0, 0x0,
		0x55, 0xf8, 0x6, 0x0})}
	tm := time.Date(2018, 4, 3, 14, 25, 32, int(time.Millisecond*123+time.Microsecond*456+789), time.UTC)
	r34 := &response{message: bytes.NewBuffer([]byte{34, 3, 0, 0, 0,
		33, 0xdb, 0xb, 0xe6, 0x8b, 0x62, 0x1, 0x0, 0x0, 0x55, 0xf8, 0x6, 0x0,
		33, 0xa3, 0x38, 0x74, 0x83, 0x6a, 0x1, 0x0, 0x0, 0x55, 0xf8, 0x6, 0x0,
		33, 0x6b, 0x1d, 0x4f, 0x85, 0x72, 0x1, 0x0, 0x0, 0x55, 0xf8, 0x6, 0x0})}
	tm1 := time.Date(2018, 4, 3, 14, 25, 32, int(time.Millisecond*123+time.Microsecond*456+789), time.UTC)
	tm2 := time.Date(2019, 5, 4, 15, 26, 33, int(time.Millisecond*123+time.Microsecond*456+789), time.UTC)
	tm3 := time.Date(2020, 6, 5, 16, 27, 34, int(time.Millisecond*123+time.Microsecond*456+789), time.UTC)
	r36 := &response{message: bytes.NewBuffer([]byte{36, 0xdb, 0x6b, 0x18, 0x3, 0x0, 0x0, 0x0, 0x0})}
	tm4 := time.Date(1, 1, 1, 14, 25, 32, int(time.Millisecond*123), time.UTC)
	r37 := &response{message: bytes.NewBuffer([]byte{37, 3, 0, 0, 0,
		36, 0xdb, 0x6b, 0x18, 0x3, 0x0, 0x0, 0x0, 0x0,
		36, 0xa3, 0x48, 0x50, 0x3, 0x0, 0x0, 0x0, 0x0,
		36, 0x6b, 0x25, 0x88, 0x3, 0x0, 0x0, 0x0, 0x0})}
	tm5 := time.Date(1, 1, 1, 14, 25, 32, int(time.Millisecond*123), time.UTC)
	tm6 := time.Date(1, 1, 1, 15, 26, 33, int(time.Millisecond*123), time.UTC)
	tm7 := time.Date(1, 1, 1, 16, 27, 34, int(time.Millisecond*123), time.UTC)
	r101 := &response{message: bytes.NewBuffer([]byte{101})}

	tests := []struct {
		name    string
		r       *response
		want    interface{}
		wantErr bool
	}{
		{
			name: "byte",
			r:    r1,
			want: byte(123),
		},
		{
			name: "short",
			r:    r2,
			want: int16(12345),
		},
		{
			name: "int",
			r:    r3,
			want: int32(1234567890),
		},
		{
			name: "long",
			r:    r4,
			want: int64(1234567890123456789),
		},
		{
			name: "float",
			r:    r5,
			want: float32(123456.789),
		},
		{
			name: "double",
			r:    r6,
			want: float64(123456789.12345),
		},
		{
			name: "char",
			r:    r7,
			want: Char('A'),
		},
		{
			name: "bool",
			r:    r8,
			want: true,
		},
		{
			name: "string",
			r:    r9,
			want: "test string",
		},
		{
			name: "UUID",
			r:    r10,
			want: uid,
		},
		{
			name: "Date",
			r:    r11,
			want: dm,
		},
		{
			name: "byte array",
			r:    r12,
			want: []byte{1, 2, 3},
		},
		{
			name: "short array",
			r:    r13,
			want: []int16{1, 2, 3},
		},
		{
			name: "int array",
			r:    r14,
			want: []int32{1, 2, 3},
		},
		{
			name: "long array",
			r:    r15,
			want: []int64{1, 2, 3},
		},
		{
			name: "float array",
			r:    r16,
			want: []float32{1.1, 2.2, 3.3},
		},
		{
			name: "double array",
			r:    r17,
			want: []float64{1.1, 2.2, 3.3},
		},
		{
			name: "char array",
			r:    r18,
			want: []Char{'A', 'B', 'Я'},
		},
		{
			name: "bool array",
			r:    r19,
			want: []bool{true, false, true},
		},
		{
			name: "string array",
			r:    r20,
			want: []string{"one", "two", "три"},
		},
		{
			name: "UUID array",
			r:    r21,
			want: []uuid.UUID{uid1, uid2, uid3},
		},
		{
			name: "date array",
			r:    r22,
			want: []time.Time{dm1, dm2, dm3},
		},
		{
			name: "Timestamp",
			r:    r33,
			want: tm,
		},
		{
			name: "Timestamp array",
			r:    r34,
			want: []time.Time{tm1, tm2, tm3},
		},
		{
			name: "Time",
			r:    r36,
			want: tm4,
		},
		{
			name: "Time array",
			r:    r37,
			want: []time.Time{tm5, tm6, tm7},
		},
		{
			name: "NULL",
			r:    r101,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadObject()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response.ReadObject() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadFrom(t *testing.T) {
	rr := bytes.NewBuffer([]byte{1, 0, 0, 0, 1})

	type args struct {
		rr io.Reader
	}
	tests := []struct {
		name    string
		r       *response
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "1",
			r:    &response{},
			args: args{
				rr: rr,
			},
			want: 4 + 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadFrom(tt.args.rr)
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}