package main

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"net/http"
)

var (
	accountBalance = map[string]float32{}
)

func main() {
	var accounts []Account

	file, err := ioutil.ReadFile("balance.json")

	if err != nil {
		log.Println("Couldn't open file", err.Error())
	}

	err =json.Unmarshal(file, &accounts)

	if err != nil {
		log.Println("Couldn't unmarshal file ", err.Error())
	}

	for _, account := range accounts {
		accountBalance[account.Account] = account.Balance
	}

	http.HandleFunc("/paymentToAccount", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			return
		}

		var incomingJson Account
		var resp Resp
		w.Header().Set("Content-type", "application/json")
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&incomingJson); err != nil {
			log.Println("Cannot decode json: ", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			resp.Statuscode = http.StatusInternalServerError
			resp.Data = "Cannot decode json"
			respByte, _ := json.Marshal(&resp)
			w.Write(respByte)
			return
		}else {
			log.Println("Decode succeed")
			var found bool
			for _, account := range accounts {
				if account.Account == incomingJson.Account {
					found = true
					if incomingJson.Balance <= 0 {
						log.Println("invalid sum")
						resp.Statuscode = http.StatusInternalServerError
						resp.Data = "invalid sum"
						respByte, _ := json.Marshal(&resp)
						w.Write(respByte)

					} else {
						accountBalance[account.Account] += incomingJson.Balance
						account.Balance = accountBalance[account.Account]
						showAcc, err := json.Marshal(account)
						if err != nil {
							log.Println("Couldn't marshal account")
							return
						}
						w.WriteHeader(http.StatusOK)
						w.Write(showAcc)
					}

				}
			}
				if !found {
					log.Println("given account doesn't exist")

					resp.Statuscode = http.StatusNotFound
					resp.Data = "account now found"
					respByte, _ := json.Marshal(&resp)
					w.Write(respByte)
				}

		}
		})

	http.HandleFunc("/payment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			return
		}

		var incomingJson Account
		var resp Resp
		w.Header().Set("Content-type", "application/json")
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&incomingJson); err != nil {
			log.Println("Cannot decode json: ", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			resp.Statuscode = http.StatusInternalServerError
			resp.Data = "Cannot decode json"
			respByte, _ := json.Marshal(&resp)
			w.Write(respByte)
			return
		}else {
			log.Println("Decode succeed")
			var found bool
			for _, account := range accounts {
				if account.Account == incomingJson.Account {
					found = true
					if accountBalance[account.Account] < incomingJson.Balance {
						log.Println("invalid sum")
						resp.Statuscode = http.StatusInternalServerError
						resp.Data = "invalid sum"
						respByte, _ := json.Marshal(&resp)
						w.Write(respByte)

					} else {
						accountBalance[account.Account] -= incomingJson.Balance
						account.Balance = accountBalance[account.Account]
						showAcc, err := json.Marshal(account)
						if err != nil {
							log.Println("Couldn't marshal account")
							return
						}
						w.WriteHeader(http.StatusOK)
						w.Write(showAcc)
					}

				}
			}
			if !found {
				log.Println("given account doesn't exist")

				resp.Statuscode = http.StatusNotFound
				resp.Data = "account now found"
				respByte, _ := json.Marshal(&resp)
				w.Write(respByte)
			}

		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("cannot start server:", err.Error())
		return
	}
}
