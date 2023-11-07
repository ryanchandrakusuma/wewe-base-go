package controller

import (
	"net/http"
	"ryanchandrakusuma/wewe-base-go/domain"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Message string `json:"message"`
}

type BookController struct {
	BUsecase domain.BookUsecase
}

func (bc *BookController) Fetch(c *gin.Context){
	books, err := bc.BUsecase.Fetch(c)
	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}