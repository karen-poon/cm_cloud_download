package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//for downloaded images
	var downloadedImageURLs = make([]string, 0)
	downloadedImages, err := os.Open("files/cloud_gallery_urls.txt.done")
	if err != nil {
		log.Fatal(err)
	}
	downloadedImagesScanner := bufio.NewScanner(downloadedImages)
	for downloadedImagesScanner.Scan() {
		//insert scanned text into the downloadedImageURLs slice
		downloadedImageURLs = append(downloadedImageURLs, downloadedImagesScanner.Text())
	}

	//for all images to be downloaded
	var allImageURLs = make([]string, 0)
	allImages, err := os.Open("files/cloud_gallery_urls.txt")
	if err != nil {
		log.Fatal(err)
	}
	allImagesScanner := bufio.NewScanner(allImages)
	for allImagesScanner.Scan() {
		//insert scanned text into the allImageURLs
		allImageURLs = append(allImageURLs, allImagesScanner.Text())
	}

	fmt.Print("\nTotal number of images: ")
	fmt.Println(len(allImageURLs))

	fmt.Print("\nNumber of Images Downloaded: ")
	fmt.Println(len(downloadedImageURLs))

	fmt.Println("\nDownloaded images are listed as follows:")
	var downloadedListofNumbers = FindDownloaded(downloadedImageURLs, allImageURLs)
	PrintDownloadedList(downloadedListofNumbers)
}
