package handler

import (
	"AuctionAPI/api/config"
	"AuctionAPI/pkg/model"
	"AuctionAPI/pkg/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func signup(service service.ClientService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error in registering client"
		var client *model.Client

		err := json.NewDecoder(r.Body).Decode(&client)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		data, err := service.FindByClientName(client.ClientName)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		if data != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Client Already Exist"))
			return
		}

		client.Id, err = service.Store(client)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		w.WriteHeader(http.StatusCreated)

	})
}

func login(service service.ClientService) http.Handler {
	cfg := config.GetConfig()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading client"
		var usr *model.Client
		err := json.NewDecoder(r.Body).Decode(&usr)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := service.FindByClientName(usr.ClientName)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Client Doesn't Exist"))
			return
		}

		clientData := data[0]

		if clientData.Password != usr.Password {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Password Doesnot Match"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != model.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		jwtmap := clientData.GenerateJWT([]byte(cfg.Secret))
		if err := json.NewEncoder(w).Encode(jwtmap); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		w.WriteHeader(http.StatusAccepted)
	})

}

//CreateUserHandlers Maps routes to http handlers
func CreateClientHandlers(r *mux.Router, n negroni.Negroni, service service.ClientService) {
	r.Handle("/login", n.With(
		negroni.Wrap(login(service)),
	)).Methods("POST", "OPTIONS").Name("login")

	r.Handle("/register", n.With(
		negroni.Wrap(signup(service)),
	)).Methods("POST", "OPTIONS").Name("signup")
}
