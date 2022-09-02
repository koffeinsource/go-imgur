package imgur

import (
	"fmt"
	"net/http"

	"github.com/koffeinsource/go-klogger"
)

// ClientAccount describe authontification
type ClientAccount struct {
	clientID    string // client ID received after registration
	accessToken string // is your secret key used to access the user's data
}

// Client used to for go-imgur
type Client struct {
	Log          klogger.KLogger
	httpClient   *http.Client
	imgurAccount ClientAccount
	rapidAPIKey  string
}

// NewClient simply creates an imgur client. RapidAPIKEY is "" if you are using the free API.
func NewClient(httpClient *http.Client, clientID string, rapidAPIKey string) (*Client, error) {
	logger := new(klogger.CLILogger)

	if len(clientID) == 0 {
		msg := "imgur client ID is empty"
		logger.Errorf(msg)
		return nil, fmt.Errorf(msg)
	}

	if len(rapidAPIKey) == 0 {
		logger.Infof("rapid api key is empty")
	}

	return &Client{
		httpClient:  httpClient,
		Log:         logger,
		rapidAPIKey: rapidAPIKey,
		imgurAccount: ClientAccount{
			clientID: clientID,
		},
	}, nil
}
