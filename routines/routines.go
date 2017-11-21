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
	fmt.Printf(" our response is: %v", *response)

	wg.Done() //calls when it's done
	return nil
}

var wg sync.WaitGroup

func main() {

	//create array of urls: urls is now an array pointer
	urls := []*Url{
		&Url{Okay: true, URL: "https://preview.beamery.com", Body: ""},
		&Url{Okay: true, URL: "https://canary.beamery.com", Body: ""},
		&Url{Okay: true, URL: "https://beamery.com", Body: ""}}

	for x := 0; x < len(urls); x++ {
		currentURL := urls[0]
		wg.Add(1)                        //add to waitgroup counter
		go printURLBody(currentURL, &wg) //pass waitgroup to go routine
	}

	wg.Wait()
}
