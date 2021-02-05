package search

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)


type GoogleResult struct {
	ResultRank int
	ResultURL string
	ResultTitle string
	ResultDesc string
}

func googleRequest(searchURL string) (*http.Response, error) {

	client := &http.Client{}

	request, _ := http.NewRequest("GET", searchURL, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	//request.Header.Set("Accept", "text/html")
	//request.Header.Set("Accept-Encoding", "gzip")
	//request.Header.Set("DNT", "1")


	response, err := client.Do(request)

	//fmt.Println(response.Header)

	if err != nil {
		return nil, err
	}
	return response, nil
}


func url(searchTerm string, countryCode string, languageCode string) string {
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	if googleBase, found := googleDomains[countryCode]; found {
		return fmt.Sprintf("%s%s&num=100&hl=%s", googleBase, searchTerm, languageCode)
	}
	fmt.Printf("%s%s&num=100&hl=%s", googleDomains["com"], searchTerm, languageCode)
	return fmt.Sprintf("%s%s&num=100&hl=%s", googleDomains["com"], searchTerm, languageCode)
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

		title := item.Find(`[aria-level="3"]`)
		titleText := title.Text()

		if (len(titleText)) == 0 {
			title = item.Find(`[role="presentation"]`)
			titleText = title.Text()
		}

		if len(titleText) > 0 {
			title.Contents().Each(func(i int, s *goquery.Selection) {


				fmt.Println("-----------------------------")
				trimmedString := strings.TrimSpace(s.Text())
				foundText := strings.Split(trimmedString, " ")
				fmt.Println(foundText[0])


			})
			//break
		}
		//descTag := item.Find("span")
		//desc := descTag.Text()
		//fmt.Println(desc)

	}
	return results, err
}


func Google(searchTerm string, countryCode string, languageCode string) ([]GoogleResult, error) {
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

