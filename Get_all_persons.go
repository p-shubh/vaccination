package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get_all_persons(c *gin.Context) {

	users := []Getallnurses_persons{}

	sqlStatement := `SELECT persons.first_name,persons.last_name,persons.date_of_birth,persons.sex,persons.email FROM persons`

	rows, err := DB.Query(sqlStatement)

	if err != nil {
		log.Println("Failed to execute query in get return request : ", err)
		return
	}

	defer rows.Close()

	user := Getallnurses_persons{}
	for rows.Next() {
		rows.Scan(&user.First_name, &user.Last_name, &user.Date_of_birth, &user.Sex, &user.Email)
		users = append(users, user)
	}

	res := gin.H{
		"nurses_directory": users,
	}

	c.JSON(http.StatusOK, res)

}
