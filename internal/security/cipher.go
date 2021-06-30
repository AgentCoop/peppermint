package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
)

type symCipher struct {
	block cipher.Block
	gcm   cipher.AEAD
	nonce []byte
}

type SymCipher interface {
	GetNonce() []byte
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
}

func NewSymCipher(key []byte, nonce []byte) (*symCipher, error) {
	// Normalize key length
	if len(key) != 32 {
		normalizedKey := sha256.Sum256(key)
		key = normalizedKey[:]
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// gcm or Galois/Counter Mode, is a mode of operation for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if nonce == nil {
		nonce = make([]byte, gcm.NonceSize())
		_, err = io.ReadFull(rand.Reader, nonce)
		if err != nil {
			return nil, err
		}
	}

	return &symCipher{block, gcm, nonce}, nil
}

func (c symCipher) GetNonce() []byte {
	return c.nonce
}

func (c symCipher) Encrypt(data []byte) []byte {
	return c.gcm.Seal(data[:0], c.nonce, data, nil)
}

func (c symCipher) Decrypt(data []byte) []byte {
	plain, err := c.gcm.Open(data[:0], c.nonce, data, nil)
	if err != nil {
		panic(err)
	}
	return plain
}
