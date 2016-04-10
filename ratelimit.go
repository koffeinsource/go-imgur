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

func extractRateLimits(h http.Header) (*RateLimit, error) {
	var rl RateLimit

	userLimitStr := h.Get("X-RateLimit-UserLimit")
	if userLimitStr != "" {
		userLimit, err := strconv.ParseInt(userLimitStr, 10, 32)
		if err != nil {
			return nil, errors.New("Problem parsing X-RateLimit-UserLimit header: " + err.Error())
		}
		rl.UserLimit = userLimit
	}

	userRemainingStr := h.Get("X-RateLimit-UserRemaining")
	if userRemainingStr != "" {
		userRemaining, err := strconv.ParseInt(userRemainingStr, 10, 32)
		if err != nil {
			return nil, errors.New("Problem parsing X-RateLimit-UserRemaining header: " + err.Error())
		}
		rl.UserRemaining = userRemaining
	}

	unixTimeStr := h.Get("X-RateLimit-UserReset")
	if unixTimeStr != "" {
		userReset, err := strconv.ParseInt(unixTimeStr, 10, 64)
		if err != nil {
			return nil, errors.New("Problem parsing X-RateLimit-UserReset header: " + err.Error())
		}
		rl.UserReset = time.Unix(userReset, 0)
	}

	clientLimitStr := h.Get("X-RateLimit-ClientLimit")
	if clientLimitStr != "" {
		clientLimit, err := strconv.ParseInt(clientLimitStr, 10, 32)
		if err != nil && h.Get("X-RateLimit-ClientLimit") != "" {
			return nil, errors.New("Problem parsing X-RateLimit-ClientLimit header: " + err.Error())
		}
		rl.ClientLimit = clientLimit
	}

	clientRemainingStr := h.Get("X-RateLimit-ClientRemaining")
	if clientRemainingStr != "" {
		clientRemaining, err := strconv.ParseInt(clientRemainingStr, 10, 32)
		if err != nil {
			return nil, errors.New("Problem parsing X-RateLimit-ClientRemaining header: " + err.Error())
		}
		rl.ClientRemaining = clientRemaining
	}

	return &rl, nil
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
