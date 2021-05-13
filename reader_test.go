package bitreader_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/looocusio/bitreader"
)

func ExampleReader_SliceToInt() {
	r := bitreader.NewReader([]byte{3, 255})
	result, err := r.SliceToInt(0, 8)
	if err != nil {
		fmt.Printf("failed slice to int: %s", err)
	}
	fmt.Println(result)
	// Output:
	// 3
}

func TestNewReader(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name     string
		args     args
		wantBits []uint8
	}{
		{
			name: "success",
			args: args{
				input: []byte{3, 255},
			},
			wantBits: []uint8{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitreader.NewReader(tt.args.input); !reflect.DeepEqual(got.ExportBits(), tt.wantBits) {
				t.Errorf("NewReader() = %v, wantBits %v", got, tt.wantBits)
			}
		})
	}
}

func TestReader_SliceToInt(t *testing.T) {
	type fields struct {
		bits []uint8
	}
	type args struct {
		offset uint64
		length uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "success from 0",
			fields: fields{
				bits: []uint8{1, 0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 0, 0, 0},
			},
			args: args{
				offset: 0,
				length: 8,
			},
			want:    181,
			wantErr: false,
		},
		{
			name: "success from offset",
			fields: fields{
				bits: []uint8{1, 0, 1, 1, 0, 1, 0, 1, 1, 0},
			},
			args: args{
				offset: 2,
				length: 5,
			},
			want:    26,
			wantErr: false,
		},
		{
			name: "failed invalid offset and length",
			fields: fields{
				bits: []uint8{1, 0, 1, 1, 0, 1, 0},
			},
			args: args{
				offset: 0,
				length: 8,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "failed empty bits",
			fields: fields{
				bits: []uint8{},
			},
			args: args{
				offset: 0,
				length: 8,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "success 64bit",
			fields: fields{
				bits: []uint8{
					1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1,
				},
			},
			args: args{
				offset: 0,
				length: 64,
			},
			want:    18446744073709551615,
			wantErr: false,
		},
		{
			name: "failed 65bit",
			fields: fields{
				bits: []uint8{
					1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1,
					1,
				},
			},
			args: args{
				offset: 0,
				length: 65,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &bitreader.Reader{}
			r.ExportBitsSet(tt.fields.bits)
			got, err := r.SliceToInt(tt.args.offset, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.SliceToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.SliceToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
