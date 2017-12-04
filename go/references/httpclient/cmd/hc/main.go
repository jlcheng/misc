package main

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
)

const API  = "http://api.open-notify.org/astros.json"

func main()  {

	//res, err := getUrl("http://www.google.com/robots.txt")
	//if err != nil {
	//	fmt.Printf("err: %v", err)
	//	return
	//}

	res := "{\"number\": 6, \"message\": \"success\", \"people\": [{\"name\": \"Sergey Ryazanskiy\", \"craft\": \"ISS\"}, {\"name\": \"Randy Bresnik\", \"craft\": \"ISS\"}, {\"name\": \"Paolo Nespoli\", \"craft\": \"ISS\"}, {\"name\": \"Alexander Misurkin\", \"craft\": \"ISS\"}, {\"name\": \"Mark Vande Hei\", \"craft\": \"ISS\"}, {\"name\": \"Joe Acaba\", \"craft\": \"ISS\"}]}"

	resultMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(res), &resultMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	peopleList := resultMap["people"].([]interface{})
	person := peopleList[0].(map[string]interface{})
	person["name"] = "Joe Schmoe"
	b, err := json.MarshalIndent(resultMap, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b[:]))
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