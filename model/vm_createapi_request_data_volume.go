// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmCreateapiRequestDataVolumeResourceTypeEnum string

// List of ResourceType
const (
    VmCreateapiRequestDataVolume_CAPEBS VmCreateapiRequestDataVolumeResourceTypeEnum = "capebs"
    VmCreateapiRequestDataVolume_SSD VmCreateapiRequestDataVolumeResourceTypeEnum = "ssd"
    VmCreateapiRequestDataVolume_SSDEBS VmCreateapiRequestDataVolumeResourceTypeEnum = "ssdebs"
)

type VmCreateapiRequestDataVolume struct {

	// 数据盘大小(G),大小需在[10,32768]间
	Size int32 `json:"size,omitempty"`

	// 是否为共享数据盘,默认为非共享盘
	IsShare bool `json:"isShare,omitempty"`

	// 数据盘类型,capebs:容量型数据盘;ssd:高性能型数据盘;ssdebs:性能优化型数据盘
	ResourceType VmCreateapiRequestDataVolumeResourceTypeEnum `json:"resourceType,omitempty"`
}
