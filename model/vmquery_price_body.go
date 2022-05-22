// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
)

type VmqueryPriceBody struct {
	position.Body
	// 订购时长，不填默认为1，包月最大12，包年最大5
	Duration int32 `json:"duration,omitempty"`

	// 云主机系统盘查询参数
	BootVolume VmqueryPriceRequestBootVolume `json:"bootVolume,omitempty"`

	// 规格名称，格式如"m1ebm.6xlarge.32"，和规格参数，规格id三选一，如都填则依据flavorId，specsName顺序取第一个为准
	SpecsName string `json:"specsName,omitempty"`

	// 订购数量，不填默认为1，不超过10
	Quantity int32 `json:"quantity,omitempty"`

	// 缴费周期类型（month/year/hour）
	FeeUnit string `json:"feeUnit,omitempty"`

	// 规格id，格式如"ecloud-BMMem-5118-00240768-2x480G-2x960G"，此项与规格参数，规格名称三选一，如都填则依据flavorId，specsName，specsParam顺序取第一个为准
	FlavorId string `json:"flavorId,omitempty"`

	// 产品类型：vm-云主机/ironic-裸金属服务器
	ProductType string `json:"productType,omitempty"`
}
