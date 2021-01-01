package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type student struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null";unique`
}

var db *gorm.DB

func main() {
	dsn := "host=localhost user=postgres password=root dbname=users port=5432 sslmode=disable"

	newLogger := logger.New(
		log.New(os.Stdout, "\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			Colorful:      true,
			LogLevel:      logger.Info,
		},
	)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	conn, _ := db.DB()
	defer conn.Close()

	if err != nil {
		log.Fatal(err)
	}

	//createTable()
	//InsertData()
	//SearchData()
	//SearchAll()
	//SearchWhere()
	//Update()
	UpdateWhere()
	DeleteData()
	SearchAll()
}

func createTable() {

	db.Migrator().DropTable(&student{})
	db.Migrator().AutoMigrate(&student{})

}

func InsertData() {

	s := []student{
		{
			Name:  "diwakar",
			Email: "diwakar@email.com",
		},
		{
			Name:  "Raj",
			Email: "raj@email.com",
		},
		{
			Name:  "dev",
			Email: "dev@email.com",
		},
	}

	err := db.Create(&s).Error
	if err != nil {
		fmt.Println(err)
	}

}

func SearchData() {

	var s student
	err := db.First(&s).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)

}

func SearchAll() {
	var s []student

	db.Find(&s)
	fmt.Println(s)

}

func SearchWhere() {

	var s student

	err := db.Where("name = ?", "dev").First(&s).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)

}

func Update() {

	var s student

	db.First(&s)

	s.Name = "abc"
	db.Save(s)

}

func UpdateWhere() {

	db.Model(&student{}).Where("email = ?", "diwakar@email.com").Update("name", "Diwakar")

}

func DeleteData() {
	var s student
	db.First(&s)
	db.Unscoped().Delete(&s)
	//db.Delete(&s)
}
