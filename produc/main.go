package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const ACCESS_TOKEN = "00HmmCrAwQChZJ3HPuBCTfo1BLDFZ-vEjsG-ctHM5zg" // replace with your actual access token

type User struct {
	Username string `json:"username"`
	UserID   string `json:"userID"`
}

func main() {
	productSlug := "multi-ai-chat" // replace with the actual product slug

	// Read users from JSON file
	users, err := readUsersFromFile("users.json")
	if err != nil {
		fmt.Println("Error reading users:", err)
		return
	}

	// Get product ID by slug
	productID, err := getProductIDBySlug(productSlug)
	if err != nil {
		fmt.Println("Error getting product ID:", err)
		return
	}

	// Check if users voted for the product
	for _, user := range users {
		voted, err := hasUserVoted(productID, user.UserID)
		if err != nil {
			fmt.Printf("Error checking if user %s voted: %v\n", user.Username, err)
			continue
		}

		if voted {
			fmt.Printf("User %s voted for the product.\n", user.Username)
		} else {
			fmt.Printf("User %s did not vote for the product.\n", user.Username)
		}
	}
}

func readUsersFromFile(filename string) ([]User, error) {
	var users []User

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func getProductIDBySlug(slug string) (string, error) {
	query := map[string]string{
		"query": fmt.Sprintf(`query { post(slug: "%s") { id } }`, slug),
	}

	queryJSON, err := json.Marshal(query)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.producthunt.com/v2/api/graphql", bytes.NewBuffer(queryJSON))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+ACCESS_TOKEN)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return "", err
	}

	data, ok := jsonResponse["data"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("unexpected response format: missing data field")
	}

	post, ok := data["post"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("unexpected response format: missing post field")
	}

	productID, ok := post["id"].(string)
	if !ok {
		return "", fmt.Errorf("unexpected response format: missing product ID")
	}

	return productID, nil
}

func hasUserVoted(productID string, userID string) (bool, error) {
	query := map[string]string{
		"query": fmt.Sprintf(`query { post(id: "%s") { votes { nodes { userId } } } }`, productID),
	}

	queryJSON, err := json.Marshal(query)
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest("POST", "https://api.producthunt.com/v2/api/graphql", bytes.NewBuffer(queryJSON))
	if err != nil {
		return false, err
	}

	req.Header.Set("Authorization", "Bearer "+ACCESS_TOKEN)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return false, err
	}

	data, ok := jsonResponse["data"].(map[string]interface{})
	if !ok {
		return false, fmt.Errorf("unexpected response format: missing data field")
	}

	post, ok := data["post"].(map[string]interface{})
	if !ok {
		return false, fmt.Errorf("unexpected response format: missing post field")
	}

	votes, ok := post["votes"].(map[string]interface{})
	if !ok {
		return false, fmt.Errorf("unexpected response format: missing votes field")
	}

	nodes, ok := votes["nodes"].([]interface{})
	if !ok {
		return false, fmt.Errorf("unexpected response format: missing nodes field")
	}

	for _, node := range nodes {
		if node.(map[string]interface{})["userId"] == userID {
			return true, nil
		}

	}

	return false, nil
}
