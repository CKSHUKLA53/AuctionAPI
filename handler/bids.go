package handler

import (
	model "../pkg/models"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/http"
)

func getAllBids(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	bids := []model.Bid{}
	db.Find(&bids)
	respondJSON(w, http.StatusOK, bids)
}

func CreateBid(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	bid := model.Bid{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&bid); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&bid).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, bid)
}
