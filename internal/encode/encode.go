package encode

type Encoder interface {
	Encode(data string) (string, error)
}

type Decoder interface {
	Decode(data string) (string, error)
}

type DecodeEncoder interface {
	Encoder
	Decoder
}
