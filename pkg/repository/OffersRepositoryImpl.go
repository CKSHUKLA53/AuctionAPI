package repository

import (
	model "DemoGO/pkg/models"
	"github.com/jinzhu/gorm"
)

type OffersRepository struct {
	DB *gorm.DB
}

func NewOffersRepository(app *gorm.DB) *OffersRepository {
	return &OffersRepository{
		DB: app,
	}
}

func (app *OffersRepository) Find(id model.ID) (*model.Offer, error) {
	bid := model.Offer{Id: id}
	err := app.DB.Find(&bid)

	if err.RecordNotFound() {
		return nil, model.ErrNotFound
	} else if err == nil {
		return &bid, nil
	} else {
		return nil, err.Error
	}
}

func (app *OffersRepository) Store(b *model.Offer) (*model.Offer, error) {

	err := app.DB.Save(&b)
	if err != nil {
		return nil, err.Error
	}
	return b, nil
}

func (app *OffersRepository) FindAll() ([]model.Offer, error) {
	offers := []model.Offer{}
	result := app.DB.Preload("Bid").Preload("Bid.Client").Find(&offers)
	if result.Error != nil {
		return nil, result.Error
	}
	return offers, nil
}

func (app *OffersRepository) Delete(id model.ID) error {

	bid := model.Offer{Id: id}
	result := app.DB.Delete(&bid)
	if result.Error != nil {
		return result.Error
	}
	return nil
}