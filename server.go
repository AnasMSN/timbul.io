package main

import (
	"encoding/json"
	"fmt"

	"timbul.io/lib/mongodb"
)

func main() {

	db := mongodb.New()

	result := db.QueryOne("sample_mflix", "movies", "title", "Back to the Future")
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}
