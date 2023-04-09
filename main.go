package main

import (
	"encoding/json"
	"net/http"
	"regexp"
)

type jsonResponse struct {
	URL     string `json:"url"`
	Turn    string `json:"turn"`
	Desc    string `json:"desc"`
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	var response jsonResponse

	githubRegexp := regexp.MustCompile(`^https:\/\/github\.com\/[A-Za-z0-9_.-]+\/[A-Za-z0-9_.-]+$`)
	youtubeRegexp := regexp.MustCompile(`^https:\/\/www\.youtube\.com\/watch\?v=[A-Za-z0-9_-]+`)

	if githubRegexp.MatchString(url) {
		response = jsonResponse{URL: "github " + url, Turn: "on", Desc: "A browser based code editor"}
	} else if youtubeRegexp.MatchString(url) {
		response = jsonResponse{URL: "youtube " + url, Turn: "on", Desc: "A browser based code editor"}
	} else {
		response = jsonResponse{URL: "unknown", Turn: "off", Desc: "Unsupported URL type"}
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

