package handlers

import (
	"net/http"
	"strings"
	"urlshortner/internal/service"
)

type UserHandler struct {
	user *service.UrlShortener
}

func NewUserHandler(u *service.UrlShortener) *UserHandler {
	return &UserHandler{user: u}
}

func (u *UserHandler) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	short, err := u.user.Create(url)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte(short))

}

func (u *UserHandler) GetLongUrl(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	parts := strings.Split(url, "/")
	long, err := u.user.Get(parts[len(parts)-1])
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	http.Redirect(w, r, long, http.StatusPermanentRedirect)
}

func (u *UserHandler) DeleteUrl(w http.ResponseWriter, r *http.Request) {

}
