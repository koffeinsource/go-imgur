package imgur

import (
	"net/http"

	"github.com/koffeinsource/go-klogger"
)

// createClient simply creates an imgur client. mashapeKey is "" if you are using the free API.
func createClient(httpClient *http.Client, imgurClientID string, mashapeKey string) *Client {
	client := new(Client)
	client.HTTPClient = httpClient
	client.Log = new(klogger.CLILogger)
	client.ImgurClientID = imgurClientID
	client.MashapeKey = mashapeKey

	return client
}
