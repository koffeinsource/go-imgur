package imgur

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/koffeinsource/go-klogger"
)

func TestGetFromURLAlbumSimulated(t *testing.T) {
	httpC, server := testHTTPClientJSON("{\"data\":{\"id\":\"VZQXk\",\"title\":\"Gianluca Gimini's bikes\",\"description\":null,\"datetime\":1460715031,\"cover\":\"CJCA0gW\",\"cover_width\":1200,\"cover_height\":786,\"account_url\":\"mrcassette\",\"account_id\":157430,\"privacy\":\"public\",\"layout\":\"blog\",\"views\":667581,\"link\":\"https:\\/\\/imgur.com\\/a\\/VZQXk\",\"favorite\":false,\"nsfw\":false,\"section\":\"pics\",\"images_count\":1,\"in_gallery\":true,\"images\":[{\"id\":\"CJCA0gW\",\"title\":null,\"description\":\"by Designer Gianluca Gimini\\nhttps:\\/\\/www.behance.net\\/gallery\\/35437979\\/Velocipedia\",\"datetime\":1460715032,\"type\":\"image\\/jpeg\",\"animated\":false,\"width\":1200,\"height\":786,\"size\":362373,\"views\":4420880,\"bandwidth\":1602007548240,\"vote\":null,\"favorite\":false,\"nsfw\":null,\"section\":null,\"account_url\":null,\"account_id\":null,\"in_gallery\":false,\"link\":\"https:\\/\\/i.imgur.com\\/CJCA0gW.jpg\"}]},\"success\":true,\"status\":200}")
	defer server.Close()

	client, _ := NewClient(httpC, "testing", "")
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

	client, _ := NewClient(new(http.Client), key, RapidAPIKey)

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
	client, _ := NewClient(httpC, "testing", "")

	_, _, err := client.GetInfoFromURL("https://imgur.com/a/")

	if err == nil {
		t.Error("GetInfoFromURL() did not failed but should have.")
		t.FailNow()
	}
}

func TestGetFromURLGalleryNoID(t *testing.T) {
	httpC, server := testHTTPClient500()
	defer server.Close()
	client, _ := NewClient(httpC, "testing", "")

	_, _, err := client.GetInfoFromURL("https://imgur.com/gallery/")

	if err == nil {
		t.Error("GetInfoFromURL() did not failed but should have.")
		t.FailNow()
	}
}

