package model

import (
	"time"
)

type Offer struct {
	Id       int `gorm:"primary_key";"AUTO_INCREMENT"`
	BidPrice float64
	GoLive   time.Time
	LifeTime int
	PhotoUrl string
	Title    string
	Sold     bool
	Bid      []Bid `gorm:"foreignkey:OfferId"` //you need to do like this
}

func (offer *Offer) Validate() bool {
	if offer.BidPrice == 0 || offer.Title == "" || offer.LifeTime < 0 {
		return false
	}
	if offer.GoLive.Before(time.Now()) {
		offer.GoLive = time.Now()
	}
	return true
}
