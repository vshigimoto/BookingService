package consumer

import (
	"booking/internal/user/server/consumer/dto"
	"encoding/json"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type UserVerificationCallback struct {
	logger *zap.SugaredLogger
}

func NewUserVerificationCallback(logger *zap.SugaredLogger) *UserVerificationCallback {
	return &UserVerificationCallback{logger: logger}
}

func (c *UserVerificationCallback) Callback(message <-chan *sarama.ConsumerMessage, error <-chan *sarama.ConsumerError) {
	for {
		select {
		case msg := <-message:
			var userCode dto.UserCode

			err := json.Unmarshal(msg.Value, &userCode)
			if err != nil {
				c.logger.Errorf("failed to unmarshall record value err: %v", err)
			} else {
				c.logger.Infof("user code: %s", userCode)
			}
		case err := <-error:
			c.logger.Errorf("failed consume err: %v", err)
		}
	}
}
