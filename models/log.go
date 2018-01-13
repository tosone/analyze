package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/tosone/worklyzer/common/constants"
)

// Log the request info wakatime log
type Log struct {
	gorm.Model
	Url         string
	QueryString string
	Type        constants.RequestType
	Method      constants.RequestMethod
	Body        []byte
	Msg         string
	Status      constants.RequestStatus
	Start       time.Time
	End         time.Time
}

// Create create a record
func (log *Log) Create() error {
	return engine.Create(log).Error
}
