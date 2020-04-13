package imgur

import (
	"net/http"
	"os"
	"testing"

	"github.com/koffeinsource/go-klogger"
)

func TestGetFromURLAlbumSimulated(t *testing.T) {
	httpC, server := testHTTPClientJSON("{\"data\":{\"id\":\"VZQXk\",\"title\":\"Gianluca Gimini's bikes\",\"description\":null,\"datetime\":1460715031,\"cover\":\"CJCA0gW\",\"cover_width\":1200,\"cover_height\":786,\"account_url\":\"mrcassette\",\"account_id\":157430,\"privacy\":\"public\",\"layout\":\"blog\",\"views\":667581,\"link\":\"https:\\/\\/imgur.com\\/a\\/VZQXk\",\"favorite\":false,\"nsfw\":false,\"section\":\"pics\",\"images_count\":1,\"in_gallery\":true,\"images\":[{\"id\":\"CJCA0gW\",\"title\":null,\"description\":\"by Designer Gianluca Gimini\\nhttps:\\/\\/www.behance.net\\/gallery\\/35437979\\/Velocipedia\",\"datetime\":1460715032,\"type\":\"image\\/jpeg\",\"animated\":false,\"width\":1200,\"height\":786,\"size\":362373,\"views\":4420880,\"bandwidth\":1602007548240,\"vote\":null,\"favorite\":false,\"nsfw\":null,\"section\":null,\"account_url\":null,\"account_id\":null,\"in_gallery\":false,\"link\":\"https:\\/\\/i.imgur.com\\/CJCA0gW.jpg\"}]},\"success\":true,\"status\":200}")
	defer server.Close()

	client := createClient(httpC, "testing", "")
	ge, status, err := client.GetInfoFromURL("https://imgur.com/a/VZQXk")

	if err != nil {
		t.Errorf("GetInfoFromURL() failed with error: %v", err)
		t.FailNow()
	}

	if ge.Album == nil || ge.GAlbum != nil || ge.GImage != nil || ge.Image != nil {
		t.Error("GetInfoFromURL() failed. Returned wrong type.")
		t.FailNow()
	}

	alb := ge.Album

	if alb.Title != "Gianluca Gimini's bikes" || alb.Cover != "CJCA0gW" || alb.CoverWidth != 1200 || alb.CoverHeight != 786 || alb.Link != "https://imgur.com/a/VZQXk" || alb.ImagesCount != 1 || alb.Images[0].ID != "CJCA0gW" {
		t.Fail()
	}

	if status != 200 {
		t.Fail()
	}
}

func TestGetFromURLAlbumReal(t *testing.T) {
	key := os.Getenv("IMGURCLIENTID")
	if key == "" {
		t.Skip("IMGURCLIENTID environment variable not set.")
	}
	RapidAPIKey := os.Getenv("RapidAPIKEY")

	client := createClient(new(http.Client), key, RapidAPIKey)

	ge, status, err := client.GetInfoFromURL("https://imgur.com/a/VZQXk")

	if err != nil {
		t.Errorf("GetInfoFromURL() failed with error: %v", err)
		t.FailNow()
	}

	if ge.Album == nil || ge.GAlbum != nil || ge.GImage != nil || ge.Image != nil {
		t.Error("GetInfoFromURL() failed. Returned wrong type.")
		t.FailNow()
	}

	alb := ge.Album

	if alb.Title != "Gianluca Gimini's bikes" || alb.Cover != "CJCA0gW" || alb.CoverWidth != 1200 || alb.CoverHeight != 786 || alb.Link != "https://imgur.com/a/VZQXk" || alb.ImagesCount != 14 || alb.Images[0].ID != "CJCA0gW" {
		t.Fail()
	}

	if status != 200 {
		t.Fail()
	}
}

