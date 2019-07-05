package routes

import (
	models2 "SQL/actType/models"
	"SQL/actType/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)
var database *gorm.DB
func Init(db *gorm.DB) {
	database = db

	r := gin.Default()

	r.POST("/insertSum", insertSum)
	r.POST("order", order)

	r.Run(":8000")
}

func insertSum(c *gin.Context)  {
	resp := models2.Request1{}

	err :=  c.ShouldBindJSON(&resp)
	if err != nil {
		log.Println("Couldn't bind json file", err.Error())
		return
	}
	users := []models2.User{}
	database.Find(&users)



	var findUser bool
	var balance int
	hash :=  utils.GetMD5Hash(&resp.Login, &resp.Password)
	log.Println(hash)
	for _, row :=range users {
		if row.Login == resp.Login && row.Password == hash{
			findUser = true
			balance = row.Balance + resp.Sum
			database.Table("users").Model(&row).Update("balance", balance)
		}
	}
	if !findUser {
		c.JSON(http.StatusBadRequest,"Couldn't find user:" + resp.Login)
		return
	}

	time1 := time.Now().Format("02.01.2006 15:04")
	log := models2.Log{User : resp.Login, Status: "Sum insert(+)", Date: time1, Amount: resp.Sum, UpdatedBalance: balance}
	database.Create(&log)

	c.JSON(http.StatusOK,"Account: " + resp.Login + "Balance: " + strconv.Itoa(balance))
}

func order(c *gin.Context) {
	resp := models2.RequestForBuy{}

	err :=  c.ShouldBindJSON(&resp)
	if err != nil {
		log.Println("Couldn't bind json file", err.Error())
		return
	}
	users := []models2.User{}
	database.Find(&users)
	sum := 0

	var rightPass bool
 	var user models2.User

	for _, row := range users {
		if row.Login == resp.Login && row.Password == resp.Password {
			rightPass = true

			goods := []models2.Goods{}
			database.Find(&goods)
			for _, merch := range resp.Purchase {
				for _, good := range goods {
					if good.Name == merch.Name{
						sum += merch.Amount * good.Price
					}
				}

			}
			if row.Balance < sum {
				c.JSON(http.StatusBadRequest, "Not enough sum:")
				return
			}
			user = row
		}
	}

	if !rightPass {
		c.JSON(http.StatusBadRequest, "invalid credentials")
		return
	}
	balance := user.Balance - sum
	database.Table("users").Model(&user).Update("balance", balance)


	time1 := time.Now().Format("02.01.2006 15:04")
	log := models2.Log{User : resp.Login, Status: "Sum withdraw(-)", Date: time1, Amount: sum, UpdatedBalance: balance}
	database.Create(&log)

	c.JSON(http.StatusOK, "Status: Success, " + "Remaining balance: " + strconv.Itoa(balance))
}
