package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type Url struct {
	Okay bool
	URL  string
	Body string
}

func printURLBody(url *Url, wg *sync.WaitGroup) error {

	response, err := http.Get(url.URL)
	if err != nil {
		return err
	}

	output, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	url.Body = string(output)
	println(string(output))
	fmt.Printf("%v", *response)

	wg.Done()
	return nil
}

func main() {

	//create array of urls
	urls := []*Url{&Url{Okay: true, URL: "https://preview.beamery.com", Body: ""},
		&Url{Okay: true, URL: "https://canary.beamery.com", Body: ""},
		&Url{Okay: true, URL: "https://beamery.com", Body: ""}} //pointer array
	//urls := &UrlLink{{Link: *url1}, {Link: *url2}, {Link: *url3}}

	//urls := []string{"https://preview.beamery.com", "https://canary.beamery.com", "https://beamery.com"}

	var wg sync.WaitGroup

	for x := 0; x < len(urls); x++ {
		currentURL := urls[0]
		wg.Add(1)
		go printURLBody(currentURL, &wg) //pass waitgroup to go routine
	}

	wg.Wait()
}
