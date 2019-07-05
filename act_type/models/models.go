package models

type User struct {
	Id int `gorm:"id"`
	Login string `gorm:"login"`
	Password string `gorm:"password"`
	Balance int `gorm:"balance"`
}

type Log struct {
	Id int `gorm:"id"`
	User string `gorm:"user"`
	Status string `gorm:"status"`
	Date string `gorm:"date"`
	Amount int `gorm:"amount"`
	UpdatedBalance int `gorm:"updated_balance"`
}

type Goods struct {
	Id int `gorm:"id"`
	Name string `gorm:"name"`
	Price int `gorm:"price"`
}

type Request1 struct {
	Insert `json:"type"`
}

type Insert struct {
	Login string `json:"login"`
	Password string `json:"password"`
	Sum int `json:"sum"`
}

type RequestForBuy struct {
	Request2 `json:"type"`
}

type Request2 struct {
	Login string `json:"login"`
	Password string `json:"password"`
	Purchase []GoodsPurch `json:"order"`
}

type GoodsPurch struct {
	Name string `json:"name"`
	Amount int `json:"amount"`
}
