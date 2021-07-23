package models

// 定义请求的参数结构体

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamSignUp 注册结构体
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录结构体
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// VoteData 投票
type ParamVoteData struct {
	PostID    string `json:"post_id" binding:"required"`       //帖子
	Direction int8   `json:"direction" binding:"oneof=-1 0 1"` //1：赞同票  -1：反对票   0:取消投票
}

// ParamPostList 获取帖子列表query string参数
type ParamPostList struct {
	Page  int64  `json:"page" form:"page"`
	Size  int64  `json:"size" form:"size"`
	Order string `json:"order" form:"order"`
}

// ParamCommunityPostList 获取帖子列表query string参数
type ParamCommunityPostList struct {
	*ParamPostList
	CommunityID int64 `json:"community_id" form:"community_id"`
}
