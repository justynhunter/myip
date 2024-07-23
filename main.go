package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("https://api.ipify.org")
	if err != nil {
		os.Stderr.Write([]byte("error contacting api"))
	}

	ip, err := io.ReadAll(response.Body)
	if err != nil {
		os.Stderr.Write([]byte("error parsing response"))
	}
	os.Stdout.Write(ip)
}
