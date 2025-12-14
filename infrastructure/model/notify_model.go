package model

type NotifyModel struct {
	ID         string `gorm:"column:notify_id;primaryKey;autoIncrement:false"`
	RecieverID string `gorm:"column:receiver_id"`
	SenderID   string `gorm:"column:sender_id"`
	Message    string `gorm:"column:message"`
	Status     string `gorm:"column:status"`
	TimeStamp  string `gorm:"column:time_stamp"`
}

func (NotifyModel) TableName() string {
	return "notify_models"
}
