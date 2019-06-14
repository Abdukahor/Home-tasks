package main

type Account struct {
	Account string `json:"account"`
	Balance float32 `json:"balance"`
}

type Resp struct {
	Statuscode int
	Data string
}
