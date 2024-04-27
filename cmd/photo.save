package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func uploadFiles(directoryPath string, endpointUrl string, authToken string) error {
	files, err := os.ReadDir(directoryPath)

	CheckError(err)

	client := &http.Client{}

	for _, file := range files {
		filePath := filepath.Join(directoryPath, file.Name())

		file, err := os.Open(filePath)
		CheckError(err)

		defer file.Close()

		req, err := http.NewRequest("POST", endpointUrl, file)
		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", "image/jpeg")
		req.Header.Set("Authorization", "Bearer "+authToken)
		fmt.Println("Content-Type " + req.Header.Get("Content-Type"))

		resp, err := client.Do(req)
		CheckError(err)

		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			fmt.Println("Response code: " + resp.Status)

			responseBody, err := io.ReadAll(resp.Body)

			CheckError(err)

			fmt.Println("Response Body: " + string(responseBody))
		}

		fmt.Printf("File %s uploaded\n\n", filePath)
		time.Sleep(1 * time.Second)
	}
	return nil
}

func CheckError(err error) error {
	if err != nil {
		return err
	}
	return nil
}
