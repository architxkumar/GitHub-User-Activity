package main

import (
	"GitHub-User-Activity/internal/helper"
	"fmt"
	"os"
)

const cacheFilePath = "cache.json"

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Invalid command usage: Invalid number of arguments")
		os.Exit(1)
	}
	username := arguments[1]
	fileContents, err := helper.ReadCacheFileContents(cacheFilePath)
	if fileContents == nil || err != nil {
		if err != nil {
			fmt.Println("Cache file not found")
		} else {
			fmt.Println("Cache file is empty")
		}
		makeRequestAndShowResult(username)
	} else {
		userCache, err := helper.GetUserEventsFromCache(fileContents, username)
		if err != nil {
			fmt.Println("Error reading user contents from cache")
			makeRequestAndShowResult(username)
		}
		isCacheValid := helper.CheckCacheValidity(userCache)
		if isCacheValid && userCache != nil {
			helper.PrintPublicUserEvents(userCache.Content)
		} else {
			response, err := helper.MakeRequest(username)
			if err != nil {
				fmt.Println("Error connecting to GitHub")
				if userCache != nil {
					fmt.Println("Show stale contents from cache")
					helper.PrintPublicUserEvents(userCache.Content)
				} else {
					os.Exit(1)
				}
			} else {
				eventList, err := helper.ParseResponseBody(response.Body)
				if err != nil {
					fmt.Println("Error parsing response body", err.Error())
					os.Exit(1)
				}
				helper.PrintPublicUserEvents(eventList)
				err = helper.WriteToCache(username, eventList, cacheFilePath)
				if err != nil {
					fmt.Println("Error writing to cache", err.Error())
				}
				defer func() {
					err := response.Body.Close()
					if err != nil {
						fmt.Println("Error closing response body")
					}
				}()
			}
		}
	}
}

func makeRequestAndShowResult(username string) {
	response, err := helper.MakeRequest(username)
	if err != nil {
		fmt.Println("Error connecting to GitHub")
		os.Exit(1)
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Println("Error closing response body")
		}
	}()
	eventList, err := helper.ParseResponseBody(response.Body)
	if err != nil {
		fmt.Println("Error parsing response body")
		os.Exit(1)
	}
	helper.PrintPublicUserEvents(eventList)
	err = helper.WriteToCache(username, eventList, cacheFilePath)
	if err != nil {
		fmt.Println("Error writing to cache")
	}
}
