package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func main() {
	// Load .env file
	error := godotenv.Load(".env")
	if error != nil {
		panic("Error loading .env file")
	}

	// Create a token to and HTTP client to access the GitHub API
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	requestURL := "https://api.github.com/octocat"
	resp, error := httpClient.Get(requestURL)

	if error != nil {
		panic(error)
	}

	defer resp.Body.Close()

	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		panic(error)
	}

	fmt.Println(string(body))

	// Use GitHub GrapQL API via githubv4 library

	client := githubv4.NewClient(httpClient)
	fmt.Println(os.Getenv("GITHUB_TOKEN"))

	// Variable to store results of GraphQL query
	var q struct {
		Viewer struct {
			Login     string
			CreatedAt time.Time
		}
	}

	error = client.Query(context.Background(), &q, nil)
	if error != nil {
		panic(error)
	}

	fmt.Println(q.Viewer.Login)
	fmt.Println(q.Viewer.CreatedAt)
}