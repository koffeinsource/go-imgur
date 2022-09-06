package imgur

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func (client *Client) createAPIURL(u string) string {
	if client.rapidAPIKey == "" {
		return apiEndpoint + u
	}
	return apiEndpointRapidAPI + u
}

// getURL returns
// - body as string
// - http status code of request, -1 if request was not made
// - RateLimit with current limits
// - error in case something broke
func (client *Client) getURL(URL string) (string, int, *RateLimit, error) {
	URL = client.createAPIURL(URL)
	client.Log.Infof("Requesting URL %v\n", URL)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return "", -1, nil, errors.New("Could not create request for " + URL + " - " + err.Error())
	}

	req.Header.Add("Authorization", "Client-ID "+client.imgurAccount.clientID)
	if client.rapidAPIKey != "" {
		req.Header.Add("x-rapidapi-host", "imgur-apiv3.p.rapidapi.com")
		req.Header.Add("x-rapidapi-key", client.rapidAPIKey)
	}

	// Make a request to the sourceURL
	res, err := client.httpClient.Do(req)
	if err != nil {
		if res == nil {
			return "", -1, nil, errors.New("Could not get " + URL + " - " + err.Error())
		}
		return "", res.StatusCode, nil, errors.New("Could not get " + URL + " - " + err.Error())
	}
	defer res.Body.Close()

	if !(res.StatusCode >= 200 && res.StatusCode <= 300) {
		return "", res.StatusCode, nil, errors.New("HTTP status indicates an error for " + URL + " - " + res.Status)
	}

	// Read the whole body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", res.StatusCode, nil, errors.New("Problem reading the body for " + URL + " - " + err.Error())
	}

	// Get RateLimit headers
	rl, err := extractRateLimits(res.Header)
	if err != nil {
		client.Log.Infof("Problem with extracting rate limits: %v", err)
	}

	return string(body[:]), res.StatusCode, rl, nil
}
