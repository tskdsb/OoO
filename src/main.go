// find a number from an array(large size) smaller than both sides(if exists)

package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

var size = 1 << 20
var data = make([]int, size)

func fillData() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < size; i++ {
		data[i] = r.Intn(1000)
	}
}

func sizeCheck() (err error) {

	if size <= 1 {
		return errors.New(fmt.Sprintf("size(%d) should > 1\n", size))
	}

	return
}

// O(1)
func special() (err error) {

	if size == 1 {
		fmt.Printf("index: %d, value: %d\n", 0, data[0])
		return
	}

	if data[0] <= data[1] {
		fmt.Printf("index: %d, value: %d\n", 0, data[0])
		return
	}

	if data[size-2] >= data[size-1] {
		fmt.Printf("index: %d, value: %d\n", size-1, data[size-1])
		return
	}

	return errors.New("no special case")
}

// O(2n)
func stupid() {

	for i := 1; i < size-1; i++ {
		if data[i] < data[i-1] && data[i] < data[i+1] {
			fmt.Printf("index: %d, value: %d\n", i, data[i])
			return
		}
	}
}

// O(n)
func lessStupid() {

	var minValue = data[0]
	var minIndex = 0

	for i := 1; i < size-1; i++ {
		if data[i] < minValue {
			minValue = data[i]
			minIndex = i
		}
	}

	fmt.Printf("index: %d, value: %d\n", minIndex, minValue)
}

// O(?)
func guess() {

	r := rand.New(rand.NewSource(time.Now().Unix()))
	for {
		i := r.Intn(size-2) + 1
		if data[i] < data[i-1] && data[i] < data[i+1] {
			fmt.Printf("index: %d, value: %d\n", i, data[i])
			return
		}
	}
}

// O(n)
func view() {

	var index = 0

	for index < size-1 {
		if data[index+1] < data[index] {
			index = index + 1
		} else {
			fmt.Printf("index: %d, value: %d\n", index, data[index])
			return
		}
	}
}

// O(log2N)
func binarySearch() {

	var leftIndex = 0
	var rightIndex = size - 1
	var midIndex = (leftIndex + rightIndex) / 2

	for leftIndex < midIndex && midIndex < rightIndex {

		if data[midIndex] < data[midIndex-1] {

			if data[midIndex] < data[midIndex+1] {
				fmt.Printf("index: %d, value: %d\n", midIndex, data[midIndex])
				return
			} else {
				leftIndex = midIndex
				midIndex = (leftIndex + rightIndex) / 2
			}

		} else {
			rightIndex = midIndex
			midIndex = (leftIndex + rightIndex) / 2
		}

	}

	if data[leftIndex] < data[rightIndex] {
		fmt.Printf("index: %d, value: %d\n", leftIndex, data[leftIndex])
		return
	} else {
		fmt.Printf("index: %d, value: %d\n", rightIndex, data[rightIndex])
		return
	}

}

func main() {

	fillData()

	err := sizeCheck()
	if err != nil {
		log.Println(err)
		return
	}

	err = special()
	if err == nil {
		return
	}

	guess()
	binarySearch()
	view()
	lessStupid()
	stupid()
}
