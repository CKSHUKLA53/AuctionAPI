package repository

import (
	"AuctionAPI/pkg/model"
	"github.com/jinzhu/gorm"
)

type ClientRepository struct {
	DB *gorm.DB
}

func NewClientRepository(app *gorm.DB) *ClientRepository {
	return &ClientRepository{
		DB: app,
	}
}

func (app *ClientRepository) Find(id int) (*model.Client, error) {
	client := model.Client{Id: id}
	err := app.DB.Find(&client)

	if err.RecordNotFound() {
		return nil, model.ErrNotFound
	} else if err.Error != nil {
		return nil, err.Error
	} else {
		return &client, nil
	}
}

func (r *ClientRepository) FindByKey(key string, val interface{}) ([]*model.Client, error) {
	var result []*model.Client
	err := r.DB.Where(key+"= ?", val).Find(&result)

	if err.RowsAffected == 0 {
		return nil, nil
	} else if err == nil {
		return result, nil
	} else {
		return result, nil
	}
}

func (app *ClientRepository) Store(b *model.Client) (int, error) {

	err := app.DB.Save(&b)
	if err.Error != nil {
		return int(0), err.Error
	}
	return b.Id, nil
}
