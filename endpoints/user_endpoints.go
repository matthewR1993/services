// User endpoints

package endpoints

import (
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/matthewR1993/services/valid"
	mod "github.com/matthewR1993/services/models"
	db "github.com/matthewR1993/services/database"
	"github.com/matthewR1993/services/crypto"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
)

type User struct {
	Passw string `json:"passw" validate:"required,max=24,min=10"`
	Email string `json:"email" validate:"required,email"`
}

/*
  New user registration endpoint
  http POST:
  {
	"passw": "qwerty",
	"email": "qwert@mail.com"
  }
*/
var RegisterNewUser = Endpoint {
	"/api/v1/user/register",
	"POST",
	"<New user registration api>",
	func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		dec := json.NewDecoder(r.Body)
	    	var user User
		if err := dec.Decode(&user); err != nil {
			log.Warning(err)
			http.Error(w, err.Error(), 500)
			return
		}

		if err := valid.Validate.Struct(&user); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		var user_info mod.UserInfo
		if db.DBCon.Where(&mod.UserInfo{Email: user.Email}).First(&user_info).RecordNotFound() {
			new_u := mod.UserInfo{}
			new_u.Email = user.Email
			new_u.PasswordHash = crypto.GeneratePDKDF2key([]byte(user.Passw))
			if err := db.DBCon.Create(&new_u).Error; err != nil {
				log.Warning(err)
				http.Error(w, err.Error(), 500)
				return
			}
			w.WriteHeader(201)
			return
		}  else {
			http.Error(w, "User with such email already exists", 409)
			return
		}

		log.Warning("Unexpected error, fucked up")
		http.Error(w, "Unexpected error - Server fucked up", 500)
	},
	[]Middleware{},
}

/*
  Authorisation endpoint
  http POST:
  {
	"passw": "qwerty",
	"email": "qwert@mail.com"
  }
*/
var GenerateAuthToken = Endpoint {
	"/api/v1/user/get_token",
	"POST",
	"<Get token>",
	func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		dec := json.NewDecoder(r.Body)
	    	var user User
		if err := dec.Decode(&user); err != nil {
			log.Warning(err)
			http.Error(w, err.Error(), 500)
			return
		}

		if err := valid.Validate.Struct(&user); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		var user_info mod.UserInfo
		if db.DBCon.Where("email = ?", user.Email).First(&user_info).RecordNotFound() {
			http.Error(w, "User with such email hasn't been found", 404)
			return
		}

		// take key from database
		key := user_info.PasswordHash
		key_split := strings.Split(string(key), "$")
		db_hash := strings.Trim(key_split[3], " ")
		db_salt := key_split[2]
		
		// generate hash using delivered password and exisitng salt
		user_hash := crypto.GetHash([]byte(user.Passw), []byte(db_salt))

		// compare hashes
		if crypto.CompareHashes([]byte(db_hash), user_hash) {
			// generate token
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"loggedInAs": "user",
				"email": user.Email,
				"nbf": time.Now().Add(time.Second * 1).Unix(),
				"exp": time.Now().Add(time.Second * 3600 * 24).Unix(), // 1 day
				"aud": "mobile",
				"iat": time.Now().Unix(),
				"sub": "identification",
			})

			tokenString, err := token.SignedString([]byte(crypto.HmacSecretKey))
			if err != nil {
				log.Warn(err)
				http.Error(w, err.Error(), 500)
				return
			}

			// send token back to user
			w.Write([]byte(fmt.Sprintf(`{"token": "%v"}`, tokenString)))

			w.Header().Set("Content-Type", "application/json")
			return
		} else {
			http.Error(w, "Invalid password, jerk.", 401)
			return
		}

		log.Warning("Unexpected error, fucked up")
		http.Error(w, "Unexpected error - Server fucked up", 500)
	},
	[]Middleware{},
}

/*
  Example with jwt middleware
  http GET: /api/v1/user/history?startdate=30.01.1995&enddate=30.02.1995
*/
var GetUserInformation = Endpoint {
	"/api/v1/user/history",
	"GET",
	"<Get user history>",
	func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Write([]byte(ps.ByName("email")))
	},
	[]Middleware{VerifyUserToken},
}

