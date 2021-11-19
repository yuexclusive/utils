package producer

import (
	"encoding/json"
	"math"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/yuexclusive/utils/logger"
)

var log = logger.Single()

// Producer 生产者
type Producer struct {
	Addrs    []string
	Config   *sarama.Config
	MaxBatch int
	Topic    string
	stack    []string
	slock    sync.Mutex
	producer sarama.SyncProducer
}

// defaultProducerConfig 默认配置
func defaultProducerConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Version = sarama.V1_0_0_0
	return config
}

// defaultProducerConfigWithDSPartitioner 默认配置
func defaultProducerConfigWithDSPartitioner() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = NewHashPartitioner
	config.Version = sarama.V1_0_0_0
	return config
}

// NewProducer 创建生产者
func NewProducer(addrs []string, topic string) (*Producer, error) {
	c := defaultProducerConfig()
	producer, err := sarama.NewSyncProducer(addrs, c)
	if err != nil {
		return nil, err
	}
	return &Producer{
		Addrs:    addrs,
		Config:   c,
		MaxBatch: MaxBatch,
		Topic:    topic,
		stack:    make([]string, 0, MaxBatch),
		producer: producer,
	}, nil
}

// NewProducerWithDSPartitioner 创建生产者
func NewProducerWithDSPartitioner(addrs []string, topic string, pnumber int32) (*Producer, error) {
	c := defaultProducerConfigWithDSPartitioner()
	producer, err := sarama.NewSyncProducer(addrs, c)
	if err != nil {
		return nil, err
	}
	return &Producer{
		Addrs:    addrs,
		Config:   c,
		MaxBatch: MaxBatch,
		Topic:    topic,
		stack:    make([]string, 0, MaxBatch),
		producer: producer,
	}, nil
}

// bigMessageProducerConfig 大消息体结构数据
func bigMessageProducerConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Producer.MaxMessageBytes = 10000000
	config.Version = sarama.V1_0_0_0
	return config
}

// NewBigMessageProducer 创建生产者
// 注意这个方法试用余消息体比较大的场景
func NewBigMessageProducer(addrs []string, topic string) (*Producer, error) {
	c := bigMessageProducerConfig()
	producer, err := sarama.NewSyncProducer(addrs, c)
	if err != nil {
		return nil, err
	}
	return &Producer{
		Addrs:    addrs,
		Config:   c,
		MaxBatch: MaxBatch,
		Topic:    topic,
		stack:    make([]string, 0, MaxBatch),
		producer: producer,
	}, nil
}

// ProduceOneString 单个生产
func (p *Producer) ProduceOneString(data string) (int32, int64, error) {
	msg := &sarama.ProducerMessage{
		Topic: p.Topic,
		Value: sarama.ByteEncoder(data),
	}

	part, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		// TODO: 这里一旦失败就会丢失数据
		return -1, -1, err
	}
	return part, offset, nil
}

// ProduceOneData 生产数据
func (p *Producer) ProduceOneData(data []byte) (int32, int64, error) {
	msg := &sarama.ProducerMessage{
		Topic: p.Topic,
		Value: sarama.ByteEncoder(data),
	}

	part, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		// TODO: 这里一旦失败就会丢失数据
		return -1, -1, err
	}
	return part, offset, nil
}

// ProduceManyString 批量生产
func (p *Producer) ProduceManyString(list []string) error {
	size := len(list)
	f := float64(size) / float64(p.MaxBatch)
	num := int(math.Ceil(f))
	start := 0
	end := p.MaxBatch
	for i := 0; i < num; i++ {
		if i == num-1 {
			end = size
		}
		sub := list[start:end]
		if err := p.produceMany(sub); err != nil {
			// TODO: 这里一旦失败就会丢失数据
			log.Error(err.Error())
		}
		start += p.MaxBatch
		end += p.MaxBatch
	}
	return nil
}

// ProduceOneJSON 单个生产
func (p *Producer) ProduceOneJSON(data interface{}) (int32, int64, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return -1, -1, err
	}
	return p.ProduceOneString(string(bytes))
}

