package encode

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
)

type Base64Manager interface {
	Encode(data string) (string, error)
	Decode(data string) (string, error)
}

type StdBase64Manager struct {
	b *bytes.Buffer
	e io.WriteCloser
	d io.Reader
}

func NewStdBase64Manager() *StdBase64Manager {
	b := new(bytes.Buffer)
	e := base64.NewEncoder(base64.StdEncoding, b)
	d := base64.NewDecoder(base64.StdEncoding, b)

	return &StdBase64Manager{b, e, d}
}

func (g *StdBase64Manager) Encode(data string) (string, error) {
	defer g.b.Reset()

	_, err := g.e.Write([]byte(data))
	if err != nil {
		return "", fmt.Errorf("- StdBase64Manager: unable to encode: %q\n  Prev:\n  - %w\n", data, err)
	}
	err = g.e.Close()
	if err != nil {
		return "", fmt.Errorf("- StdBase64Manager: unable to finish encode: %q\n  Prev:\n  - %w\n", data, err)
	}

	return g.b.String(), nil
}

func (g *StdBase64Manager) Decode(data string) (string, error) {
	defer g.b.Reset()

	g.b.WriteString(data)
	decodedContent := make([]byte, base64.StdEncoding.DecodedLen(g.b.Len()))
	size, err := g.d.Read(decodedContent)
	if err != nil {
		return "", fmt.Errorf("- StdBase64Manager: unable to finish decode: %q\n  Prev:\n  - %w\n", data, err)
	}

	return string(decodedContent[:size]), nil
}
