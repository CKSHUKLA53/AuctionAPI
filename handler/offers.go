package handler

import (
	"AuctionAPI/pkg/model"
	"AuctionAPI/pkg/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

func createOffer(service service.OfferService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var offer *model.Offer
		errorMessage := "Error occured while creating an offer"
		err := json.NewDecoder(r.Body).Decode(&offer)
		if err != nil {
			respondError(w, http.StatusInternalServerError, errorMessage)
			return
		}

		// check if offer data is valid else return error
		if !offer.Validate() {
			respondError(w, http.StatusBadRequest, "Bad Data error")
			return
		}
		offer, err = service.Store(offer)

		if err != nil {
			respondError(w, http.StatusInternalServerError, errorMessage)
			return
		}

		if err := json.NewEncoder(w).Encode(offer); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)

	})
}

func getOffer(service service.OfferService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error occured while fetching offers"
		var offers []model.Offer
		offers, err := service.FindAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		w.WriteHeader(http.StatusAccepted)
		if err := json.NewEncoder(w).Encode(offers); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

//CreateUserHandlers Maps routes to http handlers
func CreateOfferHandlers(r *mux.Router, n negroni.Negroni, service service.OfferService) {
	r.Handle("/offer", n.With(
		negroni.Wrap(createOffer(service)),
	)).Methods("POST", "OPTIONS").Name("CreateOffer")

	r.Handle("/offer", n.With(
		negroni.Wrap(getOffer(service)),
	)).Methods("GET", "OPTIONS").Name("GetOffers")
}
