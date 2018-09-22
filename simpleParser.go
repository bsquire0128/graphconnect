package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
	"strings"
)

var (
	filePath string
)

func init() {
	flag.StringVar(&filePath, "filepath", "", "Path of csv")
}

func main() {
	urls := readCSV(filePath)
	for _, url := range urls {
		// get parsedURL value
		ParseCleanURL(url)
		//send url to neo4j client
	}
}

func readCSV(filePath string) []string {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		log.Fatalf("Unable to open file, %vs", err)
	}
	urls, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatalf("Unable to read tilda sv file %v", err)
	}
	var urlArry []string
	for _, url := range urls {
		urlArry = append(urlArry, url[0])
	}
	return urlArry
}

//parseCleanURL does xyz
func ParseCleanURL(url string) string {
	url = strings.Replace(url, "//", "/", -1)
	return strings.Replace(url, "/", "~", -1)
}

//http~time~.com~/search/~q=python+programming