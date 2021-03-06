// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

type VmlistServerPortsDetailRespResponseContentImageOsTypeEnum string

// List of ImageOsType
const (
	VmlistServerPortsDetailRespResponseContent_LINUX   VmlistServerPortsDetailRespResponseContentImageOsTypeEnum = "linux"
	VmlistServerPortsDetailRespResponseContent_WINDOWS VmlistServerPortsDetailRespResponseContentImageOsTypeEnum = "windows"
	VmlistServerPortsDetailRespResponseContent_OTHER   VmlistServerPortsDetailRespResponseContentImageOsTypeEnum = "other"
)

type VmlistServerPortsDetailRespResponseContentProductTypeEnum string

// List of ProductType
const (
	VmlistServerPortsDetailRespResponseContent_NORMAL      VmlistServerPortsDetailRespResponseContentProductTypeEnum = "NORMAL"
	VmlistServerPortsDetailRespResponseContent_DE_CLOUD    VmlistServerPortsDetailRespResponseContentProductTypeEnum = "DE_CLOUD"
	VmlistServerPortsDetailRespResponseContent_AUTOSCALING VmlistServerPortsDetailRespResponseContentProductTypeEnum = "AUTOSCALING"
	VmlistServerPortsDetailRespResponseContent_VO          VmlistServerPortsDetailRespResponseContentProductTypeEnum = "VO"
	VmlistServerPortsDetailRespResponseContent_CDN         VmlistServerPortsDetailRespResponseContentProductTypeEnum = "CDN"
	VmlistServerPortsDetailRespResponseContent_PAAS_MASTER VmlistServerPortsDetailRespResponseContentProductTypeEnum = "PAAS_MASTER"
	VmlistServerPortsDetailRespResponseContent_PAAS_SLAVE  VmlistServerPortsDetailRespResponseContentProductTypeEnum = "PAAS_SLAVE"
	VmlistServerPortsDetailRespResponseContent_VCPE        VmlistServerPortsDetailRespResponseContentProductTypeEnum = "VCPE"
	VmlistServerPortsDetailRespResponseContent_EMR         VmlistServerPortsDetailRespResponseContentProductTypeEnum = "EMR"
	VmlistServerPortsDetailRespResponseContent_LOGAUDIT    VmlistServerPortsDetailRespResponseContentProductTypeEnum = "LOGAUDIT"
	VmlistServerPortsDetailRespResponseContent_MSE         VmlistServerPortsDetailRespResponseContentProductTypeEnum = "MSE"
)

type VmlistServerPortsDetailRespResponseContentBootVolumeTypeEnum string

// List of BootVolumeType
const (
	VmlistServerPortsDetailRespResponseContent_LOCAL                   VmlistServerPortsDetailRespResponseContentBootVolumeTypeEnum = "local"
	VmlistServerPortsDetailRespResponseContent_CAPACITY                VmlistServerPortsDetailRespResponseContentBootVolumeTypeEnum = "capacity"
	VmlistServerPortsDetailRespResponseContent_HIGHPERFORMANCE         VmlistServerPortsDetailRespResponseContentBootVolumeTypeEnum = "highPerformance"
	VmlistServerPortsDetailRespResponseContent_PERFORMANCEOPTIMIZATION VmlistServerPortsDetailRespResponseContentBootVolumeTypeEnum = "performanceOptimization"
)

type VmlistServerPortsDetailRespResponseContentEcStatusEnum string

