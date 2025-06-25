package main

import (
	"GitHub-User-Activity/internal/helper"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Invalid command usage: Invalid number of arguments")
		os.Exit(1)
	}
	username := arguments[1]
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error connecting to GitHub\n", err.Error())
		os.Exit(1)
	}
	parsedResponseBody, err := helper.ParseResponseBody(response.Body)
	if err != nil {
		fmt.Println("Error parsing response body\n", err.Error())
	}
	switch response.StatusCode {
	case 200:
		helper.PrintPublicUserEvents(parsedResponseBody)
	case 404:
		fmt.Println("User not found")
	default:
		fmt.Println(response.StatusCode, response.Status)
	}
	if err := response.Body.Close(); err != nil {
		log.Fatal("Unable to close response body.\n", err.Error())
	}
}
