package cmd

import (
	"crypto/sha256"
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/model/node"
	"github.com/AgentCoop/peppermint/internal/utils"
)

func BootstrapCmd(IdFromNic string, tags []string) error {
	var nodeId internal.NodeId
	switch {
	case len(IdFromNic) != 0: // Generate node ID from the given network interface
		macAddr, err := utils.Net_MacAdrr(IdFromNic)
		if err != nil { return err }
		hash := sha256.Sum256(macAddr)
		nodeId.FromByteArray(hash[:])
	default: // Random node ID
		nodeId.Rand()
	}
	return node.CreateNode(nodeId, tags)
}
