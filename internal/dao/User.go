package dao

import (
	"encoding/base64"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
	"mountain/global"
	"mountain/internal/model"
	errmessage "mountain/pkg/errcode"
)

// 查询用户名是否存在
func CheckUserName(name string) (code int) {
	var users model.User
	//数据库里寻找Username等于形参的第一个数据
	global.DBEngine.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 { //编号大于零说明找到了有重复的
		return errmessage.ERROR_USERNAME_USED
	}
	return errmessage.SUCCESS
}

// 查询用户手机号是否被注册
func CheckUserPhone(phone string) (code int) {
	var users model.User
	//数据库里寻找Username等于形参的第一个数据
	global.DBEngine.Select("id").Where("phone = ?", phone).First(&users)
	if users.ID > 0 { //编号大于零说明找到了有重复的
		return errmessage.ERROR_PHONE_USED
	}
	return errmessage.SUCCESS
}

// 创造用户
func CreatUser(data *model.User) int {
	data.Password = ScryptPassword(data.Password) //密码加密
	err := global.DBEngine.Create(data).Error
	if err != nil {
		return errmessage.ERROR
	}
	return errmessage.SUCCESS
}

// 获取用户列表
func GetUsers(pageSize int, pageNum int) ([]model.User, int64) {
	var users []model.User
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offset = -1
	}
	err := global.DBEngine.Limit(pageSize).Offset(offset).Find(&users).Count(&total).Error
	if err == gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

// 密码加密
func ScryptPassword(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{9, 42, 6, 9, 3, 8, 6, 5}
	HashPassword, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen) //N必须是2的幂次方
	if err != nil {
		log.Fatal(err)
	}
	Finalpassword := base64.StdEncoding.EncodeToString(HashPassword)
	return Finalpassword
}

// 删除用户
func DeleteUser(id int) int {
	var user model.User
	err := global.DBEngine.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return errmessage.ERROR
	}
	return errmessage.SUCCESS
}

// 编辑用户
func EditUser(id int, data *model.User) int {
	var user model.User
	maps := make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role

	err := global.DBEngine.Model(&user).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmessage.ERROR
	}
	return errmessage.SUCCESS
}

// 检查用户ID
func CheckUserID(ID int) int {
	var user model.User
	err := global.DBEngine.Where("id = ?", ID).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errmessage.ERROR_USER_NOT_EXIST // 用户不存在
		}
		return errmessage.ERROR // 查询出错
	}
	return errmessage.SUCCESS // 用户存在
}

// 检查登录
func CheckLogin(username string, password string) int {
	var user model.User
	global.DBEngine.Where("username=?", username).First(&user)
	if user.ID == 0 {
		return errmessage.ERROR_USER_NOT_EXIST
	}
	if ScryptPassword(password) != user.Password {
		return errmessage.ERROR_PASSWORD_WRONG
	}
	return errmessage.SUCCESS
}
