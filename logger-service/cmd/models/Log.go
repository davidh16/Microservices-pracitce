package models

import "gorm.io/gorm"

type Log struct {
	ID   int64
	Name string
	Data string
}

func (l *Log) CreateLog(DB *gorm.DB) error {
	result := DB.Create(&l)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
