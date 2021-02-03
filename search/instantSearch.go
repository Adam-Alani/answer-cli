package search

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
	"strings"
)

func main() {
	fetchAnswer("whats the weather in villejuif")

}

type QueryResult struct {
	Definition       string
	DefinitionSource string
	Heading          string
	AbstractText     string
	Abstract         string
	AbstractSource   string
	Type             string
	AnswerType       string
	Redirect         string
	DefinitionURL    string
	Answer           string
	AbstractURL      string
	Results          []Result
}

type Result RelatedTopic
type RelatedTopic struct {
	Result   string
	Icon	 Icon
	FirstURL string
	Text     string
}
type Icon struct {
	URL    string
	Height interface{} // can be string or number ("16" or 16)
	Width  interface{} // can be string or number ("16" or 16)
}


func processQuery(query string) string {
	query = strings.Trim(query, " ")
	query = strings.ReplaceAll(query, " ", "+")
	return query
}

func (message *QueryResult) Decode(body []byte) error {
	if err := json.Unmarshal(body, message); err != nil {
		return err
	}

	return nil
}

func fetchAnswer(query string) (*QueryResult, error){
	query = url2.QueryEscape(query)
	var baseUrl = "https://api.duckduckgo.com/?q=%s&format=json&pretty=1%s"
	resp, err := http.Get(fmt.Sprintf(baseUrl, query, ""))
	if err != nil {
		log.Println("Error on response")
	}
	data, _ := ioutil.ReadAll(resp.Body)
	message := &QueryResult{}
	if err = message.Decode(data); err != nil {
		return nil, err
	}

	json.Unmarshal(data, &message)
	fmt.Println(message)
	return message, nil

}

