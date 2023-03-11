package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type Submission struct {
	id             string
	task_id        string
	language       string
	verdict        string
	number_of_test string
	time           string
	memory         string
}

func ResultsScrape() {
	url := "https://acm.timus.ru/status.aspx?author=342187&count=100&refresh=0"

	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var arr []Submission

	// Find the review items
	doc.Find(".even").Each(func(i int, s *goquery.Selection) {
		var current Submission

		// get problem id
		s.Find(".id").Each(func(i int, q *goquery.Selection) {
			current.id = q.Text()
		})

		// get task id
		s.Find(".problem").Each(func(i int, q *goquery.Selection) {
			current.task_id = q.Text()
		})

		// get language
		s.Find(".language").Each(func(i int, q *goquery.Selection) {
			current.language = q.Text()
		})

		flag := 0
		// get verdict
		s.Find(".verdict_rj").Each(func(i int, q *goquery.Selection) {
			current.verdict = q.Text()
			flag = 1
		})

		if flag == 0 {
			current.verdict = "Accepted"
		}

		// get number of test
		s.Find(".test").Each(func(i int, q *goquery.Selection) {
			current.number_of_test = q.Text()
		})

		// get time
		s.Find(".runtime").Each(func(i int, q *goquery.Selection) {
			current.time = q.Text()
		})

		// get memory
		s.Find(".memory").Each(func(i int, q *goquery.Selection) {
			current.memory = q.Text()
		})

		arr = append(arr, current)
	})

	doc.Find(".odd").Each(func(i int, s *goquery.Selection) {
		var current Submission

		// get problem id
		s.Find(".id").Each(func(i int, q *goquery.Selection) {
			current.id = q.Text()
		})

		// get task id
		s.Find(".problem").Each(func(i int, q *goquery.Selection) {
			current.task_id = q.Text()
		})

		// get language
		s.Find(".language").Each(func(i int, q *goquery.Selection) {
			current.language = q.Text()
		})

		flag := 0
		// get verdict
		s.Find(".verdict_rj").Each(func(i int, q *goquery.Selection) {
			current.verdict = q.Text()
			flag = 1
		})

		if flag == 0 {
			current.verdict = "Accepted"
		}

		// get number of test
		s.Find(".test").Each(func(i int, q *goquery.Selection) {
			current.number_of_test = q.Text()
		})

		// get time
		s.Find(".runtime").Each(func(i int, q *goquery.Selection) {
			current.time = q.Text()
		})

		// get memory
		s.Find(".memory").Each(func(i int, q *goquery.Selection) {
			current.memory = q.Text()
		})

		arr = append(arr, current)
	})

	fmt.Print(arr)
}

func main() {
	ResultsScrape()
}
