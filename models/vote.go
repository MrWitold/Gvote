package models

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/jinzhu/gorm"
)

// Vote struct contain all informaton about vote ticktets
type Vote struct {
	gorm.Model
	Token       string
	Title       string        `json:"title"`
	StartAt     string        `json:"startat"`
	EndAt       string        `json:"endat"`
	CreatedBy   string        `json:"createdby"`
	VoteOptions []VoteOptions `gorm:"foreignkey:VoteID"`
	VoteAllowed []VoteAllowed `gorm:"foreignkey:VoteID"`
}

// InitialMigration setup sqllite database for project at iniciacion
func InitialMigration() {
	db, err := gorm.Open("sqlite3", "vote.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	db.Debug().AutoMigrate(Vote{})
	db.Debug().AutoMigrate(VoteOptions{})
}

// GetList retrun list of votes
func GetList() []*Vote {
	db, err := gorm.Open("sqlite3", "vote.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var listOfVotes []*Vote

	db.Find(&listOfVotes)

	return listOfVotes
}

// GetSingleItme retrun all information about specific vote
func GetSingleItme(Token string) []*VoteOptions {
	db, err := gorm.Open("sqlite3", "vote.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var votes Vote
	var options []*VoteOptions

	db.Where("Token = ?", Token).Find(&votes)
	db.Model(&votes).Related(&options)

	for _, o := range options {
		fmt.Println(o.Option)
	}

	return options
}

// ToJSON Parese object to json
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(i)
}
