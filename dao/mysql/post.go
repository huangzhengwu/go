package mysql

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"strings"
	"web_app/models"
)

// CreatePost 添加帖子
func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post (post_id,title,content,author_id,community_id)
	values (?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

// GetPostDetail 获取帖子详情
func GetPostDetail(pId int64) (data *models.Post, err error) {
	data = new(models.Post)
	//data = &models.Post{}
	sqlStr := `select post_id,title,content,author_id,community_id,status,create_time from post where post_id = ?`
	if err = db.Get(data, sqlStr, pId); err != nil {
		zap.L().Error("mysql post GetPostDetail err ", zap.Error(err))
		return
	}
	return
}

// GetPostList 获取帖子列表
func GetPostList(page int64, size int64) (dataList []*models.Post, err error) {
	sqlStr := `select 
	post_id,title,content,author_id,community_id,status,create_time 
	from post
	order by create_time desc 
	limit ?,?`
	dataList = make([]*models.Post, 0, 2)
	if err = db.Select(&dataList, sqlStr, (page-1)*size, size); err != nil {
		zap.L().Error("mysql post GetPostDetail err ", zap.Error(err))
		return
	}
	return
}

// 根据指定的id获取列表
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time
	from post
	where post_id in (?)
	order by FIND_IN_SET(post_id,?)
	`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}
