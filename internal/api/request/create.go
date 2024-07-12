package request

import (
	"time"

	"github.com/lavrovanton/notifications/internal/model"
)

type CreateNotification struct {
	SenderEmail   string `json:"sender_email" binding:"required,max=255,email"`
	ReceiverEmail string `json:"receiver_email" binding:"required,max=255,email"`
	Text          string `json:"text" binding:"required,max=1024"`
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
