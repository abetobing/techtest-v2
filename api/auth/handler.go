package auth

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"customer/api/middlewares"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// u := new(LoginRequest)
	var u LoginRequest
	json.NewDecoder(r.Body).Decode(&u)
	res := LoginResponse{}
	id, err := loginUser(middlewares.GetDB(r.Context()), u.User, u.Pass)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = "Invalid user or password"
		b, _ := json.Marshal(res)
		w.Write(b)
		return
	}

	// token := GenerateToken(32)
	// var claims jwt.Claims
	claims := jwt.MapClaims{"user_id": id}
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	jwtauth.SetExpiryIn(claims, 5*time.Minute)
	_, token, _ := tokenAuth.Encode(claims)
	middlewares.GetRedis(r.Context()).Set(token, id, 5*time.Minute)
	res.Token = token
	b, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(res)
	w.Write(b)
}

func GenerateToken(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
