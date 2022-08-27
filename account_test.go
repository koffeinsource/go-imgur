package imgur

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerifyEmptyInputParams(t *testing.T) {
	client, err := NewClient(http.DefaultClient, "testing", "")
	require.NoError(t, err)

	_, err = client.RefreshAccessToken("", "secret")
	require.Error(t, err)

	_, err = client.RefreshAccessToken("12345", "")
	require.Error(t, err)
}

func TestStubbedClientAuthorization(t *testing.T) {
	responseBody := `
	{
		"access_token": "argus",
		"expires_in": 315360000,
		"token_type": "bearer",
		"scope": null,
		"refresh_token": "red bull",
		"account_id": 111111111,
		"account_username": "Locker"
	}`
	httpC, server := testHTTPClientJSON(responseBody)
	defer server.Close()

	client, err := NewClient(httpC, "testing", "")
	require.NoError(t, err)

	newRefreshToken, err := client.RefreshAccessToken("12345", "to secret")
	require.NoError(t, err)
	require.Equal(t, "red bull", newRefreshToken)
	require.Equal(t, "argus", client.imgurAccount.accessToken)
}

func TestRealClientAuthorization(t *testing.T) {
	clientID := os.Getenv("IMGURCLIENTID")
	refreshToken := os.Getenv("IMGURCLIENTREFRESHTOKEN")
	clientSecret := os.Getenv("IMGURCLIENTSECRET")
	if clientID == "" || refreshToken == "" || clientSecret == "" {
		t.Skipf("Environment variables not set. ID: %t , refresh: %t , secret: %t",
			clientID == "", refreshToken == "", clientSecret == "")
	}

	client, err := NewClient(&http.Client{}, clientID, "")
	require.NoError(t, err)

	refreshToken, err = client.RefreshAccessToken(refreshToken, clientSecret)
	require.NoError(t, err)

	client.imgurAccount.accessToken = "access token"
	// update refresh token
	err = os.Setenv("IMGURCLIENTREFRESHTOKEN", refreshToken)
	require.NoError(t, err)
	require.NotEqual(t, "access token", client.imgurAccount.accessToken)
}
