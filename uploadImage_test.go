package imgur

import (
	"testing"

	"github.com/koffeinsource/go-klogger"
)

func TestUploadImageErrors(t *testing.T) {
	httpC, server := testHTTPClientJSON("")
	defer server.Close()

	client := new(Client)
	client.HTTPClient = httpC
	client.Log = new(klogger.CLILogger)
	client.ImgurClientID = "testing"

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
}
