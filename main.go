package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	//for downloaded images
	var downloadedImageURLs = make([]string, 0)
	downloadedImages, err := os.Open("cloud_gallery_urls.txt.done")
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
	allImages, err := os.Open("cloud_gallery_urls.txt")
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
	var downloadedListofNumbers = findDownloaded(downloadedImageURLs, allImageURLs)
	printDownloadedList(downloadedListofNumbers)
}

// returns a sorted list of downloaded images in numbers
func findDownloaded(downloadedImageURLs, allImageURLs []string) []int {
	var downloadedListofNumbers = make([]int, 0)
	for _, downloadedURL := range downloadedImageURLs {
		for j, allURL := range allImageURLs {
			if downloadedURL == allURL {
				downloadedListofNumbers = append(downloadedListofNumbers, j)
			}
		}
	}
	sort.Ints(downloadedListofNumbers)
	fmt.Println(downloadedListofNumbers)
	return downloadedListofNumbers
}

// prints the downloaded list of numbers in format of "someNumber - someNumber"
func printDownloadedList(downloadedListofNumbers []int) {
	var start, end int
	var scanningRange = len(downloadedListofNumbers) - 1

	for i := 0; i < scanningRange; i++ {
		//check whether the next is a consecutive number
		if downloadedListofNumbers[i]+1 != downloadedListofNumbers[i+1] {
			end = downloadedListofNumbers[i]
			fmt.Printf("%d - %d\n", start, end)
			start = downloadedListofNumbers[i+1]
		}

		//the following is to handle last set of consecutive numbers to be displayed
		//two conditions to be considered:
		//1. the last set of consecutive numbers ends with the last number
		//	(did not go into the previous if statement which makes it not able to print)
		//2. only the last number is left to be displayed
		if i == scanningRange-1 &&
			(end != downloadedListofNumbers[i] ||
				start == downloadedListofNumbers[i+1]) {
			end = downloadedListofNumbers[i+1]
			fmt.Printf("%d - %d\n", start, end)
		}
	}
	return
}
