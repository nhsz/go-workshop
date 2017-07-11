package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()
	r.GET("/calculate", calculate)
	r.Run()
}

func calculate(c *gin.Context) {
	if v := c.Query("value"); v != "" {
		if p, error := strconv.Atoi(v); error == nil {
			sum := SumDivisibleValuesInRange(p, 3, 5)

			c.JSON(http.StatusOK, gin.H{
				"sum": sum,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("Invalid arguments..., error: %s", error.Error()),
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid arguments..., must specify a value parameter",
		})
	}
}

func SumDivisibleValuesInRange(p, x, y int) int {
	sum := 0
	for v := 0; v <= p; v++ {
		if IsDivisibleBy(v, x, y) {
			sum += v
		}
	}

	return sum
}

func IsDivisibleBy(v, i, j int) bool {
	return v%i == 0 || v%j == 0
}
