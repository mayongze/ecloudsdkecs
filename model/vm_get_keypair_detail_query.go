// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
)

type VmGetKeypairDetailQuery struct {
	position.Query
	// 可用区
	Region string `json:"region,omitempty"`
}
