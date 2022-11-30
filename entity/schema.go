package entity

import (
	"gorm.io/gorm"
)

type State struct {
	gorm.Model
	Name         string
	Registration []Registration `gorm:"ForeignKey: StateID"`
}

type Student struct {
	gorm.Model
	S_ID         string `gorm:"uniqueIndex"`
	Name         string
	Phone        string `gorm:"uniqueIndex"`
	Password     string
	Registration []Registration `gorm:"foreignKey:StudentID"`
}

type Registration struct {
	gorm.Model

	SubjectID *uint
	Subject   Subject `gorm:"references:id"`
	StudentID *uint
	Student   Student `gorm:"references:id"`
	StateID   *uint
	State     State `gorm:"references:id"`
}

type Subject struct {
	gorm.Model
	CODE         string         `gorm:"uniqueIndex"`
	Name         string         `gorm:"uniqueIndex"`
	Registration []Registration `gorm:"foreignKey:SubjectID"`
}
