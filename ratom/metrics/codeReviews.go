package metrics


import (
	"fmt"
	
)


func CodeReviews(gitUrl string) float32 {
	fmt.Println(gitUrl)
	fmt.Println("codeReviews called")
	
	return 20
}
