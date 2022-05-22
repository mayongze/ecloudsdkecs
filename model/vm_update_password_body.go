// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
)

type VmUpdatePasswordBody struct {
	position.Body
	// 云主机密码
	Password string `json:"password,omitempty"`

	// 云主机id
	ServerId string `json:"serverId,omitempty"`
}
