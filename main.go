package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

// init the env variables
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func getListenAddress() string {
	port := getEnv("PORT", "2345")
	return ":" + port
}

func getProxyURL() string {
	return getEnv("PROXY_URL", "")
}

func getSuffix() string {
	return getEnv("SUFFIX", "")
}

func logSetup() {
	log.Printf("Server Listen on%s\n", getListenAddress())
}

// log request
func logRequest(r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(requestDump))
}

func handleRequestAndRedirect(w http.ResponseWriter, r *http.Request) {
	proxyURL := getProxyURL()
	suffix := getSuffix()
	serveReverseProxy(proxyURL, suffix, w, r)
}

func serveReverseProxy(proxyURL string, suffix string, w http.ResponseWriter, r *http.Request) {

	target, err := url.Parse(proxyURL)
	if err != nil {
		log.Println("error parsing proxy url!")
	}
	proxy := httputil.NewSingleHostReverseProxy(target)

	// rewrite the request according to proxy
	r.Host = target.Host
	r.URL.Host = target.Host
	r.URL.Scheme = target.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))

	// parse the suffix params and append into the query params
	query := r.URL.Query()
	suffixParams, err := url.ParseQuery(suffix)
	if err != nil {
		log.Println("error parsing suffix params!")
	}
	// append the suffix params into query params
	for param := range suffixParams {
		query.Add(param, suffixParams.Get(param))
	}

	r.URL.RawQuery = query.Encode()

	// log modified request
	logRequest(r)

	// allow CORS
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-Width")

	// serve
	proxy.ServeHTTP(w, r)
}

func main() {

	// log
	logSetup()

	// start the server
	http.HandleFunc("/", handleRequestAndRedirect)
	http.ListenAndServe(getListenAddress(), nil)
}
