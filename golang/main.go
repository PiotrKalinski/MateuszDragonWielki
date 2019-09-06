package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

type AlgorithmArguments struct {
	Arguments string
}

type AlgorithmOutput struct {
	result   int
	duration time.Duration
}

func main() {
	template := template.Must(template.ParseFiles("layout.html"))
	fmt.Print("adad")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			template.Execute(w, nil)
			return
		}

		details := AlgorithmArguments{
			Arguments: r.FormValue("input"),
		}

		executeAlgorithm(details)

		template.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":4000", nil)
}

func executeAlgorithm(args AlgorithmArguments) AlgorithmOutput {

	start := time.Now()

	arguments := strings.Split(args.Arguments, "/n")

	for _, argument := range arguments {

		url := "http://127.0.0.1:5000/rpn"

		requestBody, err := json.Marshal(map[string]string{
			"expression": argument[1:],
		})

		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(bytes.NewBuffer(requestBody))

		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))

		res, _ := http.DefaultClient.Do(req)
		var result map[string]interface{}

		defer res.Body.Close()

		json.NewDecoder(res.Body).Decode(&result)
		log.Println(result["result"])
	}

	elapsed := time.Since(start)

	records := AlgorithmOutput{
		result:   12,
		duration: elapsed,
	}
	return records
}
