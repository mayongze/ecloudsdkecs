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

	// 镜像操作系统类型
	ImageOsType VmlistServerPortsDetailRespResponseContentImageOsTypeEnum `json:"imageOsType,omitempty"`

	// 规格名称
	SpecsName string `json:"specsName,omitempty"`

	// cpu个数
	Vcpu int32 `json:"vcpu,omitempty"`

	// operationFlag
	OperationFlag int32 `json:"operationFlag,omitempty"`

	// 群组id
	GroupId string `json:"groupId,omitempty"`

	// 云主机绑定的网络信息
	PortDetail *[]VmlistServerPortsDetailRespResponsePortDetail `json:"portDetail,omitempty"`

	// 主机是否封堵
	AdminPaused bool `json:"adminPaused,omitempty"`

	// az
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// 内存大小
	Vmemory int32 `json:"vmemory,omitempty"`

	// 创建时间
	CreatedTime string `json:"createdTime,omitempty"`

	// 云主机id
	Id string `json:"id,omitempty"`

	// 云主机属于哪个产品类型
	ProductType VmlistServerPortsDetailRespResponseContentProductTypeEnum `json:"productType,omitempty"`

	// 云主机volume类型
	BootVolumeType VmlistServerPortsDetailRespResponseContentBootVolumeTypeEnum `json:"bootVolumeType,omitempty"`

	// 云主机底层状态
	EcStatus VmlistServerPortsDetailRespResponseContentEcStatusEnum `json:"ecStatus,omitempty"`

	// 系统镜像名称
	ImageName string `json:"imageName,omitempty"`

	// 是否可见
	Visible bool `json:"visible,omitempty"`

	// 密钥对名称
	KeyName string `json:"keyName,omitempty"`

	// 云主机备份策略返回体
	ServerBackupPolicyResp VmlistServerPortsDetailRespResponseServerBackupPolicyResp `json:"serverBackupPolicyResp,omitempty"`

	// 订购用户名称
	UserName string `json:"userName,omitempty"`

	// 系统盘id
	SystemDiskId string `json:"systemDiskId,omitempty"`

	// 边缘计算虚拟可用区
	Vaz string `json:"vaz,omitempty"`

	// 所属物理机编号
	HwHost string `json:"hwHost,omitempty"`

	// 是否已删除
	Deleted int32 `json:"deleted,omitempty"`

	// 网卡数量
	MaxPorts int32 `json:"maxPorts,omitempty"`

	// 底层任务描述
	Task string `json:"task,omitempty"`

	// 宿主机名称
	DeCloudServerName string `json:"deCloudServerName,omitempty"`

	// 云主机类型
	ServerType VmlistServerPortsDetailRespResponseContentServerTypeEnum `json:"serverType,omitempty"`

	// 云主机名称
	Name string `json:"name,omitempty"`

	// 资源池id
	PoolId string `json:"poolId,omitempty"`

	// 规格id
	FlavorRef string `json:"flavorRef,omitempty"`

	// 云主机类型下的主机类型
	ServerVmType VmlistServerPortsDetailRespResponseContentServerVmTypeEnum `json:"serverVmType,omitempty"`

	// 系统镜像id
	ImageRef string `json:"imageRef,omitempty"`

	// 所属物理机id
	HwHostId string `json:"hwHostId,omitempty"`

	// 云主机状态
	OpStatus VmlistServerPortsDetailRespResponseContentOpStatusEnum `json:"opStatus,omitempty"`

	// 可用区
	Region string `json:"region,omitempty"`

	// 系统盘大小
	Vdisk int32 `json:"vdisk,omitempty"`

	// 资源状态
	Status int32 `json:"status,omitempty"`
}
