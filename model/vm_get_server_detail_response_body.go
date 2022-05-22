// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

type VmGetServerDetailResponseBodyBootVolumeTypeEnum string

// List of BootVolumeType
const (
	VmGetServerDetailResponseBody_LOCAL                   VmGetServerDetailResponseBodyBootVolumeTypeEnum = "local"
	VmGetServerDetailResponseBody_CAPACITY                VmGetServerDetailResponseBodyBootVolumeTypeEnum = "capacity"
	VmGetServerDetailResponseBody_HIGHPERFORMANCE         VmGetServerDetailResponseBodyBootVolumeTypeEnum = "highPerformance"
	VmGetServerDetailResponseBody_PERFORMANCEOPTIMIZATION VmGetServerDetailResponseBodyBootVolumeTypeEnum = "performanceOptimization"
)

type VmGetServerDetailResponseBodyServerVmTypeEnum string

// List of ServerVmType
const (
	VmGetServerDetailResponseBody_HIGHIO               VmGetServerDetailResponseBodyServerVmTypeEnum = "highIO"
	VmGetServerDetailResponseBody_EXCLUSIVE            VmGetServerDetailResponseBodyServerVmTypeEnum = "exclusive"
	VmGetServerDetailResponseBody_MEMIMPROVE           VmGetServerDetailResponseBodyServerVmTypeEnum = "memImprove"
	VmGetServerDetailResponseBody_COMMON               VmGetServerDetailResponseBodyServerVmTypeEnum = "common"
	VmGetServerDetailResponseBody_GPU                  VmGetServerDetailResponseBodyServerVmTypeEnum = "gpu"
	VmGetServerDetailResponseBody_HIGHPERFORMANCE_VM   VmGetServerDetailResponseBodyServerVmTypeEnum = "highPerformance"
	VmGetServerDetailResponseBody_COMMONINTRODUCTORY   VmGetServerDetailResponseBodyServerVmTypeEnum = "commonIntroductory"
	VmGetServerDetailResponseBody_COMMONNETIMPROVE     VmGetServerDetailResponseBodyServerVmTypeEnum = "commonNetImprove"
	VmGetServerDetailResponseBody_COMPUTE              VmGetServerDetailResponseBodyServerVmTypeEnum = "compute"
	VmGetServerDetailResponseBody_COMPUTENETIMPROVE    VmGetServerDetailResponseBodyServerVmTypeEnum = "computeNetImprove"
	VmGetServerDetailResponseBody_MEMNETIMPROVE        VmGetServerDetailResponseBodyServerVmTypeEnum = "memNetImprove"
	VmGetServerDetailResponseBody_LOCALSTORAGE         VmGetServerDetailResponseBodyServerVmTypeEnum = "localStorage"
	VmGetServerDetailResponseBody_XLARGEMEMORY         VmGetServerDetailResponseBodyServerVmTypeEnum = "xlargeMemory"
	VmGetServerDetailResponseBody_HIGHFREQUENCY        VmGetServerDetailResponseBodyServerVmTypeEnum = "highFrequency"
	VmGetServerDetailResponseBody_VGPU                 VmGetServerDetailResponseBodyServerVmTypeEnum = "vgpu"
	VmGetServerDetailResponseBody_FPGA                 VmGetServerDetailResponseBodyServerVmTypeEnum = "fpga"
	VmGetServerDetailResponseBody_NORMALCOMPUTEIMPROVE VmGetServerDetailResponseBodyServerVmTypeEnum = "normalComputeImprove"
	VmGetServerDetailResponseBody_NORMALNETENHANCE     VmGetServerDetailResponseBodyServerVmTypeEnum = "normalNetEnhance"
	VmGetServerDetailResponseBody_STOREENHANCE         VmGetServerDetailResponseBodyServerVmTypeEnum = "storeEnhance"
	VmGetServerDetailResponseBody_EDGECOMMONNETIMPROVE VmGetServerDetailResponseBodyServerVmTypeEnum = "edgeCommonNetImprove"
	VmGetServerDetailResponseBody_GENERAL              VmGetServerDetailResponseBodyServerVmTypeEnum = "general"
	VmGetServerDetailResponseBody_COMPUTEG2            VmGetServerDetailResponseBodyServerVmTypeEnum = "computeG2"
	VmGetServerDetailResponseBody_NPU                  VmGetServerDetailResponseBodyServerVmTypeEnum = "npu"
	VmGetServerDetailResponseBody_MEMORYDOUBLECARD     VmGetServerDetailResponseBodyServerVmTypeEnum = "memoryDoubleCard"
	VmGetServerDetailResponseBody_MEMORYHBACARD        VmGetServerDetailResponseBodyServerVmTypeEnum = "memoryHBACard"
	VmGetServerDetailResponseBody_LOCALSTORAGELOW      VmGetServerDetailResponseBodyServerVmTypeEnum = "localStorageLow"
)

type VmGetServerDetailResponseBody struct {

	// 云主机系统盘类型
	BootVolumeType VmGetServerDetailResponseBodyBootVolumeTypeEnum `json:"bootVolumeType,omitempty"`

	// 镜像id
	ImageId string `json:"imageId,omitempty"`

	// 镜像名称
	ImageName string `json:"imageName,omitempty"`

	// 内存大小
	Memory int32 `json:"memory,omitempty"`

	// 云主机规格名称
	SpecsName string `json:"specsName,omitempty"`

	// 密钥对名称
	KeypairName string `json:"keypairName,omitempty"`

	// 系统盘id
	SystemDisk string `json:"systemDisk,omitempty"`

	// 规格id
	FlavorId string `json:"flavorId,omitempty"`

	// 云主机数据盘列表
	Volumes *[]VmGetServerDetailResponseVolumes `json:"volumes,omitempty"`

	// cpu个数
	Cpu int32 `json:"cpu,omitempty"`

	// 云主机描述
	Description string `json:"description,omitempty"`

	// 云主机网卡列表
	Ports *[]VmGetServerDetailResponsePorts `json:"ports,omitempty"`

	// 云主机备份策略返回体
	ServerBackupPolicyResp VmGetServerDetailResponseServerBackupPolicyResp `json:"serverBackupPolicyResp,omitempty"`

	// az
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// 系统盘大小
	Disk int32 `json:"disk,omitempty"`

	// 云主机内网带宽
	MaxBandwidth string `json:"maxBandwidth,omitempty"`

	// 云主机名称
	Name string `json:"name,omitempty"`

	// 创建时间
	CreatedTime string `json:"createdTime,omitempty"`

	// 云主机类型下的主机类型
	ServerVmType VmGetServerDetailResponseBodyServerVmTypeEnum `json:"serverVmType,omitempty"`

	// 云主机id
	Id string `json:"id,omitempty"`

	// 可用区
	Region string `json:"region,omitempty"`

	// 云主机状态
	Status int32 `json:"status,omitempty"`
}
