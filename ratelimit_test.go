package imgur

import (
	"net/http"
	"os"
	"testing"

	"github.com/koffeinsource/go-klogger"
)

func TestRateLimit(t *testing.T) {
	key := os.Getenv("IMGURCLIENTID")
	if key == "" {
		t.Skip("IMGURCLIENTID environment variable not set.")
	}

	client := new(Client)
	client.HTTPClient = new(http.Client)
	client.Log = new(klogger.CLILogger)
	client.ImgurClientID = key

	rl, err := client.GetRateLimit()

	if err != nil {
		t.Errorf("GetRateLimit() failed with error: %v", err)
	}

	if rl.ClientLimit != 12500 || rl.UserLimit != 500 {
		t.Error("Client/User limits are wrong. Probably something broken. Or IMGUR changed their limits. Or you are not using a free account for testing. Sorry. No real good way to test this.")
	}
}
