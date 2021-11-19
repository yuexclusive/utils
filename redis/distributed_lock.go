package redis

import (
	"fmt"
	"time"

	redisTx "github.com/go-redis/redis"
	"github.com/google/uuid"
)

// DistributedLock redis 实现的分布式锁
type DistributedLock struct {
	c         *Cao
	prefix    string
	expire    time.Duration
	randValue string // uuid相同才会删除
}

// NewDistributedLock 同一个做，new 随机值
func NewDistributedLock(prefix string, expire time.Duration) *DistributedLock {
	_client := GetClient(MKCache)
	if expire < time.Second { // 不希望出现太小的超时时间
		expire = time.Second
	}
	return &DistributedLock{
		c:         NewCao(_client),
		prefix:    prefix,
		expire:    expire,
		randValue: uuid.New().String(),
	}
}

func (l *DistributedLock) id(lockID string) string {
	return fmt.Sprintf("%s#%s", l.prefix, lockID)
}

// Lock 加锁设置随机值
func (l *DistributedLock) Lock(lockID string) (bool, error) {
	return l.c.SetNX(l.id(lockID), l.randValue, l.expire)
}

// Unlock 解锁-相同randValue才能解锁
func (l *DistributedLock) Unlock(lockID string) error {
	key := l.id(lockID)
	return l.c.Watch(func(tx *redisTx.Tx) error {
		if v, err := tx.Get(key).Result(); err != nil && err != redisTx.Nil {
			return err
		} else if v == l.randValue {
			_, err = tx.Pipelined(func(pipe redisTx.Pipeliner) error {
				pipe.Del(key)
				return nil
			})
			return err
		}
		return nil
	}, key)
}
