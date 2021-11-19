package consumer

import (
	"context"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/yuexclusive/utils/logger"
)

var log = logger.Single()

// DefaultConfig 默认配置
func DefaultConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V1_0_0_0
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Group.Session.Timeout = 20 * time.Second
	config.Consumer.Group.Heartbeat.Interval = 6 * time.Second
	config.Consumer.MaxProcessingTime = 500 * time.Millisecond
	return config
}

// Consumer 消费者结构
type Consumer struct {
	Addrs   []string
	Config  *sarama.Config
	Topic   string
	Group   string
	Handler sarama.ConsumerGroupHandler
	quit    chan struct{}         // 退出信号，保证平滑关闭
	cg      *sarama.ConsumerGroup // 消费者组，保证平滑关闭
}

// NewConsumer 创建消费者
func NewConsumer(addrs []string, topic string, group string, h sarama.ConsumerGroupHandler) *Consumer {
	return &Consumer{
		Addrs:   addrs,
		Config:  DefaultConfig(),
		Topic:   topic,
		Group:   group,
		Handler: h,
		quit:    make(chan struct{}, 1),
	}
}

// NewConsumerWithPassword 创建消费者
func NewConsumerWithPassword(addrs []string, topic string, group string, h sarama.ConsumerGroupHandler, username string, password string) *Consumer {

	conf := DefaultConfig()
	conf.Net.SASL.User = username
	conf.Net.SASL.Password = password

	return &Consumer{
		Addrs:   addrs,
		Config:  conf,
		Topic:   topic,
		Group:   group,
		Handler: h,
		quit:    make(chan struct{}, 1),
	}
}

// defaultBigMessageConfig 默认配置
func defaultBigMessageConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V1_0_0_0
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Fetch.Max = 11000000
	return config
}

// NewBigMessageConsumer 创建消费者
func NewBigMessageConsumer(addrs []string, topic string, group string, h sarama.ConsumerGroupHandler) *Consumer {
	return &Consumer{
		Addrs:   addrs,
		Config:  defaultBigMessageConfig(),
		Topic:   topic,
		Group:   group,
		Handler: h,
		quit:    make(chan struct{}, 1),
	}
}

// Consume 进行消费
func (c *Consumer) Consume() error {
	if _, err := c.resetcg(); err != nil {
		return err
	}
	for {
		select {
		case <-c.quit:
			return nil
		default:
			if c.cg == nil {
				log.Error("kafka消费异常， 正在重新创建")
				time.Sleep(3 * time.Second)
				if _, err := c.resetcg(); err != nil {
					log.Error(err.Error())
					continue
				}
			}

			ctx := context.Background()
			topics := []string{c.Topic}

			if err := (*c.cg).Consume(ctx, topics, c.Handler); err != nil {
				// 如果异常结束，先sleep 3秒，并注销当前消费者组，并重置消费者组
				time.Sleep(3 * time.Second)
				log.Error(fmt.Sprintf("kafka消费异常: %s，正在执行重连操作... ", err.Error()))
				if err := (*c.cg).Close(); err != nil {
					log.Error(fmt.Sprintf("kafka Consumer.Consume defer Close failed %v", err))
				}
				if _, err := c.resetcg(); err != nil {
					log.Error(err.Error())
				}
			} else {
				// 如果是正常结束，打印警告日志即可
				log.Warn("kafka消费者正常退出，正在重新接入消费")
			}
		}
	}
}

func (c *Consumer) resetcg() (*sarama.ConsumerGroup, error) {
	cg, err := sarama.NewConsumerGroup(c.Addrs, c.Group, c.Config)
	if err != nil {
		return nil, err
	}
	c.cg = &cg
	return c.cg, nil
}

// GracefulShutdown 平滑关闭
func (c *Consumer) GracefulShutdown() {
	log.Info("consumer close signal received")

	// 准备好退出信号
	c.quit <- struct{}{}

	// 关闭consumer
	if c != nil && c.cg != nil {
		err := (*c.cg).Close()
		if err != nil {
			log.Error(err.Error())
		}
	}

}

// GracefulShutdownWithError 平滑关闭
func (c *Consumer) GracefulShutdownWithError() error {
	log.Info("consumer close signal received")

	// 准备好退出信号
	c.quit <- struct{}{}

	// 关闭consumer
	if c == nil || c.cg == nil {
		return fmt.Errorf("client[%p] or consumer group[%p] is nil", c, c.cg)
	}

	if err := (*c.cg).Close(); err != nil {
		return err
	}
	return nil
}

func (c *Consumer) headers(payload sarama.ConsumerMessage) map[string][]byte {
	m := make(map[string][]byte)
	for _, h := range payload.Headers {
		m[string(h.Key)] = h.Value
	}
	return m
}
