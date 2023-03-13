package encode

import (
	"reflect"
	"testing"
)

func TestNewGoBase64(t *testing.T) {
	tests := []struct {
		name string
		want *StdBase64Manager
	}{
		{
			name: "it should create a new base64 manager",
			want: NewStdBase64Manager(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStdBase64Manager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStdBase64Manager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoBase64Manager_Decode(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "it should decode valid base64 string",
			args: args{
				data: "c29tZSBkYXRhIGVuY29kZWQgaW4gYmFzZTY0",
			},
			want:    "some data encoded in base64",
			wantErr: false,
		},
		{
			name: "it should fail decoding invalid base64 string",
			args: args{
				data: "LK===_)S\"?''''±`z",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewStdBase64Manager()
			got, err := g.Decode(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoBase64Manager_Encode(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "it should encode string",
			args: args{
				data: "some data encoded in base64",
			},
			want:    "c29tZSBkYXRhIGVuY29kZWQgaW4gYmFzZTY0",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewStdBase64Manager()
			got, err := g.Encode(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Encode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGoBase64Manager_Decode(b *testing.B) {
	base64Manager := NewStdBase64Manager()

	for i := 0; i < b.N; i++ {
		_, _ = base64Manager.Decode("c29tZSBkYXRhIGVuY29kZWQgaW4gYmFzZTY0")
	}
}
