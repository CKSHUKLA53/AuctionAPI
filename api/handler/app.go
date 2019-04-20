package handler

import (
	"AuctionAPI/api/config"
	"AuctionAPI/api/middleware"
	"AuctionAPI/pkg/model"
	"AuctionAPI/pkg/repository"
	"AuctionAPI/pkg/service"
	"fmt"
	"github.com/urfave/negroni"
	"log"
	"net/http"

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
	clientRepo := repository.NewClientRepository(a.DB)
	clientService := service.NewClientService(clientRepo)

	authMiddleware := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)

	apiMiddleware := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.HandlerFunc(middleware.JwtMiddleware(config.Secret)),
		negroni.HandlerFunc(middleware.LoginMiddleware(clientService)),
		negroni.NewLogger(),
	)

	CreateOfferHandlers(a.Router, *apiMiddleware, *offerService)
	CreateBidHandlers(a.Router, *apiMiddleware, *bidService, *offerService)
	CreateClientHandlers(a.Router, *authMiddleware, *clientService)

	http.Handle("/", a.Router)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
