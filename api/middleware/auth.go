package middleware

import (
	"AuctionAPI/pkg/service"
	"context"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/urfave/negroni"
	"net/http"
	"time"
)

func LoginMiddleware(clientService *service.ClientService) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		//Get and Parse Token data
		jwtToken := r.Context().Value("user").(*jwt.Token)
		created := jwtToken.Claims.(jwt.MapClaims)["created"].(float64)
		_id := jwtToken.Claims.(jwt.MapClaims)["clientId"].(float64)
		//Check if token is expired or not

		delta := time.Now().Unix() - int64(created)
		if delta > 3600 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Token Expired"))
			return
		}

		client, err := clientService.Find(int(_id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Authentication Failed"))
			return
		}

		newRequest := r.WithContext(context.WithValue(r.Context(), "me", client))
		next(w, newRequest)
	})

}

//JwtMiddleware
func JwtMiddleware(secret string) negroni.HandlerFunc {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	return jwtMiddleware.HandlerWithNext
}