// List of EcStatus
const (
	VmlistServerPortsDetailRespResponseContent_ACTIVE            VmlistServerPortsDetailRespResponseContentEcStatusEnum = "ACTIVE"
	VmlistServerPortsDetailRespResponseContent_BUILD             VmlistServerPortsDetailRespResponseContentEcStatusEnum = "BUILD"
	VmlistServerPortsDetailRespResponseContent_REBUILD           VmlistServerPortsDetailRespResponseContentEcStatusEnum = "REBUILD"
	VmlistServerPortsDetailRespResponseContent_RESCUE            VmlistServerPortsDetailRespResponseContentEcStatusEnum = "RESCUE"
	VmlistServerPortsDetailRespResponseContent_SUSPENDED         VmlistServerPortsDetailRespResponseContentEcStatusEnum = "SUSPENDED"
	VmlistServerPortsDetailRespResponseContent_PAUSED            VmlistServerPortsDetailRespResponseContentEcStatusEnum = "PAUSED"
	VmlistServerPortsDetailRespResponseContent_RESIZE            VmlistServerPortsDetailRespResponseContentEcStatusEnum = "RESIZE"
	VmlistServerPortsDetailRespResponseContent_VERIFY_RESIZE     VmlistServerPortsDetailRespResponseContentEcStatusEnum = "VERIFY_RESIZE"
	VmlistServerPortsDetailRespResponseContent_REVERT_RESIZE     VmlistServerPortsDetailRespResponseContentEcStatusEnum = "REVERT_RESIZE"
	VmlistServerPortsDetailRespResponseContent_PASSWORD          VmlistServerPortsDetailRespResponseContentEcStatusEnum = "PASSWORD"
	VmlistServerPortsDetailRespResponseContent_REBOOT            VmlistServerPortsDetailRespResponseContentEcStatusEnum = "REBOOT"
	VmlistServerPortsDetailRespResponseContent_HARD_REBOOT       VmlistServerPortsDetailRespResponseContentEcStatusEnum = "HARD_REBOOT"
	VmlistServerPortsDetailRespResponseContent_DELETED           VmlistServerPortsDetailRespResponseContentEcStatusEnum = "DELETED"
	VmlistServerPortsDetailRespResponseContent_UNKNOWN           VmlistServerPortsDetailRespResponseContentEcStatusEnum = "UNKNOWN"
	VmlistServerPortsDetailRespResponseContent_ERROR             VmlistServerPortsDetailRespResponseContentEcStatusEnum = "ERROR"
	VmlistServerPortsDetailRespResponseContent_STOPPED           VmlistServerPortsDetailRespResponseContentEcStatusEnum = "STOPPED"
	VmlistServerPortsDetailRespResponseContent_SHUTOFF           VmlistServerPortsDetailRespResponseContentEcStatusEnum = "SHUTOFF"
	VmlistServerPortsDetailRespResponseContent_MIGRATING         VmlistServerPortsDetailRespResponseContentEcStatusEnum = "MIGRATING"
	VmlistServerPortsDetailRespResponseContent_SHELVED           VmlistServerPortsDetailRespResponseContentEcStatusEnum = "SHELVED"
	VmlistServerPortsDetailRespResponseContent_SHELVED_OFFLOADED VmlistServerPortsDetailRespResponseContentEcStatusEnum = "SHELVED_OFFLOADED"
	VmlistServerPortsDetailRespResponseContent_SOFT_DELETED      VmlistServerPortsDetailRespResponseContentEcStatusEnum = "SOFT_DELETED"
	VmlistServerPortsDetailRespResponseContent_UNRECOGNIZED_EC   VmlistServerPortsDetailRespResponseContentEcStatusEnum = "UNRECOGNIZED"
)

type VmlistServerPortsDetailRespResponseContentServerTypeEnum string

// List of ServerType
const (
	VmlistServerPortsDetailRespResponseContent_VM     VmlistServerPortsDetailRespResponseContentServerTypeEnum = "VM"
	VmlistServerPortsDetailRespResponseContent_IRONIC VmlistServerPortsDetailRespResponseContentServerTypeEnum = "IRONIC"
	VmlistServerPortsDetailRespResponseContent_EBM    VmlistServerPortsDetailRespResponseContentServerTypeEnum = "EBM"
)

type VmlistServerPortsDetailRespResponseContentServerVmTypeEnum string

