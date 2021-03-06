package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

func main() {

	// If a bind address like 0.0.0.0:8080 or the like is given, we start in
	// server mode
	if len(os.Args) == 2 {
		bindAddr := os.Args[1]
		http.HandleFunc("/", serveIP)
		log.Fatal(http.ListenAndServe(bindAddr, nil))
	}

	var zoneID string
	cfAPIKey := os.Getenv("CF_API_KEY")
	cfAPIEMail := os.Getenv("CF_API_EMAIL")
	ipProvider := os.Getenv("UZH_PUBLIC_IP_PROVIDER")
	dnsZone := os.Getenv("UZH_DNS_ZONE")
	dnsNameA := os.Getenv("UZH_DNS_A_RECORD")

	// Construct a new API object
	api, err := cloudflare.New(cfAPIKey, cfAPIEMail)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err = api.ZoneIDByName(dnsZone)
	if err != nil {
		log.Fatal(err)
	}

	filter := cloudflare.DNSRecord{Name: dnsNameA}
	records, err := api.DNSRecords(zoneID, filter)
	if err != nil {
		log.Fatal(err)
	}
	if len(records) > 1 {
		log.Fatalf("Found more than 1 records for %s: %v", dnsNameA, records)
	}
	currentRecord := records[0]

	ip, err := FetchPublicIP(ipProvider)
	if err != nil {
		log.Fatalf("Could not fetch IP address: %v", err)
	}

	if net.ParseIP(ip) == nil {
		log.Fatalf("IP Address (%s) seems to be an invalid public IP Address", ip)
	}
	log.Printf("Found public IP address: %s", ip)

	newRecord := cloudflare.DNSRecord{ZoneName: dnsZone, Name: dnsNameA, Content: ip, Type: "A"}

	if newRecord.Content == currentRecord.Content {
		log.Fatalf("Public IP (%s) is already the same as found in DNS for %s (%s)", ip, currentRecord.Name, currentRecord.Content)
	}

	err = api.UpdateDNSRecord(zoneID, currentRecord.ID, newRecord)
	if err != nil {
		log.Fatalf("Could not Update: %v", err)
	}
}
