package main

import (
	"log"
	"net/http"

	"server/utils"

	"github.com/gin-gonic/gin"
)

type request struct {
	MaxCiphers int `json:"maxCiphers" binding:"required,gte=0"`
	MinCiphers int `json:"minCiphers" binding:"required,gte=0"`
}

type response struct {
	Result int `json:"result"`
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

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}

func randomMultiplyNums(c *gin.Context) {
	var req request

	var res response

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

	if err := utils.CheckMinCipher(req.MaxCiphers, req.MinCiphers); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			errorResponse{
				Error:   err.Error(),
				Message: "Invalid Request",
				Status:  http.StatusBadRequest,
			},
		)
	}

	res.Result = utils.RandIntnLimit(
		utils.GetMaxNum(req.MaxCiphers),
		utils.GetMinNum(req.MinCiphers),
	)

	c.JSON(http.StatusOK, res)
}
