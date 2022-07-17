package imgur

import (
	"fmt"
	"net/http"

	"github.com/koffeinsource/go-klogger"
)

// Client used to for go-imgur
type Client struct {
	HTTPClient    *http.Client
	Log           klogger.KLogger
	ImgurClientID string
	RapidAPIKEY   string
}

// NewClient simply creates an imgur client. RapidAPIKEY is "" if you are using the free API.
func NewClient(httpClient *http.Client, clientID string, rapidAPIKEY string) (*Client, error) {
	logger := new(klogger.CLILogger)

	if len(clientID) == 0 {
		msg := "imgur client ID is empty"
		logger.Errorf(msg)
		return nil, fmt.Errorf(msg)
	}

	if len(rapidAPIKEY) == 0 {
		logger.Infof("rapid api key is empty")
	}

	return &Client{
		HTTPClient:    httpClient,
		Log:           logger,
		ImgurClientID: clientID,
		RapidAPIKEY:   rapidAPIKEY,
	}, nil
}