func TestGetFromURLGAlbumSimulated(t *testing.T) {
	httpC, server := testHTTPClientJSON("{\"data\":{\"id\":\"VZQXk\",\"title\":\"As it turns out, most people cannot draw a bike.\",\"description\":null,\"datetime\":1460715031,\"cover\":\"CJCA0gW\",\"cover_width\":1200,\"cover_height\":786,\"account_url\":\"mrcassette\",\"account_id\":157430,\"privacy\":\"public\",\"layout\":\"blog\",\"views\":667581,\"link\":\"https:\\/\\/imgur.com\\/a\\/VZQXk\",\"ups\":13704,\"downs\":113,\"favorite\":false,\"nsfw\":false,\"section\":\"pics\",\"images_count\":1,\"in_gallery\":true,\"images\":[{\"id\":\"CJCA0gW\",\"title\":null,\"description\":\"by Designer Gianluca Gimini\\nhttps:\\/\\/www.behance.net\\/gallery\\/35437979\\/Velocipedia\",\"datetime\":1460715032,\"type\":\"image\\/jpeg\",\"animated\":false,\"width\":1200,\"height\":786,\"size\":362373,\"views\":4420880,\"bandwidth\":1602007548240,\"vote\":null,\"favorite\":false,\"nsfw\":null,\"section\":null,\"account_url\":null,\"account_id\":null,\"in_gallery\":false,\"link\":\"https:\\/\\/i.imgur.com\\/CJCA0gW.jpg\"}]},\"success\":true,\"status\":200}")
	defer server.Close()

	client, _ := NewClient(httpC, "testing", "")
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

	client, _ := NewClient(new(http.Client), key, RapidAPIKey)

	tests := []struct {
		galleryURL string
		expected   map[string]interface{}
	}{
		{
			galleryURL: "https://imgur.com/gallery/VZQXk",
			expected: map[string]interface{}{
				"title":        "As it turns out, most people cannot draw a bike.",
				"cover":        "CJCA0gW",
				"coverWidth":   1200,
				"coverHeight":  786,
				"link":         "https://imgur.com/a/VZQXk",
				"imagesCount":  14,
				"firstImageID": "CJCA0gW",
			},
		},
		{
			galleryURL: "https://imgur.com/gallery/t6l1GiW",
			expected: map[string]interface{}{
				"title":        "Funny Random Meme and Twitter Dump",
				"cover":        "60wTouU",
				"coverWidth":   1242,
				"coverHeight":  1512,
				"link":         "https://imgur.com/a/t6l1GiW",
				"imagesCount":  50,
				"firstImageID": "60wTouU",
			},
		},
	}
	for _, test := range tests {
		ge, status, err := client.GetInfoFromURL(test.galleryURL)
		if err != nil {
			t.Errorf("GetInfoFromURL() failed with error: %v", err)
			t.FailNow()
		}
		if ge.Album != nil || ge.GAlbum == nil || ge.GImage != nil || ge.Image != nil {
			t.Error("GetInfoFromURL() failed. Returned wrong type.")
			t.FailNow()
		}
		if ge.GAlbum.Title != test.expected["title"] {
			t.Errorf("title mismatch: %s != %s", ge.GAlbum.Title, test.expected["title"])
			t.Fail()
		}
		if ge.GAlbum.Cover != test.expected["cover"] {
			t.Errorf("cover mismatch: %s != %s", ge.GAlbum.Cover, test.expected["cover"])
			t.Fail()
		}
		if ge.GAlbum.CoverWidth != test.expected["coverWidth"] {
			t.Errorf("coverWidth mismatch: %d != %d", ge.GAlbum.CoverWidth, test.expected["coverWidth"])
			t.Fail()
		}
		if ge.GAlbum.CoverHeight != test.expected["coverHeight"] {
			t.Errorf("coverHeight mismatch: %d != %d", ge.GAlbum.CoverHeight, test.expected["coverHeight"])
			t.Fail()
		}
		if ge.GAlbum.Link != test.expected["link"] {
			t.Errorf("link mismatch: %s != %s", ge.GAlbum.Link, test.expected["link"])
			t.Fail()
		}
		if ge.GAlbum.ImagesCount != test.expected["imagesCount"] {
			t.Errorf("imagesCount mismatch: %d != %d", ge.GAlbum.ImagesCount, test.expected["imagesCount"])
			t.Fail()
		}
		if ge.GAlbum.Images[0].ID != test.expected["firstImageID"] {
			t.Errorf("firstImageID mismatch: %s != %s", ge.GAlbum.Images[0].ID, test.expected["firstImageID"])
			t.Fail()
		}
		if status != http.StatusOK {
			t.Errorf("status mismatch: %d != %d", status, http.StatusOK)
			t.Fail()
		}
	}
}

func TestGetURLGalleryImageReal(t *testing.T) {
	key := os.Getenv("IMGURCLIENTID")
	if key == "" {
		t.Skip("IMGURCLIENTID environment variable not set.")
	}
	RapidAPIKey := os.Getenv("RapidAPIKEY")

	client, _ := NewClient(new(http.Client), key, RapidAPIKey)

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
	g := NewWithT(t)
	RegisterFailHandler(Fail)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	client, _ := NewClient(&http.Client{}, "testing", "")
	client.Log = new(klogger.CLILogger)
	client.ImgurClientID = "testing"

	responseString := "{\"data\":{\"id\":\"ClF8rLe\",\"title\":null,\"description\":null,\"datetime\":1451248840,\"type\":\"image\\/jpeg\",\"animated\":false,\"width\":2448,\"height\":3264,\"size\":1071339,\"views\":176,\"bandwidth\":188555664,\"vote\":null,\"favorite\":false,\"nsfw\":null,\"section\":null,\"account_url\":null,\"account_id\":null,\"in_gallery\":false,\"link\":\"https:\\/\\/i.imgur.com\\/ClF8rLe.jpg\"},\"success\":true,\"status\":200}"
	MockStringResp("https://api.imgur.com/3/image/ClF8rLe", http.MethodGet, responseString, nil)

	ge, status, err := client.GetInfoFromURL("https://imgur.com/ClF8rLe")
	g.Expect(err).To(BeNil())
	g.Expect(status).To(Equal(200))

	info := httpmock.GetCallCountInfo()
	g.Expect(httpmock.GetTotalCallCount()).To(Equal(1))
	// get the amount of calls for the registered responder
	g.Expect(info["GET https://api.imgur.com/3/image/ClF8rLe"]).To(Equal(1))

	g.Expect(ge.Album).To(BeNil())
	g.Expect(ge.GAlbum).To(BeNil())

	g.Expect(ge.Image).NotTo(BeNil())
	// the mock response is for an Image, not a GalleryImage, so GImage will be nil
	g.Expect(ge.GImage).To(BeNil())

	img := ge.Image
	g.Expect(img.Animated).To(BeFalse())
	g.Expect(img.Bandwidth).To(Equal(188555664))
	g.Expect(img.Datetime).To(Equal(1451248840))
	g.Expect(img.Description).To(Equal(""))
	g.Expect(img.Height).To(Equal(3264))
	g.Expect(img.Width).To(Equal(2448))
	g.Expect(img.ID).To(Equal("ClF8rLe"))
	g.Expect(img.Link).To(Equal("https://i.imgur.com/ClF8rLe.jpg"))
	g.Expect(img.Views).To(Equal(176))
}

