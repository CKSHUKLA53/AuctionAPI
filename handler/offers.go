package handler

import (
	model "DemoGO/pkg/models"
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
)

/*
func GetAllOffers(db *gorm.DB, w http.ResponseWriter, r *http.Request, service service.Service) {
	offers := []model.Offer{}
	db.Find(&offers)
	respondJSON(w, http.StatusOK, offers)
}
*/
func CreateOffer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	offer := model.Offer{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&offer); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&offer).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, offer)
}
