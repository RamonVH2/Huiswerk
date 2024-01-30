package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	ID       int
	Username string
	Email    string
	Modules  []string
}

type UserJSON struct {
	ID       int      `json:"UserID"`
	Username string   `json:"UserName"`
	Email    string   `json:"EmailAddress"`
	Modules  []string `json:"AccessToModules"`
}

func main() {
	user1 := UserJSON{
		ID:       1,
		Username: "kees",
		Email:    "kees@che.nl",
		Modules:  []string{"FIN", "SFC", "INV"},
	}
	user2 := UserJSON{
		ID:       2,
		Username: "piet",
		Email:    "piet@che.nl",
		Modules:  []string{"MRP", "INV"},
	}
	// nu wegschrijven naar een file
	data := []UserJSON{user1, user2}
	jsonData, marshallErr := json.MarshalIndent(data, "", "  ")
	if marshallErr != nil {
		fmt.Println("Error during MarshallIndent", marshallErr)
		return
	}
	writeFileErr := os.WriteFile("user.json", jsonData, 0644)
	if writeFileErr != nil {
		fmt.Println("Error during WriteFile", writeFileErr)
		return
	}
	// nu  deze file weer in te lezen
	usersJSON, readFileErr := os.ReadFile("user.json")
	if readFileErr != nil {
		fmt.Println("Error during reading file:", readFileErr)
		return
	}
	// slice maken om Users op te slaan
	var users []UserJSON

	// Unmarshal the JSON data into the slice of structs
	unmarshallErr := json.Unmarshal(usersJSON, &users)
	if unmarshallErr != nil {
		fmt.Println("Error unmarshalling JSON:", unmarshallErr)
		return
	}

	// Print the decoded data
	fmt.Println(users)
}
