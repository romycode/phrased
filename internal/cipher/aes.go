package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/romycode/phrased/internal/encode"
)

type AesCipher struct {
	ed encode.DecodeEncoder
	c  cipher.AEAD
}

func NewAesCipher(key string, decodeEncoder encode.DecodeEncoder) (*AesCipher, error) {
	b, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, fmt.Errorf("AesCipher: unable to create encrypter with key: %q\n  Prev:\n  - %w\n", key, err)
	}

	e, _ := cipher.NewGCM(b)
	return &AesCipher{decodeEncoder, e}, nil
}

func (e *AesCipher) Encrypt(data string) (string, error) {
	nonce := e.generateNonce()

	encodedData, err := e.ed.Encode(data)
	if err != nil {
		return "", fmt.Errorf("- AesCipher: unable to encode as base64 data: %q\n  Prev:\n  - %w\n", data, err)
	}

	r := e.c.Seal(nonce, nonce, []byte(encodedData), nil)
	return hex.EncodeToString(r), nil
}

func (e *AesCipher) Decrypt(data string) (string, error) {
	input, err := hex.DecodeString(data)
	if err != nil {
		return "", fmt.Errorf("- AesCipher: unable to decode encrypted data: %s\n  Prev:\n  - %w\n", data, err)
	}

	encodedData, err := e.c.Open(nil, e.getNonce(input), e.getCipheredText(input), nil)
	if err != nil {
		return "", fmt.Errorf("- AesCipher: unable to decrypt data: %s\n  Prev:\n  - %w\n", data, err)
	}

	result, err := e.ed.Decode(string(encodedData))
	if err != nil {
		return "", fmt.Errorf("- AesCipher: unable decode base64 data: %q\n  Prev:\n  - %w\n", encodedData, err)
	}

	return result, nil
}

func (e *AesCipher) generateNonce() []byte {
	nonce := make([]byte, e.c.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)
	return nonce
}

func (e *AesCipher) getNonce(data []byte) []byte {
	return data[:e.c.NonceSize()]
}

func (e *AesCipher) getCipheredText(data []byte) []byte {
	return data[e.c.NonceSize():]
}
