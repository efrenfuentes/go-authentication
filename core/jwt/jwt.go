package jwt

import (
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/efrenfuentes/go-authentication/core/settings"
	"github.com/efrenfuentes/go-authentication/models"
)

type APIToken struct {
	Token string `json:"token"`
}

func CreateToken(user models.User) string {
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	// Create a Token that will be signed with RSA 256.
	/*
	  {
	      "typ":"JWT",
	      "alg":"RS256"
	  }
	*/
	token.Claims["email"] = user.Email
	token.Claims["name"] = user.Name
	token.Claims["exp"] = time.Now().Unix() + 36000

	jwtSettings := settings.Get()["jwt"].(map[string]interface{})

	keyfile := jwtSettings["privateKey"].(string)

	privateKey, _ := ioutil.ReadFile(keyfile)

	// The claims object allows you to store information in the actual token.
	tokenString, _ := token.SignedString(privateKey)

	return tokenString
}

func JWTTokenAuthentication(inner http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtSettings := settings.Get()["jwt"].(map[string]interface{})

		keyfile := jwtSettings["publicKey"].(string)

		publicKey, _ := ioutil.ReadFile(keyfile)

		token, err := jwt.ParseFromRequest(r,
			func(token *jwt.Token) (interface{}, error) {
				return publicKey, nil
			})

		if err == nil && token.Valid {
			inner.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}