// List of ServerVmType
const (
	VmlistServerPortsDetailRespResponseContent_HIGHIO                 VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "highIO"
	VmlistServerPortsDetailRespResponseContent_EXCLUSIVE              VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "exclusive"
	VmlistServerPortsDetailRespResponseContent_MEMIMPROVE             VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "memImprove"
	VmlistServerPortsDetailRespResponseContent_COMMON                 VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "common"
	VmlistServerPortsDetailRespResponseContent_GPU                    VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "gpu"
	VmlistServerPortsDetailRespResponseContent_HIGHPERFORMANCE_VM     VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "highPerformance"
	VmlistServerPortsDetailRespResponseContent_COMMONINTRODUCTORY     VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "commonIntroductory"
	VmlistServerPortsDetailRespResponseContent_COMMONNETIMPROVE       VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "commonNetImprove"
	VmlistServerPortsDetailRespResponseContent_COMPUTE                VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "compute"
	VmlistServerPortsDetailRespResponseContent_COMPUTENETIMPROVE      VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "computeNetImprove"
	VmlistServerPortsDetailRespResponseContent_MEMNETIMPROVE          VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "memNetImprove"
	VmlistServerPortsDetailRespResponseContent_LOCALSTORAGE           VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "localStorage"
	VmlistServerPortsDetailRespResponseContent_XLARGEMEMORY           VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "xlargeMemory"
	VmlistServerPortsDetailRespResponseContent_HIGHFREQUENCY          VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "highFrequency"
	VmlistServerPortsDetailRespResponseContent_VGPU                   VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "vgpu"
	VmlistServerPortsDetailRespResponseContent_FPGA                   VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "fpga"
	VmlistServerPortsDetailRespResponseContent_NORMALCOMPUTEIMPROVE   VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "normalComputeImprove"
	VmlistServerPortsDetailRespResponseContent_NORMALNETENHANCE       VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "normalNetEnhance"
	VmlistServerPortsDetailRespResponseContent_STOREENHANCE           VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "storeEnhance"
	VmlistServerPortsDetailRespResponseContent_EDGECOMMONNETIMPROVE   VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "edgeCommonNetImprove"
	VmlistServerPortsDetailRespResponseContent_GENERAL                VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "general"
	VmlistServerPortsDetailRespResponseContent_COMPUTEG2              VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "computeG2"
	VmlistServerPortsDetailRespResponseContent_NPU                    VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "npu"
	VmlistServerPortsDetailRespResponseContent_MEMORYDOUBLECARD       VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "memoryDoubleCard"
	VmlistServerPortsDetailRespResponseContent_MEMORYHBACARD          VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "memoryHBACard"
	VmlistServerPortsDetailRespResponseContent_LOCALSTORAGELOW        VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "localStorageLow"
	VmlistServerPortsDetailRespResponseContent_LOCALSTORAGEDOUBLECARD VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "localStorageDoubleCard"
	VmlistServerPortsDetailRespResponseContent_EDGEBALANCE            VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "edgeBalance"
	VmlistServerPortsDetailRespResponseContent_UNRECOGNIZED_VM        VmlistServerPortsDetailRespResponseContentServerVmTypeEnum = "UNRECOGNIZED"
)

type VmlistServerPortsDetailRespResponseContentOpStatusEnum string

// List of OpStatus
const (
	VmlistServerPortsDetailRespResponseContent_CREATING             VmlistServerPortsDetailRespResponseContentOpStatusEnum = "CREATING"
	VmlistServerPortsDetailRespResponseContent_OK                   VmlistServerPortsDetailRespResponseContentOpStatusEnum = "OK"
	VmlistServerPortsDetailRespResponseContent_REBOOTING            VmlistServerPortsDetailRespResponseContentOpStatusEnum = "REBOOTING"
	VmlistServerPortsDetailRespResponseContent_FROZEN               VmlistServerPortsDetailRespResponseContentOpStatusEnum = "FROZEN"
	VmlistServerPortsDetailRespResponseContent_BACKUP_REBUILDING    VmlistServerPortsDetailRespResponseContentOpStatusEnum = "BACKUP_REBUILDING"
	VmlistServerPortsDetailRespResponseContent_REBUILDING           VmlistServerPortsDetailRespResponseContentOpStatusEnum = "REBUILDING"
	VmlistServerPortsDetailRespResponseContent_RESIZING             VmlistServerPortsDetailRespResponseContentOpStatusEnum = "RESIZING"
	VmlistServerPortsDetailRespResponseContent_SHELVING             VmlistServerPortsDetailRespResponseContentOpStatusEnum = "SHELVING"
	VmlistServerPortsDetailRespResponseContent_UNSHELVING           VmlistServerPortsDetailRespResponseContentOpStatusEnum = "UNSHELVING"
	VmlistServerPortsDetailRespResponseContent_STARTING             VmlistServerPortsDetailRespResponseContentOpStatusEnum = "STARTING"
	VmlistServerPortsDetailRespResponseContent_STOPING              VmlistServerPortsDetailRespResponseContentOpStatusEnum = "STOPING"
	VmlistServerPortsDetailRespResponseContent_PASSWORD_UPDATING    VmlistServerPortsDetailRespResponseContentOpStatusEnum = "PASSWORD_UPDATING"
	VmlistServerPortsDetailRespResponseContent_BACKUPING            VmlistServerPortsDetailRespResponseContentOpStatusEnum = "BACKUPING"
	VmlistServerPortsDetailRespResponseContent_IMAGE_CREATING       VmlistServerPortsDetailRespResponseContentOpStatusEnum = "IMAGE_CREATING"
	VmlistServerPortsDetailRespResponseContent_UNRECOGNIZED_VMOP    VmlistServerPortsDetailRespResponseContentOpStatusEnum = "UNRECOGNIZED"
	VmlistServerPortsDetailRespResponseContent_BOOT_VOLUME_RESIZING VmlistServerPortsDetailRespResponseContentOpStatusEnum = "BOOT_VOLUME_RESIZING"
	VmlistServerPortsDetailRespResponseContent_DELETING             VmlistServerPortsDetailRespResponseContentOpStatusEnum = "DELETING"
	VmlistServerPortsDetailRespResponseContent_SECOND_DELETING      VmlistServerPortsDetailRespResponseContentOpStatusEnum = "SECOND_DELETING"
	VmlistServerPortsDetailRespResponseContent_MIGRATING_VMOP       VmlistServerPortsDetailRespResponseContentOpStatusEnum = "MIGRATING"
)

