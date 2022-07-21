package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get_all_persons_vaccinated(c *gin.Context) {

	users := []getallpersonsvaccinated{}

	sqlStatement := `SELECT persons.first_name,persons.last_name,persons.date_of_birth,persons.sex,persons.email,vaccinations.vaccine,vaccinations.vaccination_time,vaccinations.nurse,vaccinations.site FROM vaccinations, persons, nurses where vaccinations.recipient = persons.email and not nurses.email = persons.email`

	rows, err := DB.Query(sqlStatement)

	if err != nil {
		log.Println("Failed to execute query in get return request : ", err)
		return
	}

	defer rows.Close()

	user := getallpersonsvaccinated{}
	for rows.Next() {
		rows.Scan(&user.Getallnurses_persons.First_name, &user.Getallnurses_persons.Last_name, &user.Getallnurses_persons.Date_of_birth, &user.Getallnurses_persons.Sex, &user.Getallnurses_persons.Email, &user.Vaccine, &user.Vaccination_time, &user.Nurse, &user.Site)
		users = append(users, user)
	}

	res := gin.H{
		"result": users,
	}

	c.JSON(http.StatusOK, res)

}
