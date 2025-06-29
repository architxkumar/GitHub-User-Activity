package helper

import (
	"GitHub-User-Activity/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

// WriteToCache writes parsed HTTP response body to local file for caching.
// If the file doesn't exist, it creates a new file else overwrites the file content
// with the updated content
func WriteToCache(username string, content []model.Event, fileName string) error {
	var output []byte
	newEntry := model.UserActivity{
		Username:  username,
		Content:   content,
		Timestamp: time.Now().UTC(),
	}
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if err != nil {
		return err
	}
	fileContents, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	if len(fileContents) != 0 {
		var userActivityList []model.UserActivity
		err = json.Unmarshal(fileContents, &userActivityList)
		if err != nil {
			return err
		}
		// -1 represents lack of entry in the userActivityList
		usernameIndex := -1
		for index, content := range userActivityList {
			if content.Username == username {
				usernameIndex = index
			}
		}
		if usernameIndex != -1 {
			userActivityList[usernameIndex].Content = newEntry.Content
			userActivityList[usernameIndex].Timestamp = newEntry.Timestamp
		} else {
			userActivityList = append(userActivityList, newEntry)
		}
		output, err = json.Marshal(userActivityList)
		if err != nil {
			return err
		}
		// To overwrite the contents in case data exists
		err = file.Truncate(0)
		if err != nil {
			return err
		}
		_, err = file.Seek(0, 0)
		if err != nil {
			return err
		}
	} else {
		var userActivityList []model.UserActivity
		userActivityList = append(userActivityList, newEntry)
		output, err = json.Marshal(userActivityList)
		if err != nil {
			return err
		}
	}
	_, err = file.Write(output)
	if err != nil {
		return err
	}
	return nil
}
