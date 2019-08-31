package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"time"
)

type AlgorithmArguments struct {
	Arguments string
}

type AlgorithmOutput struct {
	result   int
	duration time.Duration
}

var count = 12

func main() {
	template := template.Must(template.ParseFiles("layout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			template.Execute(w, nil)
			return
		}

		details := AlgorithmArguments{
			Arguments: r.FormValue("input"),
		}

		fmt.Println(details.Arguments)

		log.Println(executeAlgorithm(details).duration.Seconds())

		template.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8080", nil)
}

func executeAlgorithm(args AlgorithmArguments) AlgorithmOutput {

	start := time.Now()

	formData := url.Values{
		"args": {args.Arguments},
	}
	log.Println(formData)

	resp, err := http.PostForm("https://d7fd62b6-895d-42cd-93fe-bb793d6c67b5.mock.pstmn.io/dupa", formData)

	if err != nil {
		log.Fatalln("d", err)
	}
	// do something with details
	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	log.Println(result["chuj"])

	elapsed := time.Since(start)

	records := AlgorithmOutput{
		result:   12,
		duration: elapsed,
	}
	return records
}
