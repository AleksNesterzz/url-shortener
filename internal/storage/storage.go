package storage

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
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
	storage := make(map[shortUrl]longUrl)
	file, err := os.OpenFile("../backup.txt", os.O_RDONLY, 0666)
	if err != nil {
		return &InMemory{cache: storage}
	}
	defer file.Close()
	r := bufio.NewReader(file)
	for {
		str, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		parts := strings.Split(str, ":")
		storage[shortUrl(parts[0])] = longUrl(parts[1] + ":" + parts[2])
	}
	return &InMemory{
		cache: storage,
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

// change function name
func getMD5Hash(url string) string {
	hash := sha256.Sum256(([]byte(url)))
	//hash := md5.Sum([]byte(url))
	return hex.EncodeToString(hash[:])

}
