// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model


type VmGetRebuildImageResponseBodyOsTypeEnum string

// List of OsType
const (
    VmGetRebuildImageResponseBody_LINUX VmGetRebuildImageResponseBodyOsTypeEnum = "linux"
    VmGetRebuildImageResponseBody_WINDOWS VmGetRebuildImageResponseBodyOsTypeEnum = "windows"
    VmGetRebuildImageResponseBody_OTHER VmGetRebuildImageResponseBodyOsTypeEnum = "other"
)
type VmGetRebuildImageResponseBodyStatusEnum string

// List of Status
const (
    VmGetRebuildImageResponseBody_UNRECOGNIZED VmGetRebuildImageResponseBodyStatusEnum = "unrecognized"
    VmGetRebuildImageResponseBody_ACTIVE VmGetRebuildImageResponseBodyStatusEnum = "active"
    VmGetRebuildImageResponseBody_SAVING VmGetRebuildImageResponseBodyStatusEnum = "saving"
    VmGetRebuildImageResponseBody_QUEUED VmGetRebuildImageResponseBodyStatusEnum = "queued"
    VmGetRebuildImageResponseBody_KILLED VmGetRebuildImageResponseBodyStatusEnum = "killed"
    VmGetRebuildImageResponseBody_PENDING_DELETE VmGetRebuildImageResponseBodyStatusEnum = "pending_delete"
    VmGetRebuildImageResponseBody_DELETED VmGetRebuildImageResponseBodyStatusEnum = "deleted"
    VmGetRebuildImageResponseBody_CACHING VmGetRebuildImageResponseBodyStatusEnum = "caching"
)

type VmGetRebuildImageResponseBody struct {

	// 更新时间
	ModifiedTime string `json:"modifiedTime,omitempty"`

	// 备注
	Note string `json:"note,omitempty"`

	// 源镜像ID
	SourceImageId string `json:"sourceImageId,omitempty"`

	// 镜像来源类型
	Type int32 `json:"type,omitempty"`

	// 主机规格需要的系统盘最小大小，单位G
	MinDisk int64 `json:"minDisk,omitempty"`

	// 云主机id
	ServerId string `json:"serverId,omitempty"`

	// 镜像路径
	Url string `json:"url,omitempty"`

	// 原镜像所属资源池ID
	SourceImagePoolId string `json:"sourceImagePoolId,omitempty"`

	// 镜像大小,单位M
	Size int64 `json:"size,omitempty"`

	// 镜像名称
	Name string `json:"name,omitempty"`

	// 操作系统类型:linux,windows
	OsType VmGetRebuildImageResponseBodyOsTypeEnum `json:"osType,omitempty"`

	// 创建时间
	CreatedTime string `json:"createdTime,omitempty"`

	// 镜像id
	Id string `json:"id,omitempty"`

	// 镜像状态
	Status VmGetRebuildImageResponseBodyStatusEnum `json:"status,omitempty"`
}
