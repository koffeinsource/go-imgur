package imgur

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type rateLimitDataWrapper struct {
	Rl      *RateLimit `json:"data"`
	Success bool       `json:"success"`
	Status  int        `json:"status"`
}

// RateLimit details can be found here: https://api.imgur.com/#limits
type RateLimit struct {
	// Total credits that can be allocated.
	UserLimit int
	// Total credits available.
	UserRemaining int
	// Timestamp (unix epoch) for when the credits will be reset.
	UserReset int
	// Total credits that can be allocated for the application in a day.
	ClientLimit int
	// Total credits remaining for the application in a day.
	ClientRemaining int
}

func extractRateLimits(h http.Header) (*RateLimit, error) {
	var rl RateLimit

	userLimit, err := strconv.Atoi(h.Get("X-RateLimit-UserLimit"))
	if err != nil && h.Get("X-RateLimit-UserLimit") != "" {
		return nil, errors.New("Problem parsing X-RateLimit-UserLimit header: " + err.Error())
	}
	rl.UserLimit = userLimit

	userRemaining, err := strconv.Atoi(h.Get("X-RateLimit-UserRemaining"))
	if err != nil && h.Get("X-RateLimit-UserRemaining") != "" {
		return nil, errors.New("Problem parsing X-RateLimit-UserRemaining header: " + err.Error())
	}
	rl.UserRemaining = userRemaining

	userReset, err := strconv.Atoi(h.Get("X-RateLimit-UserReset"))
	if err != nil && h.Get("X-RateLimit-UserReset") != "" {
		return nil, errors.New("Problem parsing X-RateLimit-UserReset header: " + err.Error())
	}
	rl.UserReset = userReset

	clientLimit, err := strconv.Atoi(h.Get("X-RateLimit-ClientLimit"))
	if err != nil && h.Get("X-RateLimit-ClientLimit") != "" {
		return nil, errors.New("Problem parsing X-RateLimit-ClientLimit header: " + err.Error())
	}
	rl.ClientLimit = clientLimit

	clientRemaining, err := strconv.Atoi(h.Get("X-RateLimit-ClientRemaining"))
	if err != nil && h.Get("X-RateLimit-ClientRemaining") != "" {
		return nil, errors.New("Problem parsing X-RateLimit-ClientRemaining header: " + err.Error())
	}
	rl.ClientRemaining = clientRemaining

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
	return rl.Rl, nil

}
