package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int    `json:"expiresIn"`
}

type Exchange struct {
	exchange_id   int    `json:"exchange_id"`
	exchange_name string `json:"exchange_name"`
	exchange_rate int    `json:"exchange_rate"`
}

type controller struct {
	keycloak *keycloak
}

func newController(keycloak *keycloak) *controller {
	return &controller{
		keycloak: keycloak,
	}
}

func (c *controller) login(w http.ResponseWriter, r *http.Request) {
	rq := &loginRequest{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(rq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jwt, err := c.keycloak.gocloak.Login(context.Background(),
		c.keycloak.clientId,
		c.keycloak.clientSecret,
		c.keycloak.realm,
		rq.Username,
		rq.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	rs := &loginResponse{
		AccessToken:  jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		ExpiresIn:    jwt.ExpiresIn,
	}

	rsJs, _ := json.Marshal(rs)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(rsJs)
}

func (c *controller) makeBuy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	payload := Exchange{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(payload.exchange_name)
}

func (c *controller) makeSell(w http.ResponseWriter, r *http.Request) {

}

func getExchangeRate() {
	var exc Exchange
	for range time.Tick(time.Second * 1) {
		rows := getQeury("SELECT * FROM exchange")
		for rows.Next() {
			if err := rows.Scan(&exc.exchange_id, &exc.exchange_name, &exc.exchange_rate); err != nil {
				log.Fatal(err)
			}
			fmt.Println(exc)
		}
	}
}
