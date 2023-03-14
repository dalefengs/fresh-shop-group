package system

import (
	"errors"
	"fmt"
	"fresh-shop/server/model/account"
	sysReq "fresh-shop/server/model/system/request"
	"time"

	"fresh-shop/server/global"
	"fresh-shop/server/model/system"
	"fresh-shop/server/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/likfees)
//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: userInter system.SysUser, err error

type UserService struct{}

func (userService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	// 获取账户列表
	var groupList []account.AccountGroup
	err = global.DB.Find(&groupList).Error
	if err != nil {
		global.SugarLog.Errorf("注册用户 --- 用户名：%s,账户配置获取失败,err: %s", u.Username, err.Error())
		return u, errors.New("账户配置获取失败")
	}

	var user system.SysUser
	if u.Username == "" { // 生成用户名
		u.Username = utils.GenerateUsername("U", 6)
		// 判断错误是否是由 gorm.ErrRecordNotFound 引起的
		for !errors.Is(global.DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
			name := utils.GenerateUsername("U", 6)
			global.SugarLog.Warnf("注册用户 --- 用户名：%s 已存在！重新生成新用户名：%s", u.Username, name)
			u.Username = name
		}
		global.SugarLog.Infof("注册用户 --- 用户名：%s 验证通过", u.Username)
	} else {
		if !errors.Is(global.DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
			return userInter, errors.New("用户名已注册")
		}
	}

	user = system.SysUser{}
	u.InvitationCode = utils.GenerateInviteCode(8)
	// 生成邀请码
	for !errors.Is(global.DB.Where("invitation_code = ?", u.InvitationCode).First(&user).Error, gorm.ErrRecordNotFound) {
		code := utils.GenerateInviteCode(8)
		global.SugarLog.Warnf("注册用户 --- 用户名：%s, 邀请码：%s 已存在！重新生成新邀请码：%s", u.Username, u.InvitationCode, code)
		u.InvitationCode = code
	}

	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()

	// 生成对应币种账户信息
	// 开启事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 创建用户
		err = tx.Create(&u).Error
		var groupData []account.Account
		// 生成对应币种账户信息
		for _, group := range groupList {
			groupId := group.ID
			groupData = append(groupData, account.Account{
				GroupId: &groupId,
				UserId:  &u.ID,
			})
		}
		txErr := tx.Create(groupData).Error
		if txErr != nil {
			global.SugarLog.Errorf("注册用户 --- 用户名：%s, 创建账户配置失败, 插入数据: %#v ,err: %s", u.Username, groupData, txErr.Error())
			return txErr
		}
		// 提交事务
		return nil
	})

	return u, err
}

//@author: [piexlmax](https://github.com/likfees)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func (userService *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	if nil == global.DB {
		return nil, fmt.Errorf("db not init")
	}

	var user system.SysUser
	err = global.DB.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		// 验证密码
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("用户名不存在或者密码错误")
		}
		MenuServiceApp.UserAuthorityDefaultRouter(&user)
	}
	return &user, err
}

//@author: [piexlmax](https://github.com/likfees)
//@function: ChangePassword
//@description: 修改用户密码
//@param: u *model.SysUser, newPassword string
//@return: userInter *model.SysUser,err error

func (userService *UserService) ChangePassword(u *system.SysUser, newPassword string) (userInter *system.SysUser, err error) {
	var user system.SysUser
	if err = global.DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = global.DB.Save(&user).Error
	return &user, err

}

//@author: [piexlmax](https://github.com/likfees)
//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (userService *UserService) GetUserInfoList(info sysReq.UserList, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&system.SysUser{})
	var userList []system.SysUser

	if info.Username != "" {
		db = db.Where("username like ?", "%"+info.Username+"%")
	}
	if info.NickName != "" {
		db = db.Where("nick_name like ?", "%"+info.NickName+"%")
	}
	if info.Phone != "" {
		db = db.Where("phone like ?", "%"+info.Phone+"%")
	}
	if info.InvitationCode != "" {
		db = db.Where("invitation_code like ?", "%"+info.InvitationCode+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if order != "" {
		var orderStr string
		orderMap := make(map[string]bool, 1)
		orderMap["id"] = true
		if orderMap[order] {
			if desc { // 如果为 desc 为 true则加上
				orderStr = order + " desc"
			} else {
				orderStr = order
			}
		}
		db = db.Order(orderStr)
	}
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return userList, total, err
}

//@author: [piexlmax](https://github.com/likfees)
//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: uuid uuid.UUID, authorityId string
//@return: err error

func (userService *UserService) SetUserAuthority(id uint, authorityId uint) (err error) {
	assignErr := global.DB.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityId).First(&system.SysUserAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	err = global.DB.Where("id = ?", id).First(&system.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

//@author: [piexlmax](https://github.com/likfees)
//@function: SetUserAuthorities
//@description: 设置一个用户的权限
//@param: id uint, authorityIds []string
//@return: err error

func (userService *UserService) SetUserAuthorities(id uint, authorityIds []uint) (err error) {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		var useAuthority []system.SysUserAuthority
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, system.SysUserAuthority{
				SysUserId: id, SysAuthorityAuthorityId: v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = tx.Where("id = ?", id).First(&system.SysUser{}).Update("authority_id", authorityIds[0]).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回 nil 提交事务
		return nil
	})
}

//@author: [piexlmax](https://github.com/likfees)
//@function: DeleteUser
//@description: 删除用户
//@param: id float64
//@return: err error

func (userService *UserService) DeleteUser(id int) (err error) {
	var user system.SysUser
	err = global.DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	err = global.DB.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error
	return err
}

//@author: [piexlmax](https://github.com/likfees)
//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func (userService *UserService) SetUserInfo(req system.SysUser) error {
	return global.DB.Model(&system.SysUser{}).
		Select("updated_at", "nick_name", "header_img", "phone", "email", "sideMode", "enable").
		Where("id=?", req.ID).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"nick_name":  req.NickName,
			"header_img": req.HeaderImg,
			"phone":      req.Phone,
			"email":      req.Email,
			"side_mode":  req.SideMode,
			"enable":     req.Enable,
		}).Error
}

//@author: [piexlmax](https://github.com/likfees)
//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func (userService *UserService) SetSelfInfo(req system.SysUser) error {
	return global.DB.Model(&system.SysUser{}).
		Where("id=?", req.ID).
		Updates(req).Error
}

//@author: [piexlmax](https://github.com/likfees)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetUserInfo
//@description: 获取用户信息
//@param: uuid uuid.UUID
//@return: err error, user system.SysUser

func (userService *UserService) GetUserInfo(uuid uuid.UUID) (user system.SysUser, err error) {
	var reqUser system.SysUser
	err = global.DB.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	if err != nil {
		return reqUser, err
	}
	MenuServiceApp.UserAuthorityDefaultRouter(&reqUser)
	return reqUser, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func (userService *UserService) FindUserById(id int) (user *system.SysUser, err error) {
	var u system.SysUser
	err = global.DB.Where("`id` = ?", id).First(&u).Error
	return &u, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *model.SysUser

func (userService *UserService) FindUserByUuid(uuid string) (user *system.SysUser, err error) {
	var u system.SysUser
	if err = global.DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return &u, errors.New("用户不存在")
	}
	return &u, nil
}

//@author: [piexlmax](https://github.com/likfees)
//@function: resetPassword
//@description: 修改用户密码
//@param: ID uint
//@return: err error

func (userService *UserService) ResetPassword(ID uint) (err error) {
	err = global.DB.Model(&system.SysUser{}).Where("id = ?", ID).Update("password", utils.BcryptHash("123456")).Error
	return err
}
