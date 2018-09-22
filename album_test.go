package imgur

import (
	"net/http"
	"os"
	"testing"
)

func TestAlbumImgurSimulated(t *testing.T) {
	httpC, server := testHTTPClientJSON("{\"data\":{\"id\":\"VZQXk\",\"title\":\"Gianluca Gimini's bikes\",\"description\":null,\"datetime\":1460715031,\"cover\":\"CJCA0gW\",\"cover_width\":1200,\"cover_height\":786,\"account_url\":\"mrcassette\",\"account_id\":157430,\"privacy\":\"public\",\"layout\":\"blog\",\"views\":667581,\"link\":\"https:\\/\\/imgur.com\\/a\\/VZQXk\",\"favorite\":false,\"nsfw\":false,\"section\":\"pics\",\"images_count\":1,\"in_gallery\":true,\"images\":[{\"id\":\"CJCA0gW\",\"title\":null,\"description\":\"by Designer Gianluca Gimini\\nhttps:\\/\\/www.behance.net\\/gallery\\/35437979\\/Velocipedia\",\"datetime\":1460715032,\"type\":\"image\\/jpeg\",\"animated\":false,\"width\":1200,\"height\":786,\"size\":362373,\"views\":4420880,\"bandwidth\":1602007548240,\"vote\":null,\"favorite\":false,\"nsfw\":null,\"section\":null,\"account_url\":null,\"account_id\":null,\"in_gallery\":false,\"link\":\"https:\\/\\/i.imgur.com\\/CJCA0gW.jpg\"}]},\"success\":true,\"status\":200}")
	defer server.Close()

	client := createClient(httpC, "testing", "")
	alb, status, err := client.GetAlbumInfo("VZQXk")

	if err != nil {
		t.Errorf("GetAlbumInfo() failed with error: %v", err)
		t.FailNow()
	}

	if alb.Title != "Gianluca Gimini's bikes" || alb.Cover != "CJCA0gW" || alb.CoverWidth != 1200 || alb.CoverHeight != 786 || alb.Link != "https://imgur.com/a/VZQXk" || alb.ImagesCount != 1 || alb.Images[0].ID != "CJCA0gW" {
		t.Error("Data comparision failed.")

		if alb.Title != "Gianluca Gimini's bikes" {
			t.Errorf("Title is %v.\n", alb.Title)
		}
		if alb.Cover != "CJCA0gW" {
			t.Errorf("Cover is %v.\n", alb.Cover)
		}
		if alb.CoverWidth != 1200 {
			t.Errorf("CoverWidth is %v.\n", alb.CoverWidth)
		}
		if alb.CoverHeight != 786 {
			t.Errorf("CoverHeight is %v.\n", alb.CoverHeight)
		}
		if alb.Link != "https://imgur.com/a/VZQXk" {
			t.Errorf("Link is %v.\n", alb.Link)
		}
		if alb.ImagesCount != 14 {
			t.Errorf("ImagesCount is %v.\n", alb.ImagesCount)
		}
		if alb.Images[0].ID != "CJCA0gW" {
			t.Errorf("Images is %v.\n", alb.Images)
		}
		t.Fail()
	}

	if status != 200 {
		t.Errorf("Statsu != 200. It was %v.", status)
		t.Fail()
	}
}

func TestAlbumImgurReal(t *testing.T) {
	key := os.Getenv("IMGURCLIENTID")
	if key == "" {
		t.Skip("IMGURCLIENTID environment variable not set.")
	}
	mashapKey := os.Getenv("MASHAPEKEY")

	client := createClient(new(http.Client), key, mashapKey)

	alb, status, err := client.GetAlbumInfo("VZQXk")

	if err != nil {
		t.Errorf("GetAlbumInfo() failed with error: %v", err)
		t.FailNow()
	}

	if alb.Title != "Gianluca Gimini's bikes" || alb.Cover != "CJCA0gW" || alb.CoverWidth != 1200 || alb.CoverHeight != 786 || alb.Link != "https://imgur.com/a/VZQXk" || alb.ImagesCount != 14 || alb.Images[0].ID != "CJCA0gW" {
		t.Error("Data comparision failed.")

		if alb.Title != "Gianluca Gimini's bikes" {
			t.Errorf("Title is %v.\n", alb.Title)
		}
		if alb.Cover != "CJCA0gW" {
			t.Errorf("Cover is %v.\n", alb.Cover)
		}
		if alb.CoverWidth != 1200 {
			t.Errorf("CoverWidth is %v.\n", alb.CoverWidth)
		}
		if alb.CoverHeight != 786 {
			t.Errorf("CoverHeight is %v.\n", alb.CoverHeight)
		}
		if alb.Link != "https://imgur.com/a/VZQXk" {
			t.Errorf("Link is %v.\n", alb.Link)
		}
		if alb.ImagesCount != 14 {
			t.Errorf("ImagesCount is %v.\n", alb.ImagesCount)
		}
		if alb.Images[0].ID != "CJCA0gW" {
			t.Errorf("Images is %v.\n", alb.Images)
		}
		t.Fail()
	}

	if status != 200 {
		t.Errorf("Statsu != 200. It was %v.", status)
		t.Fail()
	}
}
