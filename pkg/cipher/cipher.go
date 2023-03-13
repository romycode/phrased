package cipher

type Cipher interface {
	Encrypt(data string) (string, error)
	Decrypt(data string) (string, error)
}
