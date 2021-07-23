package redis

import "web_app/models"

// GetPostIDsInOrder 在redis中查询ids
func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// 从redis获取id
	// 1.根据用户i请求携带的order，确定要查询的redis的key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	// 2.根据用户查询的索引起始点
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	// 3.ZRevRange 查询 按照分数的从大到校查询指定数据
	return client.ZRevRange(key, start, end).Result()
}
