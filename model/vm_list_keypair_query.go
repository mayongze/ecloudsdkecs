// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
)

type VmListKeypairQuery struct {
	position.Query
	// 列表分页每页大小
	Size int32 `json:"size,omitempty"`

	// keypair名称
	KeyName string `json:"keyName,omitempty"`

	// 列表分页页数
	Page int32 `json:"page,omitempty"`
}
