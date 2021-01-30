package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

func main() {
	_, err := Search("how do you spell christmas", "com", "en")
	if err != nil {
		fmt.Println("ohno")
	}
}

type GoogleResult struct {
	ResultRank int
	ResultURL string
	ResultTitle string
	ResultDesc string
}

var googleDomains = map[string]string{
	"com": "https://www.google.com/search?q=",
	"uk": "https://www.google.co.uk/search?q=",
	"ru": "https://www.google.ru/search?q=",
	"fr": "https://www.google.fr/search?q=",
}


func url(searchTerm string, countryCode string, languageCode string) string {
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	if googleBase, found := googleDomains[countryCode]; found {
		return fmt.Sprintf("%s%s&num=100&hl=%s", googleBase, searchTerm, languageCode)
	}
	return fmt.Sprintf("%s%s&num=100&hl=%s", googleDomains["com"], searchTerm, languageCode)
}

func googleRequest(searchURL string) (*http.Response, error) {

	client := &http.Client{}

	request, _ := http.NewRequest("GET", searchURL, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	return response, nil
}

func parseQuery(response *http.Response) ([]GoogleResult, error){
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}
	results := []GoogleResult{}
	sel := doc.Find("div.mod")
	for i := range sel.Nodes {
		item := sel.Eq(i)
		//fmt.Println(item.Text())
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		titleTag := item.Find("h2")
		descTag := item.Find("span")
		desc := descTag.Text()
		title := titleTag.Text()
		fmt.Println("\033[37m", title)
		fmt.Println()
		fmt.Println(desc)
		link = strings.Trim(link, " ")


	}
	return results, err
}


func Search(searchTerm string, countryCode string, languageCode string) ([]GoogleResult, error) {
	googleUrl := url(searchTerm, countryCode, languageCode)
	response, err := googleRequest(googleUrl)
	if err != nil {
		return nil, err
	}
	result, err := parseQuery(response)
	if err != nil {
		return nil, err
	}
	return result, nil
}

