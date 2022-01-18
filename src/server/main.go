package main

import (
	"kintai/rest"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func initStaff(c *gin.Context) {
	rest.InitStaff()
	c.IndentedJSON(http.StatusOK, rest.GetStaffAll())
}
func getStaffAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, rest.GetStaffAll())
}
func getStaff(c *gin.Context) {
	uid, err := strconv.ParseInt(c.Param("uid"), 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
	} else {
		c.IndentedJSON(http.StatusOK, rest.GetStaff(uid))
	}
}
func main() {
	DemoQmgo()
	router := gin.Default()
	router.GET("/staff/init", getStaffAll)
	router.GET("/staff", getStaffAll)
	router.GET("/staff/:uid", getStaff)
	err := router.Run(":8075")
	if err != nil {
		log.Fatalf("Danger! error at router.Run() %v", err)
	}
}
