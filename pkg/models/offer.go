package model

import (
	"time"
)

type Offer struct {
	Id       ID `gorm:"primary_key";"AUTO_INCREMENT"`
	BidPrice float64
	GoLive   time.Time
	LifeTime int
	PhotoUrl string
	Title    string
	Bid      []Bid `gorm:"foreignkey:OfferId"` //you need to do like this
}
