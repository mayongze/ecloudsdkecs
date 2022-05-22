// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model



type VmGetServerChangeableResponseBody struct {

	// 云主机规格类型
	FlavorType string `json:"flavorType,omitempty"`

	// 规格名称
	SpecsName string `json:"specsName,omitempty"`

	// cpu
	Cpu int32 `json:"cpu,omitempty"`

	// ram
	Ram int32 `json:"ram,omitempty"`
}
