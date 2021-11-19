package producer

import (
	"hash/crc32"
	"sync"

	"github.com/Shopify/sarama"
)

// hashPartitioner 分区计算器
type hashPartitioner struct {
	buf     map[string]int32 // 暂存区
	bufLock sync.Mutex
}

func newHashPartitioner(topic string) *hashPartitioner {
	return &hashPartitioner{
		buf: map[string]int32{},
	}
}

// get 获取分区
func (p *hashPartitioner) get(key []byte, numPartitions int32) int32 {
	if v, ok := p.buf[string(key)]; ok {
		return v
	}
	pt := int32(crc32.ChecksumIEEE(key) % uint32(numPartitions))
	p.set(string(key), pt)
	return pt
}

// set 设置缓存
func (p *hashPartitioner) set(key string, partition int32) {
	p.bufLock.Lock()
	defer p.bufLock.Unlock()
	p.buf[key] = partition
}

// NewHashPartitioner 获取partitioner
func NewHashPartitioner(topic string) sarama.Partitioner {
	return newHashPartitioner(topic)
}

// Partition [接口实现]分区逻辑
func (p *hashPartitioner) Partition(message *sarama.ProducerMessage, numPartitions int32) (int32, error) {
	kb, err := message.Key.Encode()
	if err != nil {
		return -1, err
	}
	return p.get(kb, numPartitions), nil
}

// RequiresConsistency [接口实现]是否保证一致性
func (p *hashPartitioner) RequiresConsistency() bool {
	return false
}

// Get 获取分区，调试方法
func (p *hashPartitioner) Get(key []byte, numPartitions int32) int32 {
	return p.get(key, numPartitions)
}

// GetAllCache 获取所有缓存
func (p *hashPartitioner) GetAllCache() map[string]int32 {
	res := map[string]int32{}
	for key, value := range p.buf {
		res[key] = value
	}
	return res
}

// FlushAllCache 清理所有缓存（慎用）
func (p *hashPartitioner) FlushAllCache() {
	p.bufLock.Lock()
	defer p.bufLock.Unlock()
	p.buf = map[string]int32{}
}
