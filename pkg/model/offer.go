package model

import (
	"time"
)

type Offer struct {
	Id       int       `gorm:"primary_key";"AUTO_INCREMENT"`
	BidPrice float64   `json:"bid_price"`
	GoLive   time.Time `json:"go_live"`
	LifeTime int       `json:"life_time"`
	PhotoUrl string    `json:"photo_url"`
	Title    string    `json:"title"`
	Sold     bool      `json:"sold"`
	BidId    int       `json:"bid_id"`
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
