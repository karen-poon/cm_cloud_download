package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

// MakeSliceFromFile makes a slice by inserting the text of the given file line by line
func MakeSliceFromFile(fileName string) []string {
	var sliceOfURLs = make([]string, 0)

	// open the file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	//insert scanned text into the sliceOfURLs
	downloadedImagesScanner := bufio.NewScanner(file)
	for downloadedImagesScanner.Scan() {
		sliceOfURLs = append(sliceOfURLs, downloadedImagesScanner.Text())
	}

	return sliceOfURLs
}

// FindDownloaded returns a sorted list of downloaded images in numbers
func FindDownloaded(downloadedImageURLs, allImageURLs []string) []int {
	var downloadedListofNumbers = make([]int, 0)
	for _, downloadedURL := range downloadedImageURLs {
		for j, allURL := range allImageURLs {
			if downloadedURL == allURL {
				downloadedListofNumbers = append(downloadedListofNumbers, j)
			}
		}
	}
	sort.Ints(downloadedListofNumbers)
	downloadedListofNumbers = deleteDuplicates(downloadedListofNumbers)
	return downloadedListofNumbers
}

// PrintDownloadedList prints the downloaded list of numbers using ranges
func PrintDownloadedList(downloadedListofNumbers []int) {
	var start, end int
	var scanningRange = len(downloadedListofNumbers) - 1

	for i := 0; i < scanningRange; i++ {
		//check whether the next is a consecutive number
		if downloadedListofNumbers[i]+1 != downloadedListofNumbers[i+1] {
			end = downloadedListofNumbers[i]
			if start == end {
				fmt.Printf("%d\n", end+1)
			} else {
				fmt.Printf("%d - %d\n", start+1, end+1)
			}
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
			if start == end {
				fmt.Printf("%d\n", end+1)
			} else {
				fmt.Printf("%d - %d\n", start+1, end+1)
			}
		}
	}
	return
}

/************************************************************************************/
func deleteDuplicates(downloadedListofNumbers []int) []int {
	seen := make(map[int]struct{}, len(downloadedListofNumbers))
	i := 0
	for _, number := range downloadedListofNumbers {
		if _, ok := seen[number]; ok {
			continue
		}
		seen[number] = struct{}{}
		downloadedListofNumbers[i] = number
		i++
	}
	return downloadedListofNumbers[:i]
}