// Other tests do not correctly simulate and validate the correct path was chosen
// use this test as template for all other tests
func TestGetURLImageSimulatedWithExtension(t *testing.T) {
	g := NewWithT(t)
	RegisterFailHandler(Fail)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	client, _ := NewClient(&http.Client{}, "testing", "")
	client.Log = new(klogger.CLILogger)
	client.ImgurClientID = "testing"

	responseString := "{\"data\":{\"id\":\"ClF8rLe\",\"title\":null,\"description\":null,\"datetime\":1451248840,\"type\":\"image\\/jpeg\",\"animated\":false,\"width\":2448,\"height\":3264,\"size\":1071339,\"views\":176,\"bandwidth\":188555664,\"vote\":null,\"favorite\":false,\"nsfw\":null,\"section\":null,\"account_url\":null,\"account_id\":null,\"in_gallery\":false,\"link\":\"https:\\/\\/i.imgur.com\\/ClF8rLe.jpg\"},\"success\":true,\"status\":200}"
	MockStringResp("https://api.imgur.com/3/image/ClF8rLe", http.MethodGet, responseString, nil)

	ge, status, err := client.GetInfoFromURL("https://imgur.com/ClF8rLe.jpg")
	g.Expect(ge).NotTo(BeNil())
	g.Expect(err).To(BeNil())
	g.Expect(status).To(Equal(200))

	info := httpmock.GetCallCountInfo()
	g.Expect(httpmock.GetTotalCallCount()).To(Equal(1))
	// get the amount of calls for the registered responder
	g.Expect(info["GET https://api.imgur.com/3/image/ClF8rLe"]).To(Equal(1))

	g.Expect(ge.Album).To(BeNil())
	g.Expect(ge.GAlbum).To(BeNil())

	g.Expect(ge.Image).NotTo(BeNil())
	// the mock response is for an Image, not a GalleryImage, so GImage will be nil
	g.Expect(ge.GImage).To(BeNil())

	img := ge.Image
	g.Expect(img.Animated).To(BeFalse())
	g.Expect(img.Bandwidth).To(Equal(188555664))
	g.Expect(img.Datetime).To(Equal(1451248840))
	g.Expect(img.Description).To(Equal(""))
	g.Expect(img.Height).To(Equal(3264))
	g.Expect(img.Width).To(Equal(2448))
	g.Expect(img.ID).To(Equal("ClF8rLe"))
	g.Expect(img.Link).To(Equal("https://i.imgur.com/ClF8rLe.jpg"))
	g.Expect(img.Views).To(Equal(176))
}

func TestGetURLImageSimulatedWithExtensionNotFound(t *testing.T) {
	g := NewWithT(t)
	RegisterFailHandler(Fail)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	client, _ := NewClient(&http.Client{}, "testing", "")
	client.Log = new(klogger.CLILogger)
	client.ImgurClientID = "testing"

	MockStringResp("https://api.imgur.com/3/image/ClF8rLe", http.MethodGet, "", nil, 404)

	ge, status, err := client.GetInfoFromURL("https://imgur.com/ClF8rLe.jpg")
	g.Expect(err).NotTo(BeNil())
	g.Expect(status).To(Equal(404))
	g.Expect(ge).To(BeNil())

	info := httpmock.GetCallCountInfo()
	g.Expect(httpmock.GetTotalCallCount()).To(Equal(1))
	// get the amount of calls for the registered responder
	g.Expect(info["GET https://api.imgur.com/3/image/ClF8rLe"]).To(Equal(1))
}

