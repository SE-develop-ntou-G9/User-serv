package repositories

import (
	"golangAPI/entity"
	Model "golangAPI/infrastructure/model"

	"github.com/google/uuid"

	"gorm.io/gorm"

	"time"
)

type notifyRepository struct {
	db *gorm.DB
}

func NewNotifyRepository(db *gorm.DB) *notifyRepository {
	return &notifyRepository{db: db}
}

func (r *notifyRepository) Create(notify *entity.Notify) (*entity.Notify, error) {
	m := Model.NotifyModel{
		ID:         uuid.New().String(),
		RecieverID: notify.RecieverID,
		SenderID:   notify.SenderID,
		Message:    notify.Message,
		Status:     notify.Status,
		TimeStamp:  time.Now().UTC().String(),
	}
	err := r.db.Create(&m).Error
	if err != nil {
		return nil, err
	}
	notify.ID = m.ID
	notify.TimeStamp = m.TimeStamp
	return notify, nil
}

func (r *notifyRepository) GetByRecieverID(recieverID string) ([]entity.Notify, error) {
	var models []Model.NotifyModel
	err := r.db.Where("receiver_id = ?", recieverID).Find(&models).Error
	if err != nil {
		return nil, err
	}
	var notifies []entity.Notify
	for _, m := range models {
		notifies = append(notifies, entity.Notify{
			ID:         m.ID,
			RecieverID: m.RecieverID,
			SenderID:   m.SenderID,
			Message:    m.Message,
			Status:     m.Status,
			TimeStamp:  m.TimeStamp,
		})
	}
	return notifies, nil
}

func (r *notifyRepository) DeleteByRecieverID(id string) error {
	err := r.db.Where("receiver_id = ?", id).Delete(&Model.NotifyModel{}).Error
	return err
}

func (r *notifyRepository) DeleteByID(id string) error {
	err := r.db.Where("notify_id = ?", id).Delete(&Model.NotifyModel{}).Error
	return err
}
