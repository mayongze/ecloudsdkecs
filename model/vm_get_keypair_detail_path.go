// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
)

type VmGetKeypairDetailPath struct {
	position.Path
	// 密钥对名称
	KeypairName string `json:"keypairName,omitempty"`
}
