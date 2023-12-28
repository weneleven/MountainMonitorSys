package daonew

import (
	"fmt"
	"mountain/global"
	"mountain/internal/model"
	errmessage "mountain/pkg/errcode"
	"time"
)

// 检查登录
func CheckLogin(username string, password string) int {
	var user model.User1
	sqlStr := "SELECT id,username,password FROM u WHERE username = ?"
	stmt, err := global.Db.Prepare(sqlStr)
	stmt.QueryRow(username).Scan(&user.UserId, &user.Username, &user.Password)
	fmt.Println(err)
	fmt.Printf("id:%d name:%s password:%s\n", user.UserId, user.Username, user.Password)
	if user.UserId == 0 {
		return errmessage.ERROR_USER_NOT_EXIST
	}
	if password != user.Password {
		return errmessage.ERROR_PASSWORD_WRONG
	}
	return errmessage.SUCCESS
}

// 查询用户名是否存在
func CheckUserName(name string) (code int) {
	var user model.User1
	//数据库里寻找Username等于形参的第一个数据
	sqlStr := "SELECT id FROM u WHERE username = ?"
	stmt, err := global.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err)
	}
	stmt.QueryRow(name).Scan(&user.UserId)
	if user.UserId > 0 { //编号大于零说明找到了有重复的
		return errmessage.ERROR_USERNAME_USED
	}
	return errmessage.SUCCESS
}

// 查询用户手机号是否被注册
func CheckUserPhone(phone string) (code int) {
	var user model.User1
	//数据库里寻找Username等于形参的第一个数据
	sqlStr := "SELECT id FROM u WHERE phone = ?"
	stmt, err := global.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err)
	}
	stmt.QueryRow(phone).Scan(&user.UserId)
	if user.UserId > 0 { //编号大于零说明找到了有重复的
		return errmessage.ERROR_PHONE_USED
	}
	return errmessage.SUCCESS
}

// 创造用户
func CreatUser(data *model.User1) int {
	//data.Password = dao.ScryptPassword(data.Password) //密码加密
	now := time.Now()
	sqlStr := "insert into u(ts,id,username,password,phone) values (?,?,?,?,?)"
	stmt, err := global.Db.Prepare(sqlStr)
	_, err = stmt.Exec(now, data.UserId, data.Username, data.Password, data.Phone)
	fmt.Println(err)
	if err != nil {
		return errmessage.ERROR
	}
	return errmessage.SUCCESS
}
