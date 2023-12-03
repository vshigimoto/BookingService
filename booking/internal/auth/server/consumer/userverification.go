package consumer

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"

	"booking/internal/auth/server/consumer/dto"
)

type UserVerificationCallback struct {
}

func NewUserCallback() *UserVerificationCallback {
	return &UserVerificationCallback{}
}

func (c *UserVerificationCallback) Callback(message <-chan *sarama.ConsumerMessage, error <-chan *sarama.ConsumerError) {
	for {
		select {
		case msg := <-message:
			var userCode dto.UserCode

			err := json.Unmarshal(msg.Value, &userCode)
			if err != nil {
				//nolint:all
				fmt.Errorf("failed to unmarshall record value err: %v", err)
			} else {
				fmt.Printf("user code: %s", userCode)

			}
		case err := <-error:
			//nolint:all
			fmt.Errorf("failed consume err: %v", err)
		}
	}
}
