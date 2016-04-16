package imgur

import (
	"net/http"
	"os"
	"testing"

	"github.com/koffeinsource/go-klogger"
)

func TestRateLimitImgurSimulated(t *testing.T) {
	httpC, server := testHTTPClientJSON("{\"data\": { \"UserLimit\": 123, \"UserRemaining\": 456, \"UserReset\": 1460830093, \"ClientLimit\": 99, \"ClientRemaining\": 80 }, \"success\": true, \"status\": 200 }")
	defer server.Close()

	client := new(Client)
	client.HTTPClient = httpC
	client.Log = new(klogger.CLILogger)
	client.ImgurClientID = "testing"

	rl, err := client.GetRateLimit()

	if err != nil {
		t.Errorf("GetRateLimit() failed with error: %v", err)
		t.FailNow()
	}

	if rl.ClientLimit != 99 || rl.UserLimit != 123 || rl.UserRemaining != 456 || rl.ClientRemaining != 80 {
		t.Error("Client/User limits are wrong. Probably something broken. Or IMGUR changed their limits. Or you are not using a free account for testing. Sorry. No real good way to test this.")
	}
}

func TestRateLimitReal(t *testing.T) {
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
		t.FailNow()
	}

	if rl.ClientLimit != 12500 || rl.UserLimit != 500 {
		t.Error("Client/User limits are wrong. Probably something broken. Or IMGUR changed their limits. Or you are not using a free account for testing. Sorry. No real good way to test this.")
	}
}

func TestRateLimitImgurNotSuccess(t *testing.T) {
	httpC, server := testHTTPClientJSON("{\"data\": {}, \"success\": false, \"status\": 200 }")
	defer server.Close()

	client := new(Client)
	client.HTTPClient = httpC
	client.Log = new(klogger.CLILogger)
	client.ImgurClientID = "testing"

	_, err := client.GetRateLimit()

	if err == nil {
		t.Error("GetRateLimit() should have failed, but didn't")
	}
}

func TestRateLimitJsonError(t *testing.T) {
	httpC, server := testHTTPClientInvalidJSON()
	defer server.Close()

	client := new(Client)
	client.HTTPClient = httpC
	client.Log = new(klogger.CLILogger)
	client.ImgurClientID = "testing"

	_, err := client.GetRateLimit()

	if err == nil {
		t.Error("GetRateLimit() should have failed, but didn't")
	}
}

func TestRateLimitServerError(t *testing.T) {
	httpC, server := testHTTPClient500()
	defer server.Close()

	client := new(Client)
	client.HTTPClient = httpC
	client.Log = new(klogger.CLILogger)
	client.ImgurClientID = "testing"

	_, err := client.GetRateLimit()

	if err == nil {
		t.Error("GetRateLimit() should have failed, but didn't")
	}
}

func TestRateLimitImgurError(t *testing.T) {
	httpC, server := testHTTPClientJSON("{'data' : {}, 'success' : false, 'status'  : 500}")
	defer server.Close()

	client := new(Client)
	client.HTTPClient = httpC
	client.Log = new(klogger.CLILogger)
	client.ImgurClientID = "testing"

	_, err := client.GetRateLimit()

	if err == nil {
		t.Error("GetRateLimit() should have failed, but didn't")
	}
}
