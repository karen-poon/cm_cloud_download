package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
		allImageURLs = append(allImageURLs, allImagesScanner.Text())
	}

	fmt.Print("\nTotal number of images: ")
	fmt.Println(len(allImageURLs))

	fmt.Print("\nNumber of Images Downloaded: ")
	fmt.Println(len(downloadedImageURLs))

	fmt.Println("\nDownloaded images are listed as follows:")
	match(0, 0, downloadedImageURLs, allImageURLs)
}

func match(breakMatch, matchAgain int, downloadedImageURLs, allImageURLs []string) {
	if matchAgain == len(allImageURLs)-1 {
		if downloadedImageURLs[len(downloadedImageURLs)-1] == allImageURLs[len(allImageURLs)-1] && matchAgain != len(allImageURLs)-1 {
			fmt.Printf("From %d to %d\n", len(allImageURLs)-1, len(allImageURLs)-1)
		}
		return
	}
	var notMatchingAnymore int
	for i := 0; i < len(downloadedImageURLs)-breakMatch; i++ {
		if strings.Compare(downloadedImageURLs[breakMatch+i], allImageURLs[matchAgain+i]) != 0 {
			breakMatch = breakMatch + i
			notMatchingAnymore = matchAgain + i
			//fmt.Println(breakMatch, notMatchingAnymore)
			break
		}
	}
	if notMatchingAnymore == 0 { // if it did not go into the for loop above
		notMatchingAnymore = matchAgain + len(downloadedImageURLs) - breakMatch
	}
	fmt.Printf("From %d to %d\n", matchAgain, notMatchingAnymore-1)
	for j := 0; j < len(allImageURLs)-notMatchingAnymore; j++ {
		if strings.Compare(downloadedImageURLs[breakMatch], allImageURLs[notMatchingAnymore+j]) == 0 {
			matchAgain = notMatchingAnymore + j
			//fmt.Println(matchAgain)
			break
		} else {
			matchAgain = len(allImageURLs) - 1 //rest of the list has no matching
			//fmt.Println(matchAgain)
		}
	}
	if len(allImageURLs) == notMatchingAnymore { // if it did not go into the for loop above
		matchAgain = len(allImageURLs) - 1
	}
	match(breakMatch, matchAgain, downloadedImageURLs, allImageURLs)
}
