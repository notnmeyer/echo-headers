package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type config struct {
	port int
}

var (
	cfg         config
	portDefault = 8080
)

func init() {
	if val, ok := os.LookupEnv("PORT"); ok {
		port, err := strconv.Atoi(val)
		if err != nil {
			fmt.Errorf("couldn't parse `PORT` as int. got %s\n", val)
		}
		cfg.port = port
	} else {
		cfg.port = portDefault
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	output := ""

	fmt.Printf("%s: %s %s\n", r.RemoteAddr, r.Method, r.URL)
	for name, headers := range r.Header {
		for _, h := range headers {
			output += fmt.Sprintf("%s: %s\n", name, h)
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(output))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("Server started on :%d\n", cfg.port)
	http.ListenAndServe(":8080", nil)
}
