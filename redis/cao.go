package redis

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

// Cao 缓存访问结构
type Cao struct {
	client *Client
}

// NewCao 创建新的缓存访问对象
func NewCao(client *Client) *Cao {
	return &Cao{client}
}

// Get 获取redis缓存
func (c *Cao) Get(key string) (string, error) {
	return c.client.Get(key).Result()
}

// SetByTTL 设置 带有有效时间
func (c *Cao) SetByTTL(key string, value string, extime int64) error {
	return c.client.Set(key, value, time.Duration(extime)*time.Duration(time.Second)).Err()
}

// Set 设置
func (c *Cao) Set(key string, value string) error {
	return c.client.Set(key, value, time.Duration(-1)*time.Second).Err()
}

// Keys Keys
func (c *Cao) Keys(pattern string) ([]string, error) {
	return c.client.Keys(pattern).Result()
}

// Scan Scan
func (c *Cao) Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	return c.client.Scan(cursor, match, count).Result()
}

// ScanAll ScanAll
func (c *Cao) ScanAll(match string, count int64) ([]string, error) {
	res := make([]string, 0)
	m := make(map[string]struct{})
	var start uint64
	flag := true
	for flag {
		keys, cursor, err := c.client.Scan(start, match, count).Result()
		if err != nil {
			return res, err
		}
		start = cursor
		for i := range keys[:] {
			if _, ok := m[keys[i]]; ok {
				continue
			}
			res = append(res, keys[i])
			m[keys[i]] = struct{}{}
		}
		if cursor == 0 {
			flag = false
		}
	}
	return res, nil
}

// SetNX SetNX
func (c *Cao) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return c.client.SetNX(key, value, expiration).Result()
}

// Expire 过期
func (c *Cao) Expire(key string, extime int64) error {
	return c.client.Expire(key, time.Duration(extime)*time.Second).Err()
}

// GetTTL 过期
func (c *Cao) GetTTL(key string) (int64, error) {
	ttl, err := c.client.TTL(key).Result()

	return int64(ttl / time.Second), err
}

// ExpireAt 设置过期时间
func (c *Cao) ExpireAt(key string, ex time.Time) error {
	return c.client.ExpireAt(key, ex).Err()
}

// HSet Hset
func (c *Cao) HSet(key string, field string, value interface{}) (bool, error) {
	return c.client.HSet(key, field, value).Result()
}

// HSetJSON HSet for json
func (c *Cao) HSetJSON(key, field string, value interface{}) (bool, error) {
	btr, err := json.Marshal(value)
	if err != nil {
		return false, err
	}
	return c.client.HSet(key, field, string(btr)).Result()
}

// HMSet HMSet
func (c *Cao) HMSet(key string, fields map[string]interface{}) error {
	return c.client.HMSet(key, fields).Err()
}

// SetTTL SetTTL
func (c *Cao) SetTTL(key string, extime int64) error {
	return c.client.Expire(key, time.Duration(extime)*time.Duration(time.Second)).Err()
}

// Del Del
func (c *Cao) Del(key string) error {
	return c.client.Del(key).Err()
}

// Dels Del
func (c *Cao) Dels(key ...string) error {
	return c.client.Del(key...).Err()
}

