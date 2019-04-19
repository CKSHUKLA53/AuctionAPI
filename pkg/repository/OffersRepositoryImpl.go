package repository

import (
	"AuctionAPI/pkg/model"
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2"
)

type OffersRepository struct {
	DB *gorm.DB
}

func NewOffersRepository(app *gorm.DB) *OffersRepository {
	return &OffersRepository{
		DB: app,
	}
}

func (app *OffersRepository) Find(id int) (*model.Offer, error) {
	bid := model.Offer{Id: id}
	err := app.DB.Find(&bid)

	if err.RecordNotFound() {
		return nil, model.ErrNotFound
	} else if err.Error == nil {
		return &bid, nil
	} else {
		return nil, err.Error
	}
}

func (app *OffersRepository) Store(b *model.Offer) (*model.Offer, error) {

	err := app.DB.Save(&b)
	if err.Error != nil {
		return nil, err.Error
	}
	return b, nil
}

func (app *OffersRepository) FindAll() ([]model.Offer, error) {
	offers := []model.Offer{}
	//result := app.DB.Preload("Bid").Preload("Bid.Client").Find(&offers)
	result := app.DB.Find(&offers)
	if result.Error != nil {
		return nil, result.Error
	}
	return offers, nil
}

func (app *OffersRepository) Query(page int, size int, sortkey string) ([]*model.Offer, error) {

	if size == 0 {
		size = 10
	}

	if sortkey == "" {
		sortkey = "go_live"
	}

	var res []*model.Offer
	err := app.DB.Find(nil).Order(sortkey).Limit(size).Offset(page).Find(&res)
	if err.Error != nil {

	}
}

func (app *OffersRepository) Delete(id int) error {

	bid := model.Offer{Id: id}
	result := app.DB.Delete(&bid)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (app *OffersRepository) Update(id int, key string, value interface{}) (*model.Offer, error) {

	var offer model.Offer
	if err := app.DB.Where("id = ?", id).First(&offer).Error; err != nil {
		return nil, model.ErrNotFound
	}
	app.DB.Model(&offer).Update(key, value)
	return &offer, nil
}

func (app *OffersRepository) SoldOffers() ([]model.Offer, error) {
	offers := []model.Offer{}
	if err := app.DB.Where("sold = ?", true).Find(&offers).Error; err != nil {
		return nil, model.ErrNotFound
	}
	return offers, nil
}
