package metrics


import (
	"fmt"
	
)


func VersionPinning(gitUrl string) float32 {
	fmt.Println(gitUrl)
	fmt.Println("codeReviews called")
	
	return 20
}
