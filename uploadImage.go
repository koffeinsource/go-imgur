package imgur

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

// UploadImage uploads the image to imgur
// image                Can be a binary file, base64 data, or a URL for an image. (up to 10MB)
// album       optional The id of the album you want to add the image to.
//                      For anonymous albums, album should be the deletehash that is returned at creation.
// dtype                The type of the file that's being sent; file, base64 or URL
// title       optional The title of the image.
// description optional The description of the image.
// returns image info, status code of the upload, error
func (client *Client) UploadImage(image []byte, album string, dtype string, title string, description string) (*ImageInfo, int, error) {
	if image == nil {
		return nil, -1, errors.New("Invalid image.")
	}
	if dtype != "file" && dtype != "base64" && dtype != "URL" {
		return nil, -1, errors.New("Passed invalid dtype: " + dtype + ". Please use file/base64/URL.")
	}

	form := createUploadForm(image, album, dtype, title, description)

	URL := apiEndpoint + "image"
	req, err := http.NewRequest("POST", URL, bytes.NewBufferString(form.Encode()))
	client.Log.Infof("Posting to URL %v\n", URL)
	if err != nil {
		return nil, -1, errors.New("Could create request for " + URL + " - " + err.Error())
	}

	req.Header.Add("Authorization", "Client-ID "+client.ImgurClientID)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, -1, errors.New("Could not post " + URL + " - " + err.Error())
	}
	defer res.Body.Close()

	// Read the whole body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, -1, errors.New("Problem reading the body for " + URL + " - " + err.Error())
	}

	// client.Log.Debugf("%v\n", string(body[:]))

	dec := json.NewDecoder(bytes.NewReader(body))
	var img imageInfoDataWrapper
	if err = dec.Decode(&img); err != nil {
		return nil, -1, errors.New("Problem decoding json result from image upload - " + err.Error())
	}

	if !img.Success {
		return nil, img.Status, errors.New("Upload to imgur failed with status: " + strconv.Itoa(img.Status))
	}

	rl, err := extractRateLimits(res.Header)
	if err != nil {
		client.Log.Infof("Problem with extracting rate limits: %v", err)
	} else {
		img.Ii.Limit = rl
	}

	return img.Ii, img.Status, nil
}

func createUploadForm(image []byte, album string, dtype string, title string, description string) url.Values {
	form := url.Values{}

	form.Add("image", string(image[:]))
	form.Add("type", dtype)

	if album != "" {
		form.Add("album", album)
	}
	if title != "" {
		form.Add("title", title)
	}
	if description != "" {
		form.Add("description", description)
	}

	return form
}

// UploadImageFromFile uploads a file given by the filename string to imgur.
func (client *Client) UploadImageFromFile(filename string, album string, title string, description string) (*ImageInfo, int, error) {
	client.Log.Infof("*** IMAGE UPLOAD ***\n")
	f, err := os.Open(filename)
	if err != nil {
		return nil, 500, fmt.Errorf("Could not open file %v - Error: %v", filename, err)
	}
	defer f.Close()
	fileinfo, err := f.Stat()
	if err != nil {
		return nil, 500, fmt.Errorf("Could not stat file %v - Error: %v", filename, err)
	}
	size := fileinfo.Size()
	b := make([]byte, size)
	n, err := f.Read(b)
	if err != nil || int64(n) != size {
		return nil, 500, fmt.Errorf("Could not read file %v - Error: %v", filename, err)
	}

	//base := base64.StdEncoding.EncodeToString(b)

	return client.UploadImage(b, album, "file", title, description)
}
