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
		t.Fatalf("expected false, got %v", policy.EnforceEncryption())
	}
	mJoinHello, _ := policy.FindMethodByName(methods[0])
	if isSec := mJoinHello.CallPolicy().IsSecure(); isSec {
		t.Fatalf("expected false, got %v", isSec)
	}
	mJoin, _ := policy.FindMethodByName(methods[1])
	if isSec := mJoin.CallPolicy().IsSecure(); !isSec {
		t.Fatalf("expected true, got %v", isSec)
	}
}
