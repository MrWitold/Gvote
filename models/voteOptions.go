package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type VoteOptions struct {
	gorm.Model
	VoteID   uint
	Option   string
	IDinVote string
	VoteFor  int
}
