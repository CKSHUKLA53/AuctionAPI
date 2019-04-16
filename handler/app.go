package handler

import (
	"DemoGO/pkg/models"
	"DemoGO/pkg/repository"
	"DemoGO/pkg/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.Database.Username,
		config.Database.Password,
		config.Database.Name,
		config.Database.Charset)

	db, err := gorm.Open(config.Database.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()

	bidRepo := repository.NewBidRepository(a.DB)
	bidService := service.NewBidService(bidRepo)
	offersRepo := repository.NewOffersRepository(a.DB)
	offerService := service.NewOfferService(offersRepo)

	a.setBidsRouters(bidService, offerService)
	a.setOffersRouters(offerService)

}

func (a *App) setOffersRouters(ser *service.OfferService) {

	// Routing for handling the projects
	a.Router.HandleFunc("/offers", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		offers, err := ser.FindAll()
		if err != nil {
			respondError(w, http.StatusInternalServerError, "Error occured")
		}
		respondJSON(w, http.StatusOK, offers)
	})).Methods("GET")
	a.Router.HandleFunc("/offers", a.CreateOffer).Methods("POST")
}

func (a *App) setBidsRouters(ser *service.BidService) {
	// Routing for handling the projects
	a.Router.HandleFunc("/bid", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bids, err := ser.FindAll()
		if err != nil {
			respondError(w, http.StatusInternalServerError, "Error occured")
		}
		respondJSON(w, http.StatusOK, bids)
	})).Methods("GET")
	a.Router.HandleFunc("/bid", a.CreateOffer).Methods("POST")

	a.Router.HandleFunc("/placeBid")
}

/*func (a *App) GetAllOffers(ser *service.Service) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	});
}
*/
func (a *App) CreateOffer(w http.ResponseWriter, r *http.Request) {
	CreateOffer(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func placeBid(bidService *service.BidService, offerService *service.OfferService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var _bid *model.Bid
		errorMessage := "Error occured while Placing a Bid"
		client := r.Context().Value("me").(*model.Client)
		err := json.NewDecoder(r.Body).Decode(&_bid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error occured while Placing a Bid"))
			return
		}

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

		offer, err = offerService.Update(_bid.OfferID, "bidprice", _bid.BidPrice)
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
