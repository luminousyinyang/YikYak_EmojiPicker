// MUST HAVE A REQUEST SNIFFER TO FIND THE AUTH HEADER, CSRFTOKEN, AND LOCATION
// can add an additional logical expression to the for loop if you want a specific color
// but YikYak has specific color schemes so you must find one of their hexadecimal colors

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter Authorization header: ")
	scanner.Scan()
	authHead := scanner.Text()
	fmt.Println("Enter csrftoken: ")
	scanner.Scan()
	csrftoken := scanner.Text()

	type RespData struct {
		Data struct {
			ResetConversationIcon struct {
				Emoji string `json:"emoji"`
				Color string `json:"color"`
			} `json:"resetConversationIcon"`
		} `json:"data"`
	}

	var rData RespData


	apiLink := "https://api.yikyak.com/graphql/"
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{
		Jar: jar,
	}

	// itterations, can remove if you wish
	itter := 1
	for (rData.Data.ResetConversationIcon.Emoji != "🥡") {
		fmt.Println(rData.Data.ResetConversationIcon.Emoji)
		fmt.Println(rData.Data.ResetConversationIcon.Color)
		fmt.Println(itter)
		itter++

	jsonBuffer := bytes.NewBuffer([]byte(`{
		"operationName":"ResetConversationIcon","query":"mutation ResetConversationIcon {\n  resetConversationIcon {\n    __typename\n    emoji\n    color\n    secondaryColor\n    errors {\n      __typename\n      code\n      field\n      message\n    }\n  }\n}","variables":null
	}`))

	req, err := http.NewRequest("POST", apiLink, jsonBuffer)
	if err != nil {
		panic(err)
	}

	

	req.Header.Add("Accept", "*/*")
	req.Header.Add("apollographql-client-version", "3.0.2-1")
	// auth header edit
	req.Header.Add("Authorization", authHead)
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")
	// location header edit.
	// req.Header.Add("Location", "POINT(x y)")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-APOLLO-OPERATION-TYPE", "mutation")
	req.Header.Add("User-Agent", "Yik%20Yak/1 CFNetwork/1402.0.8 Darwin/22.2.0")
	req.Header.Add("apollographql-client-name", "com.yikyak.2-apollo-ios")
	req.Header.Add("Connection", "keep-alive")
	// cookie header edit
	req.Header.Add("Cookie", "csrftoken=" + csrftoken)
	req.Header.Add("X-APOLLO-OPERATION-NAME", "ResetConversationIcon")

		resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
		
		body, err := ioutil.ReadAll(resp.Body)
		
		if resp.StatusCode != http.StatusOK {
			println(string(body))
			break;
		}

		
		
		err = json.Unmarshal(body, &rData)

	}
	fmt.Println(rData.Data.ResetConversationIcon.Emoji)
	fmt.Println(rData.Data.ResetConversationIcon.Color)
}

