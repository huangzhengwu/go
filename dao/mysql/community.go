package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"web_app/models"
)

// CommunityHandle 获取社区列表数据
func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id,community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("have no data in db")
			err = nil
		}
	}
	return
}

// CommunityDetailHandle 根据社区id，获取详情数据
func GetCommunityDetail(id int64) (data *models.CommunityDetail, err error) {
	data = new(models.CommunityDetail)
	sqlStr := "select community_id,community_name,introduction from community where community_id = ?"
	if err := db.Get(data, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("have no data in db")
			err = nil
		}
	}
	return
}
