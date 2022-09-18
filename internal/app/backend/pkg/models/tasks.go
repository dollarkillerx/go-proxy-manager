package models

import "github.com/dollarkillerx/go-proxy-manager/proto/common"

type Tasks struct {
	BasicModel
	Domain      string           `json:"domain"`
	EnableProxy bool             `json:"enable_proxy"`
	EnableSSL   bool             `json:"enable_ssl"`
	EnableWaf   bool             `json:"enable_waf"`
	ProxyType   common.ProxyType `json:"proxy_type"`
	Payload     string           `gorm:"type:text" json:"payload"`
	Certificate string           `gorm:"type:text" json:"certificate"`
}
