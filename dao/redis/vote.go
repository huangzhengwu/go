package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"time"
)

// 本项目使用简化版的投票分数
// 投一票就加432分   86400/200 -> 需要200张赞同票才可以给你的帖子续一天

/* 投票的几种情况
direction=1时，有两种情况
	1.之前没有投过票，现在投赞同票	--> 更新分数&投票记录   差值绝对值：1	+432
	2.之前投反对票，现在投赞同票		--> 更新分数&投票记录   差值绝对值：2	+432*2
direction=0时，有两种情况
	1.之前投过赞同票，现在要取消		--> 更新分数&投票记录   差值绝对值：1	-432
	2.之前投反对票，现在要取消		--> 更新分数&投票记录   差值绝对值：1	+432
direction=-1时，有两种情况
	1.之前没有投过票，现在要投反对票	--> 更新分数&投票记录   差值绝对值：1	-432
	2.之前投过赞同票，现在要取消		--> 更新分数&投票记录   差值绝对值：2	-432*2

投票的限制：
每个帖子自发表之日起，一个星期内允许投票，超过一个星期就不允许投票了
	1.到期之后，讲redis中保存的赞同票以及反对票存储到mysql表中
	2.到期之后删除key：KeyPostVotedZSetPrefix
*/
const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 //每一票值的多少分数
)

var (
	ErrVoteTimeExpire = errors.New("投票时间以过")
)

func CreatePost(postID int64) error {
	pipeline := client.TxPipeline()
	//帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	//帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	_, err := pipeline.Exec()
	return err
}

func VoteForPost(userID, postID string, value float64) error {
	// 1.判断投票的限制
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	// 2.更新帖子分数
	// 先查询当前用户给当前帖子之前的投票记录
	ov := client.ZScore(getRedisKey(KeyPostVotedZSetPrefix+postID), userID).Val()
	var dir float64 //正负方向
	if value > ov {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(ov - value) // 计算两次投票的差值
	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), dir*diff*scorePerVote, postID)
	// 3.记录用户为该帖子投过票
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPrefix+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPrefix+postID), redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}
