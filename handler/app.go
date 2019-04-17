package handler

import (
	"AuctionAPI/middleware"
	"AuctionAPI/pkg/model"
	"AuctionAPI/pkg/repository"
	"AuctionAPI/pkg/service"
	"fmt"
	"github.com/urfave/negroni"
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

	apiMiddleware := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)

	CreateOfferHandlers(a.Router, *apiMiddleware, *offerService)
	CreateBidHandlers(a.Router, *apiMiddleware, *bidService, *offerService)

	//a.setBidsRouters(bidService, offerService)
	//a.setOffersRouters(offerService)

}

/*func (a *App) setOffersRouters(ser *service.OfferService) {

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

func (a *App) setBidsRouters(bidService *service.BidService, offerService *service.OfferService) {
	// Routing for handling the projects
	a.Router.HandleFunc("/bid", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bids, err := bidService.FindAll()
		if err != nil {
			respondError(w, http.StatusInternalServerError, "Error occured")
		}
		respondJSON(w, http.StatusOK, bids)
	})).Methods("GET")
	a.Router.HandleFunc("/bid", a.CreateOffer).Methods("POST")

	a.Router.HandleFunc("/placeBid", placeBid(bidService, offerService))
}*/

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
