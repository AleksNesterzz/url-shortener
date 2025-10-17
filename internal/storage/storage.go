package storage

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type Storage interface {
	Get(url string) (string, error)
	Create(url string) (string, error)
	Delete(url string) error
}

type shortUrl string
type longUrl string

type InMemory struct {
	cache map[shortUrl]longUrl
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

func getMD5Hash(url string) string {
	hash := md5.Sum([]byte(url))
	return hex.EncodeToString(hash[:])

}
