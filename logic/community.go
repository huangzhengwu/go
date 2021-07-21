package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
)

// CommunityHandle 获取社区列表数据
func GetCommunityList() ([]*models.Community, error) {
	//查数据库  查找到所有的community 并返回
	return mysql.GetCommunityList()
}

// CommunityDetailHandle 根据社区id，获取详情数据
func GetCommunityDetail(id int64) (data *models.CommunityDetail, err error) {
	return mysql.GetCommunityDetail(id)
}
