package imgur

import (
	"net/http"
	"os"
	"testing"
)

func TestGalleryImageImgurSimulated(t *testing.T) {
	httpC, server := testHTTPClientJSON("{\"data\":{\"id\":\"Hf6cs\",\"title\":\"The Tridge. (three way bridge)\",\"description\":null,\"datetime\":1316367003,\"type\":\"image\\/jpeg\",\"animated\":false,\"width\":1700,\"height\":1133,\"size\":268126,\"views\":1342557,\"bandwidth\":359974438182,\"vote\":null,\"favorite\":false,\"nsfw\":false,\"section\":\"pics\",\"account_url\":null,\"account_id\":null,\"in_gallery\":true,\"topic\":null,\"topic_id\":0,\"link\":\"https:\\/\\/i.imgur.com\\/Hf6cs.jpg\",\"comment_count\":90,\"ups\":585,\"downs\":3,\"points\":582,\"score\":1136,\"is_album\":false},\"success\":true,\"status\":200}")
	defer server.Close()

	client, _ := NewClient(httpC, "testing", "")
	img, status, err := client.GetGalleryImageInfo("Hf6cs")

	if err != nil {
		t.Errorf("GetImageInfo() failed with error: %v", err)
		t.FailNow()
	}

	if img.Title != "The Tridge. (three way bridge)" || img.Animated != false || img.Bandwidth != 359974438182 || img.Datetime != 1316367003 || img.Description != "" || img.Height != 1133 || img.Width != 1700 || img.ID != "Hf6cs" || img.Link != "https://i.imgur.com/Hf6cs.jpg" || img.Views != 1342557 {
		t.Fail()
	}

	if status != 200 {
		t.Fail()
	}
}

func TestGalleryImageImgurReal(t *testing.T) {
	key := os.Getenv("IMGURCLIENTID")
	if key == "" {
		t.Skip("IMGURCLIENTID environment variable not set.")
	}
	RapidAPIKey := os.Getenv("RapidAPIKEY")

	client, _ := NewClient(new(http.Client), key, RapidAPIKey)

	img, status, err := client.GetGalleryImageInfo("Hf6cs")

	if err != nil {
		t.Errorf("GetImageInfo() failed with error: %v", err)
		t.FailNow()
	}

	if img.Title != "The Tridge. (three way bridge) " || img.Animated != false || img.Description != "" || img.Height != 1133 || img.Width != 1700 || img.ID != "Hf6cs" || img.Link != "https://i.imgur.com/Hf6cs.jpg" {
		t.Fail()
	}

	if status != 200 {
		t.Fail()
	}
}
