package imgur

import (
	"net/http"
	"os"
	"testing"
)

func TestRateLimitImgurSimulated(t *testing.T) {
	httpC, server := testHTTPClientJSON("{\"success\": true, \"status\": 200 }")

	defer server.Close()

	client := createClient(httpC, "testing", "")
	rl, err := client.GetRateLimit()

	if err != nil {
		t.Errorf("GetRateLimit() failed with error: %v", err)
		t.FailNow()
	}

	if rl.ClientLimit != 40 || rl.UserLimit != 10 || rl.UserRemaining != 2 || rl.ClientRemaining != 5 {
		client.Log.Debugf("Found ClientLimit: %v and UserLimit: %v", rl.ClientLimit, rl.UserLimit)
		t.Error("Client/User limits are wrong. Probably something broken. Or IMGUR changed their limits. Or you are not using a free account for testing. Sorry. No real good way to test this.")
	}
}

func TestRateLimitRealRapidAPI(t *testing.T) {
	key := os.Getenv("IMGURCLIENTID")
	if key == "" {
		t.Skip("IMGURCLIENTID environment variable not set.")
	}
	RapidAPIKey := os.Getenv("RapidAPIKEY")
	if RapidAPIKey == "" {
		t.Skip("RapidAPIKEY environment variable not set.")
	}

	client := createClient(new(http.Client), key, RapidAPIKey)

	rl, err := client.GetRateLimit()

	if err != nil {
		t.Errorf("GetRateLimit() failed with error: %v", err)
		t.FailNow()
	}

	// There seem to be not rate limites when using the paid API
	if rl.ClientLimit != 0 || rl.UserLimit != 0 {
		client.Log.Debugf("Found ClientLimit: %v and UserLimit: %v", rl.ClientLimit, rl.UserLimit)
		t.Error("Client/User limits are wrong. Probably something broken. Or IMGUR changed their limits. Or you are using a free account for testing. Sorry. No real good way to test this.")
	}
}

func TestRateLimitRealImgur(t *testing.T) {
	key := os.Getenv("IMGURCLIENTID")
	if key == "" {
		t.Skip("IMGURCLIENTID environment variable not set.")
	}

	client := createClient(new(http.Client), key, "")

	rl, err := client.GetRateLimit()

	if err != nil {
		t.Errorf("GetRateLimit() failed with error: %v", err)
		t.FailNow()
	}

	if rl.ClientLimit != 12500 || rl.UserLimit != 500 {
		client.Log.Debugf("Found ClientLimit: %v and UserLimit: %v", rl.ClientLimit, rl.UserLimit)
		t.Error("Client/User limits are wrong. Probably something broken. Or IMGUR changed their limits. Or you are not using a free account for testing. Sorry. No real good way to test this.")
	}
}
