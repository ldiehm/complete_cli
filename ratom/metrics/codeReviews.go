package metrics


import (
	"fmt"
	
)


func codeReviews(gitUrl string) float32 {
	fmt.Println(gitUrl)
	fmt.Println("codeReviews called")
	
	return 20
}
