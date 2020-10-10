package main

import "log"

func main() {
	ip, err := FetchPublicIP()
	if err != nil {
		log.Fatalf("Could not fetch IP address: %v", err)
	}
	log.Printf("Found IP address: %s", ip)
}
