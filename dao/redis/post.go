package redis

import (
	"github.com/go-redis/redis"
	"strconv"
	"time"
	"web_app/models"
)

func getIDsFormKey(key string, page, size int64) ([]string, error) {
	start := (size - 1) * size
	end := start + size - 1
	// 3.ZRevRange 查询 按照分数的从大到校查询指定数据
	return client.ZRevRange(key, start, end).Result()
}

// GetPostIDsInOrder 在redis中查询ids
func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// 从redis获取id
	// 1.根据用户i请求携带的order，确定要查询的redis的key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	// 2.根据用户查询的索引起始点
	return getIDsFormKey(key, p.Page, p.Size)
}

// getPostVoteData 根据ids查找每篇帖子的投票数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	//data = make([]int64, 0, len(ids))
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVotedZSetPrefix + id)
	//	v := client.ZCount(key, "1", "1").Val()
	//	data = append(data, v)
	//}
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPrefix + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}

// GetCommunityPostIDsInOrder 按社区查找ids
func GetCommunityPostIDsInOrder(p *models.ParamCommunityPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}
	// 使用zinterstore 把分区的帖子set与帖子分数的zset 生成一个新的zset，针对新的zset，按照之前的逻辑取数据

	//社区的key
	cKey := getRedisKey(KeyCommunitySetPrefix + strconv.Itoa(int(p.CommunityID)))

	//利用缓存key减少zinterstore执行的次数
	key := orderKey + strconv.Itoa(int(p.CommunityID))
	if client.Exists(key).Val() < 1 {
		// 不存在，需要计算
		pipeline := client.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, cKey, orderKey)
		pipeline.Expire(key, 60*time.Second)
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	return getIDsFormKey(key, p.Page, p.Size)
}
