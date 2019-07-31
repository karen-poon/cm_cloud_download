package main

import (
	"fmt"
	"sort"
)

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
	return downloadedListofNumbers
}

// PrintDownloadedList prints the downloaded list of numbers in format of "someNumber - someNumber"
func PrintDownloadedList(downloadedListofNumbers []int) {
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
