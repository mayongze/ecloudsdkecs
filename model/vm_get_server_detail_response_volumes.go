// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmGetServerDetailResponseVolumesTypeEnum string

// List of Type
const (
    VmGetServerDetailResponseVolumes_CAPEBS VmGetServerDetailResponseVolumesTypeEnum = "capebs"
    VmGetServerDetailResponseVolumes_EBS VmGetServerDetailResponseVolumesTypeEnum = "ebs"
    VmGetServerDetailResponseVolumes_SSD VmGetServerDetailResponseVolumesTypeEnum = "ssd"
    VmGetServerDetailResponseVolumes_SSDEBS VmGetServerDetailResponseVolumesTypeEnum = "ssdebs"
    VmGetServerDetailResponseVolumes_EDGESSDEBS VmGetServerDetailResponseVolumesTypeEnum = "edgessdebs"
    VmGetServerDetailResponseVolumes_SSDIO VmGetServerDetailResponseVolumesTypeEnum = "ssdio"
    VmGetServerDetailResponseVolumes_HDD VmGetServerDetailResponseVolumesTypeEnum = "hdd"
)

type VmGetServerDetailResponseVolumes struct {

	// 数据盘大小
	Size int32 `json:"size,omitempty"`

	// 数据盘名称
	Name string `json:"name,omitempty"`

	// 数据盘id
	Id string `json:"id,omitempty"`

	// 数据盘类型
	Type VmGetServerDetailResponseVolumesTypeEnum `json:"type,omitempty"`

	// 数据盘状态
	Status string `json:"status,omitempty"`
}
