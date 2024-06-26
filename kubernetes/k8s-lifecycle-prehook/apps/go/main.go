package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var delayTerm = flag.Int("delay-term", 5000, "delay in milliseconds before SIGTERM is handled")
var port = flag.String("port", "8080", "port to listen on")
var upstream = flag.String("upstream", "http://localhost:8080", "upstream service URL")
var delayHttp = flag.Int("delay-http", 0, "delay in milliseconds before responding to HTTP requests")

func handlerHealthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request from", r.RemoteAddr)
	if *delayHttp > 0 {
		fmt.Printf("Delaying response by %d milliseconds...\n", *delayHttp)
		time.Sleep(time.Duration(*delayHttp) * time.Millisecond)
	}
	fmt.Fprintf(w, "Hello, World!")
}

func handlerChain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[chain] Received request from", r.RemoteAddr)
	// make a request to the upstream service
	resp, err := http.Get(*upstream + "/request")
	if err != nil {
		fmt.Println("[chain] Error making request to upstream service:", err)
		http.Error(w, "Error making request to upstream service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	// read the response from the upstream service
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[chain] Error reading response from upstream service:", err)
		http.Error(w, "Error reading response from upstream service", http.StatusInternalServerError)
		return
	}
	if *delayHttp > 0 {
		fmt.Printf("Delaying response by %d milliseconds...\n", *delayHttp)
		time.Sleep(time.Duration(*delayHttp) * time.Millisecond)
	}
	// write the response to the client
	fmt.Fprintf(w, "[chain] Response from upstream: %q", body)
}

func main() {
	flag.Parse()
	// Print the PID of the process
	pid := os.Getpid()
	fmt.Printf("Process ID: %d\n", pid)

	// Create a channel to receive signals
	sigChan := make(chan os.Signal, 1)
	// Notify the channel on SIGHUP and SIGINT
	signal.Notify(sigChan)

	// Run a goroutine to handle signals
	go func() {
		for {
			sig := <-sigChan
			switch sig {
			case syscall.SIGHUP:
				fmt.Println("Received SIGHUP")
			case syscall.SIGINT:
				fmt.Println("Received SIGINT, exiting...")
				// Perform cleanup if necessary
				os.Exit(0)
			case syscall.SIGTERM:
				fmt.Printf("Received SIGTERM, exiting in %d sec...\n", *delayTerm)
				time.Sleep(time.Duration(*delayTerm) * time.Millisecond)
				// Perform cleanup if necessary
				os.Exit(0)
			default:
				fmt.Printf("Received signal: %v", sig)
			}
		}
	}()

	fmt.Printf("HTTP service is running on port %s...\n", *port)

	http.HandleFunc("/", handler)
	http.HandleFunc("/chain", handlerChain)
	http.HandleFunc("/healthz", handlerHealthz)
	http.ListenAndServe(fmt.Sprintf(":%s", *port), nil)
}
