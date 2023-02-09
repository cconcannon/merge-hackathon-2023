package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func handleRequests() {
	http.HandleFunc("/", respond)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func respond(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	message := getTextValue(string(body))
	fmt.Fprintf(w, answer(message))
}

func answer(query string) string {
	message, _ := url.QueryUnescape(query)
	var response string
	switch {
	case strings.Contains(message, "batch"):
		url := "https://docs.sourcegraph.com/batch_changes/how-tos/creating_a_batch_change"
		response = fmt.Sprintf("Here's our how-to guide for creating a batch change: %s", url)
	case strings.Contains(message, "insights"):
		url := "https://docs.sourcegraph.com/code_insights/quickstart"
		response = fmt.Sprintf("Here's a link to our quickstart for code insights: %s", url)
	case strings.Contains(message, "executors"):
		url := "https://docs.sourcegraph.com/admin/deploy_executors"
		response = fmt.Sprintf("Here's a link to our guide for deploying executors: %s", url)
	default:
		response = fmt.Sprint("We don't have a specific suggestion for you, but you can try searching our documentation at https://docs.sourcegraph.com")
	}
	return response
}

func main() {
	handleRequests()
}

func getTextValue(payload string) string {
	re := regexp.MustCompile(`text=(.+)&api`)
	match := re.FindStringSubmatch(payload)
	var result string
	if len(match) > 1 {
		result = match[1]
	} else {
		result = ""
	}
	return result
}
