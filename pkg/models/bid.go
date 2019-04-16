package model

import (
	"github.com/jinzhu/gorm"
)

type Bid struct {
	Id       ID      `gorm:"primary_key";"AUTO_INCREMENT"`
	BidPrice float64 `json:"bid_price"`
	OfferId  ID
	Client   Client `gorm:"foreignkey:ClientId"`
	ClientId ID
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Bid{})
	db.AutoMigrate(&Offer{})
	db.AutoMigrate(&Client{})
	return db
}
