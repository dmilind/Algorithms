/*
Binary Search Algorithm: Implementation in golang
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	// Sorted array of numbers
	var numList = []int{11, 17, 23, 32, 41, 47, 49, 65, 78, 101}
	// number we want to find is 65
	var numNeeded = 65
	// set an index pointer i to 0 and j=len(numList)
	var i, j = 0, len(numList) - 1
	var midIndex int
	for i <= j {
		// finding a mind point of array
		midIndex = (i + j) / 2
		fmt.Println("Array we got is", numList, "and starting and ending pointers are", i, j)
		// check required number is at the mid of array or not
		if numNeeded == numList[midIndex] {
			fmt.Println("Our number", numNeeded, "is at", midIndex, "position in array")
			os.Exit(0)
		}
		if numNeeded > numList[midIndex] {
			i = midIndex + 1
		} else {
			j = midIndex - 1
		}
	}
}

/*
output
Array we got is [11 17 23 32 41 47 49 65 78 101] and starting and ending pointers are 0 9
Array we got is [11 17 23 32 41 47 49 65 78 101] and starting and ending pointers are 5 9
Our number 65 is at 7 position in array
*/
