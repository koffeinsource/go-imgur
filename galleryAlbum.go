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
	ID           string      `json:"id"`               // The ID for the album
	Title        string      `json:"title"`            // The title of the album in the gallery
	Description  string      `json:"description"`      // The description of the album in the gallery
	DateTime     int         `json:"datetime"`         // Time inserted into the gallery, epoch time
	Cover        string      `json:"cover"`            // The ID of the album cover image
	CoverWidth   int         `json:"cover_width"`      // The width, in pixels, of the album cover image
	CoverHeight  int         `json:"cover_height"`     // The height, in pixels, of the album cover image
	AccountURL   string      `json:"account_url"`      // The account username or null if it's anonymous.
	AccountID    int         `json:"account_id"`       // The account ID or null if it's anonymous.
	Privacy      string      `json:"privacy"`          // The privacy level of the album, you can only view public if not logged in as album owner
	Layout       string      `json:"layout"`           // The view layout of the album.
	Views        int         `json:"views"`            // The number of album views
	Link         string      `json:"link"`             // The URL link to the album
	Ups          int         `json:"ups"`              // Upvotes for the image
	Downs        int         `json:"downs"`            // Number of downvotes for the image
	Points       int         `json:"points"`           // Upvotes minus downvotes
	Score        int         `json:"score"`            // Imgur popularity score
	IsAlbum      bool        `json:"is_album"`         // if it's an album or not
	Vote         string      `json:"vote"`             // The current user's vote on the album. null if not signed in or if the user hasn't voted on it.
	Favorite     bool        `json:"favorite"`         // Indicates if the current user favorited the image. Defaults to false if not signed in.
	Nsfw         bool        `json:"nsfw"`             // Indicates if the image has been marked as nsfw or not. Defaults to null if information is not available.
	CommentCount int         `json:"comment_count"`    // Number of comments on the gallery album.
	Topic        string      `json:"topic"`            // Topic of the gallery album.
	TopicID      int         `json:"topic_id"`         // Topic ID of the gallery album.
	ImagesCount  int         `json:"images_count"`     // The total number of images in the album
	Images       []ImageInfo `json:"images,omitempty"` // An array of all the images in the album (only available when requesting the direct album)
	InMostViral  bool        `json:"in_most_viral"`    // Indicates if the album is in the most viral gallery or not.
	Limit        *RateLimit  // Current rate limit
}

// GetGalleryAlbumInfo queries imgur for information on a gallery album
// returns album info, status code of the request, error
func (client *Client) GetGalleryAlbumInfo(id string) (*GalleryAlbumInfo, int, error) {
	body, statusCode, rl, err := client.getURL("gallery/album/" + id)
	if err != nil {
		return nil, statusCode, errors.New("Problem getting URL for gallery album info ID " + id + " - " + err.Error())
	}
	// client.Log.Debugf("%v\n", body)

	dec := json.NewDecoder(strings.NewReader(body))
	var alb galleryAlbumInfoDataWrapper
	if err := dec.Decode(&alb); err != nil {
		return nil, statusCode, errors.New("Problem decoding json for gallery albumID " + id + " - " + err.Error())
	}
	alb.Ai.Limit = rl

	if !alb.Success {
		return nil, alb.Status, errors.New("Request to imgur failed for gallery albumID " + id + " - " + strconv.Itoa(alb.Status))
	}
	return alb.Ai, alb.Status, nil
}
