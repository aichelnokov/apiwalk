package mysql

import (
	"fmt"

	"github.com/aichelnokov/apiwalk/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New(db_config config.DBConfig) (*Storage, error) {
	const op = "storage.mysql.NewStorage"

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp=(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_config.Username, db_config.Password, db_config.Host, db_config.Port, db_config.Database),
		DefaultStringSize: 256,
		// DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
  	// DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
  	// DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}