package repository

import (
	model "DemoGO/pkg/models"
	"github.com/jinzhu/gorm"
)

type BidRepository struct {
	DB     *gorm.DB
}

func NewBidRepository(app *gorm.DB) *BidRepository {
	return &BidRepository{
		DB: app,
	}
}

func (app *BidRepository)  Find(id model.ID) (*model.Bid, error) {
	bid := model.Bid{Id: id}
	err:= app.DB.Find(&bid)

	if err.RecordNotFound() {
		return nil, model.ErrNotFound
	} else if err == nil {
		return &bid, nil;
	} else {
		return nil, err.Error
	}
}


func (app *BidRepository) Store(b *model.Bid) (*model.Bid, error) {

	err := app.DB.Save(&b)
	if err != nil {
		return nil, err.Error
	}
	return b, nil
}

func (app *BidRepository) Update(bidId *model.ID, key, ) (*model.Bid, error) {

	err := Find()
	if err != nil {
		return nil, err.Error
	}
	return b, nil
}

func (app *BidRepository) FindAll() ([]model.Bid, error) {
	bids := []model.Bid{}
	err := app.DB.Preload("Client").Find(&bids)
	if err.Error != nil {
		return nil, err.Error
	}
	return bids, nil
}

func (app *BidRepository) Delete(id model.ID) error {

	bid := model.Bid{Id: id}
	err:= app.DB.Delete(&bid)
	if err != nil {
		return err.Error;
	}
	return nil;
}