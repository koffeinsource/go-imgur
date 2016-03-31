package imgur

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

type galleryImageInfoDataWrapper struct {
	Ii      *GalleryImageInfo `json:"data"`
	Success bool              `json:"success"`
	Status  int               `json:"status"`
}

// GalleryImageInfo contains all gallery image information provided by imgur
type GalleryImageInfo struct {
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
	// The current user's vote on the album. null if not signed in or if the user hasn't voted on it.
	Vote string `json:"vote"`
	// Indicates if the current user favorited the image. Defaults to false if not signed in.
	Favorite bool `json:"favorite"`
	// Indicates if the image has been marked as nsfw or not. Defaults to null if information is not available.
	Nsfw bool `json:"nsfw"`
	// Number of comments on the gallery album.
	CommentCount int `json:"comment_count"`
	// Up to 10 top level comments, sorted by "best".
	CommentPreview []Comment `json:"comment_preview"`
	// Topic of the gallery album.
	Topic string `json:"topic"`
	// Topic ID of the gallery album.
	TopicID int `json:"topic_id"`
	// If the image has been categorized by our backend then this will contain the section the image belongs in. (funny, cats, adviceanimals, wtf, etc)
	Section string `json:"section"`
	// The username of the account that uploaded it, or null.
	AccountURL string `json:"account_url"`
	// The account ID of the account that uploaded it, or null.
	AccountID int `json:"account_id"`
	// Upvotes for the image
	Ups int `json:"ups"`
	// Number of downvotes for the image
	Downs int `json:"downs"`
	// Upvotes minus downvotes
	Points int `json:"points"`
	// Imgur popularity score
	Score int `json:"score"`
	// if it's an album or not
	IsAlbum bool `json:"is_album"`
}

// GetGalleryImageInfo queries imgur for information on a image
// returns image info, status code of the request, error
func (client *Client) GetGalleryImageInfo(id string) (*GalleryImageInfo, int, error) {
	body, err := client.getURL("gallery/image/" + id)
	if err != nil {
		return nil, -1, errors.New("Problem getting URL for gallery image info ID " + id + " - " + err.Error())
	}
	client.Log.Debugf("%v\n", body)

	dec := json.NewDecoder(strings.NewReader(body))
	var img galleryImageInfoDataWrapper
	if err := dec.Decode(&img); err != nil {
		return nil, -1, errors.New("Problem decoding json for gallery imageID " + id + " - " + err.Error())
	}

	if !img.Success {
		return nil, img.Status, errors.New("Request to imgur failed for gallery imageID " + id + " - " + strconv.Itoa(img.Status))
	}
	return img.Ii, img.Status, nil
}
