package imgur

import (
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
