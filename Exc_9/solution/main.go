package main

import (
	"exc9/mapred"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// Main function
func main() {
	// todo read file

	// Read the whole file at once
	data, err := os.ReadFile("res/meditations.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Split into lines
	lines := strings.Split(string(data), "\n")

	// todo run your mapreduce algorithm
	var mr mapred.MapReduce
	results := mr.Run(lines)

	// todo print your result to stdout
	type kv struct {
		Key   string
		Value int
	}
	sorted := make([]kv, 0, len(results))
	for k, v := range results {
		sorted = append(sorted, kv{k, v})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})

	// --- Print top 50 words ---
	fmt.Println("Top 50 most frequent words:")
	for i := 0; i < 50 && i < len(sorted); i++ {
		fmt.Printf("%-15s %d\n", sorted[i].Key, sorted[i].Value)
	}
}
