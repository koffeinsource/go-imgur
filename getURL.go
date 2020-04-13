package imgur

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func (client *Client) createAPIURL(u string) string {
	if client.MashapeKey == "" {
		return apiEndpoint + u
	}
	return apiEndpointMashape + u
}

// getURL returns
// - body as string
// - RateLimit with current limits
// - error in case something broke
func (client *Client) getURL(URL string) (string, *RateLimit, error) {
	URL = client.createAPIURL(URL)
	client.Log.Infof("Requesting URL %v\n", URL)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return "", nil, errors.New("Could not create request for " + URL + " - " + err.Error())
	}

	req.Header.Add("Authorization", "Client-ID "+client.ImgurClientID)
	if client.MashapeKey != "" {
		req.Header.Add("x-rapidapi-host", "imgur-apiv3.p.rapidapi.com")
		req.Header.Add("x-rapidapi-key", client.MashapeKey)
	}

	// Make a request to the sourceURL
	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return "", nil, errors.New("Could not get " + URL + " - " + err.Error())
	}
	defer res.Body.Close()

	if !(res.StatusCode >= 200 && res.StatusCode <= 300) {
		return "", nil, errors.New("HTTP status indicates an error for " + URL + " - " + res.Status)
	}

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
