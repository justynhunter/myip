package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://api.ipify.org")
	if err != nil {
		os.Stderr.Write([]byte("error contacting api"))
	}

	ip, err := io.ReadAll(res.Body)
	if err != nil {
		os.Stderr.Write([]byte("error parsing response"))
	}
	os.Stdout.Write(ip)
}
