// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
)

type VmupdateNameBody struct {
	position.Body
	// 云主机名称
	Name string `json:"name,omitempty"`

	// 云主机id
	ServerId string `json:"serverId,omitempty"`
}
