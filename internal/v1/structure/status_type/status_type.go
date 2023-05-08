package status_type

import (
	"time"
)

type Table struct {
	//編號
	StID string `gorm:"primary_key;column:st_id;uuid_generate_v4()type:UUID;" json:"st_id,omitempty"`
	//創建時間
	CreateTime time.Time `gorm:"column:create_time;type:TIMESTAMP;" json:"create_time"`
	//狀態
	Status string `gorm:"column:status;type:text;" json:"status,omitempty"`
}

func (a *Table) TableName() string {
	return "status_type"
}
