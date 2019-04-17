package handler

import (
	"AuctionAPI/pkg/model"
	"AuctionAPI/pkg/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/urfave/negroni"
	"net/http"
	"strconv"
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

func acceptBid(bidService service.BidService, offerService service.OfferService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		bidID, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}
		_bid, err := bidService.Update(bidID, "accepted", true)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error Placing Bid"))
			return
		}

		_, err = offerService.Update(_bid.OfferId, "sold", true)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error Placing Bid"))
			return
		}

		if err := json.NewEncoder(w).Encode(_bid); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Error"))
		}
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)
	})
}

func placeBid(bidService service.BidService, offerService service.OfferService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var _bid *model.Bid
		errorMessage := "Error occured while Placing a Bid"

		err := json.NewDecoder(r.Body).Decode(&_bid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error occured while Placing a Bid"))
			return
		}
		client := _bid.Client
		offer, err := offerService.Find(_bid.OfferId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error occured while Placing a Bid"))
			return
		}

		if offer.BidPrice >= _bid.BidPrice {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error occured while Placing a Bid. BidPrice < Previous Bid Price"))
			return
		}

		offer, err = offerService.Update(_bid.OfferId, "bid_price", _bid.BidPrice)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error occured in Placing a Bid"))
			return
		}

		_bid.ClientId = client.Id
		_bid, err = bidService.Store(_bid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error occured in Placing a Bid"))
			return
		}

		if err := json.NewEncoder(w).Encode(_bid); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)

	})
}
func CreateBidHandlers(r *mux.Router, n negroni.Negroni, bidService service.BidService, offerService service.OfferService) {
	r.Handle("/bids", n.With(
		negroni.Wrap(placeBid(bidService, offerService)),
	)).Methods("POST", "OPTIONS").Name("placeBid")

	r.Handle("/bids/{id}", n.With(
		negroni.Wrap(acceptBid(bidService, offerService)),
	)).Methods("PUT", "OPTIONS").Name("acceptBid")
}
