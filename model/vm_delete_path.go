// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
)

type VmDeletePath struct {
	position.Path
	// 云主机id
	ServerId string `json:"serverId,omitempty"`
}
