package security

import (
	"github.com/monnand/dhkx"
)

type DhKeyExchange interface {
	GetPublicKey() []byte
	ComputeKey([]byte) []byte
}

type dhKeyCtx struct {
	g       *dhkx.DHGroup
	privKey *dhkx.DHKey
}

func NewKeyExchange() (*dhKeyCtx, error) {
	ctx := &dhKeyCtx{}
	g, err := dhkx.GetGroup(0) // Default group
	if err != nil {
		return nil, err
	}
	ctx.g = g
	priv, err := g.GeneratePrivateKey(nil) // Use default random generator
	if err != nil {
		return nil, err
	}
	ctx.privKey = priv
	return ctx, nil
}

func (c *dhKeyCtx) GetPublicKey() []byte {
	return c.privKey.Bytes()
}

func (c *dhKeyCtx) ComputeKey(pubKey []byte) ([]byte, error) {
	pkey := dhkx.NewPublicKey(pubKey)
	key, err := c.g.ComputeKey(pkey, c.privKey)
	if err != nil {
		return nil, err
	}
	return key.Bytes(), nil
}
