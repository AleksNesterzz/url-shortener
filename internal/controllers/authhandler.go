package handlers

import (
	"encoding/json"
	"urlshortner/internal/service"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
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
	a.auth.Register(reg.Email, reg.Password)
}

func (a *AuthHandler) Login(c *gin.Context) {
	var log LoginRequest
	err := json.NewDecoder(c.Request.Body).Decode(&log)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
	}
	a.auth.Login(log.Email, log.Password)
}

func (a *AuthHandler) Logout(c *gin.Context) {

}
