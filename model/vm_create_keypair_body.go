// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
)

type VmCreateKeypairBody struct {
	position.Body
	// 密钥名称
	Name string `json:"name,omitempty"`

	// 可用区
	Region string `json:"region,omitempty"`
}
