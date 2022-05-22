// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmCreateapiRequestBootVolumeVolumeTypeEnum string

// List of VolumeType
const (
    VmCreateapiRequestBootVolume_HIGHPERFORMANCE VmCreateapiRequestBootVolumeVolumeTypeEnum = "highPerformance"
    VmCreateapiRequestBootVolume_PERFORMANCEOPTIMIZATION VmCreateapiRequestBootVolumeVolumeTypeEnum = "performanceOptimization"
)

type VmCreateapiRequestBootVolume struct {

	// 系统盘类型,highPerformance:高性能型;performanceOptimization:性能优化型
	VolumeType VmCreateapiRequestBootVolumeVolumeTypeEnum `json:"volumeType,omitempty"`

	// 系统盘大小(GB),大小需在[20,500]间
	Size int32 `json:"size,omitempty"`
}
