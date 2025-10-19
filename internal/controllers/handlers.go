package handlers

import (
	"encoding/json"
	"net/http"
	"urlshortner/internal/service"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Url string `json:"url"`
}
type UserHandler struct {
	user *service.UrlShortener
}

func New(u *service.UrlShortener) *UserHandler {
	return &UserHandler{user: u}
}

func (u *UserHandler) CreateShortUrl(c *gin.Context) {
	var req CreateRequest
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
	}
	short, err := u.user.Create(req.Url)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
	}
	c.Writer.Write([]byte(short))

}

func (u *UserHandler) GetLongUrl(c *gin.Context) {
	url := c.Param("id")
	long, err := u.user.Get(url)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}
	c.Redirect(http.StatusPermanentRedirect, long)
}

func (u *UserHandler) DeleteUrl(c *gin.Context) {
	url := c.Param("id")
	u.user.Delete(url)
	c.JSON(http.StatusNoContent, gin.H{})
}
