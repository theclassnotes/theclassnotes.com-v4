package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func concat(source1 []string, source2 []string) []string {
	result := make([]string, len(source1)+len(source2))
	copy(result, source1)
	copy(result[len(source1):], source2)
	return result
}

func fileList(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var filenames []string
	for _, fileinfo := range files {
		filenames = append(filenames, path+"/"+fileinfo.Name())
	}

	return filenames, nil
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
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
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func main() {
	filelist, err := fileList("_data")
	if err != nil {
		log.Fatalf("fileList:", err)
	}

	var config []string

	for _, filename := range filelist {
		fmt.Println("fetching filename:", filename)

		lines, err := readLines(filename)
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}

		config = concat(config, lines)
	}

	if err := writeLines(config, "_config.yml"); err != nil {
		log.Fatalf("writeLines: %s", err)
	}
}
