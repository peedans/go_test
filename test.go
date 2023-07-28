package main

import (
	"github.com/gin-gonic/gin"
)

type multiplyRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type multiplyResponse struct {
	Result int `json:"result"`
}

func main() {
	router := gin.Default()

	router.GET("/hello", hello)
	router.POST("/multiply", multiply)

	router.Run(":8080")
}

func hello(c *gin.Context) {
	c.String(200, "Hello, World!")
}

func multiply(c *gin.Context) {
	var req multiplyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res := calculateMultiplyResponse(req)

	c.JSON(200, res)
}

func calculateMultiplyResponse(req multiplyRequest) multiplyResponse {
	result := req.A * req.B
	return multiplyResponse{Result: result}
}
