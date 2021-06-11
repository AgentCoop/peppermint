package service

import (
	_ "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	"testing"
)

const (
	DefaultHubPort = 12001
)

func TestNewServicePolicy(t *testing.T) {
	svcName := "hub.backoffice.peppermint.Hub"
	methods := []string{"JoinHello", "Join"}
	policy := NewServicePolicy(svcName, methods)
	if policy.DefaultPort() != DefaultHubPort {
		t.Fatalf("expected port option %d, got %d", DefaultHubPort, policy.DefaultPort())
	}
	if ! policy.EnforceEncryption() {
		t.Fatalf("expected enforce encryption option to be TRUE, got %v", policy.EnforceEncryption())
	}
}
