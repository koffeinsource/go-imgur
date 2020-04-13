package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/koffeinsource/go-imgur"
	"github.com/koffeinsource/go-klogger"
)

func printRate(client *imgur.Client) {
	client.Log.Infof("*** RATE LIMIT ***\n")
	rl, err := client.GetRateLimit()
	if err != nil {
		client.Log.Errorf("Error in GetRateLimit: %v\n", err)
		return
	}
	client.Log.Infof("%v\n", *rl)
}

func printImage(client *imgur.Client, image *string) {
	client.Log.Infof("*** IMAGE ***\n")
	img, _, err := client.GetImageInfo(*image)
	if err != nil {
		client.Log.Errorf("Error in GetImageInfo: %v\n", err)
		return
	}
	client.Log.Infof("%v\n", img)
}

func printAlbum(client *imgur.Client, album *string) {
	client.Log.Infof("*** ALBUM ***\n")
	img, _, err := client.GetAlbumInfo(*album)
	if err != nil {
		client.Log.Errorf("Error in GetAlbumInfo: %v\n", err)
		return
	}
	client.Log.Infof("%v\n", img)
}

func printGImage(client *imgur.Client, gimage *string) {
	client.Log.Infof("*** GALLERY IMAGE ***\n")
	img, _, err := client.GetGalleryImageInfo(*gimage)
	if err != nil {
		client.Log.Errorf("Error in GetGalleryImageInfo: %v\n", err)
		return
	}
	client.Log.Infof("%v\n", img)
}

func printGAlbum(client *imgur.Client, galbum *string) {
	client.Log.Infof("*** GALLERY ALBUM ***\n")
	img, _, err := client.GetGalleryAlbumInfo(*galbum)
	if err != nil {
		client.Log.Errorf("Error in GetGalleryAlbumInfo: %v\n", err)
		return
	}
	client.Log.Infof("%v\n", img)
}

func printURL(client *imgur.Client, url *string) {
	client.Log.Infof("*** URL ***\n")
	img, _, err := client.GetInfoFromURL(*url)
	if err != nil {
		client.Log.Errorf("Error in GetInfoFromURL: %v\n", err)
		return
	}
	client.Log.Infof("%v\n", img)
}

func main() {
	imgurClientID := flag.String("id", "", "Your imgur client id. REQUIRED!")
	url := flag.String("url", "", "Gets information based on the URL passed.")
	upload := flag.String("upload", "", "Filepath to an image that will be uploaded to imgur.")
	image := flag.String("image", "", "The image ID to be queried.")
	album := flag.String("album", "", "The album ID to be queried.")
	gimage := flag.String("gimage", "", "The gallery image ID to be queried.")
	galbum := flag.String("galbum", "", "The gallery album ID to be queried.")
	rate := flag.Bool("rate", false, "Get the current rate limit.")
	flag.Parse()

	// Check if there is anything todo
	if flag.NFlag() >= 3 || *imgurClientID == "" {
		flag.PrintDefaults()
		return
	}

	client := new(imgur.Client)
	client.HTTPClient = new(http.Client)
	client.Log = new(klogger.CLILogger)
	client.ImgurClientID = *imgurClientID

	if *upload != "" {
		_, st, err := client.UploadImageFromFile(*upload, "", "test title", "test desc")
		if st != 200 || err != nil {
			fmt.Printf("Status: %v\n", st)
			fmt.Printf("Err: %v\n", err)
		}
	}

	if *rate {
		printRate(client)
	}

	if *image != "" {
		printImage(client, image)
	}

	if *album != "" {
		printAlbum(client, album)
	}

	if *gimage != "" {
		printGImage(client, gimage)
	}

	if *galbum != "" {
		printGAlbum(client, galbum)
	}

	if *url != "" {
		printURL(client, url)
	}
}
