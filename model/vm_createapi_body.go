// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
)

type VmCreateapiBodyVmTypeEnum string

// List of VmType
const (
	VmCreateapiBody_MEMIMPROVE         VmCreateapiBodyVmTypeEnum = "memImprove"
	VmCreateapiBody_COMMON             VmCreateapiBodyVmTypeEnum = "common"
	VmCreateapiBody_GPU                VmCreateapiBodyVmTypeEnum = "gpu"
	VmCreateapiBody_COMMONINTRODUCTORY VmCreateapiBodyVmTypeEnum = "commonIntroductory"
	VmCreateapiBody_COMMONNETIMPROVE   VmCreateapiBodyVmTypeEnum = "commonNetImprove"
	VmCreateapiBody_COMPUTE            VmCreateapiBodyVmTypeEnum = "compute"
	VmCreateapiBody_COMPUTENETIMPROVE  VmCreateapiBodyVmTypeEnum = "computeNetImprove"
	VmCreateapiBody_MEMNETIMPROVE      VmCreateapiBodyVmTypeEnum = "memNetImprove"
	VmCreateapiBody_LOCALSTORAGE       VmCreateapiBodyVmTypeEnum = "localStorage"
	VmCreateapiBody_XLARGEMEMORY       VmCreateapiBodyVmTypeEnum = "xlargeMemory"
	VmCreateapiBody_HIGHFREQUENCY      VmCreateapiBodyVmTypeEnum = "highFrequency"
	VmCreateapiBody_VGPU               VmCreateapiBodyVmTypeEnum = "vgpu"
	VmCreateapiBody_FPGA               VmCreateapiBodyVmTypeEnum = "fpga"
	VmCreateapiBody_HIGHIO             VmCreateapiBodyVmTypeEnum = "highIO"
	VmCreateapiBody_EXCLUSIVE          VmCreateapiBodyVmTypeEnum = "exclusive"
)

type VmCreateapiBodyBillingTypeEnum string

// List of BillingType
const (
	VmCreateapiBody_MONTH VmCreateapiBodyBillingTypeEnum = "MONTH"
	VmCreateapiBody_YEAR  VmCreateapiBodyBillingTypeEnum = "YEAR"
	VmCreateapiBody_HOUR  VmCreateapiBodyBillingTypeEnum = "HOUR"
)

type VmCreateapiBody struct {
	position.Body
	// 系统盘
	BootVolume VmCreateapiRequestBootVolume `json:"bootVolume,omitempty"`

	// 规格名称
	SpecsName string `json:"specsName,omitempty"`

	// 镜像名
	ImageName string `json:"imageName,omitempty"`

	// 自定义参数,长度在[1,1024]个字符间
	UserData string `json:"userData,omitempty"`

	// 订购数量,范围在[1,10]间
	Quantity int32 `json:"quantity,omitempty"`

	// 一体化订购中带宽产品
	Bandwidth VmCreateapiRequestBandwidth `json:"bandwidth,omitempty"`

	// 密钥对名称
	KeypairName string `json:"keypairName,omitempty"`

	// 一体化订购中公网IP产品
	Ip VmCreateapiRequestIp `json:"ip,omitempty"`

	// cpu
	Cpu int32 `json:"cpu,omitempty"`

	// 主机描述
	Description string `json:"description,omitempty"`

	// 网络
	Networks VmCreateapiRequestNetworks `json:"networks,omitempty"`

	// 订购时长/月(年)
	Duration int32 `json:"duration,omitempty"`

	// 云主机规格类型
	VmType VmCreateapiBodyVmTypeEnum `json:"vmType,omitempty"`

	// 磁盘大小
	Disk int32 `json:"disk,omitempty"`

	// 密码
	Password string `json:"password,omitempty"`

	// 付费方式
	BillingType VmCreateapiBodyBillingTypeEnum `json:"billingType,omitempty"`

	// 一体化订购中数据盘产品
	DataVolume *[]VmCreateapiRequestDataVolume `json:"dataVolume,omitempty"`

	// 安全组id
	SecurityGroupIds *[]string `json:"securityGroupIds,omitempty"`

	// 云主机名,名称为5~22位英文、"-"、数字的组合,"-"不可在名称开头或末尾
	Name string `json:"name,omitempty"`

	// 可用区
	Region string `json:"region,omitempty"`

	// 下单成功跳转地址，前端填写
	ReturnUrl string `json:"returnUrl,omitempty"`

	// 内存
	Ram int32 `json:"ram,omitempty"`
}
