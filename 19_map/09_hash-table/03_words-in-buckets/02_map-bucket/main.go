package main

import (
	"net/http"
	"log"
	"bufio"
	"fmt"
)


func main() {
	const nrOfBuckets = 12

	// get the book adventures of sherlock holmes
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// Set the split function fot the scanning operation.
	scanner.Split(bufio.ScanWords)

	// Create map with a key of int
	// and a value of another map
	// with a key of string, which will be the word
	// and a value of int, which will be the number of times the word occurs
	buckets := make(map[int]map[string]int)

	// Create slices to hold words
	for i := 0; i < nrOfBuckets; i++ {
		buckets[i] = make(map[string]int)
	}

	// Loop over the words
	for scanner.Scan() {
		word := scanner.Text()
		n := HashBucket(word, nrOfBuckets)
		buckets[n][word]++
	}

	// Print words in a bucket
	for k, v := range buckets[6] {
		fmt.Println(v, "\t-", k)
	}

	// Print the words in one of the buckets
	fmt.Println(buckets[6])
}

func HashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
