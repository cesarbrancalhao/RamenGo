package constants

import "os"

var apikey = os.Getenv("API_KEY_LOCAL")

func GetApiKey() string {
	return apikey
}
