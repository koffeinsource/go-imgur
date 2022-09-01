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
	ID          string      `json:"id"`                   // The ID for the album
	Title       string      `json:"title"`                // The title of the album in the gallery
	Description string      `json:"description"`          // The description of the album in the gallery
	DateTime    int         `json:"datetime"`             // Time inserted into the gallery, epoch time
	Cover       string      `json:"cover"`                // The ID of the album cover image
	CoverWidth  int         `json:"cover_width"`          // The width, in pixels, of the album cover image
	CoverHeight int         `json:"cover_height"`         // The height, in pixels, of the album cover image
	AccountURL  string      `json:"account_url"`          // The account username or null if it's anonymous.
	AccountID   int         `json:"account_id"`           // The account ID or null if it's anonymous.
	Privacy     string      `json:"privacy"`              // The privacy level of the album, you can only view public if not logged in as album owner
	Layout      string      `json:"layout"`               // The view layout of the album.
	Views       int         `json:"views"`                // The number of album views
	Link        string      `json:"link"`                 // The URL link to the album
	Favorite    bool        `json:"favorite"`             // Indicates if the current user favorited the image. Defaults to false if not signed in.
	Nsfw        bool        `json:"nsfw"`                 // Indicates if the image has been marked as nsfw or not. Defaults to null if information is not available.
	Section     string      `json:"section"`              // If the image has been categorized by our backend then this will contain the section the image belongs in. (funny, cats, adviceanimals, wtf, etc)
	Order       int         `json:"order"`                // Order number of the album on the user's album page (defaults to 0 if their albums haven't been reordered)
	Deletehash  string      `json:"deletehash,omitempty"` // OPTIONAL, the deletehash, if you're logged in as the album owner
	ImagesCount int         `json:"images_count"`         // The total number of images in the album
	Images      []ImageInfo `json:"images"`               // An array of all the images in the album (only available when requesting the direct album)
	InGallery   bool        `json:"in_gallery"`           // True if the image has been submitted to the gallery, false if otherwise.
	Limit       *RateLimit  // Current rate limit
}

// GetAlbumInfo queries imgur for information on an album
// returns album info, status code of the request or of album payload, error
// http status code of request, -1 if request was not made
func (client *Client) GetAlbumInfo(id string) (*AlbumInfo, int, error) {
	body, statusCode, rl, err := client.getURL("album/" + id)
	if err != nil {
		return nil, statusCode, errors.New("Problem getting URL for album info ID " + id + " - " + err.Error())
	}
	//client.Log.Debugf("%v\n", body)

	dec := json.NewDecoder(strings.NewReader(body))
	var alb albumInfoDataWrapper
	if err := dec.Decode(&alb); err != nil {
		return nil, statusCode, errors.New("Problem decoding json for albumID " + id + " - " + err.Error())
	}

	if !alb.Success {
		return nil, alb.Status, errors.New("Request to imgur failed for albumID " + id + " - " + strconv.Itoa(alb.Status))
	}

	alb.Ai.Limit = rl
	return alb.Ai, alb.Status, nil
}
