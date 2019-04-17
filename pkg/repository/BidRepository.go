package repository

import (
	"AuctionAPI/pkg/model"
)

//Reader interface
type BidReader interface {
	Find(id model.ID) (model.Bid, error)
	FindAll() ([]*model.Bid, error)
}

//Writer bookmark writer
type BidWriter interface {
	Store(b *model.Bid) (model.Bid, error)
	Update(b *model.Bid) (model.Bid, error)
	Delete(id model.ID) error
}

//Repository repository interface
type BidsRepository interface {
	BidReader
	BidWriter
}
