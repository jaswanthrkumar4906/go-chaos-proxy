# Go Chaos Proxy

A lightweight, concurrent HTTP reverse proxy written in **Go (Golang)** featuring configurable, randomized fault injection (artificial latency and HTTP 503 service failures) designed for system resilience testing and downstream service validation.

---

## 🚀 Key Features

* **Reverse Proxy Engine:** Built on Go's native `net/http/httputil` library to intercept, parse, and forward incoming client requests seamlessly to an upstream target (`httpbin.org`).
* **Chaos Fault Injection:** 
  * **Latency Injection:** Randomly introduces configurable delays (ranging from 500ms to 2000ms) to test application timeout handling.
  * **Failure Injection:** Programmatically drops requests with a randomized probability rate, returning HTTP `503 Service Unavailable` statuses to simulate upstream outages.
* **Real-Time Request Logging:** Captures and logs all routing actions and injected fault metrics directly to standard output.

---

## 🛠️ Tech Stack

* **Language:** Go (Golang)
* **Core Libraries:** `net/http`, `net/http/httputil`, `math/rand`, `time`
* **Version Control:** Git & GitHub

---

## 📦 Getting Started Locally

### Prerequisites
* Ensure you have [Go](https://go.dev/) installed on your machine.

### Installation & Execution

1. Clone the repository:
   ```bash
   git clone [https://github.com/jaswanthrkumar4906/go-chaos-proxy.git](https://github.com/jaswanthrkumar4906/go-chaos-proxy.git)
   cd go-chaos-proxy
