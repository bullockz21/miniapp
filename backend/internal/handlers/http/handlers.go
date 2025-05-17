package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetHandlerTest(c *gin.Context) {
	u := User{
		Name: "Susan",
		Age:  21,
	}
	c.JSON(http.StatusOK, gin.H{
		"name": u.Name,
		"age":  u.Age,
	})
}

func PostHandlerTest(c *gin.Context) {
	u := &User{
		Name: "",
		Age:  0,
	}
	if err := c.ShouldBindJSON(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name": u.Name,
		"age":  u.Age,
	})

}
