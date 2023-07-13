package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  r.GET("/address/:id", func(c *gin.Context) {
		id, error := strconv.ParseInt( c.Param("id"), 10, 0)
		if error != nil {
			renderError(c, http.StatusBadRequest, "invalid id")
			return
		}

		apiUrl := "https://jsonplaceholder.typicode.com/users/"
		resp, error := http.Get(apiUrl)
		if error != nil {
			renderError(c, http.StatusBadGateway, error.Error())
			defer resp.Body.Close()
			return
		}

		body, error := io.ReadAll(resp.Body)
		// clean up memory after execution
		defer resp.Body.Close()
		if error != nil {
			renderError(c, http.StatusBadGateway, error.Error())
			return
		}

		var usersData []User
		json.Unmarshal(body, &usersData)

		var userData User
		for i := range usersData {
			if usersData[i].Id == id {
				userData = usersData[i]
			}
		}
		if userData.Id == 0 {
			renderError(c, http.StatusNotFound, "user not found")
			return
		}

    c.JSON(http.StatusOK, gin.H{
			"id": userData.Id,
			"address": fmt.Sprintf("%s %s (%s, %s)", userData.Address.City, userData.Address.Zipcode, userData.Address.Geo.Lat, userData.Address.Geo.Lng),
		})
  })

  r.Run()
}

func renderError (c *gin.Context, statusCode int, errorMessage string) {
	c.JSON(statusCode, gin.H{
		"error": errorMessage,
	})
}

type Address struct {
	ID      string `json:"id"`
	Address string `json:"address"`
}

type User struct {
	Id int64
	Name string
	Username string
	Email string
	Address struct {
		Street string
		Suite string
		City string
		Zipcode string
		Geo struct {
			Lat string
			Lng string
		}
	}
	Phone string
	Website string
	Company struct {
		Name string
		CatchPhrase string
		Bs string
	}
}
