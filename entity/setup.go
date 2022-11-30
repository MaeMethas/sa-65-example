package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa-65-regis.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(&Student{},
		&Subject{},
		&State{},
		&Registration{})

	db = database
	phone_1, err := bcrypt.GenerateFromPassword([]byte("0935463156"), 14)
	phone_2, err := bcrypt.GenerateFromPassword([]byte("0969240925"), 14)

	// Student
	studentsa := Student{
		Name:  "B6300001",
		S_ID:  "B6300001",
		Phone: string(phone_1),
	}
	db.Model(&Student{}).Create(&studentsa)

	studentsa2 := Student{
		Name:  "B6332877",
		S_ID:  "B6332877",
		Phone: string(phone_2),
	}
	db.Model(&Student{}).Create(&studentsa2)

	//Subject
	subjectsa := Subject{
		CODE: "IST30 3001",
		Name: "ManEconomy",
	}
	db.Model(&Subject{}).Create(&subjectsa)

	subjectsa1 := Subject{
		CODE: "IST30 3002",
		Name: "ManSocial",
	}
	db.Model(&Subject{}).Create(&subjectsa1)

	//Status
	StateOnline := State{
		Name: "Online",
	}
	db.Model(&State{}).Create(&StateOnline)

	StateOnsite := State{
		Name: "Onsite",
	}
	db.Model(&State{}).Create(&StateOnsite)

	StateHybride := State{
		Name: "Hybride",
	}
	db.Model(&State{}).Create(&StateHybride)
}
