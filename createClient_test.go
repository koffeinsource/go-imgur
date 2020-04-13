package imgur

import (
	"net/http"

	"github.com/koffeinsource/go-klogger"
)

// createClient simply creates an imgur client. RapidAPIKEY is "" if you are using the free API.
func createClient(httpClient *http.Client, imgurClientID string, RapidAPIKEY string) *Client {
	client := new(Client)
	client.HTTPClient = httpClient
	client.Log = new(klogger.CLILogger)
	client.ImgurClientID = imgurClientID
	client.RapidAPIKEY = RapidAPIKEY

	return client
}
