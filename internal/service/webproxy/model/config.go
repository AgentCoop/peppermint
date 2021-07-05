package model

import "github.com/AgentCoop/peppermint/internal/runtime/node/model"

type WebProxyConfig struct {
	model.Model
	Port int `gorm:"default:443"`
	Address string `gorm:"default:localhost"`
	ServerName string `gorm:"default:peppermint.io"`
	X509CertPem []byte
	X509KeyPem []byte
}
