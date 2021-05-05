package crypto

import (
	job "github.com/AgentCoop/go-work"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type symCipher struct {
	task job.Task
	block cipher.Block
	gcm cipher.AEAD
	nonce []byte
}

func NewSymCipher(key []byte, nonce []byte, task job.Task) symCipher {
	block, err := aes.NewCipher(key)
	task.Assert(err)

	// gcm or Galois/Counter Mode, is a mode of operation for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(block)
	task.Assert(err)

	if nonce == nil {
		nonce = make([]byte, gcm.NonceSize())
		_, err = io.ReadFull(rand.Reader, nonce)
		task.Assert(err)
	}

	return symCipher{task, block, gcm, nonce}
}

func (c symCipher) GetNonce() []byte {
	return c.nonce
}

func (c symCipher) Encrypt(data []byte) []byte {
	return c.gcm.Seal(data[:0], c.nonce, data, nil)
}

func (c symCipher) Decrypt(data []byte) []byte {
	plain, err := c.gcm.Open(data[:0], c.nonce, data, nil)
	c.task.Assert(err)
	return plain
}
