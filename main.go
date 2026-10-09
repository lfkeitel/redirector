package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	serverAddr string
	serverPort int
)

const codeSeparator = " "

func init() {
	flag.StringVar(&serverAddr, "addr", "", "Server address to listen on")
	flag.IntVar(&serverPort, "port", 8080, "Server port to listen on")
}

func main() {
	flag.Parse()
	urlListFile := flag.Arg(0)
	if urlListFile == "" {
		log.Fatal("URL list file is required")
	}

	urlList, err := loadURLList(urlListFile)
	if err != nil {
		log.Fatalf("Failed to load URL list: %v", err)
	}

	log.Printf("Loaded %d URL mappings from %s", len(urlList), urlListFile)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		miniCode := strings.TrimPrefix(r.URL.Path, "/")
		targetURL, ok := urlList[miniCode]
		if !ok {
			http.NotFound(w, r)
			return
		}
		log.Printf("Redirecting %s to %s from %s", miniCode, targetURL, r.RemoteAddr)
		http.Redirect(w, r, targetURL, http.StatusFound)
	})

	addr := fmt.Sprintf("%s:%d", serverAddr, serverPort)
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

type redirectList map[string]string

func loadURLList(filePath string) (redirectList, error) {
	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(fileContents), "\n")
	redirects := make(redirectList, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, codeSeparator, 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line in URL list: %s", line)
		}

		miniCode := parts[0]
		targetURL := parts[1]

		if _, err := url.ParseRequestURI(targetURL); err != nil {
			return nil, fmt.Errorf("invalid target URL: %s", targetURL)
		}

		redirects[miniCode] = targetURL
	}

	return redirects, nil
}
