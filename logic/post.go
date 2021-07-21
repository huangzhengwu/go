package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	// 1.生成post id
	p.ID = snowflake.GetId()
	// 2.保存数据到数据库
	return mysql.CreatePost(p)
	// 3.返回

}
