package store

import "github.com/AgentCoop/peppermint/internal/runtime"

type store struct {
	fallback runtime.Store_FallbackHandler
}

func (s *store) RegisterFallback(fn runtime.Store_FallbackHandler) {
	s.fallback = fn
}
