package logic

import (
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
)

// 存放业务逻辑代码

func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 2.生成UID
	userID := snowflake.GetId()

	//构建一个user实例
	user := &models.User{
		UserId:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 3.保存进入数据库
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err = mysql.Login(user); err != nil {
		zap.L().Error("Login Login err:", zap.Error(err))
		return "", err
	}
	// 如果登录成功，就生成token
	return jwt.GenToken(user.UserId, user.Username)
}