// DelPipelined 批量删除
func (c *Cao) DelPipelined(keys []string) error {
	_, err := c.client.Pipelined(func(pipe redis.Pipeliner) error {
		for _, key := range keys {
			err := pipe.Del(key).Err()
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

// DelPipeline 批量删除
func (c *Cao) DelPipeline(keys []string) error {
	pipe := c.client.Pipeline()
	for _, key := range keys {
		err := pipe.Del(key).Err()
		if err != nil {
			return err
		}
	}
	_, err := pipe.Exec()
	return err
}

// HDel Hdel
func (c *Cao) HDel(key string, fields string) error {
	return c.client.HDel(key, fields).Err()
}

// HGet Hget
func (c *Cao) HGet(key string, field string) (string, error) {
	return c.client.HGet(key, field).Result()
}

// HGetJSON HGetJSON
func (c *Cao) HGetJSON(key, field string, T interface{}) error {
	str, err := c.client.HGet(key, field).Result()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(str), T)
	if err != nil {
		return err
	}
	return nil
}

// GetPipeline 获取keys数组的 批量获取数据
func (c *Cao) GetPipeline(keys []string, resultMap map[string]interface{}) ([]string, error) {
	unCacheKeys := make([]string, len(keys)) // 缓存未命中的key
	_, err := c.client.Pipelined(func(pipe redis.Pipeliner) error {
		for _, key := range keys {
			var value interface{}
			str, err := pipe.Get(key).Result()
			if err != nil && err.Error() != "redis: nil" {
				return err
			}
			if str == "" {
				// 没获取到数据
				unCacheKeys = append(unCacheKeys, key)
				continue
			}
			err = json.Unmarshal([]byte(str), value)
			if err != nil {
				return err
			}
			resultMap[key] = value
			// resultArr = append(resultArr, value)
		}
		return nil
	})
	return unCacheKeys, err
}

// GetPipelined test
func (c *Cao) GetPipelined(keys []string, resultMap map[string]interface{}) ([]string, error) {
	var unCacheKeys []string // 缓存未命中的key
	pipe := c.client.Pipeline()
	for _, key := range keys {
		_, _ = pipe.Get(key).Result()
	}
	resultsCmd, _ := pipe.Exec()
	var err error
	for _, cmd := range resultsCmd {
		var value interface{}
		key := ""
		args := cmd.Args()
		if v, ok := args[1].(string); ok {
			key = v
			if cmd.Err() != nil && cmd.Err().Error() == "redis: nil" {
				// 说明没查询到数据
				unCacheKeys = append(unCacheKeys, key)
				continue
			}
			if cmd.Err() != nil {
				err = cmd.Err()
				continue
			}
			if cmd, ok := cmd.(*redis.StringCmd); ok {
				str := cmd.Val() // 上面判断过是否获取到数据了，这里应该不会为nil
				s, _ := json.Marshal(str)
				err = json.Unmarshal(s, &value)
				if err != nil {
					return nil, err
				}
				resultMap[key] = value
			}
		}
	}
	return unCacheKeys, err
}

// SetPipeline 设置多个数据，根据map key作为redis key， value作为redis value
func (c *Cao) SetPipeline(values map[string]interface{}, extime int64) error {
	_, err := c.client.Pipelined(func(pipe redis.Pipeliner) error {
		for key, value := range values {
			err := pipe.Set(key, value, time.Duration(extime)*time.Second).Err()
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

// HGetAll 获取全部
func (c *Cao) HGetAll(key string) map[string]interface{} {
	return ConvertStringToMap(c.client.HGetAll(key).Val())
}

// ConvertStringToMap 转换从redis获取的数据
func ConvertStringToMap(base map[string]string) map[string]interface{} {
	resultMap := make(map[string]interface{})
	for k, v := range base {
		var dat map[string]interface{}
		if err := json.Unmarshal([]byte(v), &dat); err == nil {
			resultMap[k] = dat
		} else {
			resultMap[k] = v
		}
	}
	return resultMap
}

// 三种 发布/订阅
// 1. 标准redis pub/sub
// 2. list 模拟
// 3. 以上两种组合使用

// Subscribe 订阅
func (c *Cao) Subscribe(ch ...string) *redis.PubSub {
	return c.client.Subscribe(ch...)
}

// PSubscribe 订阅 正则
func (c *Cao) PSubscribe(ch ...string) *redis.PubSub {
	return c.client.PSubscribe(ch...)
}

// Publish 发布
func (c *Cao) Publish(ch string, message interface{}) (int64, error) {
	return c.client.Publish(ch, message).Result()
}

// BRPop timeout:0 阻塞
// 当订阅来用
func (c *Cao) BRPop(timeout time.Duration, key ...string) ([]string, error) {
	return c.client.BRPop(timeout, key...).Result()
}

// BLPop blpop
func (c *Cao) BLPop(timeout time.Duration, key ...string) ([]string, error) {
	return c.client.BLPop(timeout, key...).Result()
}

// LPush 当发布来用 从左边插入多个数据进入队列
func (c *Cao) LPush(key string, values interface{}) (int64, error) {
	return c.client.LPush(key, values).Result()
}

// RPush 当发布来用 从右边插入多个数据进入队列
func (c *Cao) RPush(key string, values interface{}) (int64, error) {
	return c.client.LPush(key, values).Result()
}

// LPop 从列表 key 左边删除一个数据并返回
func (c *Cao) LPop(key string, values interface{}) (string, error) {
	return c.client.LPop(key).Result()
}

// RPop 从列表 key 右边边删除一个数据并返回
func (c *Cao) RPop(key string, values interface{}) (string, error) {
	return c.client.RPop(key).Result()
}

// LRange 指定指定范围 [start，end] 列表 key的 元素(包含start和end的元素)
func (c *Cao) LRange(key string, start, end int64) ([]string, error) {
	return c.client.LRange(key, start, end).Result()
}

// LTrim 删除指定范围的 list数据 [start，end](包含start和end的元素)
func (c *Cao) LTrim(key string, start, end int64) (string, error) {
	return c.client.LTrim(key, start, end).Result()
}

// IncrBy  命令将 key 中储存的数字加上指定的增量值
func (c *Cao) IncrBy(key string, value int64) (int64, error) {
	return c.client.IncrBy(key, value).Result()
}

// HIncrBy 命令用于为哈希表中的字段值加上指定增量值
func (c *Cao) HIncrBy(key, field string, incr int64) (int64, error) {
	return c.client.HIncrBy(key, field, incr).Result()
}

// Watch watch
func (c *Cao) Watch(fn func(tx *redis.Tx) error, keys ...string) error {
	return c.client.Watch(fn, keys...)
}

/* bitmap相关 */

// SetBit 设置bitmap
func (c *Cao) SetBit(key string, offset int64, value int) error {
	return c.client.SetBit(key, offset, value).Err()
}

// // SetBitWithExpire 设置 bitmap 带有过期时间
// func (c *Cao) SetBitWithExpire(key string, offset int64, value int,duration string) error {
// 	// if err:= c.client.SetBit(key, offset, value).Err();err!=nil{
// 	// 	return err
// 	// }
// 	// c.client.Expire(key,c.)
// }

// GetBit get
func (c *Cao) GetBit(key string, offset int64) (int64, error) {
	return c.client.GetBit(key, offset).Result()
}

// ZAdd 添加一个或者多个成员
func (c *Cao) ZAdd(key string, values []map[string]interface{}) (int64, error) {
	zs := make([]redis.Z, 0, len(values))
	for _, v := range values {
		z := redis.Z{
			Score:  v["score"].(float64),
			Member: v["member"],
		}
		zs = append(zs, z)
	}

	return c.client.ZAdd(key, zs...).Result()
}

// ZRange 递增获取数据
func (c *Cao) ZRange(key string, start, stop int64) ([]string, error) {
	return c.client.ZRange(key, start, stop).Result()
}

// ZRank 正序获取 member 在集合中的索引(排名)
func (c *Cao) ZRank(key, member string) (int64, error) {
	return c.client.ZRank(key, member).Result()
}

// ZRangeWithScores 递增获取数据和 score
func (c *Cao) ZRangeWithScores(key string, start, stop int64) ([]redis.Z, error) {
	return c.client.ZRangeWithScores(key, start, stop).Result()
}

// ZRem 原子操作，删除数据
func (c *Cao) ZRem(key string, members ...interface{}) (int64, error) {
	return c.client.ZRem(key, members...).Result()
}

// ZRevRange 递减获取数据
func (c *Cao) ZRevRange(key string, start, stop int64) ([]string, error) {
	return c.client.ZRevRange(key, start, stop).Result()
}

// ZRevRangeWithScores 递减获取数据和 score
func (c *Cao) ZRevRangeWithScores(key string, start, stop int64) ([]redis.Z, error) {
	return c.client.ZRevRangeWithScores(key, start, stop).Result()
}

// ZRevRangeByScoreWithScores 根据分数取值(从小到大)
func (c *Cao) ZRevRangeByScoreWithScores(key string, opt redis.ZRangeBy) ([]redis.Z, error) {
	return c.client.ZRevRangeByScoreWithScores(key, opt).Result()
}

// ZRangeByScoreWithScores 根据分数取值(从大到小)
func (c *Cao) ZRangeByScoreWithScores(key string, opt redis.ZRangeBy) ([]redis.Z, error) {
	return c.client.ZRevRangeByScoreWithScores(key, opt).Result()
}

// ZPopMin 从小到大弹出
func (c *Cao) ZPopMin(key string, count int64) ([]map[string]interface{}, error) {
	list := make([]map[string]interface{}, 0)
	zlist, err := c.client.ZPopMin(key, count).Result()
	if err != nil {
		return list, err
	}
	for i := range zlist[:] {
		m := zlist[i].Member
		switch v := m.(type) {
		case string:
			data := map[string]interface{}{
				"score":  zlist[i].Score,
				"member": v,
			}
			list = append(list, data)
		}
	}
	return list, nil
}

// ZCard 获取指定key的数量
func (c *Cao) ZCard(key string) (int64, error) {
	return c.client.ZCard(key).Result()
}

// ZScore 返回有序集 key 中，成员 member 的 score 值
func (c *Cao) ZScore(key, member string) (float64, error) {
	return c.client.ZScore(key, member).Result()
}

// TxHGetAll 事物批量获取hash
func (c *Cao) TxHGetAll(keys []string) (map[string]map[string]string, error) {
	res := make(map[string]map[string]string)
	command := make(map[string]*redis.StringStringMapCmd)
	_, err := c.client.TxPipelined(func(pipe redis.Pipeliner) error {
		for i := range keys[:] {
			command[keys[i]] = pipe.HGetAll(keys[i])
		}
		return nil
	})
	if err != nil {
		return res, err
	}
	for i, v := range command {
		value, _ := v.Result()
		res[i] = value
	}
	return res, nil
}

// TxHSetAll hash 设置 fields 附带 key 的过期时间
func (c *Cao) TxHSetAll(key string, fieldMap map[string]interface{}, expire time.Duration) error {
	if expire == 0 {
		return c.client.HMSet(key, fieldMap).Err()
	}
	_, err := c.client.TxPipelined(func(pipe redis.Pipeliner) error {
		pipe.HMSet(key, fieldMap)
		pipe.Expire(key, expire)
		return nil
	})
	return err
}

// SIsMember SIsMember
func (c *Cao) SIsMember(key string, fields string) (bool, error) {
	return c.client.SIsMember(key, fields).Result()
}

// SAdd SAdd
func (c *Cao) SAdd(key string, field ...interface{}) (int64, error) {
	return c.client.SAdd(key, field...).Result()
}

// GetSet GetSet
func (c *Cao) GetSet(key string, value interface{}) (string, error) {
	return c.client.GetSet(key, value).Result()
}

// PipeHMGet hash 批量hmget
func (c *Cao) PipeHMGet(keys []string, field ...string) (map[string]map[string]interface{}, error) {
	data := make(map[string]map[string]interface{})
	k := make(map[string]*redis.SliceCmd)
	_, err := c.client.TxPipelined(func(pipe redis.Pipeliner) error {
		for i := range keys[:] {
			k[keys[i]] = pipe.HMGet(keys[i], field...)
		}
		return nil
	})
	if err != nil && err != redis.Nil {
		return data, err
	}
	for i, v := range k {
		m := make(map[string]interface{})
		r, err := v.Result()
		if err != nil && err != redis.Nil {
			return data, err
		}
		if len(field) != len(r) {
			return data, fmt.Errorf("类型错误%v", v)
		}
		for i := range field[:] {
			m[field[i]] = r[i]
		}
		data[i] = m
	}
	return data, nil
}

// HMGet hash hmget
func (c *Cao) HMGet(keys string, field ...string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	res, err := c.client.HMGet(keys, field...).Result()
	if err != nil {
		return data, err
	}
	if len(field) != len(res) {
		return data, fmt.Errorf("类型错误")
	}
	for i := range field[:] {
		if res[i] == nil {
			continue
		}
		data[field[i]] = res[i]
	}
	return data, nil
}

// Lock 锁
func (c *Cao) Lock(key string, isLock bool, t time.Duration) (bool, error) {
	if isLock {
		return c.client.SetNX(key, true, t).Result()
	}
	_, err := c.client.Del(key).Result()

	return true, err
}

// AcqLock key 键 isLock 是否加锁 bizKeep 业务时间戳 eg: 10min，expire 过期时间
func (c *Cao) AcqLock(key string, isLock bool, bizKeep, expire time.Duration) (bool, error) {
	n := time.Now()
	if !isLock {
		_, err := c.client.Del(key).Result()
		return true, err
	}
	lockCmd := new(redis.BoolCmd)
	infoCmd := new(redis.StringCmd)
	_, err := c.client.TxPipelined(func(pipe redis.Pipeliner) error {
		lockCmd = pipe.SetNX(key, n.Unix(), expire)
		infoCmd = pipe.Get(key)
		return nil
	})
	l, err := lockCmd.Result()
	if err != nil {
		return l, err
	}
	if l {
		return l, err
	}
	info, err := infoCmd.Result()
	if err != nil {
		return false, err
	}
	u, _ := strconv.Atoi(info)
	if n.Add(0-bizKeep).Unix() < int64(u) {
		return false, nil
	}
	c.client.Set(key, n.Unix(), expire)
	return true, nil
}

// HsetNx 不存在则设置 hash
func (c *Cao) HsetNx(key string, field, value string) (bool, error) {
	return c.client.HSetNX(key, field, value).Result()
}

// ScriptRun 运行 redis lua 脚本
func (c *Cao) ScriptRun(script *redis.Script, keys []string, args ...interface{}) (interface{}, error) {
	return script.Run(c.client, keys, args...).Result()
}

/* HyperLogLog 相关 */

// PFADD
func (c *Cao) PFADD(key, value string) error {
	_, err := c.client.PFAdd(key, value).Result()
	return err
}

// PFCOUNT
func (c *Cao) PFCOUNT(key string) (int64, error) {
	return c.client.PFCount(key).Result()
}
