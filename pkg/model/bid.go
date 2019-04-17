package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Bid struct {
	Id        int `gorm:"primary_key";"AUTO_INCREMENT"`
	BidPrice  float64
	OfferId   int
	Client    Client `gorm:"foreignkey:ClientId"`
	Timestamp time.Time
	Accepted  bool
	ClientId  int
}

func (bid *Bid) Validate() bool {
	if bid.BidPrice <= 0 || bid.OfferId == 0 {
		return false
	}
	bid.Timestamp = time.Now()
	return true
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Bid{})
	db.AutoMigrate(&Offer{})
	db.AutoMigrate(&Client{})
	return db
}
