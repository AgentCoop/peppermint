package security

import (
	"github.com/AgentCoop/peppermint/internal"
	"time"
)

type accessToken struct {
	issuer       internal.NodeId
	issuedAt     time.Time
	lifetimeSecs uint
	roles        []string
}
