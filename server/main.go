package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type request struct {
	Exclusions []int `json:"exclusions"`
	MaxCiphers int   `json:"maxCiphers" binding:"required"`
	MinCiphers int   `json:"minCiphers" binding:"required"`
}

type errorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/random", randomMultiplyNums)
	}
}

func randomMultiplyNums(c *gin.Context) {
	var req request

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			errorResponse{
				Error:   err.Error(),
				Message: "Invalid Request",
				Status:  http.StatusBadRequest,
			},
		)
	}
}

func randIntnLimit(max, min int) (num int) {
	//nolint:gosec
	return rand.Intn(max-min+1) + min
}
