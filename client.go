package imgur

import (
	"net/http"

	"github.com/koffeinsource/go-klogger"
)

// Client used to configure go-imgur
type Client struct {
	HTTPClient    *http.Client
	Log           klogger.KLogger
	ImgurClientID string
}
