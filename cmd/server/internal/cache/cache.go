package cache

import (
	"github.com/go-redis/redis"
)

type RedisCache struct {
	Db *redis.Client
}

func NewRedisCache(Db *redis.Client) *RedisCache {
	return &RedisCache{Db: Db}
}

func (r *RedisCache) SetThumbnail(link string, thumb []byte) error {
	return r.Db.Set(link, thumb, 0).Err()
}

func (r *RedisCache) GetThumbnail(link string) ([]byte, error) {
	res := r.Db.Get(link)
	return res.Bytes()
}
