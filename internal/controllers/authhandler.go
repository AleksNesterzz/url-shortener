package handlers

import (
	"encoding/json"
	"net/http"
	service "urlshortner/internal/service/auth"

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
type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	Status string `json:"status"`
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "request decoding error"})
	}
	resp := RegisterResponse{
		Status: "you've been succesfully registered",
	}
	err = a.auth.Register(reg.Email, reg.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "registering error"})
	}
	c.JSON(200, resp)
}

func (a *AuthHandler) Login(c *gin.Context) {
	var log LoginRequest
	err := json.NewDecoder(c.Request.Body).Decode(&log)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "request decoding error"})
	}
	token, err := a.auth.Login(log.Email, log.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "logging in error"})
	}
	logResp := &LoginResponse{
		Token: token,
	}
	c.JSON(200, logResp)
}

func (a *AuthHandler) Logout(c *gin.Context) {

}