func TestGetFromURLAlbumNoID(t *testing.T) {
	httpC, server := testHTTPClient500()
	defer server.Close()
	client := createClient(httpC, "testing", "")

	_, _, err := client.GetInfoFromURL("https://imgur.com/a/")

	if err == nil {
		t.Error("GetInfoFromURL() did not failed but should have.")
		t.FailNow()
	}
}

func TestGetFromURLGalleryNoID(t *testing.T) {
	httpC, server := testHTTPClient500()
	defer server.Close()
	client := createClient(httpC, "testing", "")

	_, _, err := client.GetInfoFromURL("https://imgur.com/gallery/")

	if err == nil {
		t.Error("GetInfoFromURL() did not failed but should have.")
		t.FailNow()
	}
}

func TestGetFromURLGAlbumSimulated(t *testing.T) {
	httpC, server := testHTTPClientJSON("{\"data\":{\"id\":\"VZQXk\",\"title\":\"As it turns out, most people cannot draw a bike.\",\"description\":null,\"datetime\":1460715031,\"cover\":\"CJCA0gW\",\"cover_width\":1200,\"cover_height\":786,\"account_url\":\"mrcassette\",\"account_id\":157430,\"privacy\":\"public\",\"layout\":\"blog\",\"views\":667581,\"link\":\"https:\\/\\/imgur.com\\/a\\/VZQXk\",\"ups\":13704,\"downs\":113,\"favorite\":false,\"nsfw\":false,\"section\":\"pics\",\"images_count\":1,\"in_gallery\":true,\"images\":[{\"id\":\"CJCA0gW\",\"title\":null,\"description\":\"by Designer Gianluca Gimini\\nhttps:\\/\\/www.behance.net\\/gallery\\/35437979\\/Velocipedia\",\"datetime\":1460715032,\"type\":\"image\\/jpeg\",\"animated\":false,\"width\":1200,\"height\":786,\"size\":362373,\"views\":4420880,\"bandwidth\":1602007548240,\"vote\":null,\"favorite\":false,\"nsfw\":null,\"section\":null,\"account_url\":null,\"account_id\":null,\"in_gallery\":false,\"link\":\"https:\\/\\/i.imgur.com\\/CJCA0gW.jpg\"}]},\"success\":true,\"status\":200}")
	defer server.Close()

	client := createClient(httpC, "testing", "")
	ge, status, err := client.GetInfoFromURL("https://imgur.com/gallery/VZQXk")

	if err != nil {
		t.Errorf("GetInfoFromURL() failed with error: %v", err)
		t.FailNow()
	}

	if ge.Album != nil || ge.GAlbum == nil || ge.GImage != nil || ge.Image != nil {
		t.Error("GetInfoFromURL() failed. Returned wrong type.")
		t.FailNow()
	}

	alb := ge.GAlbum

	if alb.Title != "As it turns out, most people cannot draw a bike." || alb.Cover != "CJCA0gW" || alb.CoverWidth != 1200 || alb.CoverHeight != 786 || alb.Link != "https://imgur.com/a/VZQXk" || alb.ImagesCount != 1 || alb.Images[0].ID != "CJCA0gW" || alb.Ups != 13704 || alb.Downs != 113 {
		t.Fail()
	}

	if status != 200 {
		t.Fail()
	}
}

func TestGetFromURLGAlbumReal(t *testing.T) {
	key := os.Getenv("IMGURCLIENTID")
	if key == "" {
		t.Skip("IMGURCLIENTID environment variable not set.")
	}
	RapidAPIKey := os.Getenv("RapidAPIKEY")

	client := createClient(new(http.Client), key, RapidAPIKey)

	ge, status, err := client.GetInfoFromURL("https://imgur.com/gallery/VZQXk")

	if err != nil {
		t.Errorf("GetInfoFromURL() failed with error: %v", err)
		t.FailNow()
	}

	if ge.Album != nil || ge.GAlbum == nil || ge.GImage != nil || ge.Image != nil {
		t.Error("GetInfoFromURL() failed. Returned wrong type.")
		t.FailNow()
	}

	alb := ge.GAlbum

	if alb.Title != "As it turns out, most people cannot draw a bike." || alb.Cover != "CJCA0gW" || alb.CoverWidth != 1200 || alb.CoverHeight != 786 || alb.Link != "https://imgur.com/a/VZQXk" || alb.ImagesCount != 14 || alb.Images[0].ID != "CJCA0gW" {
		t.Fail()
	}

	if status != 200 {
		t.Fail()
	}
}

