package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type AlgorithmArguments struct {
	Arguments string
}

type Count struct {
	count string
}

type AlgorithmOutput struct {
	result   int
	duration time.Duration
}

type ResponseData struct {
	Data       string
	Expression string
	Time       float64
}

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

		dad := executeAlgorithm(details)
		fmt.Println(dad)

		template.Execute(w, struct {
			Success  bool
			Response []ResponseData
		}{true, dad})
	})

	http.ListenAndServe(":4000", nil)
}

func executeAlgorithm(args AlgorithmArguments) []ResponseData {

	arguments := strings.Split(args.Arguments, "/n")

	var response []ResponseData

	for _, argument := range arguments {

		url := "http://localhost:5000/rpn/"

		requestBody, err := json.Marshal(map[string]string{
			"expression": argument,
		})

		if err != nil {
			fmt.Print(err)
		}

		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))

		res, _ := http.DefaultClient.Do(req)
		jsonDataFromHTTP, err := ioutil.ReadAll(res.Body)

		var jsonData []ResponseData

		err = json.Unmarshal([]byte(jsonDataFromHTTP), &jsonData) // here!

		if err != nil {
			panic(err)
		}
		response = jsonData
	}

	return response
}
