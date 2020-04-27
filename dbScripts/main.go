package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	DB "travel-planner/database"
	Config "travel-planner/utils"
)

type CountryResponse struct {
	Name string
	Code string
}

func main() {
	config := Config.LoadBasiDataScrapperConfig()
	fmt.Println("Starting the application...")
	url := "https://travelpayouts-travelpayouts-flight-data-v1.p.rapidapi.com/data/en-GB/countries.json"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", config.TravelApiHost)
	req.Header.Add("x-rapidapi-key", config.ApiKey)
	req.Header.Add("x-access-token", config.AccessToken)

	res, err := http.DefaultClient.Do(req)
	if err == nil {
		fmt.Println("Error in the GET request")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	var decoded []CountryResponse
	if err := json.Unmarshal(body, &decoded); err != nil {
		panic(err)
	}

	db, err := DB.Connect()
	for _, record := range decoded {
		fmt.Printf("Name: %s Code: %s\n", record.Name, record.Code)
		DB.Insert(db, record.Name, record.Code)

		fmt.Println("")
	}
	fmt.Println(decoded[0])

	fmt.Println("Terminating the application...")
	DB.CloseConnection(db)
}
