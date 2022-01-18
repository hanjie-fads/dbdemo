package main

import (
	"kintai/api"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getStaffAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, api.GetStaffAll())
}
func getStaff(c *gin.Context) {
	uid, err := strconv.ParseInt(c.Param("uid"), 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
	} else {
		c.IndentedJSON(http.StatusOK, api.GetStaff(uid))
	}
}
func main() {
	router := gin.Default()
	router.GET("/staff", getStaffAll)
	router.GET("/staff/:uid", getStaff)
	err := router.Run(":8075")
	if err != nil {
		log.Fatalf("Danger! error at router.Run() %v", err)
	}
}