func TestGetURLGalleryImageSimulated(t *testing.T) {
	httpC, server := testHTTPClientJSON("{\"data\":{\"id\":\"uPI76jY\",\"title\":\"The Tridge. (three way bridge)\",\"description\":null,\"datetime\":1316367003,\"type\":\"image\\/jpeg\",\"animated\":false,\"width\":1700,\"height\":1133,\"size\":268126,\"views\":1342557,\"bandwidth\":359974438182,\"vote\":null,\"favorite\":false,\"nsfw\":false,\"section\":\"pics\",\"account_url\":null,\"account_id\":null,\"in_gallery\":true,\"topic\":null,\"topic_id\":0,\"link\":\"https:\\/\\/i.imgur.com\\/uPI76jY.jpg\",\"comment_count\":90,\"ups\":585,\"downs\":3,\"points\":582,\"score\":1136,\"is_album\":false},\"success\":true,\"status\":200}")
	defer server.Close()

	client := new(Client)
	client.HTTPClient = httpC
	client.Log = new(klogger.CLILogger)
	client.ImgurClientID = "testing"

	ge, status, err := client.GetInfoFromURL("https://imgur.com/gallery/uPI76jY")

	if err != nil {
		t.Errorf("GetInfoFromURL() failed with error: %v", err)
		t.FailNow()
	}

	if ge.Album != nil || ge.GAlbum != nil || ge.GImage == nil || ge.Image != nil {
		t.Error("GetInfoFromURL() failed. Returned wrong type.")
		t.FailNow()
	}

	img := ge.GImage

	if img.Title != "The Tridge. (three way bridge)" || img.Animated != false || img.Bandwidth != 359974438182 || img.Datetime != 1316367003 || img.Description != "" || img.Height != 1133 || img.Width != 1700 || img.ID != "uPI76jY" || img.Link != "https://i.imgur.com/uPI76jY.jpg" || img.Views != 1342557 {
		t.Fail()
	}

	if status != 200 {
		t.Fail()
	}
}

func TestGetURLGalleryImageReal(t *testing.T) {
	key := os.Getenv("IMGURCLIENTID")
	if key == "" {
		t.Skip("IMGURCLIENTID environment variable not set.")
	}
	RapidAPIKey := os.Getenv("RapidAPIKEY")

	client := createClient(new(http.Client), key, RapidAPIKey)

	ge, status, err := client.GetInfoFromURL("https://imgur.com/gallery/uPI76jY")

	if err != nil {
		t.Errorf("GetInfoFromURL() failed with error: %v", err)
		t.FailNow()
	}

	if ge.Album != nil || ge.GAlbum != nil || ge.GImage == nil || ge.Image != nil {
		t.Error("GetInfoFromURL() failed. Returned wrong type.")
		t.FailNow()
	}

	img := ge.GImage

	if img.Title != "An abandoned Chinese fishing village" || img.Animated != false || img.Description != "" || img.Height != 445 || img.Width != 800 || img.ID != "uPI76jY" || img.Link != "https://i.imgur.com/uPI76jY.jpg" {
		t.Fail()
	}

	if status != 200 {
		t.Fail()
	}
}

