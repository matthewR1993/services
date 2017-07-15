package endpoints

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/dgrijalva/jwt-go"
	"github.com/matthewR1993/services/crypto"
)

/*
  Middleware for users identification through jwt.
  JWT is expected to be sent as HTTP header like this: "Authorization": "JWT.Here.yougo"
*/ 
var VerifyUserToken Middleware = func (h Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {	
		tokenString := r.Header.Get("Authorization")

		// Check whether a token is valid. (1 day expiration time)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(crypto.HmacSecretKey), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ps = append(ps, httprouter.Param{"email", claims["email"].(string)}) // pass user information further
		} else {
			http.Error(w, err.Error(), 401)
			return
		}
		
		h(w, r, ps)
	}
}

/*
  CORS standart middleware
*/
var CORSMiddle = func (h Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		h(w, r, ps)
	}
}

/*
  Middleware for tests...
*/
var RequestConsoleLog Middleware = func (h Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Println(r)

		h(w, r, ps)
	}
}
