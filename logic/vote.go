package logic

import (
	"go.uber.org/zap"
	"strconv"
	"web_app/dao/redis"
	"web_app/models"
)

// VoteForPost 为帖子投票
func VoteForPost(userId int64, p *models.ParamVoteData) error {
	zap.L().Debug("logic VoteForPost",
		zap.Int64("userId", userId),
		zap.String("post_id", p.PostID),
		zap.Int8("Direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userId)), p.PostID, float64(p.Direction))
}
