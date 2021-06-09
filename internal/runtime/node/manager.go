package node

import "github.com/AgentCoop/peppermint/internal/runtime"

type manager struct {
	encKeyStore runtime.InMemoryStore
}

func (m *manager) EncKeyStore() runtime.InMemoryStore {
	return m.encKeyStore
}
func (m *manager) InquiryHub() runtime.InMemoryStore {
	return m.encKeyStore
}