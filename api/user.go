package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Nickname string
	Email    string
}

type createUserResponse struct {
	Status string
}

func (server *Server) createUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, createUserResponse{"works"})
}
