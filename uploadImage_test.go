package imgur

import (
	"net/http"
	"os"
	"testing"
)

const (
	descr = "test upload for the go-imgr library"
	title = "go-imgur test upload"
)

func TestUploadImageErrors(t *testing.T) {
	httpC, server := testHTTPClientJSON("")
	defer server.Close()

	client := createClient(httpC, "testing", "")

	// should fail because of nil image
	ii, _, err := client.UploadImage(nil, "album", "type", "name", "desc")
	if err == nil && ii == nil {
		t.Error("UploadImage() did not result in an error even though it should have.")
		t.Fail()
	}

	img := make([]byte, 5, 5)

	// should fail because of invalid type
	ii, _, err = client.UploadImage(img, "album", "type", "name", "desc")
	if err == nil && ii == nil {
		t.Error("UploadImage() did not result in an error even though it should have.")
		t.Fail()
	}

	ii, _, err = client.UploadImageFromFile("notExistingFile.youtcantseeme", "", title, descr)
	if err == nil && ii == nil {
		t.Error("UploadImageFromFile() did not result in an error even though it should have.")
		t.Fail()
	}
}

func TestUploadImageReal(t *testing.T) {
	key := os.Getenv("IMGURCLIENTID")
	if key == "" {
		t.Skip("IMGURCLIENTID environment variable not set.")
	}
	RapidAPIKey := os.Getenv("RapidAPIKEY")

	client := createClient(new(http.Client), key, RapidAPIKey)

	ii, status, err := client.UploadImageFromFile("test_data/testImage.jpg", "", title, descr)

	if err != nil || ii == nil {
		t.Errorf("UploadImageFromFile() failed with error: %v", err)
		t.FailNow()
	}

	if ii.Description != descr || ii.Title != title {
		t.Fail()
	}

	if status != 200 {
		t.Fail()
	}
}

func TestUploadImageSimulated(t *testing.T) {
	httpC, server := testHTTPClientJSON("{\"data\":{\"id\":\"ClF8rLe\",\"title\":\"" + title + "\",\"description\":\"" + descr + "\",\"datetime\":1451248840,\"type\":\"image\\/jpeg\",\"animated\":false,\"width\":2448,\"height\":3264,\"size\":1071339,\"views\":176,\"bandwidth\":188555664,\"vote\":null,\"favorite\":false,\"nsfw\":null,\"section\":null,\"account_url\":null,\"account_id\":null,\"in_gallery\":false,\"link\":\"https:\\/\\/i.imgur.com\\/ClF8rLe.jpg\"},\"success\":true,\"status\":200}")
	defer server.Close()

	client := createClient(httpC, "testing", "")
	ii, status, err := client.UploadImageFromFile("test_data/testImage.jpg", "ALBUMID", title, descr)

	if err != nil || ii == nil {
		t.Errorf("UploadImageFromFile() failed with error: %v", err)
		t.FailNow()
	}

	if ii.Description != descr || ii.Title != title {
		t.Fail()
	}

	if status != 200 {
		t.Fail()
	}
}
