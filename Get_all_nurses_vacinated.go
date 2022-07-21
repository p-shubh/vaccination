package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get_all_nurses_vaccinated(c *gin.Context) {

	users := []getallpersonsvaccinated{}

	// reqBody := getallpersonsvaccinated{}

	sqlStatement := `SELECT persons.first_name,persons.last_name,persons.date_of_birth,persons.sex,persons.email,vaccinations.vaccine,vaccinations.vaccination_time,vaccinations.nurse,vaccinations.site FROM vaccinations, persons, nurses where nurses.email = persons.email and persons.email = vaccinations.recipient`

	r, err := DB.Query(sqlStatement)

	if err != nil {
		log.Println("Failed to execute query in get return request : ", err)
		return
	}

	defer r.Close()

	// r.Scan(&)

	user := getallpersonsvaccinated{}

	for r.Next() {
		r.Scan(&user.Getallnurses_persons.First_name, &user.Getallnurses_persons.Last_name, &user.Getallnurses_persons.Date_of_birth, &user.Getallnurses_persons.Sex, &user.Getallnurses_persons.Email, &user.Vaccine, &user.Vaccination_time, &user.Nurse, &user.Site)

		users = append(users, user)
	}

	res := gin.H{
		"vacinated nurse": users,
	}

	c.JSON(http.StatusOK, res)

}
