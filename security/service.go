package security

import (
	"context"
	"log"
	"math/rand"
)

type Service interface {
	ApiKeyExists(ctx context.Context, apiKey string) (bool, error)
}

type service struct {
	apiKeys []string
}

func NewService() Service {
	// generate random api keys
	keys := []string{}

	log.Println("Generating random API keys")
	for i := 0; i < 10; i++ {
		keys = append(keys, generateRandomString(10))
		log.Println("Generated API key: ", keys[i])
	}

	return &service{
		apiKeys: keys,
	}
}

func (s *service) ApiKeyExists(ctx context.Context, apiKey string) (bool, error) {
	for _, key := range s.apiKeys {
		if key == apiKey {
			return true, nil
		}
	}

	return false, nil
}

func generateRandomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
