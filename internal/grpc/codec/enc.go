package codec

import (
	"bytes"
	"github.com/AgentCoop/peppermint/internal/security"
)

func (p *packet) encrypt(data []byte) []byte {
	if len(p.encKey) == 0 { return data }
	var out bytes.Buffer
	cipher, _ := security.NewSymCipher(p.encKey, nil)
	encrypted := cipher.Encrypt(data)
	nonce := cipher.GetNonce()
	noncel := []byte{byte(len(nonce))}
	out.Write(noncel)
	out.Write(nonce)
	out.Write(encrypted)
	return out.Bytes()
}

func (p *packet) decrypt(data []byte) ([]byte, error) {
	if len(p.encKey) == 0 { return data, nil }
	noncel := data[0:1][0]
	nonce := data[1 : noncel+1]
	encrypted := data[1+noncel:]
	cipher, err := security.NewSymCipher(p.encKey, nonce)
	if err != nil { return nil, err }
	decrypted := cipher.Decrypt(encrypted)
	return decrypted, nil
}

