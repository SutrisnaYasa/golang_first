package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var baseURL = "http://localhost:8089"

type student struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func fetchUser(ID string) (student, error) {
	var err error
	var client = &http.Client{}
	var data student

	// Set parameter URL dengan benar
	requestURL, err := url.Parse(baseURL + "/user")
	if err != nil {
		return data, err
	}

	query := requestURL.Query()
	query.Set("id", ID)
	requestURL.RawQuery = query.Encode()

	request, err := http.NewRequest("GET", requestURL.String(), nil)
	if err != nil {
		return data, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	// Cetak body respons mentah untuk debugging
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return data, err
	}
	fmt.Println("Raw Response Body:", string(body))

	// Decode respons JSON
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func main() {
	var user1, err = fetchUser("E001")
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", user1.ID, user1.Name, user1.Grade)
}
