package store

import (
	runtime2 "github.com/AgentCoop/peppermint/pkg/runtime"
)

type store struct {
	fallback runtime2.Store_FallbackHandler
}

func (s *store) RegisterFallback(fn runtime2.Store_FallbackHandler) {
	s.fallback = fn
}
