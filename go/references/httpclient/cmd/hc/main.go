package main

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
)

const API  = "http://api.open-notify.org/astros.json"

type PeopleList struct {
	Number int `json:Number`
	People []People
}

type People struct {
	Name string `json:String`
}

func main()  {

	//res, err := getUrl("http://www.google.com/robots.txt")
	//if err != nil {
	//	fmt.Printf("err: %v", err)
	//	return
	//}

	res := "{\"number\": 6, \"message\": \"success\", \"people\": [{\"name\": \"Sergey Ryazanskiy\", \"craft\": \"ISS\"}, {\"name\": \"Randy Bresnik\", \"craft\": \"ISS\"}, {\"name\": \"Paolo Nespoli\", \"craft\": \"ISS\"}, {\"name\": \"Alexander Misurkin\", \"craft\": \"ISS\"}, {\"name\": \"Mark Vande Hei\", \"craft\": \"ISS\"}, {\"name\": \"Joe Acaba\", \"craft\": \"ISS\"}]}"

	people := &PeopleList{}
	err := json.Unmarshal([]byte(res), people)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*people)
}

func getUrl(url string) (string, error) {

	client := &http.Client{
		Timeout: time.Duration(time.Minute),
	}

	res, err := client.Get(url)
	if err != nil {
		return "", err
	}
	fmt.Printf("status: %s\n", res.Status)
	fmt.Printf("status code: %s\n", res.StatusCode)
	fmt.Printf("headers: %v\n", res.Header)

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}
	fmt.Println("")
	return fmt.Sprintf("%s\n", body), nil


}