package models

import (
	"time"

	"github.com/aichelnokov/apiwalk/internal/storage/mysql/types"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id				uint64			`gorm:"primaryKey;autoIncrement:true"`
	Name			string			`gorm:"size:255;index:idx_name,unique"`
	Password	string			`gorm:"size:64"`
	CreatedAt	time.Time		
	UpdatedAt	time.Time
	DeletedAt *time.Time		
	Walks			[]Walk			`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Walk struct {
	gorm.Model
	Id				uint64		`gorm:"primaryKey;autoIncrement:true"`
	UserId		uint64
	Coords		types.Point
	Altitude	float64
	CreatedAt	time.Time
}