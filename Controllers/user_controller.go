package Controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

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

func AddUser(c *gin.Context) {

	db := Connect()
	defer db.Close()

	var user User

	if err := c.Bind(&user); err != nil {
		fmt.Println(err)
		return
	}

	// Query
	db.Exec(`INSERT INTO 
	users(Name, Address) 
	VALUES(?, ?)`, user.Name, user.Address)

	// No Check
	c.IndentedJSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {

	db := Connect()
	defer db.Close()

	var user User

	if err := c.Bind(&user); err != nil {
		fmt.Println(err)
		return
	}

	// Query
	result, errQuery := db.Exec(`UPDATE users 
	SET Name=?, Address=? 
	WHERE ID=?`, user.Name, user.Address, user.ID)

	num, _ := result.RowsAffected() //num, err

	// Show Result
	if errQuery == nil {
		if num == 0 {
			c.AbortWithStatusJSON(400, "Update Failed")
			return
		} else {
			c.IndentedJSON(http.StatusOK, user)
		}
	}

}

func DeleteUser(c *gin.Context) {

	db := Connect()
	defer db.Close()

	id := c.Params.ByName("ID")

	// var user User

	// if err := c.Bind(&user); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(user.ID)

	// // Query
	// result, errQuery := db.Exec(`DELETE FROM users
	// WHERE ID=?`, user.ID)

	fmt.Println(id)
	// Query
	result, errQuery := db.Exec(`DELETE FROM users 
	WHERE ID=?`, id)

	num, _ := result.RowsAffected() //num, err

	// Show Result
	if errQuery == nil {
		if num == 0 {
			c.AbortWithStatusJSON(400, "Delete Failed")
			return
		} else {
			c.IndentedJSON(http.StatusOK, id)
			// c.IndentedJSON(http.StatusOK, user)
		}
	}

}
