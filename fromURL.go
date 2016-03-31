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

// GetInfoFromURL tries to query imgur based on information identified in the URL.
// returns image/album info, status code of the request, error
func (client *Client) GetInfoFromURL(url string) (*GenericInfo, int, error) {
	url = strings.TrimSpace(url)
	var ret GenericInfo

	// https://i.imgur.com/<id>.jpg -> image
	if strings.Contains(url, "://i.imgur.com/") {
		start := strings.LastIndex(url, "/") + 1
		end := strings.LastIndex(url, ".")
		if start == -1 || end == -1 || start == end {
			return nil, -1, errors.New("Could not find ID in URL " + url + ". I was going down i.imgur.com path.")
		}
		id := url[start:end]
		client.Log.Debugf("Detected imgur image ID %v. Was going down the i.imgur.com/ path.", id)
		ii, status, err := client.GetGalleryImageInfo(id)
		ret.GImage = ii
		return &ret, status, err
	}

	// https://imgur.com/a/<id> -> album
	if strings.Contains(url, "://imgur.com/a/") {
		start := strings.LastIndex(url, "/") + 1
		if start == -1 {
			return nil, -1, errors.New("Could not find ID in URL " + url + ". I was going down imgur.com/a/ path.")
		}
		id := url[start:]
		client.Log.Debugf("Detected imgur album ID %v. Was going down the imgur.com/a/ path.", id)
		ai, status, err := client.GetAlbumInfo(id)
		ret.Album = ai
		return &ret, status, err
	}

	// https://imgur.com/gallery/<id> len(id) == 5 -> gallery album
	// https://imgur.com/gallery/<id> len(id) == 7 -> gallery image
	if strings.Contains(url, "://imgur.com/gallery/") {
		start := strings.LastIndex(url, "/") + 1
		if start == -1 {
			return nil, -1, errors.New("Could not find ID in URL " + url + ". I was going down imgur.com/gallery/ path.")
		}
		id := url[start:]
		client.Log.Debugf("Detected imgur gallery ID %v. Was going down the imgur.com/gallery/ path.", id)
		if len(id) == 5 {
			client.Log.Debugf("Detected imgur gallery album.")
			ai, status, err := client.GetGalleryAlbumInfo(id)
			ret.GAlbum = ai
			return &ret, status, err
		}

		ii, status, err := client.GetGalleryImageInfo(id)
		ret.GImage = ii
		return &ret, status, err
	}

	// https://imgur.com/<id> -> image
	if strings.Contains(url, "://imgur.com/") {
		start := strings.LastIndex(url, "/") + 1
		if start == -1 {
			return nil, -1, errors.New("Could not find ID in URL " + url + ". I was going down imgur.com/ path.")
		}
		id := url[start:]
		client.Log.Debugf("Detected imgur image ID %v. Was going down the imgur.com/ path.", id)
		ii, status, err := client.GetGalleryImageInfo(id)
		ret.GImage = ii
		return &ret, status, err
	}

	return nil, -1, errors.New("URL pattern matching for URL " + url + " failed.")
}
