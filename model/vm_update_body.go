// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
)

type VmUpdateBody struct {
	position.Body
	// 密钥对名称，只能包含数字、字母，长度5~25位
	KeyName string `json:"keyName,omitempty"`

	// 主机id
	ServerId string `json:"serverId,omitempty"`
}
