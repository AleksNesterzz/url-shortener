package handlers

import (
	"encoding/json"
	"urlshortner/internal/service"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthHandler struct {
	auth *service.Auth
}

func NewAuthHandler(a *service.Auth) *AuthHandler {
	return &AuthHandler{auth: a}
}

func (a *AuthHandler) Register(c *gin.Context) {
	var reg RegisterRequest
	err := json.NewDecoder(c.Request.Body).Decode(&reg)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
	}
	a.auth.Register()
}

func (a *AuthHandler) Login(c *gin.Context) {

}

func (a *AuthHandler) Logout(c *gin.Context) {

}