// ProduceManyJSON 批量生产
func (p *Producer) ProduceManyJSON(list []interface{}) error {
	strList := make([]string, len(list))
	for index, v := range list {
		bytes, err := json.Marshal(v)
		if err != nil {
			// 任何一条数据转换错误直接返回失败
			return err
		}
		strList[index] = string(bytes)
	}
	return p.ProduceManyString(strList)
}

// produceMany 批量生产内部方法
func (p *Producer) produceMany(list []string) error {
	msgs := make([]*sarama.ProducerMessage, 0, len(list))
	for _, data := range list {
		msg := &sarama.ProducerMessage{
			Topic: p.Topic,
			Value: sarama.ByteEncoder(data),
		}
		msgs = append(msgs, msg)
	}

	if err := p.producer.SendMessages(msgs); err != nil {
		return err
	}
	return nil
}

// Message 消息结构
type Message struct {
	Headers []sarama.RecordHeader
	Value   string
	Key     sarama.ByteEncoder
}

// NewMessage 创建message
func NewMessage(headers map[string]string, value string) *Message {
	h := make([]sarama.RecordHeader, len(headers))
	index := 0
	for key, value := range headers {
		h[index] = sarama.RecordHeader{
			Key:   []byte(key),
			Value: []byte(value),
		}
		index++
	}
	return &Message{
		Headers: h,
		Value:   value,
	}
}

// NewMessageV2 创建message
func NewMessageV2(headers [][2][]byte, value string) *Message {
	h := make([]sarama.RecordHeader, len(headers))
	for index, value := range headers {
		h[index] = sarama.RecordHeader{
			Key:   value[0],
			Value: value[1],
		}
	}
	return &Message{
		Headers: h,
		Value:   value,
	}
}

// NewMessageWithKey 创建message
func NewMessageWithKey(headers map[string]string, value string, key []byte) *Message {
	h := make([]sarama.RecordHeader, len(headers))
	index := 0
	for key, value := range headers {
		h[index] = sarama.RecordHeader{
			Key:   []byte(key),
			Value: []byte(value),
		}
		index++
	}
	return &Message{
		Headers: h,
		Value:   value,
		Key:     key,
	}
}

// ProduceOneMessage 生产一条记录
func (p *Producer) ProduceOneMessage(data *Message) (int32, int64, error) {
	msg := &sarama.ProducerMessage{
		Headers: data.Headers,
		Topic:   p.Topic,
		Value:   sarama.ByteEncoder(data.Value),
		Key:     data.Key,
	}

	part, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		// TODO: 这里一旦失败就会丢失数据
		return -1, -1, err
	}
	return part, offset, nil
}

func (p *Producer) produceManyMessage(list []*Message) error {
	msgs := make([]*sarama.ProducerMessage, len(list))
	for i, data := range list {
		msg := &sarama.ProducerMessage{
			Headers: data.Headers,
			Topic:   p.Topic,
			Value:   sarama.ByteEncoder(data.Value),
			Key:     data.Key,
		}
		msgs[i] = msg
	}

	if err := p.producer.SendMessages(msgs); err != nil {
		return err
	}
	return nil
}

// ProduceManyMessage 生产多条消息
func (p *Producer) ProduceManyMessage(list []*Message) error {
	size := len(list)
	f := float64(size) / float64(p.MaxBatch)
	num := int(math.Ceil(f))
	start := 0
	end := p.MaxBatch
	for i := 0; i < num; i++ {
		if i == num-1 {
			end = size
		}
		sub := list[start:end]
		if err := p.produceManyMessage(sub); err != nil {
			// TODO: 这里一旦失败就会丢失数据
			log.Error(err.Error())
		}
		start += p.MaxBatch
		end += p.MaxBatch
	}
	return nil
}

// Close 注销生产者
func (p *Producer) Close() error {
	err := p.producer.Close()
	p = nil
	return err
}
