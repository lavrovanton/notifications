package repository

import (
	"github.com/lavrovanton/notifications/internal/model"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	Db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{db}
}

func (r *NotificationRepository) Fetch() ([]model.Notification, error) {
	var notifications []model.Notification

	result := r.Db.Find(&notifications)
	if result.Error != nil {
		return nil, result.Error
	}
	return notifications, nil
}
func (r *NotificationRepository) Store(m *model.Notification) error {
	result := r.Db.Create(&m)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
