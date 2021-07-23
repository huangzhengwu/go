package redis

const (
	KeyPrefix              = "bluebell:"
	KeyPostTimeZSet        = "post:time"   // zset： 帖子以及发帖时间
	KeyPostScoreZSet       = "post:score"  // zset:  帖子以及投票分数
	KeyPostVotedZSetPrefix = "post:voted:" // zset:  记录用户以及投票的类型,参数是post_id
	KeyCommunitySetPrefix  = "community:"  //set:保存每个分下的帖子id
)

// getRedisKey 给key加前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}
