package storage

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

type Storage interface {
	Get(url string) (string, error)
	Create(url string) (string, error)
	Delete(url string) error
	Save() error
}

type shortUrl string
type longUrl string

type InMemory struct {
	cache map[shortUrl]longUrl
}

func NewInMemory() *InMemory {
	return &InMemory{
		cache: make(map[shortUrl]longUrl),
	}
}

func (s *InMemory) Create(url string) (string, error) {
	hash := getMD5Hash(url)
	s.cache[shortUrl(hash)] = longUrl(url)
	return hash, nil
}

func (s *InMemory) Get(url string) (string, error) {
	result, ok := s.cache[shortUrl(url)]
	if !ok {
		return "", fmt.Errorf("invalid url")
	}
	return string(result), nil
}

func (s *InMemory) Delete(url string) error {
	if _, ok := s.cache[shortUrl(url)]; !ok {
		return fmt.Errorf("no such url")
	}
	delete(s.cache, shortUrl(url))
	return nil
}

func (s *InMemory) Save() error {
	file, err := os.OpenFile("../backup.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	for k, v := range s.cache {
		var str strings.Builder
		str.Grow(len(k) + len(v) + 1)
		_, err = str.WriteString(string(k) + ":" + string(v) + "\n")
		if err != nil {
			return err
		}
		file.Write([]byte(str.String()))
	}

	return nil
}

func getMD5Hash(url string) string {
	hash := md5.Sum([]byte(url))
	return hex.EncodeToString(hash[:])

}
