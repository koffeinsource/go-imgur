package imgur

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type rateLimitDataWrapper struct {
	Rl      *rateLimitInternal `json:"data"`
	Success bool               `json:"success"`
	Status  int                `json:"status"`
}

// internal representation used for the json parser
type rateLimitInternal struct {
	UserLimit       int64
	UserRemaining   int64
	UserReset       int64
	ClientLimit     int64
	ClientRemaining int64
}

// RateLimit details can be found here: https://api.imgur.com/#limits
type RateLimit struct {
	// Total credits that can be allocated.
	UserLimit int64
	// Total credits available.
	UserRemaining int64
	// Timestamp for when the credits will be reset.
	UserReset time.Time
	// Total credits that can be allocated for the application in a day.
	ClientLimit int64
	// Total credits remaining for the application in a day.
	ClientRemaining int64
}

func extractRateLimits(h http.Header) (rl *RateLimit, err error) {
	err = nil
	var r RateLimit
	rl = &r

	userLimitStr := h.Get("X-RateLimit-UserLimit")
	if userLimitStr != "" {
		rl.UserLimit, err = strconv.ParseInt(userLimitStr, 10, 32)
	}

	userRemainingStr := h.Get("X-RateLimit-UserRemaining")
	if userRemainingStr != "" {
		rl.UserRemaining, err = strconv.ParseInt(userRemainingStr, 10, 32)
	}

	unixTimeStr := h.Get("X-RateLimit-UserReset")
	if unixTimeStr != "" {
		var userReset int64
		userReset, err = strconv.ParseInt(unixTimeStr, 10, 64)
		rl.UserReset = time.Unix(userReset, 0)
	}

	clientLimitStr := h.Get("X-RateLimit-ClientLimit")
	if clientLimitStr != "" {
		rl.ClientLimit, err = strconv.ParseInt(clientLimitStr, 10, 32)
	}

	clientRemainingStr := h.Get("X-RateLimit-ClientRemaining")
	if clientRemainingStr != "" {
		rl.ClientRemaining, err = strconv.ParseInt(clientRemainingStr, 10, 32)
	}

	return
}

// GetRateLimit returns the current rate limit without doing anything else
func (client *Client) GetRateLimit() (*RateLimit, error) {
	// the strange thing here is, that imgur does not add the ratelimit http headers
	// on the credits endpoint. So we must parse the json and ignore the ratelimit
	// returned from getURL
	body, _, err := client.getURL("credits")

	if err != nil {
		return nil, errors.New("Problem getting URL for rate - " + err.Error())
	}
	// client.Log.Debugf("%v\n", body)

	dec := json.NewDecoder(strings.NewReader(body))
	var rl rateLimitDataWrapper
	if err := dec.Decode(&rl); err != nil {
		return nil, errors.New("Problem decoding json for ratelimit - " + err.Error())
	}

	if !rl.Success {
		return nil, errors.New("Request to imgur failed for ratelimit - " + strconv.Itoa(rl.Status))
	}

	var ret RateLimit
	ret.ClientLimit = rl.Rl.ClientLimit
	ret.ClientRemaining = rl.Rl.ClientRemaining
	ret.UserLimit = rl.Rl.UserLimit
	ret.UserRemaining = rl.Rl.UserRemaining
	ret.UserReset = time.Unix(rl.Rl.UserReset, 0)

	return &ret, nil

}