func TestGetURLImageSimulated(t *testing.T) {
	httpC, server := testHTTPClientJSON("{\"data\":{\"id\":\"ClF8rLe\",\"title\":null,\"description\":null,\"datetime\":1451248840,\"type\":\"image\\/jpeg\",\"animated\":false,\"width\":2448,\"height\":3264,\"size\":1071339,\"views\":176,\"bandwidth\":188555664,\"vote\":null,\"favorite\":false,\"nsfw\":null,\"section\":null,\"account_url\":null,\"account_id\":null,\"in_gallery\":false,\"link\":\"https:\\/\\/i.imgur.com\\/ClF8rLe.jpg\"},\"success\":true,\"status\":200}")
	defer server.Close()

	client := new(Client)
	client.HTTPClient = httpC
	client.Log = new(klogger.CLILogger)
	client.ImgurClientID = "testing"

	ge, status, err := client.GetInfoFromURL("https://imgur.com/ClF8rLe")

	if err != nil {
		t.Errorf("GetInfoFromURL() failed with error: %v", err)
		t.FailNow()
	}

	if ge.Album != nil || ge.GAlbum != nil {
		t.Error("GetInfoFromURL() failed. Returned wrong type.")
		t.FailNow()
	}

	if ge.Image == nil && ge.GImage == nil {
		t.FailNow()
	}

	if ge.Image != nil {
		img := ge.Image

		if img.Animated != false || img.Bandwidth != 188555664 || img.Datetime != 1451248840 || img.Description != "" || img.Height != 3264 || img.Width != 2448 || img.ID != "ClF8rLe" || img.Link != "https://i.imgur.com/ClF8rLe.jpg" || img.Views != 176 {
			t.Fail()
		}
	}

	if ge.GImage != nil {
		img := ge.GImage

		if img.Animated != false || img.Bandwidth != 188555664 || img.Datetime != 1451248840 || img.Description != "" || img.Height != 3264 || img.Width != 2448 || img.ID != "ClF8rLe" || img.Link != "https://i.imgur.com/ClF8rLe.jpg" || img.Views != 176 {
			t.Fail()
		}
	}

	if status != 200 {
		t.Fail()
	}
}

func TestGetURLImageReal(t *testing.T) {
	key := os.Getenv("IMGURCLIENTID")
	if key == "" {
		t.Skip("IMGURCLIENTID environment variable not set.")
	}
	RapidAPIKey := os.Getenv("RapidAPIKEY")

	client := createClient(new(http.Client), key, RapidAPIKey)

	ge, status, err := client.GetInfoFromURL("https://imgur.com/ClF8rLe")

	if err != nil {
		t.Errorf("GetInfoFromURL() failed with error: %v", err)
		t.FailNow()
	}

	if ge.Album != nil || ge.GAlbum != nil {
		t.Error("GetInfoFromURL() failed. Returned wrong type.")
		t.FailNow()
	}

	if ge.Image == nil && ge.GImage == nil {
		t.FailNow()
	}

	if ge.Image != nil {
		img := ge.Image

		if img.Animated != false || img.Datetime != 1451248840 || img.Description != "" || img.Height != 3264 || img.Width != 2448 || img.ID != "ClF8rLe" || img.Link != "https://i.imgur.com/ClF8rLe.jpg" {
			t.Fail()
		}
	}

	if ge.GImage != nil {
		img := ge.GImage

		if img.Animated != false || img.Datetime != 1451248840 || img.Description != "" || img.Height != 3264 || img.Width != 2448 || img.ID != "ClF8rLe" || img.Link != "https://i.imgur.com/ClF8rLe.jpg" {
			t.Fail()
		}
	}

	if status != 200 {
		t.Fail()
	}
}

func TestGetFromURLImageNoID(t *testing.T) {
	httpC, server := testHTTPClient500()
	defer server.Close()

	client := createClient(httpC, "testing", "")
	_, _, err := client.GetInfoFromURL("https://imgur.com/")

	if err == nil {
		t.Error("GetInfoFromURL() did not failed but should have.")
		t.FailNow()
	}
}

