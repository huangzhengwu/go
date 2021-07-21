package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"go.uber.org/zap"
	"web_app/models"
)

const secret = "test_go"

// 判断用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		zap.L().Error("CheckUserExist, err:%v\n", zap.Error(err))
		return err
	}
	if count > 0 {
		zap.L().Error("CheckUserExist")
		return errors.New("用户已经存在")
	}
	return
}

// InsertUser 插入数据库
func InsertUser(user *models.User) (err error) {
	//密码加密
	user.Password = encryptPassword(user.Password)
	sqlStr := `insert into user (user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserId, user.Username, user.Password)
	return err
}

// encryptPassword md5加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

//Login 获取用户信息
func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id,username,password from user where username = ?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return errors.New("用户不存在")
	}
	if err != nil {
		return err
	}
	if encryptPassword(oPassword) != user.Password {
		return errors.New("用户名或密码错误")
	}
	return
}
