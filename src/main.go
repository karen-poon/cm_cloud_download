package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Check CM Cloud Downloaded Images!")
	fmt.Println("By @karen-poon")

	var downloadedImageURLs = MakeSliceFromFile("files/cloud_gallery_urls.txt.done")
	var allImageURLs = MakeSliceFromFile("files/cloud_gallery_urls.txt")

	fmt.Print("\nTotal number of images: ")
	fmt.Println(len(allImageURLs))

	fmt.Print("\nNumber of Images Downloaded: ")
	fmt.Println(len(downloadedImageURLs))

	fmt.Println("\nDownloaded images are listed as follows:")
	var downloadedListofNumbers = FindDownloaded(downloadedImageURLs, allImageURLs)
	PrintDownloadedList(downloadedListofNumbers)

	fmt.Println("\nPress q + Enter to terminate")
	var input string
	for input != "q" {
		//waits for user's input to terminate program
		fmt.Scanf("%s", &input)
	}
}