func TestGetURLDirectImageSimulated(t *testing.T) {
	httpC, server := testHTTPClientJSON("{\"data\":{\"id\":\"ClF8rLe\",\"title\":null,\"description\":null,\"datetime\":1451248840,\"type\":\"image\\/jpeg\",\"animated\":false,\"width\":2448,\"height\":3264,\"size\":1071339,\"views\":176,\"bandwidth\":188555664,\"vote\":null,\"favorite\":false,\"nsfw\":null,\"section\":null,\"account_url\":null,\"account_id\":null,\"in_gallery\":false,\"link\":\"https:\\/\\/i.imgur.com\\/ClF8rLe.jpg\"},\"success\":true,\"status\":200}")
	defer server.Close()

	client := createClient(httpC, "testing", "")
	ge, status, err := client.GetInfoFromURL("https://i.imgur.com/ClF8rLe.jpg")

	if err != nil {
		t.Errorf("GetInfoFromURL() failed with error: %v", err)
		t.FailNow()
	}

	if ge.Album != nil || ge.GAlbum != nil {
		t.Error("GetInfoFromURL() failed. Returned wrong type.")
		t.FailNow()
	}

	if ge.Image == nil && ge.GImage == nil {
		t.FailNow()
	}

	if ge.Image != nil {
		img := ge.Image

		if img.Animated != false || img.Bandwidth != 188555664 || img.Datetime != 1451248840 || img.Description != "" || img.Height != 3264 || img.Width != 2448 || img.ID != "ClF8rLe" || img.Link != "https://i.imgur.com/ClF8rLe.jpg" || img.Views != 176 {
			t.Fail()
		}
	}

	if ge.GImage != nil {
		img := ge.GImage

		if img.Animated != false || img.Bandwidth != 188555664 || img.Datetime != 1451248840 || img.Description != "" || img.Height != 3264 || img.Width != 2448 || img.ID != "ClF8rLe" || img.Link != "https://i.imgur.com/ClF8rLe.jpg" || img.Views != 176 {
			t.Fail()
		}
	}

	if status != 200 {
		t.Fail()
	}
}

func TestGetURLDirectImageReal(t *testing.T) {
	key := os.Getenv("IMGURCLIENTID")
	if key == "" {
		t.Skip("IMGURCLIENTID environment variable not set.")
	}
	RapidAPIKey := os.Getenv("RapidAPIKEY")

	client := createClient(new(http.Client), key, RapidAPIKey)

	ge, status, err := client.GetInfoFromURL("https://i.imgur.com/ClF8rLe.jpg")

	if err != nil {
		t.Errorf("GetInfoFromURL() failed with error: %v", err)
		t.FailNow()
	}

	if ge.Album != nil || ge.GAlbum != nil {
		t.Error("GetInfoFromURL() failed. Returned wrong type.")
		t.FailNow()
	}

	if ge.Image == nil && ge.GImage == nil {
		t.FailNow()
	}

	if ge.Image != nil {
		img := ge.Image

		if img.Animated != false || img.Datetime != 1451248840 || img.Description != "" || img.Height != 3264 || img.Width != 2448 || img.ID != "ClF8rLe" || img.Link != "https://i.imgur.com/ClF8rLe.jpg" {
			t.Fail()
		}
	}

	if ge.GImage != nil {
		img := ge.GImage

		if img.Animated != false || img.Datetime != 1451248840 || img.Description != "" || img.Height != 3264 || img.Width != 2448 || img.ID != "ClF8rLe" || img.Link != "https://i.imgur.com/ClF8rLe.jpg" {
			t.Fail()
		}
	}

	if status != 200 {
		t.Fail()
	}
}

func TestGetFromURLDirectImageNoID(t *testing.T) {
	httpC, server := testHTTPClient500()
	defer server.Close()

	client := createClient(httpC, "testing", "")
	_, _, err := client.GetInfoFromURL("https://i.imgur.com/")

	if err == nil {
		t.Error("GetInfoFromURL() did not failed but should have.")
		t.FailNow()
	}
}
