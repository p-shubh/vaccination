package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

var DB *sql.DB

type Getallnurses_persons struct {
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Date_of_birth string `json:"date_of_birth"`
	Sex           string `json:"sex"`
	Email         string `json:"email"`
	City          string `json:"city" binding:"required`
	State         string `json:"state"`
}

type getallpersonsvaccinated struct {
	Getallnurses_persons Getallnurses_persons
	Vaccine              string `json:"vaccine"`
	Vaccination_time     string `json:"vaccination_time"`
	Nurse                string `json:"nurse"`
	Site                 string `json:"site"`
}

func main() {

	connection_with_db()
	defer DB.Close()

	router := gin.Default()
	setupRoutes(router)
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func setupRoutes(routes *gin.Engine) {

	routes.GET("getallnurses", Get_all_nurses)
	routes.GET("getallpersons", Get_all_persons)
	routes.GET("personsvaccinated", Get_all_persons_vaccinated)
	routes.GET("nursesvaccinated", Get_all_nurses_vaccinated)
	routes.POST("addpatients", PostPatients)
	routes.POST("addnurse", Post_nurse)
	routes.DELETE("delete", Delete)
}

// Get all person who are vaccinated
// Get all nurses who are vaccinated
