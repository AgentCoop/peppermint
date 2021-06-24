package node

import (
	runtime2 "github.com/AgentCoop/peppermint/pkg/runtime"
)

type manager struct {
	encKeyStore runtime2.InMemoryStore
}


func (m *manager) InquiryHub() runtime2.InMemoryStore {
	return m.encKeyStore
}