package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Post_nurse(c *gin.Context) {

	var count int

	reqBody := Getallnurses_persons{}

	r := c.Bind(&reqBody)

	if r != nil {
		res := gin.H{
			"error":        r.Error(),
			"message":      reqBody,
			"date_formate": "1965-12-21",
		}
		c.JSON(http.StatusBadRequest, res)
	}

	if reqBody.City == "" {
		res := gin.H{
			"city": "city cant be empty",
			// "error":        r.Error(),
			"message":      reqBody,
			"date_formate": "1965-12-21",
		}
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	// sqlStatement, err := (`select count(*) from cities where city = '$1'`. scan(&count))

	// sqlStatement, err := DB.QueryRow(`select count(*) from cities where city = '$1'`), &count

	sqlStatement := "select count(*) from cities where city = $1"

	row := DB.QueryRow(sqlStatement, reqBody.City)

	err := row.Scan(&count)

	if err != nil {
		res := gin.H{
			"message": "comming soon in this city",
			"result":  err,
		}

		c.JSON(http.StatusBadRequest, res)
	}

	fmt.Println("count", count)

	if count != 0 {
		insert := `INSERT INTO persons (email, first_name, last_name, date_of_birth, sex, city) VALUES ($1,$2,$3,$4,$5,$6);`

		_, err := DB.Exec(insert, reqBody.Email, reqBody.First_name, reqBody.Last_name, reqBody.Date_of_birth, reqBody.Sex, reqBody.City)

		insert2 := `INSERT INTO nurses (email) VALUES ($1);`

		_, err2 := DB.Exec(insert2, reqBody.Email)

		if err2 != nil {
			res := gin.H{
				"result": "fail to insert",
				"value":  err.Error(),
			}
			c.JSON(http.StatusBadRequest, res)
		} else {
			res := gin.H{
				"result": "insertion succefull",
			}
			c.JSON(http.StatusOK, res)
		}

		if err != nil {
			res := gin.H{
				"result": "fail to insert",
				"value":  err.Error(),
			}
			c.JSON(http.StatusBadRequest, res)
		} else {
			res := gin.H{
				"result": "insertion succefull",
			}
			c.JSON(http.StatusOK, res)
		}

	}
}
