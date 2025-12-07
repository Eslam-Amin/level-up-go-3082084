package main

import (
	"encoding/json"
	"log"
	"os"
)

// User represents a user record.
type User struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

const path = "users.json"

// getBiggestMarket takes in the slice of users and
// returns the biggest market.
func getBiggestMarket(users []User) (string, int) {
	marketCounts := make(map[string]int)
	for _, user := range users {
		marketCounts[user.Country]++
	}
	var maxCountry string
	maxCount := 0
	for country, count := range marketCounts {
		if count > maxCount {
			maxCount = count
			maxCountry = country
		}
	}
	return maxCountry, maxCount
}

func main() {
	users := importData()
	country, count := getBiggestMarket(users)
	log.Printf("The biggest user market is %s with %d users.\n",
		country, count)
}

// importData reads the raffle entries from file and
// creates the entries slice.
func importData() []User {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	data := []User{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
