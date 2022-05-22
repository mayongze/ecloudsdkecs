// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmqueryPriceRequestBootVolumeVolumeTypeEnum string

// List of VolumeType
const (
    VmqueryPriceRequestBootVolume_HIGHPERFORMANCE VmqueryPriceRequestBootVolumeVolumeTypeEnum = "highPerformance"
    VmqueryPriceRequestBootVolume_PERFORMANCEOPTIMIZATION VmqueryPriceRequestBootVolumeVolumeTypeEnum = "performanceOptimization"
)

type VmqueryPriceRequestBootVolume struct {

	// 系统盘类型,highPerformance:高性能型;performanceOptimization:性能优化型
	VolumeType VmqueryPriceRequestBootVolumeVolumeTypeEnum `json:"volumeType,omitempty"`

	// 系统盘大小(GB),大小需在[20,500]间
	Size int32 `json:"size,omitempty"`
}
