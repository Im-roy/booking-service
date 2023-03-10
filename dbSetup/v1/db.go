package v1

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDbClient() (*gorm.DB, error) {
	dsn := "postgres://byxbjyos:URvuR_vyb6iDTl-IHVNRAIiQjgjvnThX@arjuna.db.elephantsql.com/byxbjyos"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, err
}

type TUser struct {
	*gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

func CheckTableCreation(db *gorm.DB) {
	db.AutoMigrate(&TUser{})
	tUser := TUser{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	result := db.Create(&tUser)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println(tUser.ID)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
}
