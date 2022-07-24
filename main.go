package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User
var DB *gorm.DB

func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}

func initDB() {
	dsn := "root:MieKuah.1123#@tcp(127.0.0.1:3306)/crud_orm?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := User{Id: 1, Name: "Rara", Email: "rara@gmail.com", Password: "jklmno"}
	// render data - JSON response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		return c.JSON(400, "id on url_param is invalid format, should be numeric")
	}
	dataIsFound := false
	for i, row := range users {
		if row.Id == id {
			users = remove(users, i)
			dataIsFound = true
			break
		}
	}
	if dataIsFound {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "delete user with id is successful",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user with id is not found",
	})
}
func remove(slice []User, s int) []User {
	return append(slice[:s], slice[s+1:]...)
}

// update user by id
func UpdateUserController(c echo.Context) error {
	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		return c.JSON(400, "id on url_param is invalid format, should be numeric")
	}
	var updateUser User
	c.Bind(&updateUser)
	dataIsFound := true
	for i, row := range users {
		if row.Id == id {
			// update
			users[i].Name = updateUser.Name
			updateUser.Id = id
			// end update
			dataIsFound = true
			break
		}
	}
	if !dataIsFound == false {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "user with id is not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "update user with id is successful",
		"data":    updateUser,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}
