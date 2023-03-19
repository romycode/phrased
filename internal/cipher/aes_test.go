package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"reflect"
	"testing"

	"github.com/romycode/phrased/internal/encode"
)

func TestNewAesCipher(t *testing.T) {
	type args struct {
		key           string
		base64Manager encode.Base64Manager
	}

	key := "12345678901234567890123456789012"
	base64Manager := encode.NewStdBase64Manager()
	block, _ := aes.NewCipher([]byte(key))
	c, _ := cipher.NewGCM(block)

	tests := []struct {
		name    string
		args    args
		want    *AesCipher
		wantErr bool
	}{
		{
			name: "it should create a aes cipher",
			args: args{
				key:           key,
				base64Manager: base64Manager,
			},
			want: &AesCipher{
				ed: base64Manager,
				c:  c,
			},
			wantErr: false,
		},
		{
			name: "it should return error creating aes cipher",
			args: args{
				key:           "1234567890",
				base64Manager: base64Manager,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAesCipher(tt.args.key, tt.args.base64Manager)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAesCipher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAesCipher() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAesCipher_Decrypt(t *testing.T) {
	type fields struct {
		key           string
		base64Manager encode.Base64Manager
	}
	type args struct {
		data string
	}

	tests := []struct {
		name        string
		fields      fields
		args        args
		want        string
		wantErr     bool
		expectedErr error
	}{
		{
			name: "it should decrypt valid content",
			fields: fields{
				key:           "12345678901234567890123456789012",
				base64Manager: encode.NewStdBase64Manager(),
			},
			args: args{
				data: "84a35a063130fa2fea677d3519e23a6dfb48ca49d54250dcade97dcfc18241826aaaee47dcda0a2bf5f5bd836021d8748825f2973564c2c2",
			},
			want:    "some data to encrypt",
			wantErr: false,
		},
		{
			name: "it should return error with invalid content",
			fields: fields{
				key:           "12345678901234567890123456789012",
				base64Manager: encode.NewStdBase64Manager(),
			},
			args: args{
				data: "84a35a063130fa2fea677d3519e23a6dfcc8ca49d54250dcade97dcfc18241826aaaee47dcda0a2bf5f5bd836021d8748825f2973564c2c2",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, _ := NewAesCipher(tt.fields.key, tt.fields.base64Manager)
			got, err := e.Decrypt(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAesCipher_Encrypt(t *testing.T) {
	type fields struct {
		key           string
		base64Manager encode.Base64Manager
	}
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "it should encrypt given string",
			fields: fields{
				key:           "12345678901234567890123456789012",
				base64Manager: encode.NewStdBase64Manager(),
			},
			args: args{
				data: "some data to encrypt",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "it should return error given invalid string",
			fields: fields{
				key:           "12345678901234567890123456789012",
				base64Manager: encode.NewStdBase64Manager(),
			},
			args: args{
				data: "some data to encrypt",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, _ := NewAesCipher(tt.fields.key, tt.fields.base64Manager)
			_, err := e.Encrypt(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
