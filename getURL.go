package imgur

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// getURL returns
// - body as string
// - RateLimit with current limits
// - error in case something broke
func (client *Client) getURL(URL string) (string, *RateLimit, error) {
	URL = apiEndpoint + URL
	client.Log.Infof("Requesting URL %v\n", URL)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return "", nil, errors.New("Could not create request for " + URL + " - " + err.Error())
	}

	req.Header.Add("Authorization", "Client-ID "+client.ImgurClientID)

	// Make a request to the sourceURL
	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return "", nil, errors.New("Could not get " + URL + " - " + err.Error())
	}
	defer res.Body.Close()

	// Read the whole body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", nil, errors.New("Problem reading the body for " + URL + " - " + err.Error())
	}

	// Get RateLimit headers
	rl, err := extractRateLimits(res.Header)
	if err != nil {
		client.Log.Infof("Problem with extracting rate limits: %v", err)
	}

	return string(body[:]), rl, nil
}
