package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func main() {
	targetURL := "https://httpbin.org"
	remote, err := url.Parse(targetURL)
	if err != nil {
		log.Fatalf("Failed to parse target URL: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// --- CHAOS CONFIGURATION ---
		// 1. Artificial Latency: Add a random delay between 500ms to 2000ms
		delayMs := rand.Intn(1500) + 500
		fmt.Printf("[CHAOS] Injecting %dms latency for: %s %s\n", delayMs, r.Method, r.URL.Path)
		time.Sleep(time.Duration(delayMs) * time.Millisecond)

		// 2. Failure Injection: 20% chance to drop the request with a 503 Server Error
		if rand.Float32() < 0.20 {
			fmt.Printf("[CHAOS] 💥 Injected 503 Error for: %s %s\n", r.Method, r.URL.Path)
			http.Error(w, "Chaos Proxy: Simulated Upstream Failure (503)", http.StatusServiceUnavailable)
			return
		}
		// ---------------------------

		fmt.Printf("[PROXY] Forwarding request smoothly: %s %s -> %s\n", r.Method, r.URL.Path, targetURL)
		proxy.ServeHTTP(w, r)
	})

	port := ":8080"
	fmt.Printf("Starting Go Chaos Proxy on http://localhost%s (forwarding to %s)...\n", port, targetURL)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}