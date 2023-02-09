package ratom

import (
	"context"
	"os"
	"testing"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/paingp/ece461-project-cli/ratom/metrics"
	"golang.org/x/oauth2"
)

var file = "https://github.com/cloudinary/cloudinary_npm"
var endpoint = "https://api.github.com/repos/cloudinary/cloudinary_npm"

// Test cases 1 - 6
func TestScore(t *testing.T) {
	// Create a token to and HTTP client to access the GitHub API
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	module := Analyze(file, httpClient)

	if module.Url != file {
		t.Fatalf("File not equal")
	}

	if module.BusFactor > 1 || module.BusFactor < 0 {
		t.Fatalf("Bus Factor score out of range")
	}

	if module.Correctness > 1 || module.Correctness < 0 {
		t.Fatalf("Correctness score out of range")
	}

	if module.NetScore > 1 || module.NetScore < 0 {
		t.Fatalf("NetScore score out of range")
	}

	if module.RampUp > 1 || module.RampUp < 0 {
		t.Fatalf("Ramp Up score out of range")
	}

	if module.RespMaint > 1 || module.RespMaint < 0 {
		t.Fatalf("Responsive Maintaner score out of range")
	}
}

func TestBusFactor(t *testing.T) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	resp, error := httpClient.Get(endpoint)

	if error != nil {
		panic(error)
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			panic(error)
		}
		bodyString := string(bodyBytes)

		// Citation needed for this
		resBytes := []byte(bodyString)
		var jsonRes map[string]interface{}
		_ = json.Unmarshal(resBytes, &jsonRes)

		var bus = metrics.BusFactor(jsonRes)

		if bus < 0 || bus > 1 {
			t.Fatalf("Bus is out of range")
		}
	}

	defer resp.Body.Close()
}

func TestCorrectness(t *testing.T) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	resp, error := httpClient.Get(endpoint)

	if error != nil {
		panic(error)
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			panic(error)
		}
		bodyString := string(bodyBytes)

		// Citation needed for this
		resBytes := []byte(bodyString)
		var jsonRes map[string]interface{}
		_ = json.Unmarshal(resBytes, &jsonRes)

		var cor = metrics.Correctness(jsonRes)

		if cor < 0 || cor > 1 {
			t.Fatalf("Bus is out of range")
		}
	}

	defer resp.Body.Close()

}

// func testLicense(t *testing.T) {

// }

// func testNetScore(t *testing.T) {

// }

// func testRampUp(t *testing.T) {

// }

// func testResponsiveMaintainer(t *testing.T) {

// }
