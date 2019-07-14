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

	fmt.Print("Total number of Images: ")
	fmt.Println(len(allImageURLs))

	fmt.Print("Number of Images Downloaded: ")
	fmt.Println(len(downloadedImageURLs))

	match(0, 0, downloadedImageURLs, allImageURLs)
}

func match(breakPoint, nextStart int, downloadedImageURLs, allImageURLs []string) {
	if len(downloadedImageURLs) == breakPoint+1 || nextStart == len(allImageURLs) {
		return
	}
	var stop int
	for i := 0; i < len(downloadedImageURLs)-breakPoint; i++ {
		if strings.Compare(downloadedImageURLs[breakPoint+i], allImageURLs[nextStart+i]) != 0 {
			breakPoint = breakPoint + i
			stop = nextStart + i
			fmt.Println(breakPoint, stop)
			break
		}
	}
	for j := 0; j < len(allImageURLs)-stop; j++ {
		if strings.Compare(downloadedImageURLs[breakPoint], allImageURLs[stop+j]) == 0 {
			nextStart = stop + j
			fmt.Println(nextStart)
			break
		} else {
			nextStart = len(allImageURLs) //rest of the list has no matching
		}
	}
	match(breakPoint, nextStart, downloadedImageURLs, allImageURLs)
}
