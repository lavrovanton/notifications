package handler

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/lavrovanton/notifications/internal/model"
	"github.com/lavrovanton/notifications/internal/rabbitmq"
)

type NotificationRepository interface {
	Store(m *model.Notification) error
}

type CreateNotification struct {
	SenderEmail   string `json:"sender_email" validate:"required,max=255,email"`
	ReceiverEmail string `json:"receiver_email" validate:"required,max=255,email"`
	Text          string `json:"text" validate:"required,max=1024"`
}

func (r CreateNotification) ToModel() model.Notification {
	return model.Notification{
		Id:            0,
		CreatedAt:     time.Time{},
		SenderEmail:   r.SenderEmail,
		ReceiverEmail: r.ReceiverEmail,
		Text:          r.Text,
	}
}

func NewNotificationHandler(repo NotificationRepository) rabbitmq.MessageHandler {
	return func(message []byte) {
		request := CreateNotification{}
		if err := json.Unmarshal(message, &request); err != nil {
			log.Fatal(err)
			return
		}

		validator := validator.New(validator.WithRequiredStructEnabled())
		if err := validator.Struct(&request); err != nil {
			log.Println(err)
		}

		notification := request.ToModel()
		if err := repo.Store(&notification); err != nil {
			log.Println(err)
		}
	}
}
