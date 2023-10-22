package system

import (
	"Chinese_Learning_App/global"
	"Chinese_Learning_App/model/system"
	"Chinese_Learning_App/utils"
	"errors"

	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type UserService struct{}

// Register 用户注册服务
func (userService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	// 注册
	var user system.SysUser
	// 判断用户是否存在
	if !errors.Is(global.CLA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("用户已经注册")
	}
	//密码hash加密
	u.Password = utils.BcryptHash(u.Password)
	// 否则 附加uuid
	u.UUID = uuid.Must(uuid.NewV4())
	//注册
	err = global.CLA_DB.Create(&u).Error
	return u, err
}

// Login 用户登录服务
