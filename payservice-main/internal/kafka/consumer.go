package kafka

import (
	"fmt"
	"strings"

	"github.com/IBM/sarama"

	"payservice/internal/auth/config"
)

type ConsumerCallback interface {
	Callback(message <-chan *sarama.ConsumerMessage, error <-chan *sarama.ConsumerError)
	//Test(message <-chan *sarama.ConsumerMessage, error <-chan *sarama.ConsumerError)
}

type Consumer struct {
	topics   []string
	master   sarama.Consumer
	callback ConsumerCallback
}

func NewConsumer(
	cfg config.Kafka,
	callback ConsumerCallback,
) (*Consumer, error) {
	samaraCfg := sarama.NewConfig()
	samaraCfg.ClientID = "go-kafka-consumer"
	samaraCfg.Consumer.Return.Errors = true

	master, err := sarama.NewConsumer(cfg.Brokers, samaraCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create NewConsumer err: %w", err)
	}

	return &Consumer{
		topics:   cfg.Consumer.Topics,
		master:   master,
		callback: callback,
	}, nil
}

func (c *Consumer) Start() {
	consumers := make(chan *sarama.ConsumerMessage, 1)
	errors := make(chan *sarama.ConsumerError)

	for _, topic := range c.topics {
		if strings.Contains(topic, "__consumer_offsets") {
			continue
		}

		partitions, err := c.master.Partitions(topic)
		if err != nil {
			fmt.Printf("Failed to get partitions for topic %s: %v\n", topic, err)
			continue
		}

		consumer, err := c.master.ConsumePartition(topic, partitions[0], sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("Failed to create a consumer for topic %s: %v\n", topic, err)
			continue
		}

		fmt.Println("Start consuming topic", topic)

		go func(topic string, consumer sarama.PartitionConsumer) {
			for {
				select {
				case consumerError := <-consumer.Errors():
					errors <- consumerError

				case msg := <-consumer.Messages():
					consumers <- msg
				}
			}
		}(topic, consumer)
	}

	c.callback.Callback(consumers, errors)
}
