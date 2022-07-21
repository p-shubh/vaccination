package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get_all_nurses(c *gin.Context) {

	// reqBody := Getallnurses{}

	// sqlStament := `SELECT nurses.email,persons.first_name,persons.last_name,persons.date_of_birth,persons.sex FROM nurses INNER JOIN persons ON nurses.email = persons.email`

	// row := DB.QueryRow(sqlStament)

	// err := row.Scan(&reqBody.Email, &reqBody.First_name, &reqBody.Last_name, &reqBody.Date_of_birth, &reqBody.Sex)

	// if err != nil {
	// 	res := gin.H{
	// 		"results": err.Error(),
	// 	}
	// 	c.JSON(http.StatusBadRequest, res)
	// } else {
	// 	res := gin.H{
	// 		"results": reqBody,
	// 	}
	// 	c.JSON(http.StatusOK, res)
	// }

	users := []Getallnurses_persons{}

	sqlStatement := `SELECT nurses.email,persons.first_name,persons.last_name,persons.date_of_birth,persons.sex FROM nurses INNER JOIN persons ON nurses.email = persons.email`

	rows, err := DB.Query(sqlStatement)

	if err != nil {
		log.Println("Failed to execute query in get return request : ", err)
		return
	}

	defer rows.Close()
	user := Getallnurses_persons{}
	for rows.Next() {
		rows.Scan(&user.Email, &user.First_name, &user.Last_name, &user.Date_of_birth, &user.Sex)
		users = append(users, user)
	}

	res := gin.H{
		"nurses_directory": users,
	}

	c.JSON(http.StatusOK, res)

}
