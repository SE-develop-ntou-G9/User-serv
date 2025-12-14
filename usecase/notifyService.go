package usecase

import (
	"golangAPI/entity"
)

type NotifyUsecase struct {
	repo NotifyRepository
}

func NewNotifyUsecase(r NotifyRepository) *NotifyUsecase {
	return &NotifyUsecase{repo: r}
}

func (uc *NotifyUsecase) CreateNotification(notify entity.Notify) (*entity.Notify, error) {
	return uc.repo.Create(&notify)
}

func (uc *NotifyUsecase) GetNotificationByRecieverID(recieverID string) ([]entity.Notify, error) {
	return uc.repo.GetByRecieverID(recieverID)
}

func (uc *NotifyUsecase) DeleteNotificationByRecieverID(id string) error { //一鍵清除
	return uc.repo.DeleteByRecieverID(id)
}

func (uc *NotifyUsecase) DeleteNotificationByID(id string) error {
	return uc.repo.DeleteByID(id)
}
