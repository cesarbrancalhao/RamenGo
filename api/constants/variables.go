package constants

import "os"

var apikey = os.Getenv("API_KEY")

func GetApiKey() string {
	return apikey
}
