// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
)

type VmGetRebuildImageQuery struct {
	position.Query
	// 镜像类型; 1: 用户自定义镜像; 2: 系统镜像
	Type int32 `json:"type,omitempty"`
}
