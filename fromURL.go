package imgur

import (
	"errors"
	"strings"
)

// GenericInfo is returned from functions for which the final result type is not known beforehand.
// Only one pointer is != nil
type GenericInfo struct {
	Image  *ImageInfo
	Album  *AlbumInfo
	GImage *GalleryImageInfo
	GAlbum *GalleryAlbumInfo
	Limit  *RateLimit
}

var directURLPatterns = []string{
	"://i.imgur.com/",
	"://i.imgur.io/",
}

var albumURLPatterns = []string{
	"://imgur.com/a/",
	"://m.imgur.com/a/",
	"://imgur.io/a/",
	"://m.imgur.io/a/",
}

var galleryURLPatterns = []string{
	"://imgur.com/gallery/",
	"://m.imgur.com/gallery/",
	"://imgur.io/gallery/",
	"://m.imgur.io/gallery/",
}

var imageURLPatterns = []string{
	"://imgur.com/",
	"://m.imgur.com/",
	"://imgur.io/",
	"://m.imgur.io/",
}

func matchesSlice(url string, validFormats []string) bool {
	for _, format := range validFormats {
		if strings.Contains(url, format) {
			return true
		}
	}

	return false
}

// GetInfoFromURL tries to query imgur based on information identified in the URL.
// returns image/album info, status code of the request, error
func (client *Client) GetInfoFromURL(url string) (*GenericInfo, int, error) {
	url = strings.TrimSpace(url)

	// https://i.imgur.com/<id>.jpg -> image
	if matchesSlice(url, directURLPatterns) {
		return client.directImageURL(url)
	}

	// https://imgur.com/a/<id> -> album
	if matchesSlice(url, albumURLPatterns) {
		return client.albumURL(url)
	}

	// https://imgur.com/gallery/<id> -> gallery album
	if matchesSlice(url, galleryURLPatterns) {
		return client.galleryURL(url)
	}

	// https://imgur.com/<id> -> image
	if matchesSlice(url, imageURLPatterns) {
		return client.imageURL(url)
	}

	return nil, -1, errors.New("URL pattern matching for URL " + url + " failed.")
}

func (client *Client) directImageURL(url string) (*GenericInfo, int, error) {
	var ret GenericInfo
	start := strings.LastIndex(url, "/") + 1
	end := strings.LastIndex(url, ".")
	if start+1 >= end {
		return nil, -1, errors.New("Could not find ID in URL " + url + ". I was going down i.imgur.com path.")
	}
	id := url[start:end]
	client.Log.Debugf("Detected imgur image ID %v. Was going down the i.imgur.com/ path.", id)
	gii, status, err := client.GetGalleryImageInfo(id)
	if err == nil && status < 400 {
		ret.GImage = gii
	} else {
		var ii *ImageInfo
		ii, status, err = client.GetImageInfo(id)
		ret.Image = ii
	}
	return &ret, status, err
}

func (client *Client) albumURL(url string) (*GenericInfo, int, error) {
	var ret GenericInfo

	start := strings.LastIndex(url, "/") + 1
	end := strings.LastIndex(url, "?")
	if end == -1 {
		end = len(url)
	}
	id := url[start:end]
	if id == "" {
		return nil, -1, errors.New("Could not find ID in URL " + url + ". I was going down imgur.com/a/ path.")
	}
	client.Log.Debugf("Detected imgur album ID %v. Was going down the imgur.com/a/ path.", id)
	ai, status, err := client.GetAlbumInfo(id)
	ret.Album = ai
	return &ret, status, err
}

func (client *Client) galleryURL(url string) (*GenericInfo, int, error) {
	var ret GenericInfo

	start := strings.LastIndex(url, "/") + 1
	end := strings.LastIndex(url, "?")
	if end == -1 {
		end = len(url)
	}
	id := url[start:end]
	if id == "" {
		return nil, -1, errors.New("Could not find ID in URL " + url + ". I was going down imgur.com/gallery/ path.")
	}
	client.Log.Debugf("Detected imgur gallery ID %v. Was going down the imgur.com/gallery/ path.", id)
	ai, status, err := client.GetGalleryAlbumInfo(id)
	if err == nil && status < 400 {
		ret.GAlbum = ai
		return &ret, status, err
	}
	// fallback to GetGalleryImageInfo
	client.Log.Debugf("Failed to retrieve imgur gallery album. Attempting to retrieve imgur gallery image. err: %v status: %d", err, status)
	ii, status, err := client.GetGalleryImageInfo(id)
	ret.GImage = ii
	return &ret, status, err
}

func (client *Client) imageURL(url string) (*GenericInfo, int, error) {
	var ret GenericInfo

	start := strings.LastIndex(url, "/") + 1
	end := strings.LastIndex(url, "?")
	if end == -1 {
		end = len(url)
	}
	id := url[start:end]
	if id == "" {
		return nil, -1, errors.New("Could not find ID in URL " + url + ". I was going down imgur.com/ path.")
	}
	client.Log.Debugf("Detected imgur image ID %v. Was going down the imgur.com/ path.", id)
	ii, status, err := client.GetGalleryImageInfo(id)
	if err == nil && status < 400 {
		ret.GImage = ii

		return &ret, status, err
	}

	i, st, err := client.GetImageInfo(id)
	ret.Image = i
	return &ret, st, err
}
