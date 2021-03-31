package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	wikiRandomUrl  = "https://en.wikipedia.org/wiki/Special:Random"
	todoBackendUrl = os.Getenv("TODO_BACKEND_URL")
)

func main() {
	noRedirectClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	res, err := noRedirectClient.Get(wikiRandomUrl)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 302 {
		log.Fatal("Response was not a 302 redirect")
	}
	wikiUrl := res.Header.Get("Location")
	if wikiUrl == "" {
		log.Fatal("Location header was empty")
	}
	wikiDecodedUrl, err := url.QueryUnescape(wikiUrl)
	if err != nil {
		log.Fatal("Location decoding failed")
	}
	res.Body.Close()

	jsonBody := []byte(fmt.Sprintf("{\"title\": \"Read %s\"}", wikiDecodedUrl))
	res, err = http.Post(todoBackendUrl, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode <= 199 || res.StatusCode >= 400 {
		log.Fatal("Todo send failed. Resulted in status: " + res.Status)
	}
	res.Body.Close()

	log.Println("More stuff To Do sent successfully! :)")
}
