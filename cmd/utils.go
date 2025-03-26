package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func printResponse(resp interface{}) {
	jsonData, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error marshalling response: %v\n", err)
		return
	}
	fmt.Println(string(jsonData))
}

func prettyPrintResponse(r *http.Response) {
	if r == nil {
		fmt.Println("No response received")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	defer r.Body.Close()

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body, "", "    ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}

	fmt.Println(prettyJSON.String())
}
