package system

import (
	"fresh-shop/server/global"
	"github.com/satori/go.uuid"
	"time"
)

type SysUser struct {
	global.DbModel
	UUID           uuid.UUID      `json:"uuid" gorm:"index;comment:用户UUID"`                                                         // 用户UUID
	OpenId         string         `json:"openId" gorm:"index;comment:OpenId"`                                                       // OpenId
	Username       string         `json:"userName" gorm:"index;comment:用户登录名"`                                                      // 用户登录名
	InvitationCode string         `json:"invitationCode" gorm:"index;comment:推荐码"`                                                  // 推荐码
	Password       string         `json:"-"  gorm:"comment:用户登录密码"`                                                                 // 用户登录密码
	SafePassword   string         `json:"-"  gorm:"comment:用户安全密码"`                                                                 // 用户安全密码
	NickName       string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                                // 用户昵称
	Sex            int            `json:"sex" gorm:"default:0;comment:性别"`                                                          // 用户昵称
	SideMode       string         `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                              // 用户侧边主题
	HeaderImg      string         `json:"headerImg" gorm:"default:https://minio.kl.do/picture/images/avatar/face.png;comment:用户头像"` // 用户头像
	BaseColor      string         `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                               // 基础颜色
	ActiveColor    string         `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`                                          // 活跃颜色
	AuthorityId    uint           `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                            // 用户角色ID
	Authority      SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities    []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Phone          string         `json:"phone"  gorm:"comment:用户手机号"`                     // 用户手机号
	Email          string         `json:"email"  gorm:"comment:用户邮箱"`                      // 用户邮箱
	Enable         int            `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
	LoginIp        string         `json:"loginIp" gorm:"comment:登录IP"`                     // 登录IP
	LoginTime      time.Time      `json:"loginTime" gorm:"comment:最后登录时间"`                 // 最后登录时间

	OriginContactName  string    `json:"originContactName" gorm:"column:origin_contact_name;comment:联系人姓名"`                             // 联系人姓名
	ChangeContactName  string    `json:"changeContactName" gorm:"column:change_contact_name;comment:联系人姓名修改待审核"`                        // 联系人姓名修改待审核
	OriginCustomerName string    `json:"originCustomerName" gorm:"column:origin_customer_name;comment:客户名称、公司名称"`                       // 客户名称、公司名称
	ChangeCustomerName string    `json:"changeCustomerName" gorm:"column:change_customer_name;comment:客户名称、公司名称待审核"`                    // 客户名称、公司名称待审核
	AuditStatus        int8      `json:"auditStatus" gorm:"column:audit_status;default:0;comment:审核状态 0新注册未填写信息 1已通过 2已填写未审核 3修改信息待审核"` // 审核状态 0新注册未填写信息 1已通过 2已填写未审核 3修改信息待审核 4已拒绝
	AuditRemark        string    `json:"auditRemark" gorm:"column:audit_remark;comment:审核备注"`                                           // 审核备注
	ApplyTime          time.Time `json:"applyTime" gorm:"column:apply_time;comment:审核时间"`

	Account any `json:"account" gorm:"-"`
}

func (SysUser) TableName() string {
	return "sys_users"
}
