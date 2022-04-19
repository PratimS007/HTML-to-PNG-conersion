package main

import (
	"fmt"
	"os"
	"github.com/pdfcrowd/pdfcrowd-go"
)

func main() {
	// create the API client instance
	client := pdfcrowd.NewHtmlToImageClient( /*(your username at Pdfcrowd):*/ "demo" /*(your API Key:*/, "ce544b6ea52a5621fb9d55f8b542d14d")

	// configure the conversion
	client.SetOutputFormat("png")

	// run the conversion and write the result to a file
	err := client.ConvertFileToFile("AppMetrics.html", "out.png")

	// check for the conversion error
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		// report the error
		why, ok := err.(pdfcrowd.Error)
		if ok {
			os.Stderr.WriteString(fmt.Sprintf("Pdfcrowd Error: %s\n", why))
		} else {
			os.Stderr.WriteString(fmt.Sprintf("Generic Error: %s\n", err))
		}

		// rethrow or handle the exception
		panic(err.Error())
	}
}
