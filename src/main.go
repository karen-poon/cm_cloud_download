package main

import (
	"fmt"
)

func main() {
	var downloadedImageURLs = MakeSliceFromFile("files/cloud_gallery_urls.txt.done")
	var allImageURLs = MakeSliceFromFile("files/cloud_gallery_urls.txt")

	fmt.Print("\nTotal number of images: ")
	fmt.Println(len(allImageURLs))

	fmt.Print("\nNumber of Images Downloaded: ")
	fmt.Println(len(downloadedImageURLs))

	fmt.Println("\nDownloaded images are listed as follows:")
	var downloadedListofNumbers = FindDownloaded(downloadedImageURLs, allImageURLs)
	PrintDownloadedList(downloadedListofNumbers)
}
