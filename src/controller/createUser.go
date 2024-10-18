package controller

import (
	"git.sr.ht/~bacarin/curso-golang/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Erro!!")
	c.JSON(err.Code, err)

}