package imgur

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

type albumInfoDataWrapper struct {
	Ai      *AlbumInfo `json:"data"`
	Success bool       `json:"success"`
	Status  int        `json:"status"`
}

// AlbumInfo contains all album information provided by imgur
type AlbumInfo struct {
	// The ID for the album
	ID string `json:"id"`
	// The title of the album in the gallery
	Title string `json:"title"`
	// The description of the album in the gallery
	Description string `json:"description"`
	// Time inserted into the gallery, epoch time
	DateTime int `json:"datetime"`
	// The ID of the album cover image
	Cover string `json:"cover"`
	// The width, in pixels, of the album cover image
	CoverWidth int `json:"cover_width"`
	// The height, in pixels, of the album cover image
	CoverHeight int `json:"cover_height"`
	// The account username or null if it's anonymous.
	AccountURL string `json:"account_url"`
	// The account ID or null if it's anonymous.
	AccountID int `json:"account_id"`
	// The privacy level of the album, you can only view public if not logged in as album owner
	Privacy string `json:"privacy"`
	// The view layout of the album.
	Layout string `json:"layout"`
	// The number of album views
	Views int `json:"views"`
	// The URL link to the album
	Link string `json:"link"`
	// Indicates if the current user favorited the image. Defaults to false if not signed in.
	Favorite bool `json:"favorite"`
	// Indicates if the image has been marked as nsfw or not. Defaults to null if information is not available.
	Nsfw bool `json:"nsfw"`
	// If the image has been categorized by our backend then this will contain the section the image belongs in. (funny, cats, adviceanimals, wtf, etc)
	Section string `json:"secion"`
	// Order number of the album on the user's album page (defaults to 0 if their albums haven't been reordered)
	Order int `json:"order"`
	// OPTIONAL, the deletehash, if you're logged in as the album owner
	Deletehash string `json:"deletehash,omitempty"`
	// The total number of images in the album
	ImagesCount int `json:"image_count"`
	// An array of all the images in the album (only available when requesting the direct album)
	Images []ImageInfo `json:"images"`
	// Current rate limit
	Limit *RateLimit
}

// GetAlbumInfo queries imgur for information on a album
// returns album info, status code of the request, error
func (client *Client) GetAlbumInfo(id string) (*AlbumInfo, int, error) {
	body, rl, err := client.getURL("album/" + id)
	if err != nil {
		return nil, -1, errors.New("Problem getting URL for album info ID " + id + " - " + err.Error())
	}
	client.Log.Debugf("%v\n", body)

	dec := json.NewDecoder(strings.NewReader(body))
	var alb albumInfoDataWrapper
	if err := dec.Decode(&alb); err != nil {
		return nil, -1, errors.New("Problem decoding json for albumID " + id + " - " + err.Error())
	}
	alb.Ai.Limit = rl

	if !alb.Success {
		return nil, alb.Status, errors.New("Request to imgur failed for albumID " + id + " - " + strconv.Itoa(alb.Status))
	}
	return alb.Ai, alb.Status, nil
}
