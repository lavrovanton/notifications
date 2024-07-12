package model

import "time"

type Notification struct {
	Id            uint64    `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	SenderEmail   string    `json:"sender_email" gorm:"type:varchar"`
	ReceiverEmail string    `json:"receiver_email" gorm:"type:varchar"`
	Text          string    `json:"text" gorm:"type:text"`
}
