package main

import (
	"database/sql"
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

type CityResponse struct {
	Name         string
	Code         string
	Country_code string
}

type AirportResponse struct {
	Name      string
	Code      string
	City_code string
}

func getData(url string) (body []byte, err error) {
	config := Config.LoadBasiDataScrapperConfig()
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", config.TravelApiHost)
	req.Header.Add("x-rapidapi-key", config.ApiKey)
	req.Header.Add("x-access-token", config.AccessToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error in the GET request")
	}

	defer res.Body.Close()
	body, _ = ioutil.ReadAll(res.Body)
	return body, err
}

func getCountriesData(db *sql.DB, url string) {
	body, err := getData(url)
	if err != nil {
		panic(err)
	}

	var decoded []CountryResponse
	if err := json.Unmarshal(body, &decoded); err != nil {
		fmt.Println(body)
		panic(err)
	}

	for _, record := range decoded {
		fmt.Printf("Name: %s Code: %s\n", record.Name, record.Code)
		DB.Insert(db, "country", record.Name, record.Code)

		fmt.Println("")
	}
	fmt.Println(decoded[0])

}

func getCitiesData(db *sql.DB, url string) {
	body, err := getData(url)
	if err != nil {
		panic(err)
	}

	// fmt.Println(res)
	// fmt.Println(string(body))

	var decoded []AirportResponse
	if err := json.Unmarshal(body, &decoded); err != nil {
		fmt.Println(body)
		panic(err)
	}

	for _, record := range decoded {
		fmt.Printf("Name: %s Code: %s CityCode: %s\n", record.Name, record.Code, record.City_code)
		DB.Insert(db, "city", record.Name, record.Code)
		DB.InsertConnectingData(db, "country_cities", record.City_code, record.Code, "country", "city", "country_id", "city_id")

		fmt.Println("")
	}
	fmt.Println(decoded[0])

}

func getAirportData(db *sql.DB, url string) {
	body, err := getData(url)
	if err != nil {
		panic(err)
	}

	// fmt.Println(res)
	// fmt.Println(string(body))

	var decoded []AirportResponse
	if err := json.Unmarshal(body, &decoded); err != nil {
		fmt.Println(body)
		panic(err)
	}

	for _, record := range decoded {
		fmt.Printf("Name: %s Code: %s CityCode: %s\n", record.Name, record.Code, record.City_code)
		DB.Insert(db, "airport", record.Name, record.Code)
		DB.InsertConnectingData(db, "city_airports", record.City_code, record.Code, "city", "airport", "city_id", "airport_id")

		fmt.Println("")
	}
	fmt.Println(decoded[0])
}

func main() {
	db, err := DB.Connect()
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the application...")
	url := "https://travelpayouts-travelpayouts-flight-data-v1.p.rapidapi.com/data/en-GB/countries.json"
	getCountriesData(db, url)
	url = "https://travelpayouts-travelpayouts-flight-data-v1.p.rapidapi.com/data/en-GB/cities.json"
	getCitiesData(db, url)
	url = "https://travelpayouts-travelpayouts-flight-data-v1.p.rapidapi.com/data/en-GB/airports.json"
	getAirportData(db, url)

	fmt.Println("Terminating the application...")
	DB.CloseConnection(db)
}
