package Controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {

	db := Connect()
	defer db.Close()

	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("Err", err.Error())
	}

	var user User
	var users []User

	for results.Next() {
		err = results.Scan(&user.ID, &user.Name, &user.Address)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, users)
	}

}
