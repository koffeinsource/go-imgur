package imgur

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

type imageInfoDataWrapper struct {
	Ii      *ImageInfo `json:"data"`
	Success bool       `json:"success"`
	Status  int        `json:"status"`
}

// ImageInfo contains all image information provided by imgur
type ImageInfo struct {
	// The ID for the image
	ID string `json:"id"`
	// The title of the image.
	Title string `json:"title"`
	// Description of the image.
	Description string `json:"description"`
	// Time uploaded, epoch time
	Datetime int `json:"datetime"`
	// Image MIME type.
	MimeType string `json:"type"`
	// is the image animated
	Animated bool `json:"animated"`
	// The width of the image in pixels
	Width int `json:"width"`
	// The height of the image in pixels
	Height int `json:"height"`
	// The size of the image in bytes
	Size int `json:"size"`
	// The number of image views
	Views int `json:"views"`
	// Bandwidth consumed by the image in bytes
	Bandwidth int `json:"bandwidth"`
	// OPTIONAL, the deletehash, if you're logged in as the image owner
	Deletehash string `json:"deletehash,omitempty"`
	// OPTIONAL, the original filename, if you're logged in as the image owner
	Name string `json:"name,omitempty"`
	// If the image has been categorized by our backend then this will contain the section the image belongs in. (funny, cats, adviceanimals, wtf, etc)
	Section string `json:"section"`
	// The direct link to the the image. (Note: if fetching an animated GIF that was over 20MB in original size, a .gif thumbnail will be returned)
	Link string `json:"link"`
	// OPTIONAL, The .gifv link. Only available if the image is animated and type is 'image/gif'.
	Gifv string `json:"gifv,omitempty"`
	// OPTIONAL, The direct link to the .mp4. Only available if the image is animated and type is 'image/gif'.
	Mp4 string `json:"mp4,omitempty"`
	// OPTIONAL, The direct link to the .webm. Only available if the image is animated and type is 'image/gif'.
	Webm string `json:"webm,omitempty"`
	// OPTIONAL, Whether the image has a looping animation. Only available if the image is animated and type is 'image/gif'.
	Looping bool `json:"looping,omitempty"`
	// Indicates if the current user favorited the image. Defaults to false if not signed in.
	Favorite bool `json:"favorite"`
	// Indicates if the image has been marked as nsfw or not. Defaults to null if information is not available.
	Nsfw bool `json:"nsfw"`
	// The current user's vote on the album. null if not signed in, if the user hasn't voted on it, or if not submitted to the gallery.
	Vote string `json:"vote"`
	// Current rate limit
	Limit *RateLimit
}

// GetImageInfo queries imgur for information on a image
// returns image info, status code of the request, error
func (client *Client) GetImageInfo(id string) (*ImageInfo, int, error) {
	body, rl, err := client.getURL("image/" + id)
	if err != nil {
		return nil, -1, errors.New("Problem getting URL for image info ID " + id + " - " + err.Error())
	}
	client.Log.Debugf("%v\n", body)

	dec := json.NewDecoder(strings.NewReader(body))
	var img imageInfoDataWrapper
	if err := dec.Decode(&img); err != nil {
		return nil, -1, errors.New("Problem decoding json for imageID " + id + " - " + err.Error())
	}
	img.Ii.Limit = rl

	if !img.Success {
		return nil, img.Status, errors.New("Request to imgur failed for imageID " + id + " - " + strconv.Itoa(img.Status))
	}
	return img.Ii, img.Status, nil
}