type VmlistServerPortsDetailRespResponseContent struct {

	// ????????????????????????
	ImageOsType VmlistServerPortsDetailRespResponseContentImageOsTypeEnum `json:"imageOsType,omitempty"`

	// ????????????
	SpecsName string `json:"specsName,omitempty"`

	// cpu??????
	Vcpu int32 `json:"vcpu,omitempty"`

	// operationFlag
	OperationFlag int32 `json:"operationFlag,omitempty"`

	// ??????id
	GroupId string `json:"groupId,omitempty"`

	// ??????????????????????????????
	PortDetail *[]VmlistServerPortsDetailRespResponsePortDetail `json:"portDetail,omitempty"`

	// ??????????????????
	AdminPaused bool `json:"adminPaused,omitempty"`

	// az
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// ????????????
	Vmemory int32 `json:"vmemory,omitempty"`

	// ????????????
	CreatedTime string `json:"createdTime,omitempty"`

	// ?????????id
	Id string `json:"id,omitempty"`

	// ?????????????????????????????????
	ProductType VmlistServerPortsDetailRespResponseContentProductTypeEnum `json:"productType,omitempty"`

	// ?????????volume??????
	BootVolumeType VmlistServerPortsDetailRespResponseContentBootVolumeTypeEnum `json:"bootVolumeType,omitempty"`

	// ?????????????????????
	EcStatus VmlistServerPortsDetailRespResponseContentEcStatusEnum `json:"ecStatus,omitempty"`

	// ??????????????????
	ImageName string `json:"imageName,omitempty"`

	// ????????????
	Visible bool `json:"visible,omitempty"`

	// ???????????????
	KeyName string `json:"keyName,omitempty"`

	// ??????????????????????????????
	ServerBackupPolicyResp VmlistServerPortsDetailRespResponseServerBackupPolicyResp `json:"serverBackupPolicyResp,omitempty"`

	// ??????????????????
	UserName string `json:"userName,omitempty"`

	// ?????????id
	SystemDiskId string `json:"systemDiskId,omitempty"`

	// ???????????????????????????
	Vaz string `json:"vaz,omitempty"`

	// ?????????????????????
	HwHost string `json:"hwHost,omitempty"`

	// ???????????????
	Deleted int32 `json:"deleted,omitempty"`

	// ????????????
	MaxPorts int32 `json:"maxPorts,omitempty"`

	// ??????????????????
	Task string `json:"task,omitempty"`

	// ???????????????
	DeCloudServerName string `json:"deCloudServerName,omitempty"`

	// ???????????????
	ServerType VmlistServerPortsDetailRespResponseContentServerTypeEnum `json:"serverType,omitempty"`

	// ???????????????
	Name string `json:"name,omitempty"`

	// ?????????id
	PoolId string `json:"poolId,omitempty"`

	// ??????id
	FlavorRef string `json:"flavorRef,omitempty"`

	// ?????????????????????????????????
	ServerVmType VmlistServerPortsDetailRespResponseContentServerVmTypeEnum `json:"serverVmType,omitempty"`

	// ????????????id
	ImageRef string `json:"imageRef,omitempty"`

	// ???????????????id
	HwHostId string `json:"hwHostId,omitempty"`

	// ???????????????
	OpStatus VmlistServerPortsDetailRespResponseContentOpStatusEnum `json:"opStatus,omitempty"`

	// ?????????
	Region string `json:"region,omitempty"`

	// ???????????????
	Vdisk int32 `json:"vdisk,omitempty"`

	// ????????????
	Status int32 `json:"status,omitempty"`
}
