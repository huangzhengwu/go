package logic

import (
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	// 1.生成post id
	p.ID = snowflake.GetId()
	// 2.保存数据到数据库
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID)
	return err
}

func GetPostDetail(pId int64) (data *models.PostDetail, err error) {
	post, err := mysql.GetPostDetail(pId)
	if err != nil {
		zap.L().Error("logic mysql.GetPostDetail(pId) err", zap.Error(err))
		return
	}
	user, err := mysql.GetUserDetail(post.AuthorID)
	if err != nil {
		zap.L().Error("logic mysql.GetUserDetail(post.AuthorID) err", zap.Error(err))
		return
	}
	community, err := mysql.GetCommunityDetail(post.CommunityID)
	if err != nil {
		zap.L().Error("logic mysql.GetCommunityDetail(post.CommunityID) err", zap.Error(err))
		return
	}
	data = &models.PostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

// GetPostList 获取帖子列表
func GetPostList(page int64, size int64) (data []*models.PostDetail, err error) {
	postList, err := mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic mysql.GetPostList(pId) err", zap.Error(err))
		return
	}
	data = make([]*models.PostDetail, 0, len(postList))
	for _, post := range postList {
		user, err := mysql.GetUserDetail(post.AuthorID)
		if err != nil {
			zap.L().Error("logic mysql.GetUserDetail(post.AuthorID) err", zap.Error(err))
			return nil, err
		}
		community, err := mysql.GetCommunityDetail(post.CommunityID)
		if err != nil {
			zap.L().Error("logic mysql.GetCommunityDetail(post.CommunityID) err", zap.Error(err))
			return nil, err
		}
		postDetail := &models.PostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

// GetPostListRedis 获取帖子列表 根据redis
func GetPostListRedis(p *models.ParamPostList) (data []*models.PostDetail, err error) {
	// 2.去redis查询id列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
		return
	}
	// 3.根据id去数据库查询帖子详情信息
	// 返回的数据还要按照给定的数据顺序返回
	postList, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	//将帖子的作者和分区信息填充到帖子中
	for _, post := range postList {
		user, err := mysql.GetUserDetail(post.AuthorID)
		if err != nil {
			zap.L().Error("logic mysql.GetUserDetail(post.AuthorID) err", zap.Error(err))
			return nil, err
		}
		community, err := mysql.GetCommunityDetail(post.CommunityID)
		if err != nil {
			zap.L().Error("logic mysql.GetCommunityDetail(post.CommunityID) err", zap.Error(err))
			return nil, err
		}
		postDetail := &models.PostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return

}
