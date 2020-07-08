package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//VoteAllowed struct conrains all permited emails to vote
type VoteAllowed struct {
	gorm.Model
	VoteID      uint
	Email       string
	AccessToken string
	Status      int
}
