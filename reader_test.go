package bitreader

import (
	"reflect"
	"testing"
)

func TestNewReader(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name string
		args args
		want *Reader
	}{
		{
			name: "success",
			args: args{
				input: []byte{3, 255},
			},
			want: &Reader{
				bits: []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReader(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bitsToInt(t *testing.T) {
	type args struct {
		bits []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				bits: []int{1, 1, 0, 1, 0},
			},
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitsToInt(tt.args.bits); got != tt.want {
				t.Errorf("bitsToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReader_SliceToInt(t *testing.T) {
	type fields struct {
		bits []int
	}
	type args struct {
		offset int
		length int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "success from 0",
			fields: fields{
				bits: []int{1, 0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 0, 0, 0},
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
				bits: []int{1, 0, 1, 1, 0, 1, 0, 1, 1, 0},
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
				bits: []int{1, 0, 1, 1, 0, 1, 0},
			},
			args: args{
				offset: 0,
				length: 8,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				bits: tt.fields.bits,
			}
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
