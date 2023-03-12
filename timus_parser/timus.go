package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"os"
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
	url_ := "https://acm.timus.ru/status.aspx?author=342187&count=100&refresh=0"

	// Request the HTML page.
	res, err := http.Get(url_)
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

func SendSubmission() {
	url_ := "https://acm.timus.ru/submit.aspx"

	r := url.Values{
		"action":     {"submit"},
		"SpaceID":    {"1"},
		"JudgeID":    {"342187EL"},
		"Language":   {"65"},
		"ProblemNum": {"1000"},
		"Source":     {"abacaba"},
	}

	resp, err := http.PostForm(url_, r)

	if err != nil {
		return
	}

	defer resp.Body.Close()
}

func GetTaskHtml(url_ string) {
	// Load the HTML document
	doc, err := goquery.NewDocument(url_)
	if err != nil {
		log.Fatal(err)
	}

	// Find the div with class "problem_content"
	problemContent := doc.Find("div.problem_content")

	// Get the HTML content of the div
	htmlContent, err := problemContent.Html()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new file to write the content to
	file, err := os.Create("output.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write the content to the file
	_, err = file.WriteString(htmlContent)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Content written to output.html")
}

func main() {
	// ResultsScrape()
	// SendSubmission()
	// GetTaskHtml("https://acm.timus.ru/problem.aspx?space=1&num=1000")
}
