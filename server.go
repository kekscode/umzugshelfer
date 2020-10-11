package main

import (
	"io"
	"net/http"
	"strings"
)

// Returns the requesting client's IP address without the Port
func serveIP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, strings.Split(r.RemoteAddr, ":")[0])
}
