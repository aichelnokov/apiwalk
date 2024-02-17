package mysql

import (
	"fmt"

	"github.com/aichelnokov/apiwalk/internal/config"
	"github.com/aichelnokov/apiwalk/internal/storage/mysql/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New(DBConfig config.DBConfig) (*Storage, error) {
	const op = "storage.mysql.NewStorage"
	const dsnString = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf(dsnString, DBConfig.Username, DBConfig.Password, DBConfig.Host, DBConfig.Port, DBConfig.Database),
		DefaultStringSize: 256,
		// DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
  	// DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
  	// DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	db.AutoMigrate(&models.User{}, &models.Walk{})

	// h := md5.Sum([]byte("123qwe!"))
	// db.Create(&models.User{
	// 	Name:      "aichelnokov@gmail.com",
	// 	Password:  fmt.Sprintf("%x", h),
	// })

	// db.Create(&models.User{
	// 	Name:      "aichelnokov@gmail.com",
	// 	Password:  fmt.Sprintf("%x", h),
	// 	CreatedAt: time.Time{},
	// 	UpdatedAt: time.Time{},
	// 	Walks:     []models.Walk{{Coords: types.Point{X:54.494258, Y:36.207713}, Altitude: 200.000000}},
	// })

	var user models.User
	// db.First(&user)
	// db.Preload("Walks").First(&user, "users.name = ?", "aichelnokov@gmail.com")
	// db.Model(&models.User{}).Preload("Walks").First(&user, "users.name = ?", "aichelnokov@gmail.com")
	db.Model(&models.User{}).Joins("Walks").First(&user, "users.name = ?", "aichelnokov@gmail.com")

	// db.Joins("RIGHT JOIN walks ON walks.user_id=users.id").Find(&user, "users.name = ?", "aichelnokov@gmail.com")
	// fmt.Printf("%d %s %s %s", user.Id, user.Name, user.Password, user.CreatedAt.Format(constants.ISO8601))
	fmt.Printf("Walk Id:%d X:%f Y:%f Altitude:%f", user.Walks[0].Id, user.Walks[0].Coords.X, user.Walks[0].Coords.Y, user.Walks[0].Altitude)
	// fmt.Printf("Walk Id:%d Altitude:%f", user.Walks[0].Id, user.Walks[0].Altitude)


	return &Storage{db: db}, nil
}