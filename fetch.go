package main

import (
	"io/ioutil"
	"net/http"
)

// FetchPublicIP find out the public IP address on the internet
func FetchPublicIP() (string, error) {
	const remote string = "https://ifconfig.me/ip"

	resp, err := http.Get(remote)
	if err != nil {
		return "", err
	}
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(ip), nil
}
