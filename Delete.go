package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {

	var count int
	var count2 int

	reqBody := Getallnurses_persons{}

	c.Bind(&reqBody)

	if reqBody.Email == "" {
		res := gin.H{
			"err":    "email id can't be empty",
			"result": reqBody,
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	sqlStatement2 := "select count(*) from nurses where email = $1"

	row2 := DB.QueryRow(sqlStatement2, reqBody.Email)

	err2 := row2.Scan(&count2)

	if err2 != nil {
		res := gin.H{
			"message": "scan error",
			"result":  err2,
		}

		c.JSON(http.StatusBadRequest, res)

	} else if count2 == 0 {
		res := gin.H{
			"error":   "this email is not been registered in nurses",
			"reqBody": reqBody.Email,
		}
		c.JSON(http.StatusBadRequest, res)

	}

	if count2 != 0 {

		DELETE2 := `DELETE FROM NURSES WHERE email = $1;`

		_, err := DB.Exec(DELETE2, reqBody.Email)

		if err != nil {
			res := gin.H{
				"result": "fail to DELETE in nurses",
				"value":  err.Error(),
			}
			c.JSON(http.StatusBadRequest, res)
		} else {
			res := gin.H{
				"result": "DELETED succefull from nurse",
			}
			c.JSON(http.StatusOK, res)
		}

	}

	sqlStatement := "select count(*) from persons where email = $1"

	row := DB.QueryRow(sqlStatement, reqBody.Email)

	err := row.Scan(&count)

	fmt.Println("count", count)

	if err != nil {
		res := gin.H{
			"message": "scan error",
			"result":  err,
		}

		c.JSON(http.StatusBadRequest, res)

		c.Abort()
		return

	} else if count == 0 {
		res := gin.H{
			"error":   "this email is not been registered in persons",
			"reqBody": reqBody.Email,
		}
		c.JSON(http.StatusBadRequest, res)

		c.Abort()
		return

	}
	if count != 0 {
		DELETE := `DELETE FROM PERSONS WHERE email = $1;`

		_, err := DB.Exec(DELETE, reqBody.Email)

		if err != nil {
			res := gin.H{
				"result": "fail to DELETE in persons",
				"value":  err.Error(),
			}
			c.JSON(http.StatusBadRequest, res)
		} else {
			res := gin.H{
				"result": "DELETED succefull from persons",
			}
			c.JSON(http.StatusOK, res)
		}
	}

}
