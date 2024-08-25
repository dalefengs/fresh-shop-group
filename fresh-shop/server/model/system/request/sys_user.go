package request

import (
	"fresh-shop/server/model/common/request"
	"fresh-shop/server/model/system"
	"time"
)

// Register User register structure
type Register struct {
	Username     string `json:"userName" file:"用户名"`
	Password     string `json:"passWord" file:"密码"`
	NickName     string `json:"nickName" file:"昵称"`
	HeaderImg    string `json:"headerImg" file:"头像链接"`
	AuthorityId  uint   `json:"authorityId" swaggertype:"string" file:"int 角色id"`
	Enable       int    `json:"enable" swaggertype:"string" file:"int 是否启用"`
	AuthorityIds []uint `json:"authorityIds" swaggertype:"string" file:"[]uint 角色id"`
	Phone        string `json:"phone" file:"电话号码"`
	Email        string `json:"email" file:"电子邮箱"`
}

// User wechat structure
type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// Modify password structure
type ChangePasswordReq struct {
	ID          uint   `json:"-"`           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// Modify  user's auth structure
type SetUserAuth struct {
	AuthorityId uint `json:"authorityId"` // 角色ID
}

// Modify  user's auth structure
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []uint `json:"authorityIds"` // 角色ID
}

type ChangeUserInfo struct {
	ID                 uint      `gorm:"primarykey"`                                                                                    // 主键ID
	NickName           string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                                     // 用户昵称
	Phone              string    `json:"phone"  gorm:"comment:用户手机号"`                                                                   // 用户手机号
	AuthorityIds       []uint    `json:"authorityIds" gorm:"-"`                                                                         // 角色ID
	Email              string    `json:"email"  gorm:"comment:用户邮箱"`                                                                    // 用户邮箱
	HeaderImg          string    `json:"headerImg" gorm:"default:https://minio.kl.do/picture/images/avatar/face.png;comment:用户头像"`      // 用户头像
	SideMode           string    `json:"sideMode"  gorm:"comment:用户侧边主题"`                                                               // 用户侧边主题
	Enable             int       `json:"enable" gorm:"comment:冻结用户"`                                                                    //冻结用户
	OriginContactName  string    `json:"originContactName" gorm:"column:origin_contact_name;comment:联系人姓名"`                             // 联系人姓名
	ChangeContactName  string    `json:"changeContactName" gorm:"column:change_contact_name;comment:联系人姓名修改待审核"`                        // 联系人姓名修改待审核
	OriginCustomerName string    `json:"originCustomerName" gorm:"column:origin_customer_name;comment:客户名称、公司名称"`                       // 客户名称、公司名称
	ChangeCustomerName string    `json:"changeCustomerName" gorm:"column:change_customer_name;comment:客户名称、公司名称待审核"`                    // 客户名称、公司名称待审核
	AuditStatus        int8      `json:"auditStatus" gorm:"column:audit_status;default:1;comment:审核状态 0新注册未填写信息 1已通过 2已填写未审核 3修改信息待审核"` // 审核状态 0新注册未填写信息 1已通过 2已填写未审核 3修改信息待审核 4已拒绝
	AuditRemark        string    `json:"auditRemark" gorm:"column:audit_remark;default:'';comment:审核备注"`                                // 审核备注
	ApplyTime          time.Time `json:"applyTime" gorm:"column:apply_time;comment:审核时间"`                                               // 审核备注

	Authorities []system.SysAuthority `json:"-" gorm:"many2many:sys_user_authority;"`
}

type UserList struct {
	request.PageInfo
	Username       string `json:"username"`
	NickName       string `json:"nickName"`
	ContactName    string `json:"contactName"`
	CustomerName   string `json:"customerName"`
	InvitationCode string `json:"invitationCode"`
	Phone          string `json:"phone"`
	OrderKey       string `json:"orderKey"` // 排序
	Desc           bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
	AuditStatus    *int8  `json:"auditStatus"`
}
