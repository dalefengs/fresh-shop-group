package system

import (
	"fresh-shop/server/global"
	"github.com/satori/go.uuid"
	"time"
)

type SysUser struct {
	global.DbModel
	UUID           uuid.UUID      `json:"uuid" gorm:"index;comment:用户UUID"`                                                     // 用户UUID
	OpenId         string         `json:"openId" gorm:"index;comment:OpenId"`                                                   // OpenId
	Username       string         `json:"userName" gorm:"index;comment:用户登录名"`                                                  // 用户登录名
	InvitationCode string         `json:"invitationCode" gorm:"index;comment:推荐码"`                                              // 推荐码
	Password       string         `json:"-"  gorm:"comment:用户登录密码"`                                                             // 用户登录密码
	SafePassword   string         `json:"-"  gorm:"comment:用户安全密码"`                                                             // 用户安全密码
	NickName       string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	Sex            int            `json:"sex" gorm:"default:0;comment:性别"`                                                      // 用户昵称
	SideMode       string         `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                          // 用户侧边主题
	HeaderImg      string         `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	BaseColor      string         `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                           // 基础颜色
	ActiveColor    string         `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`                                      // 活跃颜色
	AuthorityId    uint           `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                        // 用户角色ID
	Authority      SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities    []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Phone          string         `json:"phone"  gorm:"comment:用户手机号"`                     // 用户手机号
	Email          string         `json:"email"  gorm:"comment:用户邮箱"`                      // 用户邮箱
	Enable         int            `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
	LoginIp        string         `json:"loginIp" gorm:"comment:登录IP"`                     // 登录IP
	LoginTime      *time.Time     `json:"loginTime" gorm:"comment:最后登录时间"`                 // 最后登录时间
	Account        any            `json:"account" gorm:"-"`
}

func (SysUser) TableName() string {
	return "sys_users"
}
