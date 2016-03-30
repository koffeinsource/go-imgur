package imgur

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func getURL(URL string, client *Client) (string, error) {
	URL = apiEndpoint + URL
	client.Log.Infof("Requesting URL %v\n", URL)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return "", errors.New("Could create request for " + URL + " - " + err.Error())
	}

	req.Header.Add("Authorization", "Client-ID "+client.ImgurClientID)

	// Make a request to the sorceURL
	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return "", errors.New("Could not get " + URL + " - " + err.Error())
	}
	defer res.Body.Close()

	// Read the whole body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.New("Problem reading the body for " + URL + " - " + err.Error())
	}

	return string(body[:]), nil
}
