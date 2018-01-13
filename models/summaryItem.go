package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// SummaryItem the detail info for summary
type SummaryItem struct {
	gorm.Model
	UserID uint
	Data   []byte
	Start  time.Time
	End    time.Time
}

// Create create a record
func (item *SummaryItem) Create() error {
	return engine.Create(item).Error
}

// FindByStarter Find a record by start
func (item *SummaryItem) FindByStarter() (exist bool) {
	var num uint
	var err error
	err = engine.Model(new(SummaryItem)).Where(SummaryItem{Start: item.Start, UserID: item.UserID}).Count(&num).Error
	if err != nil {
		return
	}
	if num == 0 {
		return
	}
	exist = true
	return
}
