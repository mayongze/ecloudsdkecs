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

	// ????????????????????????
	BootVolumeType VmGetServerDetailResponseBodyBootVolumeTypeEnum `json:"bootVolumeType,omitempty"`

	// ??????id
	ImageId string `json:"imageId,omitempty"`

	// ????????????
	ImageName string `json:"imageName,omitempty"`

	// ????????????
	Memory int32 `json:"memory,omitempty"`

	// ?????????????????????
	SpecsName string `json:"specsName,omitempty"`

	// ???????????????
	KeypairName string `json:"keypairName,omitempty"`

	// ?????????id
	SystemDisk string `json:"systemDisk,omitempty"`

	// ??????id
	FlavorId string `json:"flavorId,omitempty"`

	// ????????????????????????
	Volumes *[]VmGetServerDetailResponseVolumes `json:"volumes,omitempty"`

	// cpu??????
	Cpu int32 `json:"cpu,omitempty"`

	// ???????????????
	Description string `json:"description,omitempty"`

	// ?????????????????????
	Ports *[]VmGetServerDetailResponsePorts `json:"ports,omitempty"`

	// ??????????????????????????????
	ServerBackupPolicyResp VmGetServerDetailResponseServerBackupPolicyResp `json:"serverBackupPolicyResp,omitempty"`

	// az
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// ???????????????
	Disk int32 `json:"disk,omitempty"`

	// ?????????????????????
	MaxBandwidth string `json:"maxBandwidth,omitempty"`

	// ???????????????
	Name string `json:"name,omitempty"`

	// ????????????
	CreatedTime string `json:"createdTime,omitempty"`

	// ?????????????????????????????????
	ServerVmType VmGetServerDetailResponseBodyServerVmTypeEnum `json:"serverVmType,omitempty"`

	// ?????????id
	Id string `json:"id,omitempty"`

	// ?????????
	Region string `json:"region,omitempty"`

	// ???????????????
	Status int32 `json:"status,omitempty"`
}
