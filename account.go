package imgur

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GenerateAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token"` // The refresh token returned from the authorization code exchange
	ClientID     string `json:"client_id"`     // The client_id obtained during application registration
	ClientSecret string `json:"client_secret"` // The client secret obtained during application registration
	GrandType    string `json:"grant_type"`    // As defined in the OAuth2 specification, this field must contain a value of: refresh_token
}

type GenerateAccessTokenResponse struct {
	AccessToken     string `json:"access_token"` // TNew access token to use
	ExpiresIn       uint64 `json:"expires_in"`   // These parameters describe the lifetime of the token in seconds, and the kind of token that is being returned
	TokenType       string `json:"token_type"`
	Scope           string `json:"scope,omitempty"`            // Scope which were provided earlier during creation access_token
	RefreshToken    string `json:"refresh_token"`              // New refresh token
	AccountID       int    `json:"account_id,omitempty"`       // not specified in documentation
	AccountUserName string `json:"account_username,omitempty"` // not specified in documentation
}

// RefreshAccessToken let you reissue expired access_token
func (c *Client) RefreshAccessToken(refreshToken string, clientSecret string) (string, error) {
	if len(refreshToken) == 0 {
		msg := "Refresh token is empty"
		c.Log.Errorf(msg)
		return "", fmt.Errorf(msg)
	}

	if len(clientSecret) == 0 {
		msg := "Client secret is empty"
		c.Log.Errorf(msg)
		return "", fmt.Errorf(msg)
	}

	rawBody, err := json.Marshal(
		GenerateAccessTokenRequest{
			RefreshToken: refreshToken,
			ClientID:     c.imgurAccount.clientID,
			ClientSecret: clientSecret,
			GrandType:    "refresh_token",
		})
	if err != nil {
		c.Log.Errorf("Failed to marshal GenerateAccessToken. %v", err)
		return "", err
	}

	c.Log.Debugf("Prepared body %v", string(rawBody))

	url := apiEndpointGenerateAccessToken
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(rawBody))
	if err != nil {
		c.Log.Errorf("Failed to create new request for refresh access token. %v", err)
		return "", err
	}

	c.Log.Infof("Sending request to refresh access token")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.Log.Errorf("HTTP request was failed. %v", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Log.Errorf("Reading response body was failed. %v", err)
		return "", err
	}

	response := GenerateAccessTokenResponse{}
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.UseNumber()
	if err = decoder.Decode(&response); err != nil {
		c.Log.Errorf("Decoding response was failed. %v", err)
		return "", err
	}

	c.Log.Infof("Token was success updated and it will be relevant within next %v seconds", response.ExpiresIn)
	c.Log.Debugf("New token: %v New refresh token: %v", response.AccessToken, response.RefreshToken)

	c.imgurAccount.accessToken = response.AccessToken
	return response.RefreshToken, nil
}
