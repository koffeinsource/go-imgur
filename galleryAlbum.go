package imgur

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

type galleryAlbumInfoDataWrapper struct {
	Ai      *GalleryAlbumInfo `json:"data"`
	Success bool              `json:"success"`
	Status  int               `json:"status"`
}

// GalleryAlbumInfo contains all information provided by imgur of a gallery album
type GalleryAlbumInfo struct {
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
	// The total number of images in the album
	ImagesCount int `json:"images_count"`
	// An array of all the images in the album (only available when requesting the direct album)
	Images []ImageInfo `json:"images,omitempty"`
	// Current rate limit
	Limit *RateLimit
}

// GetGalleryAlbumInfo queries imgur for information on a gallery album
// returns album info, status code of the request, error
func (client *Client) GetGalleryAlbumInfo(id string) (*GalleryAlbumInfo, int, error) {
	body, rl, err := client.getURL("gallery/album/" + id)
	if err != nil {
		return nil, -1, errors.New("Problem getting URL for gallery album info ID " + id + " - " + err.Error())
	}
	// client.Log.Debugf("%v\n", body)

	dec := json.NewDecoder(strings.NewReader(body))
	var alb galleryAlbumInfoDataWrapper
	if err := dec.Decode(&alb); err != nil {
		return nil, -1, errors.New("Problem decoding json for gallery albumID " + id + " - " + err.Error())
	}
	alb.Ai.Limit = rl

	if !alb.Success {
		return nil, alb.Status, errors.New("Request to imgur failed for gallery albumID " + id + " - " + strconv.Itoa(alb.Status))
	}
	return alb.Ai, alb.Status, nil
}
