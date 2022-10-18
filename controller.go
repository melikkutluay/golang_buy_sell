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

type buyRequest struct {
	Kaynak_id    int    `json:"Kaynak_id"`
	Hedef_id     int    `json:"Hedef_id"`
	Kaynak_hesap string `json:"Kaynak_hesap"`
	Hedef_hesap  string `json:"Hedef_hesap"`
	Miktar       int    `json:"Miktar"`
	User_id      string `json:"User_id"`
}

type loginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int    `json:"expiresIn"`
}

type Exchange struct {
	Exchange_id   int    `json:"exchange_id"`
	Exchange_name string `json:"exchange_name"`
	Exchange_rate int    `json:"exchange_rate"`
}

type Account struct {
	//Account_id   int    `json:"account_id"`
	//Account_name string `json:"account_name"`
	Balance int `json:"balance"`
	//User_id      int    `json:"User_id"`
}

type Buy struct {
	kaynak_id    int    `json:"kaynak_id"`
	hedef_id     int    `json:"kaynak_id"`
	kaynak_hesap string `json:"kaynak_hesap"`
	hedef_hesap  string `json:"kaynak_hesap"`
	miktar       int    `json:"exchange_rate"`
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
	rq := &buyRequest{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(rq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("server: could not read request body: %s\n", err)
	}

	fmt.Printf("server: request body: %s\n", reqBody)
	ss := fmt.Sprintf("select balance from accounts where user_id = '%s' and account_id = %d", rq.User_id, rq.Kaynak_id)
	fmt.Println(ss)
	getMoney := getQeury(ss)
	fmt.Println(getMoney)
	rows := makingBuy(ss)
	var AccountList []Account
	for rows.Next() {
		var account Account
		if err := rows.Scan(&account.Balance); err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		AccountList = append(AccountList, account)
	}
	rsJs, _ := json.Marshal(AccountList)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(rsJs)

}

func (c *controller) makeSell(w http.ResponseWriter, r *http.Request) {

}

func getExchangeRate() {
	var exc Exchange
	for range time.Tick(time.Second * 10) {
		rows := getQeury("SELECT * FROM exchange")
		for rows.Next() {
			if err := rows.Scan(&exc.Exchange_id, &exc.Exchange_name, &exc.Exchange_rate); err != nil {
				log.Fatal(err)
			}
			fmt.Println(exc)
		}
	}
}
