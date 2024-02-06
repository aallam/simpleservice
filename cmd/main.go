package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	port      = getEnv("PORT0", "8080")
	version   = getEnv("SIMPLE_SERVICE_VERSION", "0.0.0")
	healthMin = getEnvAsInt("HEALTH_MIN", 0)
	healthMax = getEnvAsInt("HEALTH_MAX", 0)
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/info", infoHandler)
	http.HandleFunc("/env", envHandler)

	log.Printf("This is simple service in version v%s listening on port %s", version, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// rootHandler handles `/` resource
func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("/ serving from %s has been invoked from %s\n", r.Host, r.RemoteAddr)
	response := map[string]string{
		"version": version,
		"host":    r.Host,
		"result":  "Welcome to simple service",
	}
	writeJSONResponse(w, response)
}

// healthHandler handles `/health` resource
func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("/health serving from %s has been invoked from %s\n", r.Host, r.RemoteAddr)
	if healthMax > healthMin && healthMin >= 0 {
		delay := healthMin + rand.Intn(healthMax-healthMin)
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
	response := map[string]bool{
		"healthy": true,
	}
	writeJSONResponse(w, response)
}

// infoHandler handles `/info` resource
func infoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("/info serving from %s has been invoked from %s\n", r.Host, r.RemoteAddr)
	response := map[string]string{
		"version": version,
		"host":    r.Host,
		"from":    r.RemoteAddr,
	}
	writeJSONResponse(w, response)
}

// envHandler handles `/env` resource
func envHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("/env serving from %s has been invoked from %s\n", r.Host, r.RemoteAddr)
	response := map[string]interface{}{
		"version": version,
		"env":     os.Environ(),
	}
	writeJSONResponse(w, response)
}

// writeJSONResponse writes JSON response
func writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON: %v\n", err)
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

// getEnv returns environment variable value or default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// getEnvAsInt returns environment variable value as int or default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}
