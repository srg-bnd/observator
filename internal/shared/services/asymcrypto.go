package services

type PublicKey struct {
	key string
}

type PrivateKey struct {
	key string
}

func NewPublicKey(key string) *PublicKey {
	return &PublicKey{
		key: key,
	}
}

func NewPrivateKey(key string) *PrivateKey {
	return &PrivateKey{
		key: key,
	}
}

// Encrypts data with a public key
func (s *PublicKey) Encrypt(data []byte) ([]byte, error) {
	// TODO: realize
	return data, nil
}

// Decrypts data with a private key
func (s *PrivateKey) Decrypt(data []byte) ([]byte, error) {
	// TODO: realize
	return data, nil
}
