package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func sendErrorResponse(c echo.Context, message string) error {
	response := ErrorResponse{
		Status:  400,
		Message: message,
	}
	return c.JSON(http.StatusBadRequest, response)

}

func GetAllUsers(c echo.Context) error {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM users"

	name := c.QueryParam("name")
	age := c.QueryParam("age")

	if name != "" {
		fmt.Println(name)
		query += " WHERE name= '" + name + " ', "
	}

	if age != "" {
		if name != "" {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " age= '" + age + "' "
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return sendErrorResponse(c, "Something went wrong, please try again.")
	}

	var user User
	var users []User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.userType); err != nil {
			log.Println(err)
			return sendErrorResponse(c, "Something went wrong, please try again.")
		} else {
			users = append(users, user)
		}
	}

	var response UsersResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = users

	return c.JSON(http.StatusOK, response)
}

func InsertUser(c echo.Context) error {
	db := connect()
	defer db.Close()

	name := c.FormValue("name")
	age, _ := strconv.Atoi(c.FormValue("age"))
	address := c.FormValue("address")
	user_type, _ := strconv.Atoi(c.FormValue("user_type"))

	_, errQuery := db.Exec("INSERT INTO users(name, age, address, user_type) values (?,?,?,?)",
		name,
		age,
		address,
		user_type,
	)

	if errQuery == nil {
		response := UserResponse{
			Status:  200,
			Message: "Success",
			Data: User{
				Name:     name,
				Age:      age,
				Address:  address,
				userType: user_type,
			},
		}
		return c.JSON(http.StatusOK, response)
	} else {
		fmt.Println(errQuery)
		response := ErrorResponse{
			Status:  400,
			Message: "insert Failed!",
		}
		return c.JSON(http.StatusBadRequest, response)
	}
}

func DeleteUser(c echo.Context) error {
	db := connect()
	defer db.Close()

	userId := c.Param("user_id")

	_, err := db.Exec("DELETE FROM users WHERE id=?", userId)
	if err != nil {
		sendErrorResponse(c, "Delete failed")
	}

	response := UserResponse{
		Status:  200,
		Message: "Success",
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateUser(c echo.Context) error {
	db := connect()
	defer db.Close()

	userId := c.Param("user_id")

	name := c.FormValue("name")
	age, _ := strconv.Atoi(c.FormValue("age"))
	address := c.FormValue("address")
	userType, _ := strconv.Atoi(c.FormValue("user_type"))

	_, errQuery := db.Exec("UPDATE users SET name=?, age=?, address=?, user_type=? WHERE id=?",
		name,
		age,
		address,
		userType,
		userId,
	)

	var response UserResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success Update"
	} else {
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "Update Failed"
	}

	return c.JSON(http.StatusOK, response)
}