func TestGetURLImageSimulatedWithExtensionMoved302(t *testing.T) {
	g := NewWithT(t)
	RegisterFailHandler(Fail)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	client, _ := NewClient(&http.Client{}, "testing", "")
	client.Log = new(klogger.CLILogger)
	client.ImgurClientID = "testing"

	headers := make(map[string]string)
	headers["Location"] = "moved_url"
	MockStringResp("https://api.imgur.com/3/image/ClF8rLe", http.MethodGet, "", headers, 302)
	responseString := "{\"data\":{\"id\":\"ClF8rLe\",\"title\":null,\"description\":null,\"datetime\":1451248840,\"type\":\"image\\/jpeg\",\"animated\":false,\"width\":2448,\"height\":3264,\"size\":1071339,\"views\":176,\"bandwidth\":188555664,\"vote\":null,\"favorite\":false,\"nsfw\":null,\"section\":null,\"account_url\":null,\"account_id\":null,\"in_gallery\":false,\"link\":\"https:\\/\\/i.imgur.com\\/ClF8rLe.jpg\"},\"success\":true,\"status\":200}"
	MockStringResp("https://api.imgur.com/3/image/moved_url", http.MethodGet, responseString, nil)

	ge, status, err := client.GetInfoFromURL("https://imgur.com/ClF8rLe.jpg")
	g.Expect(ge).NotTo(BeNil())
	g.Expect(err).To(BeNil())
	g.Expect(status).To(Equal(200))

	info := httpmock.GetCallCountInfo()
	g.Expect(httpmock.GetTotalCallCount()).To(Equal(2))
	// get the amount of calls for the registered responder
	g.Expect(info["GET https://api.imgur.com/3/image/ClF8rLe"]).To(Equal(1))
	g.Expect(info["GET https://api.imgur.com/3/image/moved_url"]).To(Equal(1))

	g.Expect(ge.Image).NotTo(BeNil())
	// the mock response is for an Image, not a GalleryImage, so GImage will be nil
	g.Expect(ge.GImage).To(BeNil())

	img := ge.Image
	g.Expect(img.Animated).To(BeFalse())
	g.Expect(img.Bandwidth).To(Equal(188555664))
	g.Expect(img.Datetime).To(Equal(1451248840))
	g.Expect(img.Description).To(Equal(""))
	g.Expect(img.Height).To(Equal(3264))
	g.Expect(img.Width).To(Equal(2448))
	g.Expect(img.ID).To(Equal("ClF8rLe"))
	g.Expect(img.Link).To(Equal("https://i.imgur.com/ClF8rLe.jpg"))
	g.Expect(img.Views).To(Equal(176))
}

func TestGetURLImageReal(t *testing.T) {
	key := os.Getenv("IMGURCLIENTID")
	if key == "" {
		t.Skip("IMGURCLIENTID environment variable not set.")
	}
	RapidAPIKey := os.Getenv("RapidAPIKEY")

	client, _ := NewClient(new(http.Client), key, RapidAPIKey)

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

	client, _ := NewClient(httpC, "testing", "")
	_, _, err := client.GetInfoFromURL("https://imgur.com/")

	if err == nil {
		t.Error("GetInfoFromURL() did not failed but should have.")
		t.FailNow()
	}
}

func TestGetURLDirectImageSimulated(t *testing.T) {
	httpC, server := testHTTPClientJSON("{\"data\":{\"id\":\"ClF8rLe\",\"title\":null,\"description\":null,\"datetime\":1451248840,\"type\":\"image\\/jpeg\",\"animated\":false,\"width\":2448,\"height\":3264,\"size\":1071339,\"views\":176,\"bandwidth\":188555664,\"vote\":null,\"favorite\":false,\"nsfw\":null,\"section\":null,\"account_url\":null,\"account_id\":null,\"in_gallery\":false,\"link\":\"https:\\/\\/i.imgur.com\\/ClF8rLe.jpg\"},\"success\":true,\"status\":200}")
	defer server.Close()

	client, _ := NewClient(httpC, "testing", "")
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

	client, _ := NewClient(new(http.Client), key, RapidAPIKey)

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

	client, _ := NewClient(httpC, "testing", "")
	_, _, err := client.GetInfoFromURL("https://i.imgur.com/")

	if err == nil {
		t.Error("GetInfoFromURL() did not failed but should have.")
		t.FailNow()
	}
}
