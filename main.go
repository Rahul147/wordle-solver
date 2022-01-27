package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func contains(s string, whitelist []string, blacklist []string, mapping map[int]string) bool {
	// Eliminate all blacklisted chars
	for _, char := range blacklist {
		if strings.Contains(s, char) {
			return false
		}
	}

	// Eliminate all non whitelisted chars
	for _, char := range whitelist {
		if !strings.Contains(s, char) {
			return false
		}
	}

	// Ensure mapping is correct for remaining words
	for k, v := range mapping {
		if string(s[k]) != v {
			return false
		}
	}

	return true
}

func genChunks(xs []string, chunkSize int) [][]string {
	if len(xs) == 0 {
		return nil
	}
	chunks := make([][]string, (len(xs)+chunkSize-1)/chunkSize)
	prev := 0
	i := 0
	till := len(xs) - chunkSize
	for prev < till {
		next := prev + chunkSize
		chunks[i] = xs[prev:next]
		prev = next
		i++
	}
	chunks[i] = xs[prev:]
	return chunks
}

func main() {
	lines, err := readLines("5_letter_word_list.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	whitelist := []string{}
	blacklist := []string{}
	mappings := map[int]string{}

	var wg sync.WaitGroup
	chunks := genChunks(lines, runtime.NumCPU())

	for _, _line := range chunks {
		wg.Add(1)
		go func(_line []string) {
			defer wg.Done()
			for _, line := range _line {
				if contains(line, whitelist, blacklist, mappings) {
					fmt.Println(line)
				}
			}
		}(_line)
	}
	wg.Wait()
}
