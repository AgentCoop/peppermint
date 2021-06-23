package codec

import (
	"bytes"
	"github.com/AgentCoop/peppermint/internal/crypto"
)

func encrypt(data []byte, encKey []byte) []byte {
	var out bytes.Buffer
	cipher, _ := crypto.NewSymCipher(encKey, nil)
	encrypted := cipher.Encrypt(data)
	nonce := cipher.GetNonce()
	noncel := []byte{byte(len(nonce))}
	out.Write(noncel)
	out.Write(nonce)
	out.Write(encrypted)
	return out.Bytes()
}

func decrypt(data []byte, encKey []byte) ([]byte, error) {
	if len(encKey) == 0 { return data, nil }
	noncel := data[0:1][0]
	nonce := data[1 : noncel+1]
	encrypted := data[1+noncel:]
	cipher, err := crypto.NewSymCipher(encKey, nonce)
	if err != nil { return nil, err }
	decrypted := cipher.Decrypt(encrypted)
	return decrypted, nil
}

