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
	ID          string     `json:"id"`                   // The ID for the image
	Title       string     `json:"title"`                // The title of the image.
	Description string     `json:"description"`          // Description of the image.
	Datetime    int        `json:"datetime"`             // Time uploaded, epoch time
	MimeType    string     `json:"type"`                 // Image MIME type.
	Animated    bool       `json:"animated"`             // is the image animated
	Width       int        `json:"width"`                // The width of the image in pixels
	Height      int        `json:"height"`               // The height of the image in pixels
	Size        int        `json:"size"`                 // The size of the image in bytes
	Views       int        `json:"views"`                // The number of image views
	Bandwidth   int        `json:"bandwidth"`            // Bandwidth consumed by the image in bytes
	Deletehash  string     `json:"deletehash,omitempty"` // OPTIONAL, the deletehash, if you're logged in as the image owner
	Name        string     `json:"name,omitempty"`       // OPTIONAL, the original filename, if you're logged in as the image owner
	Section     string     `json:"section"`              // If the image has been categorized by our backend then this will contain the section the image belongs in. (funny, cats, adviceanimals, wtf, etc)
	Link        string     `json:"link"`                 // The direct link to the the image. (Note: if fetching an animated GIF that was over 20MB in original size, a .gif thumbnail will be returned)
	Gifv        string     `json:"gifv,omitempty"`       // OPTIONAL, The .gifv link. Only available if the image is animated and type is 'image/gif'.
	Mp4         string     `json:"mp4,omitempty"`        // OPTIONAL, The direct link to the .mp4. Only available if the image is animated and type is 'image/gif'.
	Mp4Size     int        `json:"mp4_size,omitempty"`   // OPTIONAL, The Content-Length of the .mp4. Only available if the image is animated and type is 'image/gif'. Note that a zero value (0) is possible if the video has not yet been generated
	Looping     bool       `json:"looping,omitempty"`    // OPTIONAL, Whether the image has a looping animation. Only available if the image is animated and type is 'image/gif'.
	Favorite    bool       `json:"favorite"`             // Indicates if the current user favorited the image. Defaults to false if not signed in.
	Nsfw        bool       `json:"nsfw"`                 // Indicates if the image has been marked as nsfw or not. Defaults to null if information is not available.
	Vote        string     `json:"vote"`                 // The current user's vote on the album. null if not signed in, if the user hasn't voted on it, or if not submitted to the gallery.
	InGallery   bool       `json:"in_gallery"`           // True if the image has been submitted to the gallery, false if otherwise.
	HasSound    bool       `json:"has_sound"`            // Indicates if the video has sound.
	Limit       *RateLimit // Current rate limit
}

// GetImageInfo queries imgur for information on a image
// returns image info, status code of the request, error
func (client *Client) GetImageInfo(id string) (*ImageInfo, int, error) {
	body, rl, err := client.getURL("image/" + id)
	if err != nil {
		return nil, -1, errors.New("Problem getting URL for image info ID " + id + " - " + err.Error())
	}

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
