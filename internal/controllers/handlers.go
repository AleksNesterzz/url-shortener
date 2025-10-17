package handlers

import (
	"net/http"
	"strings"
	"urlshortner/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	user *service.UrlShortener
}

func NewUserHandler(u *service.UrlShortener) *UserHandler {
	return &UserHandler{user: u}
}

func (u *UserHandler) CreateShortUrl(c *gin.Context) {
	val := c.Request.URL.Query()
	url := val.Get("url")
	short, err := u.user.Create(url)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
	}
	c.Writer.Write([]byte(short))

}

func (u *UserHandler) GetLongUrl(c *gin.Context) {
	url := c.Request.URL.String()
	parts := strings.Split(url, "/")
	long, err := u.user.Get(parts[len(parts)-1])
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
	}
	c.Redirect(http.StatusPermanentRedirect, long)
}

func (u *UserHandler) DeleteUrl(c *gin.Context) {

}
