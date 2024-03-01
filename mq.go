package config

type Mq struct {
	Driver string `json:"driver" yaml:"driver" desc:"消息队列类型, nsq, kafka, rabbit"`

	// nsq
	ConsumeAddress string `json:"consumeAddress" yaml:"consumeAddress" desc:"nsq消费者链接地址"` // nsq
	ProductAddress string `json:"productAddress" yaml:"productAddress" desc:"nsq生产者链接地址"` // nsq

	// kafka
	Brokers []string `json:"brokers" yaml:"brokers" desc:"kafka brokers"`
}
