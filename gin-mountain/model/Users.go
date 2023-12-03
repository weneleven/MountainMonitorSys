package model

import (
	"encoding/base64"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
	"mountain/utils/errmessage"
)

type User struct {
	gorm.Model        //gorm提供的基础模型定义
	Username   string `gorm:"type:varchar(20);not null" json:"username"`
	Password   string `gorm:"type:varchar(20);not null" json:"password"`
	Role       int    `gorm:"type:int" json:"role"` //身份
}

// 查询用户是否存在
func CheckUserName(name string) (code int) {
	var users User
	//数据库里寻找Username等于形参的第一个数据
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 { //编号大于零说明找到了有重复的
		return errmessage.ERROR_USERNAME_USED
	}
	return errmessage.SUCCESS
}

// 创造用户
func CreatUser(data *User) int {
	data.Password = ScryptPassword(data.Password) //密码加密
	err := db.Create(data).Error
	if err != nil {
		return errmessage.ERROR
	}
	return errmessage.SUCCESS
}

// 获取用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offset = -1
	}
	err = db.Limit(pageSize).Offset(offset).Find(&users).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return users
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
	var user User
	err = db.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return errmessage.ERROR
	}
	return errmessage.SUCCESS
}

// 编辑用户
func EditUser(id int, data *User) int {
	var user User
	maps := make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role

	err = db.Model(&user).Where("id=?", id).Update(maps).Error
	if err != nil {
		return errmessage.ERROR
	}
	return errmessage.SUCCESS
}

// 检查用户ID
func CheckUserID(ID int) int {
	var user User
	err := db.Where("id = ?", ID).First(&user).Error
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
	var user User
	db.Where("username=?", username).First(&user)
	if user.ID == 0 {
		return errmessage.ERROR_USER_NOT_EXIST
	}
	if ScryptPassword(password) != user.Password {
		return errmessage.ERROR_PASSWORD_WRONG
	}
	return errmessage.SUCCESS
}
