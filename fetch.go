package main

import (
	"io"
	"net/http"
)

// FetchPublicIP find out the public IP address on the internet
func FetchPublicIP(provider string) (string, error) {
	resp, err := http.Get(provider)
	if err != nil {
		return "", err
	}
	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(ip), nil
}
