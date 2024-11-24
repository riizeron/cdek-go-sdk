package v2

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func createTestClient() Client {
	wd, _ := os.Getwd()
	godotenv.Load(fmt.Sprintf("%s/.env.local", wd))

	clientId := os.Getenv("CDEK_CLIENT_ID")
	clientSecretId := os.Getenv("CDEK_SECRET_ID")

	// public cdek test credentials
	if clientId == "" {
		clientId = "wqGwiQx0gg8mLtiEKsUinjVSICCjtTEP"
	}
	if clientSecretId == "" {
		clientSecretId = "RmAmgvSgSl1yirlz9QupbzOJVqhCxcP5"
	}

	return NewClient(EndpointTest, clientId, clientSecretId)
}
