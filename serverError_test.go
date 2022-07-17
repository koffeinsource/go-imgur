package imgur

import "testing"

func TestImgurNotSuccess(t *testing.T) {
	httpC, server := testHTTPClientJSON("{\"data\": {}, \"success\": false, \"status\": 200 }")
	defer server.Close()

	client, _ := NewClient(httpC, "testing", "")

	_, err := client.GetRateLimit()

	if err == nil {
		t.Error("GetRateLimit() should have failed, but didn't")
	}

	_, _, err = client.GetImageInfo("asd")

	if err == nil {
		t.Error("GetImageInfo() should have failed, but didn't")
	}

	_, _, err = client.GetAlbumInfo("asd")

	if err == nil {
		t.Error("GetAlbumInfo() should have failed, but didn't")
	}

	_, _, err = client.GetGalleryAlbumInfo("asd")

	if err == nil {
		t.Error("GetGalleryAlbumInfo() should have failed, but didn't")
	}

	_, _, err = client.GetGalleryImageInfo("asd")

	if err == nil {
		t.Error("GetGalleryImageInfo() should have failed, but didn't")
	}

	_, _, err = client.GetInfoFromURL("asd")

	if err == nil {
		t.Error("GetInfoFromURL() should have failed, but didn't")
	}

	var im []byte
	_, _, err = client.UploadImage(im, "", "file", "t", "d")

	if err == nil {
		t.Error("UploadImage() should have failed, but didn't")
	}
}

func TestJsonError(t *testing.T) {
	httpC, server := testHTTPClientInvalidJSON()
	defer server.Close()

	client, _ := NewClient(httpC, "testing", "")

	img, _, err := client.GetImageInfo("asd")

	if err == nil || img != nil {
		t.Error("GetImageInfo() should have failed, but didn't")
	}

	ab, _, err := client.GetAlbumInfo("asd")

	if err == nil || ab != nil {
		t.Error("GetAlbumInfo() should have failed, but didn't")
	}

	gab, _, err := client.GetGalleryAlbumInfo("asd")

	if err == nil || gab != nil {
		t.Error("GetGalleryAlbumInfo() should have failed, but didn't")
	}

	gim, _, err := client.GetGalleryImageInfo("asd")

	if err == nil || gim != nil {
		t.Error("GetGalleryImageInfo() should have failed, but didn't")
	}

	ge, _, err := client.GetInfoFromURL("asd")

	if err == nil || ge != nil {
		t.Error("GetInfoFromURL() should have failed, but didn't")
	}

	var im []byte
	img, _, err = client.UploadImage(im, "", "file", "t", "d")

	if err == nil || img != nil {
		t.Error("UploadImage() should have failed, but didn't")
	}

}

func TestServerError(t *testing.T) {
	httpC, server := testHTTPClient500()
	defer server.Close()

	client, _ := NewClient(httpC, "testing", "")

	_, err := client.GetRateLimit()

	if err == nil {
		t.Error("GetRateLimit() should have failed, but didn't")
	}

	_, _, err = client.GetImageInfo("asd")

	if err == nil {
		t.Error("GetImageInfo() should have failed, but didn't")
	}

	_, _, err = client.GetAlbumInfo("asd")

	if err == nil {
		t.Error("GetAlbumInfo() should have failed, but didn't")
	}

	_, _, err = client.GetGalleryAlbumInfo("asd")

	if err == nil {
		t.Error("GetGalleryAlbumInfo() should have failed, but didn't")
	}

	_, _, err = client.GetGalleryImageInfo("asd")

	if err == nil {
		t.Error("GetGalleryImageInfo() should have failed, but didn't")
	}

	_, _, err = client.GetInfoFromURL("asd")

	if err == nil {
		t.Error("GetInfoFromURL() should have failed, but didn't")
	}

	var im []byte
	_, _, err = client.UploadImage(im, "", "file", "t", "d")

	if err == nil {
		t.Error("UploadImage() should have failed, but didn't")
	}
}

func TestImgurError(t *testing.T) {
	httpC, server := testHTTPClientJSON("{'data' : {}, 'success' : false, 'status'  : 500}")
	defer server.Close()

	client, _ := NewClient(httpC, "testing", "")
	_, err := client.GetRateLimit()

	if err == nil {
		t.Error("GetRateLimit() should have failed, but didn't")
	}

	_, _, err = client.GetImageInfo("asd")

	if err == nil {
		t.Error("GetImageInfo() should have failed, but didn't")
	}

	_, _, err = client.GetAlbumInfo("asd")

	if err == nil {
		t.Error("GetAlbumInfo() should have failed, but didn't")
	}

	_, _, err = client.GetGalleryAlbumInfo("asd")

	if err == nil {
		t.Error("GetGalleryAlbumInfo() should have failed, but didn't")
	}

	_, _, err = client.GetGalleryImageInfo("asd")

	if err == nil {
		t.Error("GetGalleryImageInfo() should have failed, but didn't")
	}

	_, _, err = client.GetInfoFromURL("asd")

	if err == nil {
		t.Error("GetInfoFromURL() should have failed, but didn't")
	}

	var im []byte
	_, _, err = client.UploadImage(im, "", "file", "t", "d")

	if err == nil {
		t.Error("UploadImage() should have failed, but didn't")
	}
}

func TestServerDown(t *testing.T) {
	httpC, server := testHTTPClient500()
	server.Close()

	client, _ := NewClient(httpC, "testing", "")
	_, err := client.GetRateLimit()

	if err == nil {
		t.Error("GetRateLimit() should have failed, but didn't")
	}

	_, _, err = client.GetImageInfo("asd")

	if err == nil {
		t.Error("GetImageInfo() should have failed, but didn't")
	}

	_, _, err = client.GetAlbumInfo("asd")

	if err == nil {
		t.Error("GetAlbumInfo() should have failed, but didn't")
	}

	_, _, err = client.GetGalleryAlbumInfo("asd")

	if err == nil {
		t.Error("GetGalleryAlbumInfo() should have failed, but didn't")
	}

	_, _, err = client.GetGalleryImageInfo("asd")

	if err == nil {
		t.Error("GetGalleryImageInfo() should have failed, but didn't")
	}

	_, _, err = client.GetInfoFromURL("asd")

	if err == nil {
		t.Error("GetInfoFromURL() should have failed, but didn't")
	}

	var im []byte
	_, _, err = client.UploadImage(im, "", "file", "t", "d")

	if err == nil {
		t.Error("UploadImage() should have failed, but didn't")
	}
}
