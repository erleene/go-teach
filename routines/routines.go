package main

import (
	"fmt"
	"net/http"
)

type Url struct {
	Okay bool
	URL  string
	Body string
}

func printURLBody(url *UrlLink) error {

	body, err := http.Get(url.Link)
	if err != nil {
		return err
	}
	fmt.Println(url.Body)
	//ioutil
	return nil
}

type UrlLink struct {
	Link *Url
}

func main() {

	url1 := &Url{Okay: true, URL: "https://preview.beamery.com", Body: ""}
	url2 := &Url{Okay: true, URL: "https://canary.beamery.com", Body: ""}
	url3 := &Url{Okay: true, URL: "https://beamery.com", Body: ""}

	//create array struct
	urls := []UrlLink{{Link: url1}, {Link: url2}, {Link: url3}}
	//urls := &UrlLink{{Link: *url1}, {Link: *url2}, {Link: *url3}}

	//urls := []string{"https://preview.beamery.com", "https://canary.beamery.com", "https://beamery.com"}

	printURLBody(urls)
	//printURLBody(url1)
	//printURLBody(url1)Url(urls)

	//wait for completion?

}
