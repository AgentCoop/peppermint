package crypto

import (
	job "github.com/AgentCoop/go-work"
	"github.com/monnand/dhkx"
)

type DhKeyExchange interface {
	GetPublicKey() []byte
	ComputeKey([]byte) []byte
}

type dhKeyCtx struct {
	task job.Task
	g *dhkx.DHGroup
	privKey *dhkx.DHKey
}

func NewKeyExchange(task job.Task) *dhKeyCtx {
	 ctx := &dhKeyCtx{task: task}
	 g, err := dhkx.GetGroup(0) // Default group
	 task.Assert(err)
	 ctx.g = g

	 priv, err := g.GeneratePrivateKey(nil) // Use default random generator
	 task.Assert(err)
	 ctx.privKey = priv
	 return ctx
}

func (c *dhKeyCtx) GetPublicKey() []byte {
	return c.privKey.Bytes()
}

func (c *dhKeyCtx) ComputeKey(pubKey []byte) []byte {
	pkey := dhkx.NewPublicKey(pubKey)
	key, err := c.g.ComputeKey(pkey, c.privKey)
	c.task.Assert(err)
	return key.Bytes()
}
