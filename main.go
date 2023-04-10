package main

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strconv"

	//"sort"

	"ece461-project-cli/ratom"
	"ece461-project-cli/ratom/metrics"
	"golang.org/x/oauth2"
	// "github.com/ldiehm/complete_cli/ratom"
	// "github.com/ldiehm/complete_cli/ratom/metrics"
	// "golang.org/x/oauth2"
)

// // Function to get the current line number - citation needed
// func main() {
// 	// Check if input file can be opened for read
// 	filePath := os.Args[1]

// 	file, err := os.Open(filePath)

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	scanner.Split(bufio.ScanLines)

// 	// // Load environment variables from .env file
// 	// err = godotenv.Load(".env")
// 	// if err != nil {
// 	// 	panic("Error loading .env file")
// 	// }

// 	// Create a token to and HTTP client to access the GitHub API
// 	src := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
// 	)
// 	httpClient := oauth2.NewClient(context.Background(), src)

// 	// Create temp directory to store the directories for each of the repositories
// 	_, err = os.Stat("temp")
//     if !os.IsNotExist(err) {
//         os.RemoveAll("temp")
//     }

// 	err = os.Mkdir("temp", 0770)		// read, write, and execute access for file owner and group owner
// 	if err != nil {
// 		panic("Error creating temp directory")
// 	}

// 	// Determine log level
// 	logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))

// 	if err != nil {
// 		fmt.Println("Error during string to int conversion")
// 		return
// 	}

// 	var modules []ratom.Module

// 	// Read file line by line
// 	for scanner.Scan() {
// 		url := scanner.Text()
// 		metrics.Functions = append(metrics.Functions, "-------------------------------------")
// 		metrics.Functions = append(metrics.Functions, "For URL: "+url)
// 		module := ratom.Analyze(url, httpClient)
// 		lineNumb := metrics.File_line()
// 		metrics.Functions = append(metrics.Functions, "Function: ratom.Analyze called on main.go at line "+lineNumb)
// 		//fmt.Println(module)
// 		modules = append(modules, module)
// 		//urls = append(urls, url)
// 	}

// 	os.RemoveAll("temp")

// 	// Sorting the modules by decreasing net score
// 	sort.SliceStable(modules, func(i, j int) bool {
// 		return modules[i].NetScore > modules[j].NetScore
// 	})
// 	lineNumb := metrics.File_line()
// 	metrics.Functions = append(metrics.Functions, "Function: SliceStable called on main.go at line "+lineNumb)

// 	for a := 0; a < len(modules); a++ {
// 		var b = 0
// 		if modules[a].License {
// 			b = 1
// 		}
// 		fmt.Printf("{\"URL\":\"%s\", \"NET_SCORE\":%1.1f, \"RAMP_UP_SCORE\":%1.1f, \"CORRECTNESS_SCORE\":%1.1f, \"BUS_FACTOR_SCORE\":%1.1f, \"RESPONSIVE_MAINTAINER_SCORE\":%1.1f, \"LICENSE_SCORE\":%d}\n", modules[a].Url, modules[a].NetScore, modules[a].RampUp, modules[a].Correctness, modules[a].BusFactor, modules[a].RespMaint, b)
// 		//fmt.Println(modules[a])
// 	}

// 	if logLevel == 1 {
// 		ratom.LoggerVerbOne(modules)
// 	} else if logLevel == 2 {
// 		ratom.LoggerVerbTwo(modules)
// 	}

// }

func main() {
	// Check if input file can be opened for read
	inputString_url := os.Args[1]

	// // Load environment variables from .env file
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	panic("Error loading .env file")
	// }

	// Create a token to and HTTP client to access the GitHub API
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	// Create temp directory to store the directories for each of the repositories
	_, err := os.Stat("temp")
	if !os.IsNotExist(err) {
		os.RemoveAll("temp")
	}

	err = os.Mkdir("temp", 0770) // read, write, and execute access for file owner and group owner
	if err != nil {
		panic("Error creating temp directory")
	}

	// Determine log level
	logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))

	if err != nil {
		fmt.Println("Error during string to int conversion")
		return
	}

	var modules []ratom.Module

	// Read file line by line
	// for scanner.Scan() {
		// url := scanner.Text()
		metrics.Functions = append(metrics.Functions, "-------------------------------------")
		metrics.Functions = append(metrics.Functions, "For URL: "+inputString_url)
		module := ratom.Analyze(inputString_url, httpClient)
		lineNumb := metrics.File_line()
		metrics.Functions = append(metrics.Functions, "Function: ratom.Analyze called on main.go at line "+lineNumb)
		//fmt.Println(module)
		modules = append(modules, module)
		//urls = append(urls, url)
	// }

	os.RemoveAll("temp")

	// Sorting the modules by decreasing net score
	sort.SliceStable(modules, func(i, j int) bool {
		return modules[i].NetScore > modules[j].NetScore
	})
	lineNumb = metrics.File_line()
	metrics.Functions = append(metrics.Functions, "Function: SliceStable called on main.go at line "+lineNumb)

	for a := 0; a < len(modules); a++ {
		var b = 0
		if modules[a].License {
			b = 1
		}
		fmt.Printf("{\"URL\":\"%s\", \"NET_SCORE\":%1.1f, \"RAMP_UP_SCORE\":%1.1f, \"CORRECTNESS_SCORE\":%1.1f, \"BUS_FACTOR_SCORE\":%1.1f, \"RESPONSIVE_MAINTAINER_SCORE\":%1.1f, \"LICENSE_SCORE\":%d, \"VERSION PINNING SCORE\":%d}\n", modules[a].Url, modules[a].NetScore, modules[a].RampUp, modules[a].Correctness, modules[a].BusFactor, modules[a].RespMaint, modules[a].versionPinning_score, b)
		//fmt.Println(modules[a])
	}

	if logLevel == 1 {
		ratom.LoggerVerbOne(modules)
	} else if logLevel == 2 {
		ratom.LoggerVerbTwo(modules)
	}

}
