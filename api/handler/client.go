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
		errorMessage := "Error sigining up user"
		var client *model.Client

		err := json.NewDecoder(r.Body).Decode(&client)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		data, err := service.FindByUsername(client.ClientName)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		if data != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("User Already Exist"))
			return
		}
		//usr.Password = user.SaltPassowrd(usr.Password)
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
		errorMessage := "Error reading user"
		var usr *model.Client
		err := json.NewDecoder(r.Body).Decode(&usr)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := service.FindByUsername(usr.ClientName)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("User Doesn't Exist"))
			return
		}

		userDatum := data[0]

		if userDatum.Password != usr.Password {
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
		jwtmap := userDatum.GenerateJWT([]byte(cfg.Secret))
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

	r.Handle("/signup", n.With(
		negroni.Wrap(signup(service)),
	)).Methods("POST", "OPTIONS").Name("signup")
}
